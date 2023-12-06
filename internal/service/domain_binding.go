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

	domainBindingUseCase  *biz.DomainBindingUseCase
	networkMappingUseCase *biz.NetworkMappingUseCase
}

func NewDomainBindingService(domainBindingUseCase *biz.DomainBindingUseCase, networkMappingUseCase *biz.NetworkMappingUseCase) *DomainBindingService {
	return &DomainBindingService{
		domainBindingUseCase:  domainBindingUseCase,
		networkMappingUseCase: networkMappingUseCase,
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
func (s *DomainBindingService) UpdateDomainBinding(_ context.Context, _ *pb.UpdateDomainBindingRequest) (*pb.UpdateDomainBindingReply, error) {
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
func (s *DomainBindingService) GetDomainBinding(_ context.Context, _ *pb.GetDomainBindingRequest) (*pb.GetDomainBindingReply, error) {

	return &pb.GetDomainBindingReply{}, nil
}
func (s *DomainBindingService) ListDomainBinding(ctx context.Context, req *pb.ListDomainBindingRequest) (*pb.ListDomainBindingReply, error) {
	user, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}

	networkMappingId, err := uuid.Parse(req.GetNetworkMappingId())
	if err != nil {
		return nil, errors.New("networkMapping is not uuid")
	}
	page, err := s.domainBindingUseCase.List(ctx, user.GetUserId(), networkMappingId, req.Page, req.Size)
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

	networkMappingId, err := uuid.Parse(req.GetNetworkMappingId())
	if err != nil {
		return nil, errors.New("networkMappingId is not uuid")
	}

	ip, err := s.networkMappingUseCase.GetNetworkMappingIP(ctx, networkMappingId)
	if err != nil {
		return nil, err
	}

	ips, err := net.LookupIP(req.Domain)
	if err != nil {
		fmt.Printf("Error looking up IP for %s: %s\n", req.Domain, err)
		return nil, err
	}

	return &pb.NsLookupReply{
		Code:    200,
		Message: SUCCESS,
		Data: lo.ContainsBy(ips, func(item net.IP) bool {
			return item.String() == ip
		}),
	}, err
}
