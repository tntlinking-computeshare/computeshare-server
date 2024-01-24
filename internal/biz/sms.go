package biz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/data/model"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"strings"
)

// SmsUseCase 短信通知
type SmsUseCase struct {
	userRepo UserRepo
	dispose  conf.Dispose
	logger   log.Logger
}

func NewSmsUseCase(userRepo UserRepo, confDispose *conf.Dispose, logger log.Logger) *SmsUseCase {
	return &SmsUseCase{
		userRepo: userRepo,
		dispose:  *confDispose,
		logger:   logger,
	}
}

// SendVerificationCode 发送验证码
func (su *SmsUseCase) SendVerificationCode(phone string, vCode string) (err error) {
	dh3tCfg := su.dispose.Dh3T
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
		su.logger.Log(log.LevelInfo, "dh3t seed SendVerificationCode message result", string(responseBody))
		return nil
	} else {
		su.logger.Log(log.LevelError, "dh3t seed SendVerificationCode message result", string(responseBody))
		return fmt.Errorf(string(responseBody))
	}

}

// InsufficientBalance 余额不足
func (su *SmsUseCase) InsufficientBalance(resourceName string, cycle decimal.Decimal, userId uuid.UUID) error {
	fmt.Printf("您的资源[%s] 因余额不足，续费扣款失败", resourceName)
	user, err := su.userRepo.GetUser(context.Background(), userId)
	if err != nil {
		return err
	}
	dh3tCfg := su.dispose.Dh3T
	template := &model.Template{
		Id: dh3tCfg.DeductionsButInsufficientBalanceTemplateId,
	}
	dh3t := &model.Dh3t{
		Account:  dh3tCfg.Account,
		Password: dh3tCfg.Password,
		Phones:   user.TelephoneNumber,
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
		su.logger.Log(log.LevelInfo, "dh3t seed InsufficientBalance message result", string(responseBody))
		return nil
	} else {
		su.logger.Log(log.LevelError, "dh3t seed InsufficientBalance message result", string(responseBody))
		return fmt.Errorf(string(responseBody))
	}
	return nil
}

// ChargingSuccess 扣款成功
func (su *SmsUseCase) ChargingSuccess(resourceName string, cycle decimal.Decimal, userId uuid.UUID) error {
	fmt.Printf("您的资源[%s] 购买/续费 扣款成功", resourceName)
	user, err := su.userRepo.GetUser(context.Background(), userId)
	if err != nil {
		return err
	}
	dh3tCfg := su.dispose.Dh3T
	template := &model.Template{
		Id: dh3tCfg.DeductionsSuccessfulTemplateId,
	}
	dh3t := &model.Dh3t{
		Account:  dh3tCfg.Account,
		Password: dh3tCfg.Password,
		Phones:   user.TelephoneNumber,
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
		su.logger.Log(log.LevelInfo, "dh3t seed ChargingSuccess message result", string(responseBody))
		return nil
	} else {
		su.logger.Log(log.LevelError, "dh3t seed ChargingSuccess message result", string(responseBody))
		return fmt.Errorf(string(responseBody))
	}
	return nil
}

func (su *SmsUseCase) ResourceBecomeDue(resourceName string, userId uuid.UUID) error {
	fmt.Printf("您的资源[%s] 即将过期，请及时处理", resourceName)
	user, err := su.userRepo.GetUser(context.Background(), userId)
	if err != nil {
		return err
	}
	dh3tCfg := su.dispose.Dh3T
	template := &model.Template{
		Id: dh3tCfg.ThreeDaysBeforeExpirationTemplateId,
	}
	dh3t := &model.Dh3t{
		Account:  dh3tCfg.Account,
		Password: dh3tCfg.Password,
		Phones:   user.TelephoneNumber,
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
		su.logger.Log(log.LevelInfo, "dh3t seed ResourceBecomeDue message result", string(responseBody))
		return nil
	} else {
		su.logger.Log(log.LevelError, "dh3t seed ResourceBecomeDue message result", string(responseBody))
		return fmt.Errorf(string(responseBody))
	}
	return nil
}
