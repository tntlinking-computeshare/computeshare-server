package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/samber/lo"
)

type ComputeInstanceService struct {
	pb.UnimplementedComputeInstanceServer
	uc      *biz.ComputeInstanceUsercase
	dispose *conf.Dispose
	log     *log.Helper
}

func NewComputeInstanceService(uc *biz.ComputeInstanceUsercase, dispose *conf.Dispose, logger log.Logger) *ComputeInstanceService {
	return &ComputeInstanceService{
		uc:      uc,
		log:     log.NewHelper(logger),
		dispose: dispose,
	}
}

func (s *ComputeInstanceService) ListComputeSpec(ctx context.Context, req *pb.ListComputeSpecRequest) (*pb.ListComputeSpecReply, error) {
	list, err := s.uc.ListComputeSpec(ctx)
	return &pb.ListComputeSpecReply{
		Code:    200,
		Message: SUCCESS,
		Data: lo.Map(list, func(item *biz.ComputeSpec, _ int) *pb.ComputeSpec {
			return &pb.ComputeSpec{
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
		Code:    200,
		Message: SUCCESS,
		Data: lo.Map(list, func(item *biz.ComputeImage, _ int) *pb.ComputeImage {
			return &pb.ComputeImage{
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
		Code:    200,
		Message: SUCCESS,
		Data: []*pb.ComputeDuration{
			{
				Name:     "一个月",
				Duration: 1,
			},
		},
	}, nil
}
func (s *ComputeInstanceService) Create(ctx context.Context, req *pb.CreateInstanceRequest) (*pb.CreateInstanceReply, error) {
	instance, err := s.uc.Create(ctx, &biz.ComputeInstanceCreate{
		SpecId:        req.GetSpecId(),
		ImageId:       req.GetImageId(),
		Duration:      req.Duration,
		Name:          req.Name,
		PublicKey:     req.PublicKey,
		Password:      req.Password,
		DockerCompose: req.DockerCompose,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateInstanceReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CreateInstanceReply_Data{
			Id:   instance.ID.String(),
			Name: instance.Name,
		},
	}, err
}
func (s *ComputeInstanceService) Delete(ctx context.Context, req *pb.DeleteInstanceRequest) (*pb.CommonReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.uc.Delete(ctx, id)
	return &pb.CommonReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *ComputeInstanceService) Get(ctx context.Context, req *pb.GetInstanceRequest) (*pb.GetInstanceReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	instance, err := s.uc.Get(ctx, id)
	return &pb.GetInstanceReply{
		Code:    200,
		Message: SUCCESS,
		Data:    s.toReply(instance, 0),
	}, err
}
func (s *ComputeInstanceService) List(ctx context.Context, req *pb.ListInstanceRequest) (*pb.ListInstanceReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("cannot get user id")
	}
	list, err := s.uc.ListComputeInstance(ctx, claim.UserID)
	return &pb.ListInstanceReply{
		Code:    200,
		Message: SUCCESS,
		Data:    lo.Map(list, s.toReply),
	}, err
}
func (s *ComputeInstanceService) StopInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.CommonReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.uc.Stop(ctx, id)
	return &pb.CommonReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *ComputeInstanceService) StartInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.CommonReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.uc.Start(ctx, id)
	return &pb.CommonReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}

func (s *ComputeInstanceService) toReply(p *biz.ComputeInstance, _ int) *pb.Instance {
	if p == nil {
		return nil
	}
	return &pb.Instance{
		Id:             p.ID.String(),
		Name:           p.Name,
		Status:         int32(p.Status),
		ExpirationTime: p.ExpirationTime.UnixMilli(),
		ImageName:      p.Image,
		Core:           p.Core,
		Memory:         p.Memory,
		Stats: lo.Map(p.Stats, func(item *biz.ComputeInstanceRds, _ int) *pb.InstanceStats {
			if item == nil {
				return nil
			}
			return &pb.InstanceStats{
				Id:          item.ID,
				CpuUsage:    item.CpuUsage,
				MemoryUsage: item.MemoryUsage,
				StatsTime:   item.StatsTime.UnixMilli(),
			}
		}),
	}
}

func (s *ComputeInstanceService) GetInstanceConsole(ctx context.Context, id string, userId string) (string, error) {
	instanceId, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}
	return s.uc.GetVncConsole(ctx, instanceId, userId)
}

func (s *ComputeInstanceService) RestartInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.CommonReply, error) {
	instanceId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.uc.Reboot(ctx, instanceId)
	if err != nil {
		return nil, err
	}

	return &pb.CommonReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}

func (s *ComputeInstanceService) ReCreateInstance(ctx context.Context, req *pb.RecreateInstanceRequest) (*pb.CommonReply, error) {
	instanceId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.uc.Recreate(ctx, instanceId, &biz.ComputeInstanceCreate{
		ImageId:       req.ImageId,
		PublicKey:     req.PublicKey,
		Password:      req.Password,
		DockerCompose: req.DockerCompose,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CommonReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}

func (s *ComputeInstanceService) GetInstanceVncURL(ctx context.Context, req *pb.GetInstanceRequest) (*pb.GetInstanceVncURLReply, error) {
	return &pb.GetInstanceVncURLReply{
		Code:    200,
		Message: SUCCESS,
		Data:    fmt.Sprintf("%s/vnc_lite.html?host=%s&instanceId=%s", s.dispose.Domain.VncHost, s.dispose.Domain.ApiHost, req.GetId()),
	}, nil
}
