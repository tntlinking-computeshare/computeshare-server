// Code generated by ent, DO NOT EDIT.

package agent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldID, id))
}

// MAC applies equality check predicate on the "mac" field. It's identical to MACEQ.
func MAC(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldMAC, v))
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldActive, v))
}

// LastUpdateTime applies equality check predicate on the "last_update_time" field. It's identical to LastUpdateTimeEQ.
func LastUpdateTime(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldLastUpdateTime, v))
}

// Hostname applies equality check predicate on the "hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldHostname, v))
}

// TotalCPU applies equality check predicate on the "total_cpu" field. It's identical to TotalCPUEQ.
func TotalCPU(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldTotalCPU, v))
}

// TotalMemory applies equality check predicate on the "total_memory" field. It's identical to TotalMemoryEQ.
func TotalMemory(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldTotalMemory, v))
}

// OccupiedCPU applies equality check predicate on the "occupied_cpu" field. It's identical to OccupiedCPUEQ.
func OccupiedCPU(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldOccupiedCPU, v))
}

// OccupiedMemory applies equality check predicate on the "occupied_memory" field. It's identical to OccupiedMemoryEQ.
func OccupiedMemory(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldOccupiedMemory, v))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldIP, v))
}

// MACEQ applies the EQ predicate on the "mac" field.
func MACEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldMAC, v))
}

// MACNEQ applies the NEQ predicate on the "mac" field.
func MACNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldMAC, v))
}

// MACIn applies the In predicate on the "mac" field.
func MACIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldMAC, vs...))
}

// MACNotIn applies the NotIn predicate on the "mac" field.
func MACNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldMAC, vs...))
}

// MACGT applies the GT predicate on the "mac" field.
func MACGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldMAC, v))
}

// MACGTE applies the GTE predicate on the "mac" field.
func MACGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldMAC, v))
}

// MACLT applies the LT predicate on the "mac" field.
func MACLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldMAC, v))
}

// MACLTE applies the LTE predicate on the "mac" field.
func MACLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldMAC, v))
}

// MACContains applies the Contains predicate on the "mac" field.
func MACContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldMAC, v))
}

// MACHasPrefix applies the HasPrefix predicate on the "mac" field.
func MACHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldMAC, v))
}

// MACHasSuffix applies the HasSuffix predicate on the "mac" field.
func MACHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldMAC, v))
}

// MACEqualFold applies the EqualFold predicate on the "mac" field.
func MACEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldMAC, v))
}

// MACContainsFold applies the ContainsFold predicate on the "mac" field.
func MACContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldMAC, v))
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldActive, v))
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldActive, v))
}

// LastUpdateTimeEQ applies the EQ predicate on the "last_update_time" field.
func LastUpdateTimeEQ(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldLastUpdateTime, v))
}

// LastUpdateTimeNEQ applies the NEQ predicate on the "last_update_time" field.
func LastUpdateTimeNEQ(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldLastUpdateTime, v))
}

// LastUpdateTimeIn applies the In predicate on the "last_update_time" field.
func LastUpdateTimeIn(vs ...time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldLastUpdateTime, vs...))
}

// LastUpdateTimeNotIn applies the NotIn predicate on the "last_update_time" field.
func LastUpdateTimeNotIn(vs ...time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldLastUpdateTime, vs...))
}

// LastUpdateTimeGT applies the GT predicate on the "last_update_time" field.
func LastUpdateTimeGT(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldLastUpdateTime, v))
}

// LastUpdateTimeGTE applies the GTE predicate on the "last_update_time" field.
func LastUpdateTimeGTE(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldLastUpdateTime, v))
}

// LastUpdateTimeLT applies the LT predicate on the "last_update_time" field.
func LastUpdateTimeLT(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldLastUpdateTime, v))
}

// LastUpdateTimeLTE applies the LTE predicate on the "last_update_time" field.
func LastUpdateTimeLTE(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldLastUpdateTime, v))
}

// HostnameEQ applies the EQ predicate on the "hostname" field.
func HostnameEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldHostname, v))
}

// HostnameNEQ applies the NEQ predicate on the "hostname" field.
func HostnameNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldHostname, v))
}

// HostnameIn applies the In predicate on the "hostname" field.
func HostnameIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldHostname, vs...))
}

// HostnameNotIn applies the NotIn predicate on the "hostname" field.
func HostnameNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldHostname, vs...))
}

// HostnameGT applies the GT predicate on the "hostname" field.
func HostnameGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldHostname, v))
}

// HostnameGTE applies the GTE predicate on the "hostname" field.
func HostnameGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldHostname, v))
}

// HostnameLT applies the LT predicate on the "hostname" field.
func HostnameLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldHostname, v))
}

// HostnameLTE applies the LTE predicate on the "hostname" field.
func HostnameLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldHostname, v))
}

// HostnameContains applies the Contains predicate on the "hostname" field.
func HostnameContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldHostname, v))
}

// HostnameHasPrefix applies the HasPrefix predicate on the "hostname" field.
func HostnameHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldHostname, v))
}

// HostnameHasSuffix applies the HasSuffix predicate on the "hostname" field.
func HostnameHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldHostname, v))
}

