package biz

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SmsUseCase 短信通知
type SmsUseCase struct {
}

func NewSmsUseCase() *SmsUseCase {
	return &SmsUseCase{}
}

// InsufficientBalance 余额不足
func (uc *SmsUseCase) InsufficientBalance(resourceName string, cycle decimal.Decimal, userId uuid.UUID) error {
	fmt.Printf("您的资源[%s] 因余额不足，续费扣款失败", resourceName)
	return nil
}

// ChargingSuccess 扣款成功
func (uc *SmsUseCase) ChargingSuccess(resourceName string, cycle decimal.Decimal, userId uuid.UUID) error {
	fmt.Printf("您的资源[%s] 购买/续费 扣款成功", resourceName)
	return nil
}

func (uc *SmsUseCase) ResourceBecomeDue(resourceName string, userId uuid.UUID) error {
	fmt.Printf("您的资源[%s] 即将过期，请及时处理", resourceName)
	return nil
}
