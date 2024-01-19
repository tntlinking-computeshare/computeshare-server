// Code generated by ent, DO NOT EDIT.

package cyclerenewal

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldID, id))
}

// FkUserID applies equality check predicate on the "fk_user_id" field. It's identical to FkUserIDEQ.
func FkUserID(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldFkUserID, v))
}

// ResourceID applies equality check predicate on the "resource_id" field. It's identical to ResourceIDEQ.
func ResourceID(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldResourceID, v))
}

// ResourceType applies equality check predicate on the "resource_type" field. It's identical to ResourceTypeEQ.
func ResourceType(v int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldResourceType, v))
}

// ProductName applies equality check predicate on the "product_name" field. It's identical to ProductNameEQ.
func ProductName(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldProductName, v))
}

// ProductDesc applies equality check predicate on the "product_desc" field. It's identical to ProductDescEQ.
func ProductDesc(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldProductDesc, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldState, v))
}

// ExtendDay applies equality check predicate on the "extend_day" field. It's identical to ExtendDayEQ.
func ExtendDay(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldExtendDay, v))
}

// ExtendPrice applies equality check predicate on the "extend_price" field. It's identical to ExtendPriceEQ.
func ExtendPrice(v float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldExtendPrice, v))
}

// DueTime applies equality check predicate on the "due_time" field. It's identical to DueTimeEQ.
func DueTime(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldDueTime, v))
}

// RenewalTime applies equality check predicate on the "renewal_time" field. It's identical to RenewalTimeEQ.
func RenewalTime(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldRenewalTime, v))
}

// AutoRenewal applies equality check predicate on the "auto_renewal" field. It's identical to AutoRenewalEQ.
func AutoRenewal(v bool) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldAutoRenewal, v))
}

// FkUserIDEQ applies the EQ predicate on the "fk_user_id" field.
func FkUserIDEQ(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldFkUserID, v))
}

// FkUserIDNEQ applies the NEQ predicate on the "fk_user_id" field.
func FkUserIDNEQ(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldFkUserID, v))
}

// FkUserIDIn applies the In predicate on the "fk_user_id" field.
func FkUserIDIn(vs ...uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldFkUserID, vs...))
}

// FkUserIDNotIn applies the NotIn predicate on the "fk_user_id" field.
func FkUserIDNotIn(vs ...uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldFkUserID, vs...))
}

// FkUserIDGT applies the GT predicate on the "fk_user_id" field.
func FkUserIDGT(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldFkUserID, v))
}

// FkUserIDGTE applies the GTE predicate on the "fk_user_id" field.
func FkUserIDGTE(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldFkUserID, v))
}

// FkUserIDLT applies the LT predicate on the "fk_user_id" field.
func FkUserIDLT(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldFkUserID, v))
}

// FkUserIDLTE applies the LTE predicate on the "fk_user_id" field.
func FkUserIDLTE(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldFkUserID, v))
}

// ResourceIDEQ applies the EQ predicate on the "resource_id" field.
func ResourceIDEQ(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldResourceID, v))
}

// ResourceIDNEQ applies the NEQ predicate on the "resource_id" field.
func ResourceIDNEQ(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldResourceID, v))
}

// ResourceIDIn applies the In predicate on the "resource_id" field.
func ResourceIDIn(vs ...uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldResourceID, vs...))
}

// ResourceIDNotIn applies the NotIn predicate on the "resource_id" field.
func ResourceIDNotIn(vs ...uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldResourceID, vs...))
}

// ResourceIDGT applies the GT predicate on the "resource_id" field.
func ResourceIDGT(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldResourceID, v))
}

// ResourceIDGTE applies the GTE predicate on the "resource_id" field.
func ResourceIDGTE(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldResourceID, v))
}

// ResourceIDLT applies the LT predicate on the "resource_id" field.
func ResourceIDLT(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldResourceID, v))
}

// ResourceIDLTE applies the LTE predicate on the "resource_id" field.
func ResourceIDLTE(v uuid.UUID) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldResourceID, v))
}

// ResourceTypeEQ applies the EQ predicate on the "resource_type" field.
func ResourceTypeEQ(v int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldResourceType, v))
}

// ResourceTypeNEQ applies the NEQ predicate on the "resource_type" field.
func ResourceTypeNEQ(v int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldResourceType, v))
}

// ResourceTypeIn applies the In predicate on the "resource_type" field.
func ResourceTypeIn(vs ...int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldResourceType, vs...))
}

// ResourceTypeNotIn applies the NotIn predicate on the "resource_type" field.
func ResourceTypeNotIn(vs ...int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldResourceType, vs...))
}

// ResourceTypeGT applies the GT predicate on the "resource_type" field.
func ResourceTypeGT(v int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldResourceType, v))
}

// ResourceTypeGTE applies the GTE predicate on the "resource_type" field.
func ResourceTypeGTE(v int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldResourceType, v))
}

// ResourceTypeLT applies the LT predicate on the "resource_type" field.
func ResourceTypeLT(v int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldResourceType, v))
}

