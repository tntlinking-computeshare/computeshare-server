package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/samber/lo"
	"time"

	pb "github.com/mohaijiang/computeshare-server/api/system/v1"
)

const SUCCESS = "success"

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
		PwdConfig:         req.Password != "",
		CreateDate:        time.Now(),
		ValidateCode:      req.ValidateCode,
	}
	err := s.uc.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CreateUserReply_Data{
			Id: user.ID.String(),
		},
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	err := s.uc.Update(ctx, claim.GetUserId(), &biz.User{
		Name: req.GetName(),
		Icon: req.GetIcon(),
	})
	return &pb.UpdateUserReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *UserService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	err := s.uc.UpdateUserPassword(ctx, claim.GetUserId(), req.GetOldPassword(), req.GetNewPassword())
	return &pb.UpdateUserPasswordReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *UserService) UpdateUserTelephone(ctx context.Context, req *pb.UpdateUserTelephoneRequest) (*pb.UpdateUserTelephoneReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	err := s.uc.UpdateUserTelephone(ctx, claim.GetUserId(), &biz.User{
		CountryCallCoding: req.GetCountryCallCoding(),
		TelephoneNumber:   req.GetTelephoneNumber(),
		ValidateCode:      req.GetValidateCode(),
	})
	return &pb.UpdateUserTelephoneReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	id, _ := uuid.Parse(req.Id)
	err := s.uc.Delete(ctx, id)
	return &pb.DeleteUserReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("Unauthorized")
	}
	u, err := s.uc.Get(ctx, claim.GetUserId())
	return &pb.GetUserReply{
		Code:    200,
		Message: SUCCESS,
		Data:    toUserReply(u, 0),
	}, err
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	list, err := s.uc.List(ctx, biz.User{
		CountryCallCoding: req.GetCountryCallCoding(),
		TelephoneNumber:   req.GetTelephoneNumber(),
	})

	return &pb.ListUserReply{
		Code:    200,
		Message: SUCCESS,
		Data:    lo.Map(list, toUserReply),
	}, err
}

func toUserReply(user *biz.User, _ int) *pb.UserReply {

	return &pb.UserReply{
		Id:                user.ID.String(),
		CountryCallCoding: user.CountryCallCoding,
		TelephoneNumber:   user.TelephoneNumber,
		CreateDate:        user.CreateDate.UnixMilli(),
		Name:              user.Name,
		LastLoginDate:     user.LastLoginDate.UnixMilli(),
		Icon:              user.Icon,
		PwdConfig:         user.PwdConfig,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.uc.Login(ctx, &biz.User{
		CountryCallCoding: req.GetCountryCallCoding(),
		TelephoneNumber:   req.GetTelephoneNumber(),
		Password:          req.GetPassword(),
	})
	return &pb.LoginReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.LoginReply_Data{
			Token: token,
		},
	}, err
}

func (s *UserService) LoginWithClient(ctx context.Context, req *pb.LoginWithClientRequest) (*pb.LoginReply, error) {
	token, err := s.uc.LoginWithClient(ctx, req.GetUsername(), req.GetPassword())
	return &pb.LoginReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.LoginReply_Data{
			Token: token,
		},
	}, err
}

func (s *UserService) LoginWithValidateCode(ctx context.Context, req *pb.LoginWithValidateCodeRequest) (*pb.LoginReply, error) {
	token, err := s.uc.LoginWithValidateCode(ctx, &biz.User{
		CountryCallCoding: req.GetCountryCallCoding(),
		TelephoneNumber:   req.GetTelephoneNumber(),
		ValidateCode:      req.GetValidateCode(),
	})
	return &pb.LoginReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.LoginReply_Data{
			Token: token,
		},
	}, err
}
func (s *UserService) SendValidateCode(ctx context.Context, req *pb.SendValidateCodeRequest) (*pb.SendValidateCodeReply, error) {
	err := s.uc.SendValidateCode(ctx, biz.User{
		CountryCallCoding: req.CountryCallCoding,
		TelephoneNumber:   req.TelephoneNumber,
	})
	return &pb.SendValidateCodeReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
func (s *UserService) VerifyCode(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.VerifyCodeReply, error) {
	err := s.uc.VerifyCode(ctx, biz.User{
		CountryCallCoding: req.CountryCallCoding,
		TelephoneNumber:   req.TelephoneNumber,
	}, req.GetValidateCode())
	return &pb.VerifyCodeReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
