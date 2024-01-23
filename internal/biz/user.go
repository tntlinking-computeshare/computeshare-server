package biz

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/data/model"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID uuid.UUID `json:"id,omitempty"`
	// 登录名
	Username string `json:"username"`
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
	//是否配置过密码
	PwdConfig bool
}

type UserResourceLimit struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户id
	FkUserID uuid.UUID `json:"fk_user_id,omitempty"`
	// MaxCPU holds the value of the "max_cpu" field.
	MaxCPU int32 `json:"max_cpu,omitempty"`
	// MaxMemory holds the value of the "max_memory" field.
	MaxMemory int32 `json:"max_memory,omitempty"`
	// MaxNetworkMapping holds the value of the "max_network_mapping" field.
	MaxNetworkMapping int32 `json:"max_network_mapping,omitempty"`
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
	SetValidateCode(ctx context.Context, entity User, vCode string) error
	GetValidateCode(ctx context.Context, telephone string) (string, error)
	DeleteValidateCode(ctx context.Context, user User)
	SetResendVerification(ctx context.Context, telephoneNumber string) error
	GetResendVerification(ctx context.Context, telephoneNumber string) (string, error)
	FindUserByFullTelephone(ctx context.Context, countryCallCoding string, telephone string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type UserResourceLimitRepo interface {
	GetByUserId(ctx context.Context, userId uuid.UUID) (*UserResourceLimit, error)
	Update(ctx context.Context, id uuid.UUID, limit *UserResourceLimit) error
}

type UserUsercase struct {
	repo    UserRepo
	key     []byte
	logger  log.Logger
	dispose conf.Dispose
}

func NewUserUsecase(conf *conf.Auth, repo UserRepo, logger log.Logger, confDispose *conf.Dispose) *UserUsercase {
	return &UserUsercase{
		repo:    repo,
		logger:  logger,
		key:     []byte(conf.ApiKey),
		dispose: *confDispose,
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

func isValidUsername(username string) bool {
	// 包含小写字母、数字、中划线
	// 大于3个英文字母并小于等于32个字符
	re := regexp.MustCompile("^[a-z0-9-]{3,32}$")
	return re.MatchString(username)
}

func (uc *UserUsercase) Update(ctx context.Context, id uuid.UUID, user *User) error {
	if isValidUsername(user.Name) {
		return uc.repo.UpdateUser(ctx, id, user)
	} else {
		return errors.New(400, "USERNAME_INVALID", "用户名不符合规范")
	}
}

func (uc *UserUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUsercase) SendValidateCode(ctx context.Context, entity User) (err error) {
	_, err = uc.repo.GetResendVerification(ctx, entity.TelephoneNumber)
	if err != redis.Nil {
		return fmt.Errorf("请勿频繁的发送验证码")
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	err = uc.repo.SetResendVerification(ctx, entity.TelephoneNumber)
	if err != nil {
		return err
	}
	err = uc.repo.SetValidateCode(ctx, entity, vCode)
	if err != nil {
		return err
	}
	return uc.SendMessageFromDh3t(entity.TelephoneNumber, vCode)
}

func (uc *UserUsercase) GetValidateCode(ctx context.Context, user User) (string, error) {
	return uc.repo.GetValidateCode(ctx, user.GetFullTelephone())
}

func (uc *UserUsercase) Login(ctx context.Context, user *User) (string, error) {
	u, err := uc.repo.FindUserByFullTelephone(ctx, user.CountryCallCoding, user.TelephoneNumber)
	if err != nil {
		return "", errors.New(400, "USER_NOT_FOUND", "用户不存在")
	}

	encodedPassword := md5.Sum([]byte(user.Password))
	if hex.EncodeToString(encodedPassword[:]) != u.Password {
		return "", errors.New(400, "PASSWORD_ERROR", "密码错误")
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
	code, err := uc.repo.GetValidateCode(ctx, user.GetFullTelephone())

	if err != nil || code != user.ValidateCode {
		return "", errors.New(400, "USER_NOT_FOUND", "用户不存在")
	}

	u, err := uc.repo.FindUserByFullTelephone(ctx, user.CountryCallCoding, user.TelephoneNumber)
	if err != nil {
		user.Password = "Not ALLOW PASSWORD LOGIN"
		user.PwdConfig = false
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

func (uc *UserUsercase) VerifyCode(ctx context.Context, user User, validateCode string) error {
	code, err := uc.repo.GetValidateCode(ctx, user.GetFullTelephone())
	if err != nil {
		return err
	}
	if code == validateCode {
		return nil
	} else {
		return errors.New(400, "Incorrect_ValidateCode", "Incorrect_ValidateCode")
	}

}

func (uc *UserUsercase) UpdateUserPassword(ctx context.Context, id uuid.UUID, oldPassword, newPassword string) error {
	u, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}
	if u.PwdConfig == true {
		encodedPassword := md5.Sum([]byte(oldPassword))
		if hex.EncodeToString(encodedPassword[:]) != u.Password {
			return errors.New(400, "PASSWORD_ERROR", "密码错误")
		}
	}

	encodedPassword := md5.Sum([]byte(newPassword))

	return uc.repo.UpdateUserPassword(ctx, id, &User{
		Password:  hex.EncodeToString(encodedPassword[:]),
		PwdConfig: true,
	})

}

func (uc *UserUsercase) UpdateUserTelephone(ctx context.Context, id uuid.UUID, user *User) error {
	code, err := uc.repo.GetValidateCode(ctx, user.GetFullTelephone())

	if err != nil || code != user.ValidateCode {
		return errors.New(400, "VALIDATE_CODE_ERROR", "验证码校验失败")
	}

	return uc.repo.UpdateUserTelephone(ctx, id, user)
}

func (uc *UserUsercase) SendMessageFromDh3t(phone string, vCode string) (err error) {
	dh3tCfg := uc.dispose.Dh3T
	var dh3tTemplateContents []model.Dh3tTemplateContent
	dh3tTemplateContents = append(dh3tTemplateContents, model.Dh3tTemplateContent{
		Name:  "1",
		Value: vCode,
	})
	template := &model.Template{
		Id:        dh3tCfg.VerificationCodeTemplateId,
		Variables: dh3tTemplateContents,
	}
	dh3t := &model.Dh3t{
		Account:  dh3tCfg.Account,
		Password: dh3tCfg.Password,
		Phones:   phone,
		Template: *template,
	}

	json, err := json.Marshal(dh3t)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(json)
	req, _ := http.NewRequest("POST", dh3tCfg.SendUrl, reader)
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	responseBody, _ := io.ReadAll(response.Body)
	if strings.Contains(string(responseBody), "提交成功") {
		uc.logger.Log(log.LevelInfo, "dh3t seed message result", string(responseBody))
		return nil
	} else {
		uc.logger.Log(log.LevelError, "dh3t seed message result", string(responseBody))
		return fmt.Errorf(string(responseBody))
	}

}

func (uc *UserUsercase) LoginWithClient(ctx context.Context, username, password string) (string, error) {
	u, err := uc.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	encodedPassword := md5.Sum([]byte(password))
	if hex.EncodeToString(encodedPassword[:]) != u.Password {
		return "", errors.New(400, "PASSWORD_ERROR", "密码错误")
	}

	tokenHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, &global.ComputeServerClaim{
		UserID: u.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(20, 0, 0)),
		},
	})

	return tokenHeader.SignedString(uc.key)
}
