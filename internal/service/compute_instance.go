package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/samber/lo"

	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
)

type ComputeInstanceService struct {
	pb.UnimplementedComputeInstanceServer
	uc  *biz.ComputeInstanceUsercase
	log *log.Helper
}

func NewComputeInstanceService(uc *biz.ComputeInstanceUsercase, logger log.Logger) *ComputeInstanceService {
	return &ComputeInstanceService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *ComputeInstanceService) ListComputeSpec(ctx context.Context, req *pb.ListComputeSpecRequest) (*pb.ListComputeSpecReply, error) {
	list, err := s.uc.ListComputeSpec(ctx)
	return &pb.ListComputeSpecReply{
		Result: lo.Map(list, func(item *biz.ComputeSpec, _ int) *pb.ComputeSpecReply {
			return &pb.ComputeSpecReply{
				Id:     item.ID,
				Core:   item.Core,
				Memory: item.Memory,
			}
		}),
	}, err
}
func (s *ComputeInstanceService) ListComputeImage(ctx context.Context, req *pb.ListComputeImageRequest) (*pb.ListComputeImageReply, error) {
	list, err := s.uc.ListComputeImage(ctx)
	return &pb.ListComputeImageReply{
		Result: lo.Map(list, func(item *biz.ComputeImage, _ int) *pb.ComputeImageReply {
			return &pb.ComputeImageReply{
				Id:    item.ID,
				Name:  item.Name,
				Image: item.Image,
				Tag:   item.Tag,
				Port:  item.Port,
			}
		}),
	}, err
}
func (s *ComputeInstanceService) ListComputeInstanceDuration(ctx context.Context, req *pb.ListComputeDurationRequest) (*pb.ListComputeDurationReply, error) {
	return &pb.ListComputeDurationReply{
		Result: []*pb.ComputeDurationReply{
			{
				Name:     "一个月",
				Duration: 1,
			},
		},
	}, nil
}
func (s *ComputeInstanceService) Create(ctx context.Context, req *pb.CreateInstanceRequest) (*pb.CreateInstanceReply, error) {
	instance, err := s.uc.Create(ctx, &biz.ComputeInstanceCreate{
		SpecId:   req.GetSpecId(),
		ImageId:  req.GetImageId(),
		Duration: req.Duration,
		Name:     req.Name,
	})

	return &pb.CreateInstanceReply{
		Id:   instance.ID.String(),
		Name: instance.Name,
	}, err
}
func (s *ComputeInstanceService) Delete(ctx context.Context, req *pb.DeleteInstanceRequest) (*pb.DeleteInstanceReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.uc.Delete(ctx, id)
	return &pb.DeleteInstanceReply{}, err
}
func (s *ComputeInstanceService) Get(ctx context.Context, req *pb.GetInstanceRequest) (*pb.GetInstanceReply, error) {
	return &pb.GetInstanceReply{}, nil
}
func (s *ComputeInstanceService) List(ctx context.Context, req *pb.ListInstanceRequest) (*pb.ListInstanceReply, error) {
	return &pb.ListInstanceReply{}, nil
}
func (s *ComputeInstanceService) StopInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.StopInstanceReply, error) {
	return &pb.StopInstanceReply{}, nil
}
func (s *ComputeInstanceService) StartInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.StartInstanceReply, error) {
	return &pb.StartInstanceReply{}, nil
}
func (s *ComputeInstanceService) SSHInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.SSHInstanceReply, error) {
	return &pb.SSHInstanceReply{}, nil
}
