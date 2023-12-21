package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3bucket"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3user"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"
	"time"
)

func NewS3UserRepo(data *Data, logger log.Logger) biz.S3UserRepo {

	return &s3UserRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type s3UserRepoImpl struct {
	data *Data
	log  *log.Helper
}

func (r *s3UserRepoImpl) CreateS3User(ctx context.Context, user *biz.S3User) (*biz.S3User, error) {
	entity, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(user.FkUserID)).First(ctx)
	if err == nil && entity != nil {
		entity.FkUserID = user.FkUserID
		entity.AccessKey = user.AccessKey
		entity.SecretKey = user.SecretKey
		entity.CreateTime = user.CreateTime
		entity.UpdateTime = user.UpdateTime
		err := r.data.getS3UserClient(ctx).UpdateOne(entity).Exec(ctx)
		return r.toBiz(entity, 0), err
	}

	entity, err = r.data.getS3UserClient(ctx).Create().
		SetFkUserID(user.FkUserID).
		SetAccessKey(user.AccessKey).
		SetSecretKey(user.SecretKey).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBiz(entity, 0), err
}
func (r *s3UserRepoImpl) GetUserS3User(ctx context.Context, userId uuid.UUID) ([]*biz.S3User, error) {
	s3Users, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(userId), s3user.Type(int8(consts.UserCreation))).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(s3Users, r.toBiz), err
}
func (r *s3UserRepoImpl) GetPlatformS3User(ctx context.Context, userId uuid.UUID) (*biz.S3User, error) {
	entity, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(userId), s3user.Type(int8(consts.PlatformCreation))).First(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBiz(entity, 0), err
}
func (r *s3UserRepoImpl) CreateBucket(ctx context.Context, s3User *biz.S3User, bucketName string) (*biz.S3Bucket, error) {
	s3Bucket, err := r.data.getS3BucketClient(ctx).Create().SetFkUserID(s3User.FkUserID).SetBucketName(bucketName).SetCreatedTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBucketBiz(s3Bucket, 0), err
}
func (r *s3UserRepoImpl) DeleteBucket(ctx context.Context, user *biz.S3User, bucketName string) error {
	first, err := r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(user.FkUserID), s3bucket.BucketName(bucketName)).First(ctx)
	if err == nil {
		return r.data.getS3BucketClient(ctx).DeleteOne(first).Exec(ctx)
	}
	log.Log(log.LevelInfo, "DeleteBucketInDB", bucketName+"不存在于DB")
	return nil
}
func (r *s3UserRepoImpl) ListBucket(ctx context.Context, userId uuid.UUID) ([]*biz.S3Bucket, error) {
	buckets, err := r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(userId)).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(buckets, r.toBucketBiz), err
}

func (r *s3UserRepoImpl) toBiz(item *ent.S3User, _ int) *biz.S3User {
	return &biz.S3User{
		ID:         item.ID,
		FkUserID:   item.FkUserID,
		AccessKey:  item.AccessKey,
		SecretKey:  item.SecretKey,
		CreateTime: item.CreateTime,
		UpdateTime: item.UpdateTime,
		Endpoint:   "computeshare.newtouch.com:8333",
	}
}

func (r *s3UserRepoImpl) toBucketBiz(item *ent.S3Bucket, _ int) *biz.S3Bucket {
	return &biz.S3Bucket{
		ID:          item.ID,
		BucketName:  item.BucketName,
		CreatedTime: item.CreatedTime,
	}
}
