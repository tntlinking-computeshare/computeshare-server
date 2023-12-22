package biz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/mohaijiang/computeshare-server/internal/utils"
	"github.com/samber/lo"
	"io"
	"os"
	"strings"
	"time"
)

type S3User struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"FkUserID,omitempty"`
	Type     int8      `json:"type,omitempty"`
	// accessKey
	AccessKey string `json:"accessKey,omitempty"`
	// secretKey
	SecretKey  string    `json:"secret_key,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	// Endpoint
	Endpoint string `json:"endpoint"`
}

type S3Bucket struct {
	ID          uuid.UUID `json:"id,omitempty"`
	FkUserID    uuid.UUID `json:"FkUserID,omitempty"`
	BucketName  string    `json:"bucketName,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
}

type S3UserRepo interface {
	SetValidateCode(ctx context.Context, entity User, vCode string) error
	GetValidateCode(ctx context.Context, telephone string) (string, error)
	DeleteValidateCode(ctx context.Context, user User)
	CreateS3User(ctx context.Context, user *S3User) (*S3User, error)
	GetS3User(ctx context.Context, id uuid.UUID) (*S3User, error)
	DeleteS3User(ctx context.Context, id uuid.UUID) error
	GetUserS3User(ctx context.Context, userId uuid.UUID) ([]*S3User, error)
	GetPlatformS3User(ctx context.Context, userId uuid.UUID) (*S3User, error)
	CreateBucket(ctx context.Context, user *S3User, bucket string) (*S3Bucket, error)
	DeleteBucket(ctx context.Context, user *S3User, bucketName string) error
	ListBucket(ctx context.Context, userId uuid.UUID) ([]*S3Bucket, error)
}

type StorageS3UseCase struct {
	repo     S3UserRepo
	userRepo UserRepo
	log      *log.Helper
	dispose  conf.Dispose
}

func NewStorageS3UseCase(
	repo S3UserRepo,
	userRepo UserRepo,
	confDispose *conf.Dispose,
	logger log.Logger) *StorageS3UseCase {
	return &StorageS3UseCase{
		repo:     repo,
		userRepo: userRepo,
		dispose:  *confDispose,
		log:      log.NewHelper(logger),
	}
}

func (c *StorageS3UseCase) createS3User(ctx context.Context, userId uuid.UUID, creator int8) (*S3User, error) {
	_, err := c.userRepo.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	accessKey, err := utils.GenerateRandomString(20)
	if err != nil {
		return nil, err
	}
	secretKey, err := utils.GenerateRandomString(40)
	if err != nil {
		return nil, err
	}
	s3User := &S3User{
		FkUserID:   userId,
		AccessKey:  accessKey,
		SecretKey:  secretKey,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Type:       creator,
	}
	return c.repo.CreateS3User(ctx, s3User)
}

func (c *StorageS3UseCase) CreateS3Key(ctx context.Context, userId uuid.UUID) error {
	_, err := c.createS3User(ctx, userId, int8(consts.UserCreation))
	if err != nil {
		return err
	}
	err = c.RefreshUserS3UserPermissions(ctx, userId)
	return err
}

func (c *StorageS3UseCase) GetUserS3UserList(ctx context.Context, userId uuid.UUID) ([]*S3User, error) {
	return c.repo.GetUserS3User(ctx, userId)
}

func (c *StorageS3UseCase) GetUserS3User(ctx context.Context, id string, countryCallCoding, telephoneNumber, validateCode string) (*S3User, error) {
	code, err := c.repo.GetValidateCode(ctx, strings.Join([]string{countryCallCoding, telephoneNumber}, ""))
	if err != nil || code != validateCode {
		return nil, errors.New("VALIDATE_CODE_ERROR")
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return c.repo.GetS3User(ctx, uuid)
}

func (c *StorageS3UseCase) DeleteUserS3User(ctx context.Context, id string, countryCallCoding, telephoneNumber, validateCode string) error {
	code, err := c.repo.GetValidateCode(ctx, strings.Join([]string{countryCallCoding, telephoneNumber}, ""))
	if err != nil || code != validateCode {
		return errors.New("VALIDATE_CODE_ERROR")
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	user, err := c.repo.GetS3User(ctx, uuid)
	if err != nil {
		return err
	}
	s3Buckets, err := c.ListBucket(ctx, user.FkUserID)
	if err != nil {
		return err
	}

	buckets := lo.Reduce(s3Buckets, func(agg string, item *S3Bucket, index int) string {
		if agg == "" {
			return item.BucketName
		}
		return strings.Join([]string{agg, item.BucketName}, ",")
	}, "")
	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.WithHost(c.dispose.S3.TargetDockerHost))
	if err != nil {
		return err
	}
	action := "Read,Write,List,Tagging"
	cmd := fmt.Sprintf(`echo "s3.configure -access_key=%s -secret_key=%s -buckets=%s -user=%s -actions=%s -delete -apply" | weed shell`,
		user.AccessKey, user.SecretKey, buckets, user.ID.String(), action)
	c.S3UserConfigure(cli, cmd)
	return c.repo.DeleteS3User(ctx, uuid)
}

func (c *StorageS3UseCase) CreateBucket(ctx context.Context, userId uuid.UUID, bucket string) (*S3Bucket, error) {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		if ent.IsNotFound(err) {
			platformS3User, err = c.createS3User(ctx, userId, int8(consts.PlatformCreation))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	s3Bucket, err := c.repo.CreateBucket(ctx, platformS3User, bucket)
	if err != nil {
		return nil, err
	}
	s3Buckets, err := c.ListBucket(ctx, userId)
	if err != nil {
		return nil, err
	}

	buckets := lo.Reduce(s3Buckets, func(agg string, item *S3Bucket, index int) string {
		if agg == "" {
			return item.BucketName
		}
		return strings.Join([]string{agg, item.BucketName}, ",")
	}, "")

	action := "Read,Write,List,Tagging,Admin"
	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.WithHost(c.dispose.S3.TargetDockerHost))
	cmd := fmt.Sprintf(`echo "s3.configure -access_key=%s -secret_key=%s -buckets=%s -user=%s -actions=%s -apply " | weed shell`,
		platformS3User.AccessKey, platformS3User.SecretKey, buckets, platformS3User.FkUserID.String(), action)
	err = c.S3UserConfigure(cli, cmd)
	if err != nil {
		return nil, err
	}

	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(session)
	if err != nil {
		return nil, err
	}
	_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}
	return s3Bucket, err
}

func (c *StorageS3UseCase) DeleteBucket(ctx context.Context, userId uuid.UUID, bucketName string) error {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return err
	}
	err = c.repo.DeleteBucket(ctx, platformS3User, bucketName)
	if err != nil {
		return err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return err
	}
	s3Client := s3.New(session)
	if err != nil {
		return err
	}
	_, err = s3Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *StorageS3UseCase) EmptyBucket(ctx context.Context, userId uuid.UUID, bucketName string) error {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return err
	}
	s3Client := s3.New(session)
	if err != nil {
		return err
	}
	resp, err := s3Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucketName)})
	if len(resp.Contents) < 1 {
		return nil
	}
	var ObjectIdentifierList []*s3.ObjectIdentifier
	for _, content := range resp.Contents {
		var ObjectIdentifier s3.ObjectIdentifier
		ObjectIdentifier.Key = content.Key
		ObjectIdentifierList = append(ObjectIdentifierList, &ObjectIdentifier)
	}
	deleteObjectsInput := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3.Delete{
			Objects: ObjectIdentifierList,
			Quiet:   aws.Bool(false),
		},
	}
	_, err = s3Client.DeleteObjects(deleteObjectsInput)
	if err != nil {
		return err
	}
	return nil
}

func (c *StorageS3UseCase) ListBucket(ctx context.Context, userId uuid.UUID) ([]*S3Bucket, error) {
	return c.repo.ListBucket(ctx, userId)
}

func (c *StorageS3UseCase) S3StorageInBucketList(ctx context.Context, userId uuid.UUID, bucketName, prefix string) ([]*pb.S3Object, error) {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return nil, err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(session)
	if err != nil {
		return nil, err
	}
	resp, err := s3Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucketName), Prefix: aws.String(prefix)})
	if err != nil {
		return nil, err
	}
	var s3ObjectList []*pb.S3Object
	for _, object := range resp.Contents {
		var s3Object pb.S3Object
		key := *object.Key
		eTag := *object.ETag
		if strings.Contains(eTag, "\"") {
			eTag = eTag[1 : len(eTag)-1]
		}
		if prefix == "" {
			splitN := strings.SplitN(key, "/", 2)
			if len(splitN) > 1 && !containsDir(s3ObjectList, splitN[0]+"/") {
				s3Object.Name = splitN[0] + "/"
			} else if len(splitN) == 1 {
				if splitN[0] == ".keep" {
					continue
				}
				s3Object.Etag = eTag
				s3Object.LastModify = object.LastModified.UnixMilli()
				s3Object.Size = int32(*object.Size)
				s3Object.Name = splitN[0]
				s3Object.S3Url = c.dispose.S3.S3UrlPrefix + bucketName + "/" + key
				s3Object.Url = c.dispose.S3.UrlPrefix + bucketName + "/" + key
			} else {
				continue
			}
		} else {
			dir := prefix + "/"
			if key[:len(dir)] == dir {
				splitN := strings.SplitN(key[len(dir):], "/", 2)
				if len(splitN) > 1 && !containsDir(s3ObjectList, splitN[0]+"/") {
					s3Object.Name = splitN[0] + "/"
				} else if len(splitN) == 1 {
					if splitN[0] == ".keep" {
						continue
					}
					s3Object.Etag = eTag
					s3Object.LastModify = object.LastModified.UnixMilli()
					s3Object.Size = int32(*object.Size)
					s3Object.Name = splitN[0]
					s3Object.S3Url = c.dispose.S3.S3UrlPrefix + bucketName + "/" + key
					s3Object.Url = c.dispose.S3.UrlPrefix + bucketName + "/" + key
				} else {
					continue
				}
			}
		}
		s3ObjectList = append(s3ObjectList, &s3Object)
	}
	return s3ObjectList, nil
}

func containsDir(slice []*pb.S3Object, target string) bool {
	for _, s := range slice {
		if s.Name == target {
			return true
		}
	}
	return false
}

func (c *StorageS3UseCase) S3StorageUploadFile(ctx context.Context, userId uuid.UUID, bucketName, key string, fileByte []byte) (*s3.PutObjectOutput, error) {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return nil, err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(session)
	if err != nil {
		return nil, err
	}
	putObject, err := s3Client.PutObjectWithContext(context.Background(), &s3.PutObjectInput{
		Body:   bytes.NewReader(fileByte),
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return putObject, nil
}

func (c *StorageS3UseCase) S3StorageMkdir(ctx context.Context, userId uuid.UUID, bucketName, dirName string) error {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return err
	}
	s3Client := s3.New(session)
	if err != nil {
		return err
	}
	_, err = s3Client.PutObjectWithContext(context.Background(), &s3.PutObjectInput{
		Body:   bytes.NewReader([]byte("")),
		Bucket: aws.String(bucketName),
		Key:    aws.String(dirName + "/.keep"),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *StorageS3UseCase) S3StorageDeleteMkdir(ctx context.Context, userId uuid.UUID, bucketName, dirName string) error {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return err
	}
	s3Client := s3.New(session)
	if err != nil {
		return err
	}
	resp, err := s3Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucketName), Prefix: aws.String(dirName)})
	if len(resp.Contents) < 1 {
		return nil
	}
	var ObjectIdentifierList []*s3.ObjectIdentifier
	for _, content := range resp.Contents {
		var ObjectIdentifier s3.ObjectIdentifier
		ObjectIdentifier.Key = content.Key
		ObjectIdentifierList = append(ObjectIdentifierList, &ObjectIdentifier)
	}
	deleteObjectsInput := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3.Delete{
			Objects: ObjectIdentifierList,
			Quiet:   aws.Bool(false),
		},
	}
	_, err = s3Client.DeleteObjects(deleteObjectsInput)
	if err != nil {
		return err
	}
	return nil
}

func (c *StorageS3UseCase) S3StorageDownload(ctx context.Context, userId uuid.UUID, bucketName, key string) (*s3.GetObjectOutput, error) {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return nil, err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(session)
	if err != nil {
		return nil, err
	}
	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		//在bucket中的完整路径
		Key: aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *StorageS3UseCase) S3StorageDelete(ctx context.Context, userId uuid.UUID, bucketName, key string) error {
	platformS3User, err := c.repo.GetPlatformS3User(ctx, userId)
	if err != nil {
		return err
	}
	config := &aws.Config{
		Region:           aws.String(c.dispose.S3.Region),
		Endpoint:         aws.String(c.dispose.S3.Endpoint),
		Credentials:      credentials.NewStaticCredentials(platformS3User.AccessKey, platformS3User.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}
	session, err := session.NewSession(config)
	if err != nil {
		return err
	}
	s3Client := s3.New(session)
	if err != nil {
		return err
	}
	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *StorageS3UseCase) RefreshUserS3UserPermissions(ctx context.Context, userId uuid.UUID) error {
	userS3User, err := c.repo.GetUserS3User(ctx, userId)
	if err != nil {
		return err
	}
	s3Buckets, err := c.ListBucket(ctx, userId)
	if err != nil {
		return err
	}

	buckets := lo.Reduce(s3Buckets, func(agg string, item *S3Bucket, index int) string {
		if agg == "" {
			return item.BucketName
		}
		return strings.Join([]string{agg, item.BucketName}, ",")
	}, "")
	action := "Read,Write,List,Tagging"
	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.WithHost(c.dispose.S3.TargetDockerHost))
	if err != nil {
		return err
	}
	for _, user := range userS3User {
		// 构建命令
		cmd := fmt.Sprintf(`echo "s3.configure -access_key=%s -secret_key=%s -buckets=%s -user=%s -actions=%s -apply " | weed shell`,
			user.AccessKey, user.SecretKey, buckets, user.ID.String(), action)
		c.S3UserConfigure(cli, cmd)
	}
	return nil
}

func (c *StorageS3UseCase) S3UserConfigure(cli *client.Client, cmd string) error {
	// 准备执行命令
	cmdArgs := []string{"sh", "-c", cmd}

	// 创建容器执行命令的配置
	execConfig := types.ExecConfig{
		Cmd:          cmdArgs,
		AttachStdout: true,
		AttachStderr: true,
	}

	// 创建容器执行命令的请求
	execResp, err := cli.ContainerExecCreate(context.Background(), c.dispose.S3.TargetDockerContainerName, execConfig)
	if err != nil {
		log.Log(log.LevelError, "无法创建容器执行命令", err)
		return err
	}

	// 执行命令
	resp, err := cli.ContainerExecAttach(context.Background(), execResp.ID, types.ExecStartCheck{})
	if err != nil {
		log.Log(log.LevelError, "无法执行容器命令", err)
		return err
	}
	defer resp.Close()

	// 输出命令的标准输出和标准错误
	_, err = io.Copy(os.Stdout, resp.Reader)
	if err != nil {
		log.Log(log.LevelError, "无法输出命令的标准输出和标准错误", err)
		return err
	}
	// 等待命令完成
	status, err := cli.ContainerExecInspect(context.Background(), execResp.ID)
	if err != nil {
		log.Log(log.LevelError, "无法检查命令状态", err)
		return err
	}

	if status.ExitCode != 0 {
		log.Log(log.LevelError, "命令执行失败，退出码", status.ExitCode)
		return err
	}
	log.Log(log.LevelInfo, "s3.configure", "success")
	return nil
}
