package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AlipayOrderRollback holds the schema definition for the Cycle entity.
type AlipayOrderRollback struct {
	ent.Schema
}

// Fields of the Cycle.
func (AlipayOrderRollback) Fields() []ent.Field {
	return []ent.Field{
		field.String("notify_id").Unique().Comment("通知校验ID"),
		field.String("notify_type").Comment("通知类型"),
		field.String("notify_time").Comment("通知时间"),
		field.String("charset").Comment("编码格式，如 utf-8、gbk、gb2312 等"),
		field.String("version").Comment("调用的接口版本，固定为：1.0"),
		field.String("sign_type").Comment("签名类型"),
		field.Text("sign").Comment("签名"),
		field.Text("fund_bill_list").Comment("支付成功的各个渠道金额信息。详情可查看 资金明细信息说明"),
		field.String("receipt_amount").Comment("实收金额"),
		field.String("invoice_amount").Comment("用户在交易中支付的可开发票的金额"),
		field.String("buyer_pay_amount").Comment("付款金额"),
		field.String("point_amount").Comment("集分宝金额"),
		field.Text("voucher_detail_list").Comment("本交易支付时所有优惠券信息，详情可查看 优惠券信息说明"),
		field.Text("passback_params").Comment("公共回传参数，如果请求时传递了该参数，则返回给商家时会在异步通知时将该参数原样返回。本参数必须进行 UrlEncode 之后才可以发送给支付宝。"),
		field.String("trade_no").Unique().Comment("支付宝交易号"),
		field.String("app_id").Comment("开发者id"),
		field.String("out_trade_no").Unique().Comment("商户订单号"),
		field.String("out_biz_no").Comment("商户业务号"),
		field.String("buyer_id").Comment("买家支付宝ID"),
		field.String("seller_id").Comment("卖家支付宝id"),
		field.String("trade_status").Comment("交易状态"),
		field.String("total_amount").Comment("订单金额"),
		field.String("refund_fee").Comment("总退款金额"),
		field.String("subject").Comment("订单标题"),
		field.String("body").Comment("订单的备注、描述、明细等。对应请求时的 body 参数，原样通知回来"),
		field.String("gmt_create").Comment("交易创建时间"),
		field.String("gmt_payment").Comment("交易付款时间"),
		field.String("gmt_close").Comment("交易关闭时间"),
		field.Time("create_time").Comment("创建时间"),
		field.Time("update_time").Comment("更新时间"),
	}
}

// Edges of the Cycle.
func (AlipayOrderRollback) Edges() []ent.Edge {
	//}
	return nil
}
