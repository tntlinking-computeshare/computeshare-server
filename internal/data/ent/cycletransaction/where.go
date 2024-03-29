// Code generated by ent, DO NOT EDIT.

package cycletransaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldID, id))
}

// FkCycleID applies equality check predicate on the "fk_cycle_id" field. It's identical to FkCycleIDEQ.
func FkCycleID(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkCycleID, v))
}

// FkUserID applies equality check predicate on the "fk_user_id" field. It's identical to FkUserIDEQ.
func FkUserID(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkUserID, v))
}

// FkCycleOrderID applies equality check predicate on the "fk_cycle_order_id" field. It's identical to FkCycleOrderIDEQ.
func FkCycleOrderID(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkCycleOrderID, v))
}

// FkCycleRechargeID applies equality check predicate on the "fk_cycle_recharge_id" field. It's identical to FkCycleRechargeIDEQ.
func FkCycleRechargeID(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkCycleRechargeID, v))
}

// Operation applies equality check predicate on the "operation" field. It's identical to OperationEQ.
func Operation(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldOperation, v))
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldSymbol, v))
}

// Cycle applies equality check predicate on the "cycle" field. It's identical to CycleEQ.
func Cycle(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldCycle, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldBalance, v))
}

// OperationTime applies equality check predicate on the "operation_time" field. It's identical to OperationTimeEQ.
func OperationTime(v time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldOperationTime, v))
}

// FkCycleIDEQ applies the EQ predicate on the "fk_cycle_id" field.
func FkCycleIDEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkCycleID, v))
}

// FkCycleIDNEQ applies the NEQ predicate on the "fk_cycle_id" field.
func FkCycleIDNEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldFkCycleID, v))
}

// FkCycleIDIn applies the In predicate on the "fk_cycle_id" field.
func FkCycleIDIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldFkCycleID, vs...))
}

// FkCycleIDNotIn applies the NotIn predicate on the "fk_cycle_id" field.
func FkCycleIDNotIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldFkCycleID, vs...))
}

// FkCycleIDGT applies the GT predicate on the "fk_cycle_id" field.
func FkCycleIDGT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldFkCycleID, v))
}

// FkCycleIDGTE applies the GTE predicate on the "fk_cycle_id" field.
func FkCycleIDGTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldFkCycleID, v))
}

// FkCycleIDLT applies the LT predicate on the "fk_cycle_id" field.
func FkCycleIDLT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldFkCycleID, v))
}

// FkCycleIDLTE applies the LTE predicate on the "fk_cycle_id" field.
func FkCycleIDLTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldFkCycleID, v))
}

// FkUserIDEQ applies the EQ predicate on the "fk_user_id" field.
func FkUserIDEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkUserID, v))
}

// FkUserIDNEQ applies the NEQ predicate on the "fk_user_id" field.
func FkUserIDNEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldFkUserID, v))
}

// FkUserIDIn applies the In predicate on the "fk_user_id" field.
func FkUserIDIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldFkUserID, vs...))
}

// FkUserIDNotIn applies the NotIn predicate on the "fk_user_id" field.
func FkUserIDNotIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldFkUserID, vs...))
}

// FkUserIDGT applies the GT predicate on the "fk_user_id" field.
func FkUserIDGT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldFkUserID, v))
}

// FkUserIDGTE applies the GTE predicate on the "fk_user_id" field.
func FkUserIDGTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldFkUserID, v))
}

// FkUserIDLT applies the LT predicate on the "fk_user_id" field.
func FkUserIDLT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldFkUserID, v))
}

// FkUserIDLTE applies the LTE predicate on the "fk_user_id" field.
func FkUserIDLTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldFkUserID, v))
}

// FkCycleOrderIDEQ applies the EQ predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkCycleOrderID, v))
}

// FkCycleOrderIDNEQ applies the NEQ predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDNEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldFkCycleOrderID, v))
}

// FkCycleOrderIDIn applies the In predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldFkCycleOrderID, vs...))
}

// FkCycleOrderIDNotIn applies the NotIn predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDNotIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldFkCycleOrderID, vs...))
}

// FkCycleOrderIDGT applies the GT predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDGT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldFkCycleOrderID, v))
}

// FkCycleOrderIDGTE applies the GTE predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDGTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldFkCycleOrderID, v))
}

// FkCycleOrderIDLT applies the LT predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDLT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldFkCycleOrderID, v))
}

// FkCycleOrderIDLTE applies the LTE predicate on the "fk_cycle_order_id" field.
func FkCycleOrderIDLTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldFkCycleOrderID, v))
}

// FkCycleRechargeIDEQ applies the EQ predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldFkCycleRechargeID, v))
}

// FkCycleRechargeIDNEQ applies the NEQ predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDNEQ(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldFkCycleRechargeID, v))
}

// FkCycleRechargeIDIn applies the In predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldFkCycleRechargeID, vs...))
}

