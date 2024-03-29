// Code generated by ent, DO NOT EDIT.

package userresourcelimit

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLTE(FieldID, id))
}

// FkUserID applies equality check predicate on the "fk_user_id" field. It's identical to FkUserIDEQ.
func FkUserID(v uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldFkUserID, v))
}

// MaxCPU applies equality check predicate on the "max_cpu" field. It's identical to MaxCPUEQ.
func MaxCPU(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldMaxCPU, v))
}

// MaxMemory applies equality check predicate on the "max_memory" field. It's identical to MaxMemoryEQ.
func MaxMemory(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldMaxMemory, v))
}

// MaxNetworkMapping applies equality check predicate on the "max_network_mapping" field. It's identical to MaxNetworkMappingEQ.
func MaxNetworkMapping(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldMaxNetworkMapping, v))
}

// FkUserIDEQ applies the EQ predicate on the "fk_user_id" field.
func FkUserIDEQ(v uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldFkUserID, v))
}

// FkUserIDNEQ applies the NEQ predicate on the "fk_user_id" field.
func FkUserIDNEQ(v uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNEQ(FieldFkUserID, v))
}

// FkUserIDIn applies the In predicate on the "fk_user_id" field.
func FkUserIDIn(vs ...uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldIn(FieldFkUserID, vs...))
}

// FkUserIDNotIn applies the NotIn predicate on the "fk_user_id" field.
func FkUserIDNotIn(vs ...uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNotIn(FieldFkUserID, vs...))
}

// FkUserIDGT applies the GT predicate on the "fk_user_id" field.
func FkUserIDGT(v uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGT(FieldFkUserID, v))
}

// FkUserIDGTE applies the GTE predicate on the "fk_user_id" field.
func FkUserIDGTE(v uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGTE(FieldFkUserID, v))
}

// FkUserIDLT applies the LT predicate on the "fk_user_id" field.
func FkUserIDLT(v uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLT(FieldFkUserID, v))
}

// FkUserIDLTE applies the LTE predicate on the "fk_user_id" field.
func FkUserIDLTE(v uuid.UUID) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLTE(FieldFkUserID, v))
}

// MaxCPUEQ applies the EQ predicate on the "max_cpu" field.
func MaxCPUEQ(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldMaxCPU, v))
}

// MaxCPUNEQ applies the NEQ predicate on the "max_cpu" field.
func MaxCPUNEQ(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNEQ(FieldMaxCPU, v))
}

// MaxCPUIn applies the In predicate on the "max_cpu" field.
func MaxCPUIn(vs ...int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldIn(FieldMaxCPU, vs...))
}

// MaxCPUNotIn applies the NotIn predicate on the "max_cpu" field.
func MaxCPUNotIn(vs ...int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNotIn(FieldMaxCPU, vs...))
}

// MaxCPUGT applies the GT predicate on the "max_cpu" field.
func MaxCPUGT(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGT(FieldMaxCPU, v))
}

// MaxCPUGTE applies the GTE predicate on the "max_cpu" field.
func MaxCPUGTE(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGTE(FieldMaxCPU, v))
}

// MaxCPULT applies the LT predicate on the "max_cpu" field.
func MaxCPULT(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLT(FieldMaxCPU, v))
}

// MaxCPULTE applies the LTE predicate on the "max_cpu" field.
func MaxCPULTE(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLTE(FieldMaxCPU, v))
}

// MaxMemoryEQ applies the EQ predicate on the "max_memory" field.
func MaxMemoryEQ(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldMaxMemory, v))
}

// MaxMemoryNEQ applies the NEQ predicate on the "max_memory" field.
func MaxMemoryNEQ(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNEQ(FieldMaxMemory, v))
}

// MaxMemoryIn applies the In predicate on the "max_memory" field.
func MaxMemoryIn(vs ...int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldIn(FieldMaxMemory, vs...))
}

// MaxMemoryNotIn applies the NotIn predicate on the "max_memory" field.
func MaxMemoryNotIn(vs ...int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNotIn(FieldMaxMemory, vs...))
}

// MaxMemoryGT applies the GT predicate on the "max_memory" field.
func MaxMemoryGT(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGT(FieldMaxMemory, v))
}

// MaxMemoryGTE applies the GTE predicate on the "max_memory" field.
func MaxMemoryGTE(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGTE(FieldMaxMemory, v))
}

// MaxMemoryLT applies the LT predicate on the "max_memory" field.
func MaxMemoryLT(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLT(FieldMaxMemory, v))
}

// MaxMemoryLTE applies the LTE predicate on the "max_memory" field.
func MaxMemoryLTE(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLTE(FieldMaxMemory, v))
}

// MaxNetworkMappingEQ applies the EQ predicate on the "max_network_mapping" field.
func MaxNetworkMappingEQ(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldEQ(FieldMaxNetworkMapping, v))
}

// MaxNetworkMappingNEQ applies the NEQ predicate on the "max_network_mapping" field.
func MaxNetworkMappingNEQ(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNEQ(FieldMaxNetworkMapping, v))
}

// MaxNetworkMappingIn applies the In predicate on the "max_network_mapping" field.
func MaxNetworkMappingIn(vs ...int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldIn(FieldMaxNetworkMapping, vs...))
}

// MaxNetworkMappingNotIn applies the NotIn predicate on the "max_network_mapping" field.
func MaxNetworkMappingNotIn(vs ...int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldNotIn(FieldMaxNetworkMapping, vs...))
}

// MaxNetworkMappingGT applies the GT predicate on the "max_network_mapping" field.
func MaxNetworkMappingGT(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGT(FieldMaxNetworkMapping, v))
}

// MaxNetworkMappingGTE applies the GTE predicate on the "max_network_mapping" field.
func MaxNetworkMappingGTE(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldGTE(FieldMaxNetworkMapping, v))
}

// MaxNetworkMappingLT applies the LT predicate on the "max_network_mapping" field.
func MaxNetworkMappingLT(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLT(FieldMaxNetworkMapping, v))
}

// MaxNetworkMappingLTE applies the LTE predicate on the "max_network_mapping" field.
func MaxNetworkMappingLTE(v int32) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(sql.FieldLTE(FieldMaxNetworkMapping, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserResourceLimit) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserResourceLimit) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(func(s *sql.Selector) {
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
func Not(p predicate.UserResourceLimit) predicate.UserResourceLimit {
	return predicate.UserResourceLimit(func(s *sql.Selector) {
		p(s.Not())
	})
}
