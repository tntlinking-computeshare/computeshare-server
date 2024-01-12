package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CycleTransaction holds the schema definition for the CycleTransaction entity.
type CycleRecharge struct {
	ent.Schema
}

// Fields of the CycleTransaction.
func (CycleRecharge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_user_id", uuid.UUID{}).Comment("用户id"),
		field.String("out_trade_no").Unique().Comment("商家订单号"),
		field.String("alipay_trade_no").Comment("支付宝订单号"),
		field.Int("recharge_channel").Comment("充值渠道"),
		field.String("redeem_code").Comment("兑换码"),
		field.String("state").Comment("状态"),
		field.Float("pay_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}).Comment("支付的钱"),
		field.Float("total_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}).Comment("收到的钱"),
		field.Float("buy_cycle").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}).Comment("购买的周期"),
		field.Time("create_time").Comment("创建时间"),
		field.Time("update_time").Comment("更新时间"),
	}
}

// Edges of the CycleTransaction.
func (CycleRecharge) Edges() []ent.Edge {
	//}
	return nil
}
