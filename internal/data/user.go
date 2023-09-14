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

func likeKey(telephone string) string {
	return fmt.Sprintf("telephone:%s", telephone)
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) ListUser(ctx context.Context, entity biz.User) ([]*biz.User, error) {
	ps, err := ur.data.db.User.Query().
		Where(user.CountryCallCodingContains(entity.CountryCallCoding), user.TelephoneNumberContains(entity.TelephoneNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(ps, ur.toBiz), err
}

func (ur *userRepo) GetUser(ctx context.Context, id uuid.UUID) (*biz.User, error) {
	p, err := ur.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return ur.toBiz(p, 0), nil
}

func (ur *userRepo) GetUserPassword(ctx context.Context, id uuid.UUID) (*biz.User, error) {
	p, err := ur.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return ur.toBiz(p, 0), nil
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) error {

	code, err := ur.GetValidateCode(ctx, *user)

	if user.Name == "" {
		user.Name = user.GetFullTelephone()
	}

	if err == nil && code == user.ValidateCode {
		encodePassword := md5.Sum([]byte(user.Password))
		result, err := ur.data.db.User.
			Create().
			SetCountryCallCoding(user.CountryCallCoding).
			SetTelephoneNumber(user.TelephoneNumber).
			SetPassword(hex.EncodeToString(encodePassword[:])).
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
	p, err := ur.data.db.User.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetLastLoginDate(user.LastLoginDate).
		SetIcon(user.Icon).
		SetName(user.Name).
		Save(ctx)
	return err
}
func (ur *userRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return ur.data.db.User.DeleteOneID(id).Exec(ctx)
}

func (ur *userRepo) SendValidateCode(ctx context.Context, entity biz.User) error {
	_, err := ur.data.rdb.Set(ctx, likeKey(entity.GetFullTelephone()), "000000", time.Minute*10).Result()
	return err
}

func (ur *userRepo) GetValidateCode(ctx context.Context, user biz.User) (string, error) {
	get := ur.data.rdb.Get(ctx, likeKey(user.GetFullTelephone()))
	return get.Result()
}

func (ur *userRepo) FindUserByFullTelephone(ctx context.Context, countryCallCoding string, telephone string) (*biz.User, error) {

	p, err := ur.data.db.User.Query().Where(user.CountryCallCodingEQ(countryCallCoding), user.TelephoneNumberEQ(telephone)).First(ctx)
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
	}
}
