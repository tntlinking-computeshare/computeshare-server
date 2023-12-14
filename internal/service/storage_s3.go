package service

import (
	"context"
	"errors"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/samber/lo"
	"io"
	"path/filepath"
	"time"

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
	bucket, err := s.uc.CreateBucket(ctx, userId, req.GetBucket(), req.GetSecretKey())
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
	err := s.uc.DeleteBucket(ctx, userId, req.BucketName)
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
func (s *StorageS3Service) S3StorageInBucketList(ctx context.Context, req *pb.S3StorageInBucketListRequest) (*pb.S3StorageInBucketListReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	objects, err := s.uc.S3StorageInBucketList(ctx, userId, req.BucketName, req.Prefix)
	if err != nil {
		return nil, err
	}
	return &pb.S3StorageInBucketListReply{
		Code:    200,
		Message: SUCCESS,
		Data:    objects,
	}, nil
}
func (s *StorageS3Service) S3StorageUploadFile(ctx context.Context, req *pb.S3StorageUploadFileRequest) (*pb.S3StorageUploadFileReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	var key string
	if len(req.Prefix) > 0 {
		key = "/" + req.Prefix + "/" + req.FileName
	} else {
		key = req.FileName
	}
	uploadFile, err := s.uc.S3StorageUploadFile(ctx, userId, req.BucketName, key, req.Body)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &pb.S3StorageUploadFileReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.S3Object{
			Etag:       *uploadFile.ETag,
			Name:       key,
			Size:       int32(len(req.Body)),
			LastModify: time.Now().UnixMilli(),
		},
	}, nil
}
func (s *StorageS3Service) S3StorageDownload(ctx context.Context, req *pb.S3StorageDownloadRequest) (*pb.S3StorageDownloadReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	objectOutput, err := s.uc.S3StorageDownload(ctx, userId, req.BucketName, req.Key)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(objectOutput.Body)
	if err != nil {
		return nil, err
	}

	return &pb.S3StorageDownloadReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.S3StorageDownloadReply_Data{
			Body: data,
			Name: filepath.Base(req.Key),
		},
	}, nil
}
func (s *StorageS3Service) S3StorageDelete(ctx context.Context, req *pb.S3StorageDeleteRequest) (*pb.S3StorageDeleteReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	err := s.uc.S3StorageDelete(ctx, userId, req.BucketName, req.Key)
	if err != nil {
		return nil, err
	}
	return &pb.S3StorageDeleteReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
