package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/shopspring/decimal"
	"time"
)

type CycleRenewalRepo interface {
	GetById(ctx context.Context, id uuid.UUID) (*CycleRenewal, error)
	Create(ctx context.Context, renewal *CycleRenewal) (*CycleRenewal, error)
	Update(ctx context.Context, id uuid.UUID, renewal *CycleRenewal) error
	PageByUserId(ctx context.Context, id uuid.UUID, page, size int) (*global2.Page[*CycleRenewal], error)
	QueryDailyRenew(ctx context.Context) ([]*CycleRenewal, error)
}

type CycleRenewalUseCase struct {
	log                  *log.Helper
	repo                 CycleRenewalRepo
	cycleRepo            CycleRepo
	cycleOrderRepo       CycleOrderRepo
	cycleTransactionRepo CycleTransactionRepo
	computeInstanceRepo  ComputeInstanceRepo
}

func NewCycleRenewalUseCase(
	logger log.Logger,
	repo CycleRenewalRepo,
	cycleRepo CycleRepo,
	cycleOrderRepo CycleOrderRepo,
	cycleTransactionRepo CycleTransactionRepo,
	computeInstanceRepo ComputeInstanceRepo,
) *CycleRenewalUseCase {
	return &CycleRenewalUseCase{
		log:                  log.NewHelper(logger),
		repo:                 repo,
		cycleRepo:            cycleRepo,
		cycleOrderRepo:       cycleOrderRepo,
		cycleTransactionRepo: cycleTransactionRepo,
		computeInstanceRepo:  computeInstanceRepo,
	}
}

func (c *CycleRenewalUseCase) PageByUser(ctx context.Context, page, size int32) (*global2.Page[*CycleRenewal], error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()
	return c.repo.PageByUserId(ctx, userId, int(page), int(size))
}

func (c *CycleRenewalUseCase) OpenRenewal(ctx context.Context, renewalId uuid.UUID) error {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()

	renewal, err := c.repo.GetById(ctx, renewalId)

	if err != nil {
		return errors.New(400, "not_found", "renewal not exists")
	}

	if renewal.FkUserID != userId {
		return errors.New(400, "unauthorized", "unauthorized")
	}

	instance, err := c.computeInstanceRepo.Get(ctx, renewal.ResourceID)
	if err != nil {
		return err
	}
	renewal.AutoRenewal = true

	renewalTime := instance.ExpirationTime.AddDate(0, 0, -9)
	renewalTime = time.Date(renewalTime.Year(), renewalTime.Month(), renewalTime.Day(), 23, 0, 0, 0, renewalTime.Location())

	// 获取当前时间
	currentTime := time.Now()

	// 将当前时间设置为当天的 23 点整
	zeroTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 0, 0, 0, currentTime.Location())

	if zeroTime.After(renewalTime) {
		renewal.RenewalTime = &zeroTime
	} else {
		renewal.RenewalTime = &renewalTime
	}

	return c.repo.Update(ctx, renewalId, renewal)

}

func (c *CycleRenewalUseCase) CloseRenewal(ctx context.Context, renewalId uuid.UUID) error {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()

	renewal, err := c.repo.GetById(ctx, renewalId)

	if err != nil {
		return errors.New(400, "not_found", "renewal not exists")
	}

	if renewal.FkUserID != userId {
		return errors.New(400, "unauthorized", "unauthorized")
	}

	renewal.AutoRenewal = false
	renewal.RenewalTime = nil
	return c.repo.Update(ctx, renewalId, renewal)
}

func (c *CycleRenewalUseCase) Get(ctx context.Context, renewalId uuid.UUID) (*CycleRenewal, error) {
	return c.repo.GetById(ctx, renewalId)
}

func (c *CycleRenewalUseCase) Detail(ctx context.Context, renewalId uuid.UUID) (*CycleRenewalDetail, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()

	renewal, err := c.repo.GetById(ctx, renewalId)
	if err != nil {
		return nil, err
	}

	cycle, err := c.cycleRepo.FindByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}
	balance, _ := cycle.Cycle.Float64()

	instance, err := c.computeInstanceRepo.Get(ctx, renewal.ResourceID)
	if err != nil {
		return nil, err
	}

	detail := &CycleRenewalDetail{
		ID:           renewal.ID,
		FkUserID:     renewal.FkUserID,
		ResourceID:   renewal.ResourceID,
		ResourceType: renewal.ResourceType,
		ProductName:  renewal.ProductName,
		ProductDesc:  renewal.ProductDesc,
		State:        renewal.State,
		ExtendDay:    renewal.ExtendDay,
		ExtendPrice:  renewal.ExtendPrice,
		DueTime:      renewal.DueTime,
		RenewalTime:  renewal.RenewalTime,
		AutoRenewal:  renewal.AutoRenewal,
		InstanceId:   instance.ID,
		InstanceName: instance.Name,
		InstanceSpec: fmt.Sprintf("%s核 %sGB", instance.Core, instance.Memory),
		Image:        instance.Image,
		Balance:      float32(balance),
	}
	return detail, nil
}