// FkCycleRechargeIDNotIn applies the NotIn predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDNotIn(vs ...uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldFkCycleRechargeID, vs...))
}

// FkCycleRechargeIDGT applies the GT predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDGT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldFkCycleRechargeID, v))
}

// FkCycleRechargeIDGTE applies the GTE predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDGTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldFkCycleRechargeID, v))
}

// FkCycleRechargeIDLT applies the LT predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDLT(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldFkCycleRechargeID, v))
}

// FkCycleRechargeIDLTE applies the LTE predicate on the "fk_cycle_recharge_id" field.
func FkCycleRechargeIDLTE(v uuid.UUID) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldFkCycleRechargeID, v))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldOperation, vs...))
}

// OperationGT applies the GT predicate on the "operation" field.
func OperationGT(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldOperation, v))
}

// OperationGTE applies the GTE predicate on the "operation" field.
func OperationGTE(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldOperation, v))
}

// OperationLT applies the LT predicate on the "operation" field.
func OperationLT(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldOperation, v))
}

// OperationLTE applies the LTE predicate on the "operation" field.
func OperationLTE(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldOperation, v))
}

// OperationContains applies the Contains predicate on the "operation" field.
func OperationContains(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldContains(FieldOperation, v))
}

// OperationHasPrefix applies the HasPrefix predicate on the "operation" field.
func OperationHasPrefix(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldHasPrefix(FieldOperation, v))
}

// OperationHasSuffix applies the HasSuffix predicate on the "operation" field.
func OperationHasSuffix(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldHasSuffix(FieldOperation, v))
}

// OperationEqualFold applies the EqualFold predicate on the "operation" field.
func OperationEqualFold(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEqualFold(FieldOperation, v))
}

// OperationContainsFold applies the ContainsFold predicate on the "operation" field.
func OperationContainsFold(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldContainsFold(FieldOperation, v))
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldSymbol, v))
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldSymbol, v))
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldSymbol, vs...))
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldSymbol, vs...))
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldSymbol, v))
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldSymbol, v))
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldSymbol, v))
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldSymbol, v))
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldContains(FieldSymbol, v))
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldHasPrefix(FieldSymbol, v))
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldHasSuffix(FieldSymbol, v))
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEqualFold(FieldSymbol, v))
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldContainsFold(FieldSymbol, v))
}

// CycleEQ applies the EQ predicate on the "cycle" field.
func CycleEQ(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldCycle, v))
}

// CycleNEQ applies the NEQ predicate on the "cycle" field.
func CycleNEQ(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldCycle, v))
}

// CycleIn applies the In predicate on the "cycle" field.
func CycleIn(vs ...float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldCycle, vs...))
}

// CycleNotIn applies the NotIn predicate on the "cycle" field.
func CycleNotIn(vs ...float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldCycle, vs...))
}

// CycleGT applies the GT predicate on the "cycle" field.
func CycleGT(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldCycle, v))
}

// CycleGTE applies the GTE predicate on the "cycle" field.
func CycleGTE(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldCycle, v))
}

// CycleLT applies the LT predicate on the "cycle" field.
func CycleLT(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldCycle, v))
}

// CycleLTE applies the LTE predicate on the "cycle" field.
func CycleLTE(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldCycle, v))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float64) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldBalance, v))
}

// OperationTimeEQ applies the EQ predicate on the "operation_time" field.
func OperationTimeEQ(v time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldEQ(FieldOperationTime, v))
}

// OperationTimeNEQ applies the NEQ predicate on the "operation_time" field.
func OperationTimeNEQ(v time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNEQ(FieldOperationTime, v))
}

// OperationTimeIn applies the In predicate on the "operation_time" field.
func OperationTimeIn(vs ...time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldIn(FieldOperationTime, vs...))
}

// OperationTimeNotIn applies the NotIn predicate on the "operation_time" field.
func OperationTimeNotIn(vs ...time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldNotIn(FieldOperationTime, vs...))
}

// OperationTimeGT applies the GT predicate on the "operation_time" field.
func OperationTimeGT(v time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGT(FieldOperationTime, v))
}

// OperationTimeGTE applies the GTE predicate on the "operation_time" field.
func OperationTimeGTE(v time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldGTE(FieldOperationTime, v))
}

// OperationTimeLT applies the LT predicate on the "operation_time" field.
func OperationTimeLT(v time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLT(FieldOperationTime, v))
}

// OperationTimeLTE applies the LTE predicate on the "operation_time" field.
func OperationTimeLTE(v time.Time) predicate.CycleTransaction {
	return predicate.CycleTransaction(sql.FieldLTE(FieldOperationTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CycleTransaction) predicate.CycleTransaction {
	return predicate.CycleTransaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CycleTransaction) predicate.CycleTransaction {
	return predicate.CycleTransaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CycleTransaction) predicate.CycleTransaction {
	return predicate.CycleTransaction(func(s *sql.Selector) {
		p(s.Not())
	})
}
