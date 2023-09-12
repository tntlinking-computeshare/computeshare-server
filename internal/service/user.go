package service

import (
	"computeshare-server/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"time"

	pb "computeshare-server/api/system/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUsercase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsercase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user := &biz.User{
		CountryCallCoding: req.CountryCallCoding,
		TelephoneNumber:   req.TelephoneNumber,
		Password:          req.Password,
		CreateDate:        time.Now(),
		ValidateCode:      req.ValidateCode,
	}
	err := s.uc.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: user.ID.String(),
	}, nil
}
func (s *UserService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordReply, error) {
	id, _ := uuid.Parse(req.Id)
	lastLoginDate := time.Now()
	err := s.uc.Update(ctx, id, &biz.User{
		LastLoginDate: &lastLoginDate,
	})
	return &pb.UpdateUserPasswordReply{}, err
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	id, _ := uuid.Parse(req.Id)
	err := s.uc.Delete(ctx, id)
	return &pb.DeleteUserReply{}, err
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	id, _ := uuid.Parse(req.Id)
	u, err := s.uc.Get(ctx, id)
	return toUserReply(u, 0), err
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	list, err := s.uc.List(ctx, biz.User{
		CountryCallCoding: req.GetCountryCallCoding(),
		TelephoneNumber:   req.GetTelephoneNumber(),
	})

	return &pb.ListUserReply{
		Result: lo.Map(list, toUserReply),
	}, err
}

func toUserReply(user *biz.User, _ int) *pb.GetUserReply {
	var lastLoginDate *int64

	if user.LastLoginDate != nil {
		loginTime := user.LastLoginDate.Unix()
		lastLoginDate = &loginTime
	}

	return &pb.GetUserReply{
		Id:                user.ID.String(),
		CountryCallCoding: user.CountryCallCoding,
		TelephoneNumber:   user.TelephoneNumber,
		CreateDate:        user.CreateDate.Unix(),
		LastLoginDate:     lastLoginDate,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.uc.Login(ctx, &biz.User{
		CountryCallCoding: req.GetCountryCallCoding(),
		TelephoneNumber:   req.GetTelephoneNumber(),
		Password:          req.GetPassword(),
	})
	return &pb.LoginReply{
		Token: token,
	}, err
}
func (s *UserService) LoginWithValidateCode(ctx context.Context, req *pb.LoginWithValidateCodeRequest) (*pb.LoginReply, error) {
	token, err := s.uc.LoginWithValidateCode(ctx, &biz.User{
		CountryCallCoding: req.GetCountryCallCoding(),
		TelephoneNumber:   req.GetTelephoneNumber(),
		ValidateCode:      req.GetValidateCode(),
	})
	return &pb.LoginReply{
		Token: token,
	}, err
}
func (s *UserService) SendValidateCode(ctx context.Context, req *pb.SendValidateCodeRequest) (*pb.SendValidateCodeReply, error) {
	err := s.uc.SendValidateCode(ctx, biz.User{
		CountryCallCoding: req.CountryCallCoding,
		TelephoneNumber:   req.TelephoneNumber,
	})
	return &pb.SendValidateCodeReply{}, err
}
