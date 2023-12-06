package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/samber/lo"

	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
)

type StorageProviderService struct {
	pb.UnimplementedStorageProviderServer
	uc biz.StorageProviderUseCase
}

func NewStorageProviderService(uc biz.StorageProviderUseCase) *StorageProviderService {
	return &StorageProviderService{
		uc: uc,
	}
}

func (s *StorageProviderService) CreateStorageProvider(ctx context.Context, req *pb.CreateStorageProviderRequest) (*pb.CreateStorageProviderReply, error) {
	agentId, err := uuid.Parse(req.GetAgentId())
	if err != nil {
		return nil, err
	}
	sp, err := s.uc.CreateStorageProvider(ctx, agentId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateStorageProviderReply{
		Code:    200,
		Message: SUCCESS,
		Data:    s.toBiz(sp, 0),
	}, nil
}
func (s *StorageProviderService) DeleteStorageProvider(ctx context.Context, req *pb.DeleteStorageProviderRequest) (*pb.DeleteStorageProviderReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.uc.DeleteStorageProvider(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStorageProviderReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
func (s *StorageProviderService) GetStorageProvider(ctx context.Context, req *pb.GetStorageProviderRequest) (*pb.GetStorageProviderReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	sp, err := s.uc.GetStorageProvider(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStorageProviderReply{
		Code:    200,
		Message: SUCCESS,
		Data:    s.toBiz(sp, 0),
	}, nil
}
func (s *StorageProviderService) ListStorageProvider(ctx context.Context, req *pb.ListStorageProviderRequest) (*pb.ListStorageProviderReply, error) {
	providers, err := s.uc.ListStorageProvider(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListStorageProviderReply{
		Code:    200,
		Message: SUCCESS,
		Data:    lo.Map(providers, s.toBiz),
	}, nil
}

func (s *StorageProviderService) toBiz(item *biz.StorageProvider, _ int) *pb.StorageProviderInfo {
	return &pb.StorageProviderInfo{
		Id:           item.ID.String(),
		AgentId:      item.AgentID.String(),
		Status:       int32(item.Status),
		MasterServer: item.MasterServer,
		PublicIp:     item.PublicIP,
		PublicPort:   item.PublicPort,
		GrpcPort:     item.GrpcPort,
	}
}
