package biz

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/api/compute"
	"time"
)

type ComputeSpec struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Core holds the value of the "core" field.
	Core int `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory int `json:"memory,omitempty"`
}

type ComputeSpecPrice struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// 资源规格id
	FkComputeSpecID int32 `json:"fkComputeSpecId,omitempty"`
	// 天数
	Day int32 `json:"day,omitempty"`
	// 此天数的价格
	Price float32 `json:"price,omitempty"`
}

func (c *ComputeSpec) GetCore() int64 {
	return int64(c.Core)
}

func (c *ComputeSpec) GetMemory() int64 {
	return int64(c.Memory)
}

type ComputeInstance struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Core holds the value of the "core" field.
	Core int `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory int `json:"memory,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// 镜像id
	ImageId int32  `json:"omitempty"`
	Port    string `json:"port,omitempty"`
	// ExpirationTime holds the value of the "expiration_time" field.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// 0: 启动中,1:运行中,2:连接中断, 3:过期
	Status compute.InstanceStatus `json:"status,omitempty"`
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

	CreateTime time.Time `json:"create_time"`
}

func (i *ComputeInstance) GetCore() int64 {
	return int64(i.Core)
}

func (i *ComputeInstance) GetMemory() int64 {
	return int64(i.Memory)
}

type ComputeInstanceCreate struct {
	SpecId        int32
	ImageId       int32
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
	Status string              `json:"status"`
	Data   PrometheusQueryData `json:"data"`
}

type PrometheusQueryData struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}
