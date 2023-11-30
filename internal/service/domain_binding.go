package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/samber/lo"
	"net"

	pb "github.com/mohaijiang/computeshare-server/api/network_mapping/v1"
)

type DomainBindingService struct {
	pb.UnimplementedDomainBindingServer

	domainBindingUseCase *biz.DomainBindingUseCase
}

func NewDomainBindingService(domainBindingUseCase *biz.DomainBindingUseCase) *DomainBindingService {
	return &DomainBindingService{
		domainBindingUseCase: domainBindingUseCase,
	}
}

func (s *DomainBindingService) CreateDomainBinding(ctx context.Context, req *pb.CreateDomainBindingRequest) (*pb.CreateDomainBindingReply, error) {

	user, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}

	networkMappingId, err := uuid.Parse(req.NetworkMappingId)
	if err != nil {
		return nil, err
	}
	domainBinding := &biz.DomainBinding{
		UserID:             user.GetUserId(),
		FkNetworkMappingID: networkMappingId,
		Name:               req.Name,
		Domain:             req.Domain,
	}
	err = s.domainBindingUseCase.CreateDomainBinding(ctx, domainBinding)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDomainBindingReply{
		Code:    200,
		Message: SUCCESS,
		Data:    domainBinding.ID.String(),
	}, nil
}
func (s *DomainBindingService) UpdateDomainBinding(ctx context.Context, req *pb.UpdateDomainBindingRequest) (*pb.UpdateDomainBindingReply, error) {
	return &pb.UpdateDomainBindingReply{}, nil
}
func (s *DomainBindingService) DeleteDomainBinding(ctx context.Context, req *pb.DeleteDomainBindingRequest) (*pb.DeleteDomainBindingReply, error) {
	user, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.domainBindingUseCase.DeleteDomainBinding(ctx, id, user.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDomainBindingReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
func (s *DomainBindingService) GetDomainBinding(ctx context.Context, req *pb.GetDomainBindingRequest) (*pb.GetDomainBindingReply, error) {

	return &pb.GetDomainBindingReply{}, nil
}
func (s *DomainBindingService) ListDomainBinding(ctx context.Context, req *pb.ListDomainBindingRequest) (*pb.ListDomainBindingReply, error) {
	user, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	page, err := s.domainBindingUseCase.List(ctx, user.GetUserId(), req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &pb.ListDomainBindingReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.DomainBindingPageResponse{
			Page:  page.PageNum,
			Size:  page.PageSize,
			Total: page.Total,
			List:  page.Data,
		},
	}, nil
}

func (s *DomainBindingService) NsLookup(ctx context.Context, req *pb.NsLookupRequest) (*pb.NsLookupReply, error) {
	ips, err := net.LookupIP(req.Domain)
	if err != nil {
		fmt.Printf("Error looking up IP for %s: %s\n", req.Domain, err)
		return nil, err
	}

	return &pb.NsLookupReply{
		Code:    200,
		Message: SUCCESS,
		Data: lo.Map(ips, func(item net.IP, index int) string {
			return item.String()
		}),
	}, err
}
