package biz

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
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
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// accessKey
	AccessKey string `json:"access_key,omitempty"`
	// secretKey
	SecretKey string `json:"secret_key,omitempty"`
	// Endpoint
	Endpoint string `json:"endpoint"`
}

type S3Bucket struct {
	ID          uuid.UUID
	Bucket      string
	CreatedTime time.Time
}

type S3UserRepo interface {
	CreateS3User(ctx context.Context, user *S3User) (*S3User, error)
	GetS3User(ctx context.Context, userId uuid.UUID) (*S3User, error)
	CreateBucket(ctx context.Context, user *S3User, bucket string) (*S3Bucket, error)
	DeleteBucket(ctx context.Context, user *S3User, id uuid.UUID) error
	ListBucket(ctx context.Context, userId uuid.UUID) ([]*S3Bucket, error)
}

type StorageS3UseCase struct {
	repo     S3UserRepo
	userRepo UserRepo
	log      *log.Helper
}

func NewStorageS3UseCase(repo S3UserRepo, userRepo UserRepo, logger log.Logger) *StorageS3UseCase {
	return &StorageS3UseCase{
		repo:     repo,
		userRepo: userRepo,
		log:      log.NewHelper(logger),
	}
}

func (c *StorageS3UseCase) createS3User(ctx context.Context, userId uuid.UUID, secretKey string) (*S3User, error) {
	user, err := c.userRepo.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	s3User := &S3User{
		FkUserID:  userId,
		AccessKey: user.TelephoneNumber,
		SecretKey: secretKey,
	}

	return c.repo.CreateS3User(ctx, s3User)
}

func (c *StorageS3UseCase) GetS3User(ctx context.Context, userId uuid.UUID) (*S3User, error) {

	return c.repo.GetS3User(ctx, userId)
}

func (c *StorageS3UseCase) CreateBucket(ctx context.Context, userId uuid.UUID, bucket, accessKey string) (*S3Bucket, error) {
	s3User, err := c.createS3User(ctx, userId, accessKey)
	if err != nil {
		return nil, err
	}

	s3Bucket, err := c.repo.CreateBucket(ctx, s3User, bucket)
	if err != nil {
		return nil, err
	}
	s3Buckets, err := c.ListBucket(ctx, userId)
	if err != nil {
		return nil, err
	}

	buckets := lo.Reduce(s3Buckets, func(agg string, item *S3Bucket, index int) string {
		if agg == "" {
			return item.Bucket
		}
		return strings.Join([]string{agg, item.Bucket}, ",")
	}, "")

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"))

	// 构建命令
	cmd := fmt.Sprintf(`echo "s3.configure -access_key=%s -secret_key=%s -buckets=%s -user=%s -actions=Read,Write,List,Tagging,Admin -apply " | weed shell`,
		s3User.AccessKey, s3User.SecretKey, buckets, s3User.FkUserID)

	// 准备执行命令
	cmdArgs := []string{"sh", "-c", cmd}

	// 创建容器执行命令的配置
	execConfig := types.ExecConfig{
		Cmd:          cmdArgs,
		AttachStdout: true,
		AttachStderr: true,
	}

	// 创建容器执行命令的请求
	execResp, err := cli.ContainerExecCreate(context.Background(), "seaweedfs-master-1", execConfig)
	if err != nil {
		fmt.Printf("无法创建容器执行命令: %v\n", err)
		return nil, err
	}

	// 执行命令
	resp, err := cli.ContainerExecAttach(context.Background(), execResp.ID, types.ExecStartCheck{})
	if err != nil {
		fmt.Printf("无法执行容器命令: %v\n", err)
		return nil, err
	}
	defer resp.Close()

	// 输出命令的标准输出和标准错误
	_, err = io.Copy(os.Stdout, resp.Reader)
	if err != nil {
		fmt.Printf("无法输出命令的标准输出和标准错误: %v\n", err)
		return nil, err
	}
	// 等待命令完成
	status, err := cli.ContainerExecInspect(context.Background(), execResp.ID)
	if err != nil {
		fmt.Printf("无法检查命令状态: %v\n", err)
		return nil, err
	}

	if status.ExitCode != 0 {
		fmt.Printf("命令执行失败，退出码: %d\n", status.ExitCode)
		return nil, err
	}

	fmt.Println("命令执行成功")

	return s3Bucket, err
}

func (c *StorageS3UseCase) DeleteBucket(ctx context.Context, userId uuid.UUID, bucketId uuid.UUID) error {
	s3User, err := c.GetS3User(ctx, userId)
	if err != nil {
		return err
	}
	return c.repo.DeleteBucket(ctx, s3User, bucketId)
}

func (c *StorageS3UseCase) ListBucket(ctx context.Context, userId uuid.UUID) ([]*S3Bucket, error) {
	return c.repo.ListBucket(ctx, userId)
}