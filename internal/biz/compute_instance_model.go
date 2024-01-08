package biz

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"strconv"
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

func (c *ComputeSpec) GetCore() int64 {
	core, err := strconv.Atoi(c.Core)
	if err != nil {
		return 1
	}
	return int64(core)
}

func (c *ComputeSpec) GetMemory() int64 {
	memory, err := strconv.Atoi(c.Memory)
	if err != nil {
		return 1024
	}
	return int64(memory)
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
	// 镜像id
	ImageId int32  `json:"omitempty"`
	Port    string `json:"port,omitempty"`
	// ExpirationTime holds the value of the "expiration_time" field.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// 0: 启动中,1:运行中,2:连接中断, 3:过期
	Status consts.InstanceStatus `json:"status,omitempty"`
	// 容器id
	ContainerID string `json:"container_id,omitempty"`
	// p2p agent Id
	AgentId string `json:"agent_id,omitempty"`
	// vnc 内网链接ip
	VncIP string `json:"vnc_ip,omitempty"`
	// vnc 内网链接端口号
	VncPort       int32                 `json:"vnc_port,omitempty"`
	DockerCompose string                `json:"docker_compose"`
	Stats         []*ComputeInstanceRds `json:"stats"`
}

func (i *ComputeInstance) GetCore() int64 {
	core, err := strconv.Atoi(i.Core)
	if err != nil {
		return 0
	}
	return int64(core)
}

func (i *ComputeInstance) GetMemory() int64 {
	memory, err := strconv.Atoi(i.Memory)
	if err != nil {
		return 0
	}
	return int64(memory)
}

type ComputeInstanceCreate struct {
	SpecId        int32
	ImageId       int32
	Duration      int32
	Name          string
	PublicKey     string
	Password      string
	DockerCompose string
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

func (c *ComputeImage) GetImageTag() string {
	return fmt.Sprintf("%s:%s", c.Image, c.Tag)
}

type ComputeInstanceRds struct {
	ID          string    `json:"id"`
	CpuUsage    float32   `json:"cpuUsage"`
	MemoryUsage float32   `json:"memoryUsage"`
	StatsTime   time.Time `json:"statsTime"`
}

func (m *ComputeInstanceRds) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *ComputeInstanceRds) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type InstanceCreateParam struct {
	PublicKey      string
	Password       string
	GatewayIP      string
	GatewayPort    int32
	VncConnectIP   string
	VncConnectPort int32
	DockerCompose  string
}

type Metric struct {
	Instance string `json:"instance"`
}

type Result struct {
	Metric Metric          `json:"metric"`
	Values [][]interface{} `json:"values"`
}

type PrometheusQueryResult struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}
