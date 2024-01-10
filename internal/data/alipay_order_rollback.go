package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/alipayorderrollback"
	"time"
)

type alipayOrderRollbackRepo struct {
	data *Data
	log  *log.Helper
}

func NewAlipayOrderRollbackRepo(data *Data, logger log.Logger) biz.AlipayOrderRollbackRepo {
	return &alipayOrderRollbackRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *alipayOrderRollbackRepo) FindByOutTradeNo(ctx context.Context, outTradeNo string) (*biz.AlipayOrderRollback, error) {
	alipayOrderRollback, err := a.data.getAlipayOrderRollback(ctx).Query().Where(alipayorderrollback.OutTradeNo(outTradeNo)).First(ctx)
	if err != nil {
		return nil, err
	}
	return a.toBiz(alipayOrderRollback, 0), nil
}

func (a *alipayOrderRollbackRepo) SaveAlipayOrderRollback(ctx context.Context, rollback *biz.AlipayOrderRollback) (*biz.AlipayOrderRollback, error) {
	orderRollback, err := a.data.getAlipayOrderRollback(ctx).Query().Where(alipayorderrollback.OutTradeNo(rollback.OutTradeNo)).First(ctx)
	if err == nil && orderRollback != nil {
		err := a.data.getAlipayOrderRollback(ctx).UpdateOne(orderRollback).SetNotifyID(rollback.NotifyID).SetNotifyTime(rollback.NotifyType).
			SetNotifyTime(rollback.NotifyTime).SetCharset(rollback.Charset).SetVersion(rollback.Version).SetSignType(rollback.SignType).
			SetSign(rollback.Sign).SetFundBillList(rollback.FundBillList).SetReceiptAmount(rollback.ReceiptAmount).SetInvoiceAmount(rollback.InvoiceAmount).
			SetBuyerPayAmount(rollback.BuyerPayAmount).SetPointAmount(rollback.PointAmount).SetVoucherDetailList(rollback.VoucherDetailList).
			SetPassbackParams(rollback.PassbackParams).SetTradeNo(rollback.TradeNo).SetAppID(rollback.AppID).SetOutTradeNo(rollback.OutTradeNo).
			SetOutBizNo(rollback.OutBizNo).SetBuyerID(rollback.BuyerID).SetSellerID(rollback.SellerID).SetTradeStatus(rollback.TradeStatus).
			SetTotalAmount(rollback.TotalAmount).SetRefundFee(rollback.RefundFee).SetSubject(rollback.Subject).SetBody(rollback.Body).
			SetGmtCreate(rollback.GmtCreate).SetGmtPayment(rollback.GmtPayment).SetGmtClose(rollback.GmtClose).SetUpdateTime(time.Now()).Exec(ctx)
		if err != nil {
			return nil, err
		}
	} else if err != nil && orderRollback == nil {
		err := a.data.getAlipayOrderRollback(ctx).Create().SetNotifyID(rollback.NotifyID).SetNotifyTime(rollback.NotifyType).
			SetNotifyTime(rollback.NotifyTime).SetCharset(rollback.Charset).SetVersion(rollback.Version).SetSignType(rollback.SignType).
			SetSign(rollback.Sign).SetFundBillList(rollback.FundBillList).SetReceiptAmount(rollback.ReceiptAmount).SetInvoiceAmount(rollback.InvoiceAmount).
			SetBuyerPayAmount(rollback.BuyerPayAmount).SetPointAmount(rollback.PointAmount).SetVoucherDetailList(rollback.VoucherDetailList).
			SetPassbackParams(rollback.PassbackParams).SetTradeNo(rollback.TradeNo).SetAppID(rollback.AppID).SetOutTradeNo(rollback.OutTradeNo).
			SetOutBizNo(rollback.OutBizNo).SetBuyerID(rollback.BuyerID).SetSellerID(rollback.SellerID).SetTradeStatus(rollback.TradeStatus).
			SetTotalAmount(rollback.TotalAmount).SetRefundFee(rollback.RefundFee).SetSubject(rollback.Subject).SetBody(rollback.Body).
			SetGmtCreate(rollback.GmtCreate).SetGmtPayment(rollback.GmtPayment).SetGmtClose(rollback.GmtClose).
			SetCreateTime(time.Now()).SetUpdateTime(time.Now()).Exec(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	return rollback, nil
}

func (a *alipayOrderRollbackRepo) toBiz(p *ent.AlipayOrderRollback, _ int) *biz.AlipayOrderRollback {
	if p == nil {
		return nil
	}
	var alipayOrderRollback biz.AlipayOrderRollback
	copier.Copy(alipayOrderRollback, p)
	return &alipayOrderRollback
}