// HostnameEqualFold applies the EqualFold predicate on the "hostname" field.
func HostnameEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldHostname, v))
}

// HostnameContainsFold applies the ContainsFold predicate on the "hostname" field.
func HostnameContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldHostname, v))
}

// TotalCPUEQ applies the EQ predicate on the "total_cpu" field.
func TotalCPUEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldTotalCPU, v))
}

// TotalCPUNEQ applies the NEQ predicate on the "total_cpu" field.
func TotalCPUNEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldTotalCPU, v))
}

// TotalCPUIn applies the In predicate on the "total_cpu" field.
func TotalCPUIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldTotalCPU, vs...))
}

// TotalCPUNotIn applies the NotIn predicate on the "total_cpu" field.
func TotalCPUNotIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldTotalCPU, vs...))
}

// TotalCPUGT applies the GT predicate on the "total_cpu" field.
func TotalCPUGT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldTotalCPU, v))
}

// TotalCPUGTE applies the GTE predicate on the "total_cpu" field.
func TotalCPUGTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldTotalCPU, v))
}

// TotalCPULT applies the LT predicate on the "total_cpu" field.
func TotalCPULT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldTotalCPU, v))
}

// TotalCPULTE applies the LTE predicate on the "total_cpu" field.
func TotalCPULTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldTotalCPU, v))
}

// TotalMemoryEQ applies the EQ predicate on the "total_memory" field.
func TotalMemoryEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldTotalMemory, v))
}

// TotalMemoryNEQ applies the NEQ predicate on the "total_memory" field.
func TotalMemoryNEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldTotalMemory, v))
}

// TotalMemoryIn applies the In predicate on the "total_memory" field.
func TotalMemoryIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldTotalMemory, vs...))
}

// TotalMemoryNotIn applies the NotIn predicate on the "total_memory" field.
func TotalMemoryNotIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldTotalMemory, vs...))
}

// TotalMemoryGT applies the GT predicate on the "total_memory" field.
func TotalMemoryGT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldTotalMemory, v))
}

// TotalMemoryGTE applies the GTE predicate on the "total_memory" field.
func TotalMemoryGTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldTotalMemory, v))
}

// TotalMemoryLT applies the LT predicate on the "total_memory" field.
func TotalMemoryLT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldTotalMemory, v))
}

// TotalMemoryLTE applies the LTE predicate on the "total_memory" field.
func TotalMemoryLTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldTotalMemory, v))
}

// OccupiedCPUEQ applies the EQ predicate on the "occupied_cpu" field.
func OccupiedCPUEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldOccupiedCPU, v))
}

// OccupiedCPUNEQ applies the NEQ predicate on the "occupied_cpu" field.
func OccupiedCPUNEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldOccupiedCPU, v))
}

// OccupiedCPUIn applies the In predicate on the "occupied_cpu" field.
func OccupiedCPUIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldOccupiedCPU, vs...))
}

// OccupiedCPUNotIn applies the NotIn predicate on the "occupied_cpu" field.
func OccupiedCPUNotIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldOccupiedCPU, vs...))
}

// OccupiedCPUGT applies the GT predicate on the "occupied_cpu" field.
func OccupiedCPUGT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldOccupiedCPU, v))
}

// OccupiedCPUGTE applies the GTE predicate on the "occupied_cpu" field.
func OccupiedCPUGTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldOccupiedCPU, v))
}

// OccupiedCPULT applies the LT predicate on the "occupied_cpu" field.
func OccupiedCPULT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldOccupiedCPU, v))
}

// OccupiedCPULTE applies the LTE predicate on the "occupied_cpu" field.
func OccupiedCPULTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldOccupiedCPU, v))
}

// OccupiedMemoryEQ applies the EQ predicate on the "occupied_memory" field.
func OccupiedMemoryEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldOccupiedMemory, v))
}

// OccupiedMemoryNEQ applies the NEQ predicate on the "occupied_memory" field.
func OccupiedMemoryNEQ(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldOccupiedMemory, v))
}

// OccupiedMemoryIn applies the In predicate on the "occupied_memory" field.
func OccupiedMemoryIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldOccupiedMemory, vs...))
}

// OccupiedMemoryNotIn applies the NotIn predicate on the "occupied_memory" field.
func OccupiedMemoryNotIn(vs ...int32) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldOccupiedMemory, vs...))
}

// OccupiedMemoryGT applies the GT predicate on the "occupied_memory" field.
func OccupiedMemoryGT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldOccupiedMemory, v))
}

// OccupiedMemoryGTE applies the GTE predicate on the "occupied_memory" field.
func OccupiedMemoryGTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldOccupiedMemory, v))
}

// OccupiedMemoryLT applies the LT predicate on the "occupied_memory" field.
func OccupiedMemoryLT(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldOccupiedMemory, v))
}

// OccupiedMemoryLTE applies the LTE predicate on the "occupied_memory" field.
func OccupiedMemoryLTE(v int32) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldOccupiedMemory, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldIP, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Agent) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Agent) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
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
func Not(p predicate.Agent) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		p(s.Not())
	})
}
