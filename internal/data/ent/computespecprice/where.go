// Code generated by ent, DO NOT EDIT.

package computespecprice

import (
	"entgo.io/ent/dialect/sql"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLTE(FieldID, id))
}

// FkComputeSpecID applies equality check predicate on the "fk_compute_spec_id" field. It's identical to FkComputeSpecIDEQ.
func FkComputeSpecID(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldFkComputeSpecID, v))
}

// Day applies equality check predicate on the "day" field. It's identical to DayEQ.
func Day(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldDay, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldPrice, v))
}

// FkComputeSpecIDEQ applies the EQ predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDEQ(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldFkComputeSpecID, v))
}

// FkComputeSpecIDNEQ applies the NEQ predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDNEQ(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNEQ(FieldFkComputeSpecID, v))
}

// FkComputeSpecIDIn applies the In predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDIn(vs ...int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldIn(FieldFkComputeSpecID, vs...))
}

// FkComputeSpecIDNotIn applies the NotIn predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDNotIn(vs ...int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNotIn(FieldFkComputeSpecID, vs...))
}

// FkComputeSpecIDGT applies the GT predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDGT(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGT(FieldFkComputeSpecID, v))
}

// FkComputeSpecIDGTE applies the GTE predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDGTE(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGTE(FieldFkComputeSpecID, v))
}

// FkComputeSpecIDLT applies the LT predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDLT(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLT(FieldFkComputeSpecID, v))
}

// FkComputeSpecIDLTE applies the LTE predicate on the "fk_compute_spec_id" field.
func FkComputeSpecIDLTE(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLTE(FieldFkComputeSpecID, v))
}

// DayEQ applies the EQ predicate on the "day" field.
func DayEQ(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldDay, v))
}

// DayNEQ applies the NEQ predicate on the "day" field.
func DayNEQ(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNEQ(FieldDay, v))
}

// DayIn applies the In predicate on the "day" field.
func DayIn(vs ...int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldIn(FieldDay, vs...))
}

// DayNotIn applies the NotIn predicate on the "day" field.
func DayNotIn(vs ...int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNotIn(FieldDay, vs...))
}

// DayGT applies the GT predicate on the "day" field.
func DayGT(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGT(FieldDay, v))
}

// DayGTE applies the GTE predicate on the "day" field.
func DayGTE(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGTE(FieldDay, v))
}

// DayLT applies the LT predicate on the "day" field.
func DayLT(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLT(FieldDay, v))
}

// DayLTE applies the LTE predicate on the "day" field.
func DayLTE(v int32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLTE(FieldDay, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float32) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(sql.FieldLTE(FieldPrice, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ComputeSpecPrice) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ComputeSpecPrice) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(func(s *sql.Selector) {
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
func Not(p predicate.ComputeSpecPrice) predicate.ComputeSpecPrice {
	return predicate.ComputeSpecPrice(func(s *sql.Selector) {
		p(s.Not())
	})
}