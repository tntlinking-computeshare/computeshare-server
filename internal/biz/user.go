package biz

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"strings"
	"time"
)

type User struct {
	ID uuid.UUID `json:"id,omitempty"`
	// CountryCallCoding holds the value of the "country_call_coding" field.
	CountryCallCoding string `json:"country_call_coding,omitempty"`
	// TelephoneNumber holds the value of the "telephone_number" field.
	TelephoneNumber string `json:"telephone_number,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// CreateDate holds the value of the "create_date" field.
	CreateDate time.Time `json:"create_date,omitempty"`
	// LastLoginDate holds the value of the "last_login_date" field.
	LastLoginDate time.Time `json:"last_login_date,omitempty"`
	ValidateCode  string    `json:"validate_code"`
	// 用户名
	Name string `json:"name,omitempty"`
	// 头像地址
	Icon string `json:"icon,omitempty"`
}

func (u *User) GetFullTelephone() string {
	return strings.Join([]string{u.CountryCallCoding, u.TelephoneNumber}, "")
}

type UserRepo interface {
	ListUser(ctx context.Context, entity User) ([]*User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id uuid.UUID, user *User) error
	UpdateUserTelephone(ctx context.Context, id uuid.UUID, user *User) error
	UpdateUserPassword(ctx context.Context, id uuid.UUID, user *User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	SendValidateCode(ctx context.Context, entity User) error
	GetValidateCode(ctx context.Context, user User) (string, error)
	DeleteValidateCode(ctx context.Context, user User)
	FindUserByFullTelephone(ctx context.Context, countryCallCoding string, telephone string) (*User, error)
}

type UserUsercase struct {
	repo   UserRepo
	key    []byte
	logger log.Logger
}

func NewUserUsecase(conf *conf.Auth, repo UserRepo, logger log.Logger) *UserUsercase {
	return &UserUsercase{
		repo:   repo,
		logger: logger,
		key:    []byte(conf.ApiKey),
	}
}

func (uc *UserUsercase) List(ctx context.Context, entity User) (ps []*User, err error) {
	ps, err = uc.repo.ListUser(ctx, entity)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsercase) Get(ctx context.Context, id uuid.UUID) (p *User, err error) {
	token, ok := global.FromContext(ctx)
	if ok {
		fmt.Println(token)
	}
	p, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsercase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsercase) Update(ctx context.Context, id uuid.UUID, user *User) error {
	return uc.repo.UpdateUser(ctx, id, user)
}

func (uc *UserUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUsercase) SendValidateCode(ctx context.Context, entity User) error {
	return uc.repo.SendValidateCode(ctx, entity)
}

func (uc *UserUsercase) GetValidateCode(ctx context.Context, user User) (string, error) {
	return uc.repo.GetValidateCode(ctx, user)
}

func (uc *UserUsercase) Login(ctx context.Context, user *User) (string, error) {
	u, err := uc.repo.FindUserByFullTelephone(ctx, user.CountryCallCoding, user.TelephoneNumber)
	if err != nil {
		return "", errors.New("telephone or password does not match")
	}

	encodedPassword := md5.Sum([]byte(user.Password))
	if hex.EncodeToString(encodedPassword[:]) != u.Password {
		return "", errors.New("telephone or password does not match")
	}

	tokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, &global.ComputeServerClaim{
		UserID: u.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(global.ExpiresTime)),
		},
	})

	return tokenHeader.SignedString(uc.key)
}

func (uc *UserUsercase) LoginWithValidateCode(ctx context.Context, user *User) (string, error) {
	code, err := uc.repo.GetValidateCode(ctx, *user)

	if err != nil || code != user.ValidateCode {
		return "", errors.New("telephone or password does not match")
	}

	u, err := uc.repo.FindUserByFullTelephone(ctx, user.CountryCallCoding, user.TelephoneNumber)
	if err != nil {
		user.Password = "Not ALLOW PASSWORD LOGIN"
		err = uc.Create(ctx, user)
		if err != nil {
			return "", err
		}
		u, err = uc.Get(ctx, user.ID)
		if err != nil {
			return "", err
		}
	}

	// 删除使用过的验证码
	uc.repo.DeleteValidateCode(ctx, *user)

	tokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, &global.ComputeServerClaim{
		UserID: u.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(global.ExpiresTime)),
		},
	})

	return tokenHeader.SignedString(uc.key)
}

func (uc *UserUsercase) UpdateUserPassword(ctx context.Context, id uuid.UUID, oldPassword, newPassword string) error {
	u, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}
	encodedPassword := md5.Sum([]byte(oldPassword))
	if hex.EncodeToString(encodedPassword[:]) != u.Password {
		return errors.New("telephone or password does not match")
	}

	return uc.repo.UpdateUserPassword(ctx, id, &User{
		Password: newPassword,
	})

}

func (uc *UserUsercase) UpdateUserTelephone(ctx context.Context, id uuid.UUID, user *User) error {
	code, err := uc.repo.GetValidateCode(ctx, *user)

	if err != nil || code != user.ValidateCode {
		return errors.New("telephone or password does not match")
	}

	if code != user.ValidateCode {
		return errors.New("validate code does no match")
	}

	return uc.repo.UpdateUserTelephone(ctx, id, user)
}
