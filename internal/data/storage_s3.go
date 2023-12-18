package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3bucket"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/s3user"
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
func (r *s3UserRepoImpl) GetS3User(ctx context.Context, userId uuid.UUID) (*biz.S3User, error) {
	entity, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(userId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBiz(entity, 0), err
}
func (r *s3UserRepoImpl) CreateBucket(ctx context.Context, user *biz.S3User, bucket string) (*biz.S3Bucket, error) {
	s3User, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(user.FkUserID)).First(ctx)
	if err != nil {
		return nil, err
	}

	s3Bucket, err := r.data.getS3BucketClient(ctx).Create().SetS3User(s3User).SetBucket(bucket).SetCreatedTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBucketBiz(s3Bucket, 0), err
}
func (r *s3UserRepoImpl) DeleteBucket(ctx context.Context, user *biz.S3User, bucketName string) error {
	s3User, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(user.FkUserID)).First(ctx)

	if err != nil {
		return err
	}
	first, err := s3User.QueryBuckets().Where(s3bucket.BucketEQ(bucketName)).First(ctx)
	if err != nil {
		return err
	}
	return r.data.getS3BucketClient(ctx).DeleteOne(first).Exec(ctx)

}
func (r *s3UserRepoImpl) ListBucket(ctx context.Context, userId uuid.UUID) ([]*biz.S3Bucket, error) {
	s3User, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(userId)).First(ctx)

	if err != nil {
		return nil, err
	}
	list, err := s3User.QueryBuckets().All(ctx)

	return lo.Map(list, r.toBucketBiz), err
}

func (r *s3UserRepoImpl) toBiz(item *ent.S3User, _ int) *biz.S3User {
	return &biz.S3User{
		ID:        item.ID,
		FkUserID:  item.FkUserID,
		AccessKey: item.AccessKey,
		SecretKey: item.SecretKey,
		Endpoint:  "computeshare.newtouch.com:8333",
	}
}

func (r *s3UserRepoImpl) toBucketBiz(item *ent.S3Bucket, _ int) *biz.S3Bucket {
	return &biz.S3Bucket{
		ID:          item.ID,
		Bucket:      item.Bucket,
		CreatedTime: item.CreatedTime,
	}
}