func (c *CycleRenewalUseCase) ManualRenew(ctx context.Context, renewalId uuid.UUID) error {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()

	renewal, err := c.repo.GetById(ctx, renewalId)

	if err != nil {
		return errors.New(400, "not_found", "renewal not exists")
	}

	cycle, err := c.cycleRepo.FindByUserID(ctx, userId)
	if err != nil {
		return errors.New(400, "not_found", "cycle not exists")
	}

	// 判断余额
	extendPrice := decimal.NewFromFloat(renewal.ExtendPrice)

	if cycle.Cycle.LessThan(extendPrice) {
		return errors.New(400, "insufficient balance", "Insufficient balance")
	}

	var orderNo string
	// 判断订单号不重复
	var exists bool
	for {
		orderNo = NewOrderNo()
		exists = c.cycleOrderRepo.CheckOrderNoExists(ctx, orderNo)
		if !exists {
			break
		}
	}

	// 创建订单
	cycleOrder := &CycleOrder{
		OrderNo:     orderNo,
		FkUserID:    userId,
		ProductName: renewal.ProductName,
		ProductDesc: renewal.ProductDesc,
		Symbol:      "-",
		Cycle:       renewal.ExtendPrice,
		CreateTime:  time.Now(),
	}

	cycleOrder, err = c.cycleOrderRepo.Create(ctx, cycleOrder)
	if err != nil {
		return errors.New(400, "internal error", err.Error())
	}

	balance, _ := cycle.Cycle.Sub(extendPrice).Float64()
	// 创建交易流水
	cycleTransaction := &CycleTransaction{
		FkCycleID:         cycle.ID,
		FkUserID:          userId,
		FkCycleOrderID:    cycleOrder.ID,
		FkCycleRechargeID: uuid.Nil,
		Operation:         "租用云服务器",
		Symbol:            cycleOrder.Symbol,
		Cycle:             cycleOrder.Cycle,
		Balance:           balance,
		OperationTime:     time.Now(),
	}

	cycleTransaction, err = c.cycleTransactionRepo.Create(ctx, cycleTransaction)
	if err != nil {
		return err
	}

	// 记录余额变动
	cycle.Cycle = cycle.Cycle.Sub(extendPrice)
	cycle.UpdateTime = time.Now()
	err = c.cycleRepo.Update(ctx, cycle)

	// 续费成功，延长虚拟机到期时间
	instance, err := c.computeInstanceRepo.Get(ctx, renewal.ResourceID)
	if err != nil {
		return err
	}

	instance.ExpirationTime = instance.ExpirationTime.AddDate(0, 0, int(renewal.ExtendDay))
	err = c.computeInstanceRepo.Update(ctx, instance.ID, instance)
	if err != nil {
		return err
	}

	// 续费成功，推长续费时间和到期时间
	renewal.DueTime = &instance.ExpirationTime
	renewalTime := instance.ExpirationTime.AddDate(0, 0, -9)
	renewalTime = time.Date(renewalTime.Year(), renewalTime.Month(), renewalTime.Day(), 23, 0, 0, 0, renewalTime.Location())

	// 获取当前时间
	currentTime := time.Now()

	// 将当前时间设置为当天的 0 点整
	zeroTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 0, 0, 0, currentTime.Location())

	if zeroTime.After(renewalTime) {
		renewal.RenewalTime = &zeroTime
	} else {
		renewal.RenewalTime = &renewalTime
	}

	err = c.repo.Update(ctx, renewalId, renewal)
	return err
}

func (c *CycleRenewalUseCase) DailyCheck(db *ent.Client) {
	ctx := context.Background()
	//查询当天应当续费续费
	list, err := c.repo.QueryDailyRenew(ctx)
	if err != nil {
		c.log.Error("查询当天应当续费续费失败： ", err.Error())
		return
	}

	for _, renewal := range list {
		// 开启事物
		tx, _ := db.Tx(ctx)
		tctx := context.WithValue(ctx, "tx", tx)
		err := c.ManualRenew(tctx, renewal.ID)
		if err != nil {
			// 如何回滚，续费失败，延长续费时间
			_ = tx.Rollback()
			rollbackCtx := context.Background()
			tx, _ = db.Tx(rollbackCtx)
			rctx := context.WithValue(ctx, "tx", tx)
			newRenewDate := renewal.RenewalTime.AddDate(0, 0, 1)
			renewal.RenewalTime = &newRenewDate
			err = c.repo.Update(rctx, renewal.ID, renewal)
			if err != nil {
				_ = tx.Rollback()
			} else {
				_ = tx.Commit()
			}
		} else {
			// 续费成功
			_ = tx.Commit()

		}
	}
}
