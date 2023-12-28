package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
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

func likeKey(telephone string) string {
	return fmt.Sprintf("telephone:%s", telephone)
}

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

func (r *s3UserRepoImpl) SetValidateCode(ctx context.Context, entity biz.User, vCode string) error {
	_, err := r.data.rdb.Set(ctx, likeKey(entity.GetFullTelephone()), vCode, time.Minute*10).Result()
	return err
}

func (r *s3UserRepoImpl) GetValidateCode(ctx context.Context, telephone string) (string, error) {
	get := r.data.rdb.Get(ctx, likeKey(telephone))
	return get.Result()
}
func (r *s3UserRepoImpl) DeleteValidateCode(ctx context.Context, user biz.User) {
	// 删除使用过的验证码
	_, _ = r.data.rdb.Del(ctx, likeKey(user.GetFullTelephone())).Result()
}
func (r *s3UserRepoImpl) GetS3User(ctx context.Context, id uuid.UUID) (*biz.S3User, error) {
	first, err := r.data.getS3UserClient(ctx).Query().Where(s3user.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBiz(first, 0), err
}
func (r *s3UserRepoImpl) GetS3UserType(ctx context.Context, id uuid.UUID, creator int8) (*biz.S3User, error) {
	first, err := r.data.getS3UserClient(ctx).Query().Where(s3user.ID(id), s3user.Type(creator)).First(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBiz(first, 0), err
}
func (r *s3UserRepoImpl) DeleteS3User(ctx context.Context, id uuid.UUID) error {
	first, err := r.data.getS3UserClient(ctx).Query().Where(s3user.ID(id)).First(ctx)
	if err == nil {
		return r.data.getS3UserClient(ctx).DeleteOne(first).Exec(ctx)
	}
	log.Log(log.LevelInfo, "DeleteS3User", id.String()+"不存在于DB")
	return err
}
func (r *s3UserRepoImpl) CreateS3User(ctx context.Context, user *biz.S3User) (*biz.S3User, error) {
	if int8(consts.PlatformCreation) == user.Type {
		entity, err := r.data.getS3UserClient(ctx).Query().Where(s3user.FkUserID(user.FkUserID), s3user.Type(int8(consts.PlatformCreation))).First(ctx)
		if err == nil && entity != nil {
			//entity.FkUserID = user.FkUserID
			//entity.AccessKey = user.AccessKey
			//entity.SecretKey = user.SecretKey
			//entity.CreateTime = user.CreateTime
			//entity.UpdateTime = user.UpdateTime
			//entity.Type = user.Type
			//err := r.data.getS3UserClient(ctx).UpdateOne(entity).Exec(ctx)
			//return r.toBiz(entity, 0), err
			return nil, errors.New("已经存在Platforms3User")
		}
	}
	entity, err := r.data.getS3UserClient(ctx).Create().
		SetFkUserID(user.FkUserID).
		SetAccessKey(user.AccessKey).
		SetSecretKey(user.SecretKey).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		SetType(user.Type).
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
	return err
}
func (r *s3UserRepoImpl) GetBucket(ctx context.Context, userId uuid.UUID, bucketName string) (*biz.S3Bucket, error) {
	s3Bucket, err := r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(userId), s3bucket.BucketName(bucketName)).First(ctx)
	if err == nil {
		return r.toBucketBiz(s3Bucket, 0), err
	}
	return nil, err
}
func (r *s3UserRepoImpl) ListBucket(ctx context.Context, userId uuid.UUID) ([]*biz.S3Bucket, error) {
	buckets, err := r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(userId)).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(buckets, r.toBucketBiz), err
}
func (r *s3UserRepoImpl) BucketPage(ctx context.Context, userId uuid.UUID, name string, page, size int32) ([]*biz.S3Bucket, int, error) {
	var buckets []*ent.S3Bucket
	var err error
	var offset int32
	var count int
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * size
	}
	if name != "" {
		count, err = r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(userId), s3bucket.BucketNameContains(name)).Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		buckets, err = r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(userId), s3bucket.BucketNameContains(name)).Order(s3bucket.ByCreatedTime(sql.OrderDesc())).
			Offset(int(offset)).Limit(int(size)).All(ctx)
	} else {
		count, err = r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(userId)).Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		buckets, err = r.data.getS3BucketClient(ctx).Query().Where(s3bucket.FkUserID(userId)).
			Order(s3bucket.ByCreatedTime(sql.OrderDesc())).Offset(int(offset)).Limit(int(size)).All(ctx)
	}
	if err != nil {
		return nil, 0, err
	}
	return lo.Map(buckets, r.toBucketBiz), count, err
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
