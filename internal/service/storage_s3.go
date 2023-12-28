package service

import (
	"context"
	"errors"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/utils"
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
func (s *StorageS3Service) CreateS3Key(ctx context.Context, req *pb.CreateS3KeyRequest) (*pb.CreateS3KeyReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	err := s.uc.CreateS3Key(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateS3KeyReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
func (s *StorageS3Service) GetUserS3UserList(ctx context.Context, req *pb.GetUserS3UserListRequest) (*pb.GetUserS3UserListReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	users, err := s.uc.GetUserS3UserList(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserS3UserListReply{
		Code:    200,
		Message: SUCCESS,
		Data: lo.Map(users, func(item *biz.S3User, _ int) *pb.S3User {
			return &pb.S3User{
				Id:         item.ID.String(),
				AccessKey:  utils.StringKeyDesensitization(item.AccessKey),
				SecretKey:  utils.StringKeyDesensitization(item.SecretKey),
				CreateTime: item.CreateTime.UnixMilli(),
				UpdateTime: item.CreateTime.UnixMilli(),
				Endpoint:   item.Endpoint,
			}
		}),
	}, nil
}
func (s *StorageS3Service) GetUserS3User(ctx context.Context, req *pb.GetUserS3UserRequest) (*pb.GetUserS3UserReply, error) {
	s3Users, err := s.uc.GetUserS3User(ctx, req.GetId(), req.GetCountryCallCoding(), req.GetTelephoneNumber(), req.GetValidateCode())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserS3UserReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.S3User{
			Id:         s3Users.ID.String(),
			AccessKey:  s3Users.AccessKey,
			SecretKey:  s3Users.SecretKey,
			CreateTime: s3Users.CreateTime.UnixMilli(),
			UpdateTime: s3Users.CreateTime.UnixMilli(),
			Endpoint:   s3Users.Endpoint,
		},
	}, nil
}
func (s *StorageS3Service) DeleteUserS3User(ctx context.Context, req *pb.DeleteUserS3UserRequest) (*pb.DeleteUserS3UserReply, error) {
	err := s.uc.DeleteUserS3User(ctx, req.GetId(), req.GetCountryCallCoding(), req.GetTelephoneNumber(), req.GetValidateCode())
	return &pb.DeleteUserS3UserReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *StorageS3Service) CreateBucket(ctx context.Context, req *pb.CreateBucketRequest) (*pb.CreateBucketReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	bucket, err := s.uc.CreateBucket(ctx, userId, req.GetBucketName())
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
func (s *StorageS3Service) EmptyBucket(ctx context.Context, req *pb.EmptyBucketRequest) (*pb.EmptyBucketReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	err := s.uc.EmptyBucket(ctx, userId, req.BucketName)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyBucketReply{
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

	buckets, count, err := s.uc.ListBucketPage(ctx, userId, req.Name, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &pb.ListBucketReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.ListBucketReply_Data{
			List: lo.Map(buckets, func(item *biz.S3Bucket, _ int) *pb.ListBucketReply_BucketVo {
				return &pb.ListBucketReply_BucketVo{
					Id:          item.ID.String(),
					Bucket:      item.BucketName,
					CreatedTime: item.CreatedTime.UnixMilli(),
				}
			}),
			Total: int32(count),
			Page:  req.Page,
			Size:  req.Size,
		},
	}, nil
}
func (s *StorageS3Service) S3StorageInBucketList(ctx context.Context, req *pb.S3StorageInBucketListRequest) (*pb.S3StorageInBucketListReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	objects, count, err := s.uc.S3StorageInBucketList(ctx, userId, req.BucketName, req.Prefix, req.Name, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &pb.S3StorageInBucketListReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.S3StorageInBucketListReply_Data{
			List:  objects,
			Total: int32(count),
			Page:  req.Page,
			Size:  req.Size,
		},
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
func (s *StorageS3Service) S3StorageMkdir(ctx context.Context, req *pb.S3StorageMkdirRequest) (*pb.S3StorageMkdirReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	err := s.uc.S3StorageMkdir(ctx, userId, req.BucketName, req.Prefix, req.DirName)
	if err != nil {
		return nil, err
	}
	return &pb.S3StorageMkdirReply{
		Code:    200,
		Message: SUCCESS,
		Data:    req.DirName,
	}, nil
}
func (s *StorageS3Service) S3StorageDeleteMkdir(ctx context.Context, req *pb.S3StorageDeleteMkdirRequest) (*pb.S3StorageDeleteMkdirReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userId := claim.GetUserId()
	err := s.uc.S3StorageDeleteMkdir(ctx, userId, req.BucketName, req.Prefix, req.DirName)
	return &pb.S3StorageDeleteMkdirReply{
		Code:    200,
		Message: SUCCESS,
	}, err
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
