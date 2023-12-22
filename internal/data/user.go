package data

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/user"
	"github.com/samber/lo"
	"time"
)

func likeResendVerificationKey(telephone string) string {
	return fmt.Sprintf("telephone_send:%s", telephone)
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	lg := log.NewHelper(logger)
	return &userRepo{
		data: data,
		log:  lg,
	}
}

func (ur *userRepo) ListUser(ctx context.Context, entity biz.User) ([]*biz.User, error) {
	ps, err := ur.data.getUserClient(ctx).Query().
		Where(user.CountryCallCodingContains(entity.CountryCallCoding), user.TelephoneNumberContains(entity.TelephoneNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(ps, ur.toBiz), err
}

func (ur *userRepo) GetUser(ctx context.Context, id uuid.UUID) (*biz.User, error) {
	p, err := ur.data.getUserClient(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return ur.toBiz(p, 0), nil
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) error {

	code, err := ur.GetValidateCode(ctx, user.GetFullTelephone())

	if user.Name == "" {
		user.Name = user.GetFullTelephone()
	}

	if err == nil && code == user.ValidateCode {
		encodePassword := md5.Sum([]byte(user.Password))
		result, err := ur.data.getUserClient(ctx).
			Create().
			SetCountryCallCoding(user.CountryCallCoding).
			SetTelephoneNumber(user.TelephoneNumber).
			SetPassword(hex.EncodeToString(encodePassword[:])).
			SetPwdConfig(user.PwdConfig).
			SetLastLoginDate(time.Now()).
			SetName(user.Name).
			SetIcon(user.Icon).
			Save(ctx)

		if err != nil {
			return err
		}

		user.ID = result.ID

		// 删除使用过的验证码
		_, _ = ur.data.rdb.Del(ctx, likeKey(user.GetFullTelephone())).Result()

		return err
	} else {
		return errors.New("validate code is not match")
	}

}
func (ur *userRepo) UpdateUser(ctx context.Context, id uuid.UUID, user *biz.User) error {
	p, err := ur.data.getUserClient(ctx).Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetIcon(user.Icon).
		SetName(user.Name).
		Save(ctx)
	return err
}

func (ur *userRepo) UpdateUserTelephone(ctx context.Context, id uuid.UUID, updateUser *biz.User) error {
	first, err := ur.data.getUserClient(ctx).Query().Where(user.TelephoneNumber(updateUser.TelephoneNumber), user.CountryCallCoding(updateUser.CountryCallCoding)).First(ctx)
	if err != nil {
		return err
	}
	if first != nil {
		return errors.New("该手机号已经被注册")
	}
	p, err := ur.data.getUserClient(ctx).Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetCountryCallCoding(updateUser.CountryCallCoding).
		SetTelephoneNumber(updateUser.TelephoneNumber).
		Save(ctx)
	return err
}

func (ur *userRepo) UpdateUserPassword(ctx context.Context, id uuid.UUID, user *biz.User) error {
	p, err := ur.data.getUserClient(ctx).Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetPassword(user.Password).
		SetPwdConfig(user.PwdConfig).
		Save(ctx)
	return err
}

func (ur *userRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return ur.data.getUserClient(ctx).DeleteOneID(id).Exec(ctx)
}

func (ur *userRepo) SetValidateCode(ctx context.Context, entity biz.User, vCode string) error {
	_, err := ur.data.rdb.Set(ctx, likeKey(entity.GetFullTelephone()), vCode, time.Minute*10).Result()
	return err
}

func (ur *userRepo) GetValidateCode(ctx context.Context, telephone string) (string, error) {
	get := ur.data.rdb.Get(ctx, likeKey(telephone))
	return get.Result()
}

func (ur *userRepo) SetResendVerification(ctx context.Context, telephoneNumber string) error {
	_, err := ur.data.rdb.Set(ctx, likeResendVerificationKey(telephoneNumber), "Cooling", time.Minute).Result()
	return err
}

func (ur *userRepo) GetResendVerification(ctx context.Context, telephoneNumber string) (string, error) {
	get := ur.data.rdb.Get(ctx, likeResendVerificationKey(telephoneNumber))
	return get.Result()
}

func (ur *userRepo) FindUserByFullTelephone(ctx context.Context, countryCallCoding string, telephone string) (*biz.User, error) {

	p, err := ur.data.getUserClient(ctx).Query().Where(user.CountryCallCodingEQ(countryCallCoding), user.TelephoneNumberEQ(telephone)).First(ctx)
	if err != nil {
		return nil, err
	}
	return ur.toBiz(p, 0), err
}

func (ur *userRepo) DeleteValidateCode(ctx context.Context, user biz.User) {
	// 删除使用过的验证码
	_, _ = ur.data.rdb.Del(ctx, likeKey(user.GetFullTelephone())).Result()
}

func (ur *userRepo) toBiz(p *ent.User, _ int) *biz.User {
	return &biz.User{
		ID:                p.ID,
		CountryCallCoding: p.CountryCallCoding,
		TelephoneNumber:   p.TelephoneNumber,
		Password:          p.Password,
		CreateDate:        p.CreateDate,
		LastLoginDate:     p.LastLoginDate,
		Name:              p.Name,
		Icon:              p.Icon,
		PwdConfig:         p.PwdConfig,
	}
}