// ResourceTypeLTE applies the LTE predicate on the "resource_type" field.
func ResourceTypeLTE(v int) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldResourceType, v))
}

// ProductNameEQ applies the EQ predicate on the "product_name" field.
func ProductNameEQ(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldProductName, v))
}

// ProductNameNEQ applies the NEQ predicate on the "product_name" field.
func ProductNameNEQ(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldProductName, v))
}

// ProductNameIn applies the In predicate on the "product_name" field.
func ProductNameIn(vs ...string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldProductName, vs...))
}

// ProductNameNotIn applies the NotIn predicate on the "product_name" field.
func ProductNameNotIn(vs ...string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldProductName, vs...))
}

// ProductNameGT applies the GT predicate on the "product_name" field.
func ProductNameGT(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldProductName, v))
}

// ProductNameGTE applies the GTE predicate on the "product_name" field.
func ProductNameGTE(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldProductName, v))
}

// ProductNameLT applies the LT predicate on the "product_name" field.
func ProductNameLT(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldProductName, v))
}

// ProductNameLTE applies the LTE predicate on the "product_name" field.
func ProductNameLTE(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldProductName, v))
}

// ProductNameContains applies the Contains predicate on the "product_name" field.
func ProductNameContains(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldContains(FieldProductName, v))
}

// ProductNameHasPrefix applies the HasPrefix predicate on the "product_name" field.
func ProductNameHasPrefix(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldHasPrefix(FieldProductName, v))
}

// ProductNameHasSuffix applies the HasSuffix predicate on the "product_name" field.
func ProductNameHasSuffix(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldHasSuffix(FieldProductName, v))
}

// ProductNameEqualFold applies the EqualFold predicate on the "product_name" field.
func ProductNameEqualFold(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEqualFold(FieldProductName, v))
}

// ProductNameContainsFold applies the ContainsFold predicate on the "product_name" field.
func ProductNameContainsFold(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldContainsFold(FieldProductName, v))
}

// ProductDescEQ applies the EQ predicate on the "product_desc" field.
func ProductDescEQ(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldProductDesc, v))
}

// ProductDescNEQ applies the NEQ predicate on the "product_desc" field.
func ProductDescNEQ(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldProductDesc, v))
}

// ProductDescIn applies the In predicate on the "product_desc" field.
func ProductDescIn(vs ...string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldProductDesc, vs...))
}

// ProductDescNotIn applies the NotIn predicate on the "product_desc" field.
func ProductDescNotIn(vs ...string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldProductDesc, vs...))
}

// ProductDescGT applies the GT predicate on the "product_desc" field.
func ProductDescGT(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldProductDesc, v))
}

// ProductDescGTE applies the GTE predicate on the "product_desc" field.
func ProductDescGTE(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldProductDesc, v))
}

// ProductDescLT applies the LT predicate on the "product_desc" field.
func ProductDescLT(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldProductDesc, v))
}

// ProductDescLTE applies the LTE predicate on the "product_desc" field.
func ProductDescLTE(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldProductDesc, v))
}

// ProductDescContains applies the Contains predicate on the "product_desc" field.
func ProductDescContains(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldContains(FieldProductDesc, v))
}

// ProductDescHasPrefix applies the HasPrefix predicate on the "product_desc" field.
func ProductDescHasPrefix(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldHasPrefix(FieldProductDesc, v))
}

// ProductDescHasSuffix applies the HasSuffix predicate on the "product_desc" field.
func ProductDescHasSuffix(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldHasSuffix(FieldProductDesc, v))
}

// ProductDescEqualFold applies the EqualFold predicate on the "product_desc" field.
func ProductDescEqualFold(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEqualFold(FieldProductDesc, v))
}

// ProductDescContainsFold applies the ContainsFold predicate on the "product_desc" field.
func ProductDescContainsFold(v string) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldContainsFold(FieldProductDesc, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldState, v))
}

// ExtendDayEQ applies the EQ predicate on the "extend_day" field.
func ExtendDayEQ(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldExtendDay, v))
}

// ExtendDayNEQ applies the NEQ predicate on the "extend_day" field.
func ExtendDayNEQ(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldExtendDay, v))
}

// ExtendDayIn applies the In predicate on the "extend_day" field.
func ExtendDayIn(vs ...int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldExtendDay, vs...))
}

// ExtendDayNotIn applies the NotIn predicate on the "extend_day" field.
func ExtendDayNotIn(vs ...int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldExtendDay, vs...))
}

// ExtendDayGT applies the GT predicate on the "extend_day" field.
func ExtendDayGT(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldExtendDay, v))
}

// ExtendDayGTE applies the GTE predicate on the "extend_day" field.
func ExtendDayGTE(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldExtendDay, v))
}

// ExtendDayLT applies the LT predicate on the "extend_day" field.
func ExtendDayLT(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldExtendDay, v))
}

// ExtendDayLTE applies the LTE predicate on the "extend_day" field.
func ExtendDayLTE(v int8) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldExtendDay, v))
}

