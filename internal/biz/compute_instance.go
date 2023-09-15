package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/ipfs/kubo/core"
	clientcomputev1 "github.com/mohaijiang/computeshare-client/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/third_party/p2p"
	ma "github.com/multiformats/go-multiaddr"
	"golang.org/x/exp/rand"
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
	PeerID string `json:"peer_id,omitempty"`
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
	Port int32 `json:"port,omitempty"`
}

type ComputeSpecRepo interface {
	List(ctx context.Context) ([]*ComputeSpec, error)
	Get(ctx context.Context, id int32) (*ComputeSpec, error)
}

type ComputeInstanceRepo interface {
	List(ctx context.Context, owner string) ([]*ComputeInstance, error)
	Create(ctx context.Context, instance *ComputeInstance) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, instance *ComputeInstance) error
}

type ComputeImageRepo interface {
	List(ctx context.Context) ([]*ComputeImage, error)
	Get(ctx context.Context, id int32) (*ComputeImage, error)
}

type ComputeInstanceUsercase struct {
	specRepo     ComputeSpecRepo
	instanceRepo ComputeInstanceRepo
	imageRepo    ComputeImageRepo
	agentRepo    AgentRepo
	ipfsNode     *core.IpfsNode
	log          *log.Helper
}

func NewComputeInstanceUsercase(
	specRepo ComputeSpecRepo,
	instanceRepo ComputeInstanceRepo,
	imageRepo ComputeImageRepo,
	ipfsNode *core.IpfsNode,
	agentRepo AgentRepo,
	logger log.Logger) *ComputeInstanceUsercase {
	return &ComputeInstanceUsercase{
		specRepo:     specRepo,
		instanceRepo: instanceRepo,
		imageRepo:    imageRepo,
		ipfsNode:     ipfsNode,
		agentRepo:    agentRepo,
		log:          log.NewHelper(logger),
	}
}

func (uc *ComputeInstanceUsercase) ListComputeSpec(ctx context.Context) ([]*ComputeSpec, error) {
	return uc.specRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) ListComputeImage(ctx context.Context) ([]*ComputeImage, error) {
	return uc.imageRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) Create(ctx context.Context, cic *ComputeInstanceCreate) (*ComputeInstance, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("cannot get user ID")
	}

	computeSpec, err := uc.specRepo.Get(ctx, cic.SpecId)
	if err != nil {
		return nil, err
	}
	computeImage, err := uc.imageRepo.Get(ctx, cic.ImageId)
	if err != nil {
		return nil, err
	}

	instance := &ComputeInstance{
		Owner:          claim.UserID,
		Name:           cic.Name,
		Core:           computeSpec.Core,
		Memory:         computeSpec.Memory,
		Port:           string(computeImage.Port),
		Image:          fmt.Sprintf("%s:%s", computeImage.Image, computeImage.Tag),
		ExpirationTime: time.Now().AddDate(0, int(cic.Duration), 0),
		Status:         InstanceStatusStarting,
	}

	err = uc.instanceRepo.Create(ctx, instance)

	// 选择一个agent节点进行通信

	agent, err := uc.agentRepo.FindOneActiveAgent(ctx, instance.Core, instance.Memory)
	if err != nil {
		return nil, err
	}
	go uc.CreateInstanceOnAgent(agent.PeerId, instance)

	return instance, err
}

func (uc *ComputeInstanceUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.instanceRepo.Delete(ctx, id)
}

func (uc *ComputeInstanceUsercase) CreateInstanceOnAgent(peerId string, instance *ComputeInstance) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*20)
	p2pService := p2p.NewP2pService(uc.ipfsNode)
	pingOk := p2pService.Ping(ctx, peerId)

	fmt.Println("pingOk: ", pingOk)
	if !pingOk {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Errorf("无法与%s完成ping", peerId)
		return
	}

	listenIp := "127.0.0.1"
	listenPort := rand.Intn(9999) + 30000

	listenOpt := fmt.Sprintf("/ip4/%s/tcp/%d", listenIp, listenPort)
	listen, err := ma.NewMultiaddr(listenOpt)
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}
	targetOpt := fmt.Sprintf("/p2p/%s", peerId)
	proto := "/x/ssh"

	err = p2pService.CheckPort(listen)
	if err != nil {
		_ = p2pService.CloseListen(ctx, proto, listenOpt, targetOpt)
	}
	err = p2pService.CreateForward(ctx, proto, listenOpt, targetOpt)
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}

	defer p2pService.CloseListen(ctx, proto, listenOpt, targetOpt)

	client, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(fmt.Sprintf("%s:%d", listenIp, listenPort)),
	)

	defer client.Close()

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}

	vmClient := clientcomputev1.NewVmHTTPClient(client)

	reply, err := vmClient.CreateVm(ctx, &clientcomputev1.CreateVmRequest{
		Image: instance.Image,
		Port:  instance.Port,
	})
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}
	fmt.Println("containerId:", reply.GetId())

	instance.Status = InstanceStatusRunning
	instance.PeerID = peerId
	instance.ContainerID = reply.GetId()

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}
}
