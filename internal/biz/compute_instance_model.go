package biz

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type ComputeSpec struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Core holds the value of the "core" field.
	Core string `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
}

type ComputeInstance struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Core holds the value of the "core" field.
	Core string `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	Port  string `json:"port,omitempty"`
	// ExpirationTime holds the value of the "expiration_time" field.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// 0: 启动中,1:运行中,2:连接中断, 3:过期
	Status int8 `json:"status,omitempty"`
	// 容器id
	ContainerID string `json:"container_id,omitempty"`
	// p2p agent Id
	PeerID  string                `json:"peer_id,omitempty"`
	Command string                `json:"command,omitempty"`
	Stats   []*ComputeInstanceRds `json:"stats"`
}

const (
	InstanceStatusStarting int8 = iota
	InstanceStatusRunning
	InstanceStatusTerminal
	InstanceStatusExpire
)

type ComputeInstanceCreate struct {
	SpecId   int32
	ImageId  int32
	Duration int32
	Name     string
}

type ComputeImage struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// 显示名
	Name string `json:"name,omitempty"`
	// 镜像名
	Image string `json:"image,omitempty"`
	// 版本名
	Tag string `json:"tag,omitempty"`
	// 端口号
	Port    int32 `json:"port,omitempty"`
	Command string
}

type ComputeInstanceRds struct {
	ID          string    `json:"id"`
	CpuUsage    uint64    `json:"cpuUsage"`
	MemoryUsage uint64    `json:"memoryUsage"`
	StatsTime   time.Time `json:"statsTime"`
}

func (m *ComputeInstanceRds) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *ComputeInstanceRds) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)

}