// ExtendPriceEQ applies the EQ predicate on the "extend_price" field.
func ExtendPriceEQ(v float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldExtendPrice, v))
}

// ExtendPriceNEQ applies the NEQ predicate on the "extend_price" field.
func ExtendPriceNEQ(v float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldExtendPrice, v))
}

// ExtendPriceIn applies the In predicate on the "extend_price" field.
func ExtendPriceIn(vs ...float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldExtendPrice, vs...))
}

// ExtendPriceNotIn applies the NotIn predicate on the "extend_price" field.
func ExtendPriceNotIn(vs ...float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldExtendPrice, vs...))
}

// ExtendPriceGT applies the GT predicate on the "extend_price" field.
func ExtendPriceGT(v float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldExtendPrice, v))
}

// ExtendPriceGTE applies the GTE predicate on the "extend_price" field.
func ExtendPriceGTE(v float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldExtendPrice, v))
}

// ExtendPriceLT applies the LT predicate on the "extend_price" field.
func ExtendPriceLT(v float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldExtendPrice, v))
}

// ExtendPriceLTE applies the LTE predicate on the "extend_price" field.
func ExtendPriceLTE(v float64) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldExtendPrice, v))
}

// DueTimeEQ applies the EQ predicate on the "due_time" field.
func DueTimeEQ(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldDueTime, v))
}

// DueTimeNEQ applies the NEQ predicate on the "due_time" field.
func DueTimeNEQ(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldDueTime, v))
}

// DueTimeIn applies the In predicate on the "due_time" field.
func DueTimeIn(vs ...time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldDueTime, vs...))
}

// DueTimeNotIn applies the NotIn predicate on the "due_time" field.
func DueTimeNotIn(vs ...time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldDueTime, vs...))
}

// DueTimeGT applies the GT predicate on the "due_time" field.
func DueTimeGT(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldDueTime, v))
}

// DueTimeGTE applies the GTE predicate on the "due_time" field.
func DueTimeGTE(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldDueTime, v))
}

// DueTimeLT applies the LT predicate on the "due_time" field.
func DueTimeLT(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldDueTime, v))
}

// DueTimeLTE applies the LTE predicate on the "due_time" field.
func DueTimeLTE(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldDueTime, v))
}

// DueTimeIsNil applies the IsNil predicate on the "due_time" field.
func DueTimeIsNil() predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIsNull(FieldDueTime))
}

// DueTimeNotNil applies the NotNil predicate on the "due_time" field.
func DueTimeNotNil() predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotNull(FieldDueTime))
}

// RenewalTimeEQ applies the EQ predicate on the "renewal_time" field.
func RenewalTimeEQ(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldRenewalTime, v))
}

// RenewalTimeNEQ applies the NEQ predicate on the "renewal_time" field.
func RenewalTimeNEQ(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldRenewalTime, v))
}

// RenewalTimeIn applies the In predicate on the "renewal_time" field.
func RenewalTimeIn(vs ...time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIn(FieldRenewalTime, vs...))
}

// RenewalTimeNotIn applies the NotIn predicate on the "renewal_time" field.
func RenewalTimeNotIn(vs ...time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotIn(FieldRenewalTime, vs...))
}

// RenewalTimeGT applies the GT predicate on the "renewal_time" field.
func RenewalTimeGT(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGT(FieldRenewalTime, v))
}

// RenewalTimeGTE applies the GTE predicate on the "renewal_time" field.
func RenewalTimeGTE(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldGTE(FieldRenewalTime, v))
}

// RenewalTimeLT applies the LT predicate on the "renewal_time" field.
func RenewalTimeLT(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLT(FieldRenewalTime, v))
}

// RenewalTimeLTE applies the LTE predicate on the "renewal_time" field.
func RenewalTimeLTE(v time.Time) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldLTE(FieldRenewalTime, v))
}

// RenewalTimeIsNil applies the IsNil predicate on the "renewal_time" field.
func RenewalTimeIsNil() predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldIsNull(FieldRenewalTime))
}

// RenewalTimeNotNil applies the NotNil predicate on the "renewal_time" field.
func RenewalTimeNotNil() predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNotNull(FieldRenewalTime))
}

// AutoRenewalEQ applies the EQ predicate on the "auto_renewal" field.
func AutoRenewalEQ(v bool) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldEQ(FieldAutoRenewal, v))
}

// AutoRenewalNEQ applies the NEQ predicate on the "auto_renewal" field.
func AutoRenewalNEQ(v bool) predicate.CycleRenewal {
	return predicate.CycleRenewal(sql.FieldNEQ(FieldAutoRenewal, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CycleRenewal) predicate.CycleRenewal {
	return predicate.CycleRenewal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CycleRenewal) predicate.CycleRenewal {
	return predicate.CycleRenewal(func(s *sql.Selector) {
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
func Not(p predicate.CycleRenewal) predicate.CycleRenewal {
	return predicate.CycleRenewal(func(s *sql.Selector) {
		p(s.Not())
	})
}