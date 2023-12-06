package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/samber/lo"

	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
)

type StorageS3Service struct {
	pb.UnimplementedStorageS3Server
	uc *biz.StorageS3UseCase
}

func NewStorageS3Service(uc *biz.StorageS3UseCase) *StorageS3Service {
	return &StorageS3Service{
		uc: uc,
	}
}

func (s *StorageS3Service) GetS3User(ctx context.Context, req *pb.GetS3UserRequest) (*pb.GetS3UserReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	user, err := s.uc.GetS3User(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &pb.GetS3UserReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.S3User{
			AccessKey: user.AccessKey,
			SecretKey: user.SecretKey,
			Endpoint:  user.Endpoint,
		},
	}, nil
}
func (s *StorageS3Service) CreateBucket(ctx context.Context, req *pb.CreateBucketRequest) (*pb.CreateBucketReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	bucket, err := s.uc.CreateBucket(ctx, userId, req.GetBucket(), req.GetAccessKey())
	if err != nil {
		return nil, err
	}
	return &pb.CreateBucketReply{
		Code:    200,
		Message: SUCCESS,
		Data:    bucket.ID.String(),
	}, nil
}
func (s *StorageS3Service) DeleteBucket(ctx context.Context, req *pb.DeleteBucketRequest) (*pb.DeleteBucketReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	bucketId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.uc.DeleteBucket(ctx, userId, bucketId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBucketReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
func (s *StorageS3Service) ListBucket(ctx context.Context, req *pb.ListBucketRequest) (*pb.ListBucketReply, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()

	buckets, err := s.uc.ListBucket(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &pb.ListBucketReply{
		Code:    200,
		Message: SUCCESS,
		Data: lo.Map(buckets, func(item *biz.S3Bucket, _ int) *pb.ListBucketReply_BucketVo {
			return &pb.ListBucketReply_BucketVo{
				Id:          item.ID.String(),
				Bucket:      item.Bucket,
				CreatedTime: item.CreatedTime.UnixMilli(),
			}
		}),
	}, nil
}
