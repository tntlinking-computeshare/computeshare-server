package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	clientcomputev1 "github.com/mohaijiang/computeshare-client/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/global"
	goipfsp2p "github.com/mohaijiang/go-ipfs-p2p"
	"net/http"
	"strings"
	"time"
)

type ComputeSpecRepo interface {
	List(ctx context.Context) ([]*ComputeSpec, error)
	Get(ctx context.Context, id int32) (*ComputeSpec, error)
}

type ComputeInstanceRepo interface {
	List(ctx context.Context, owner string) ([]*ComputeInstance, error)
	ListByPeerId(ctx context.Context, peerId string) ([]*ComputeInstance, error)
	ListAll(ctx context.Context) ([]*ComputeInstance, error)
	Create(ctx context.Context, instance *ComputeInstance) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, instance *ComputeInstance) error
	Get(ctx context.Context, id uuid.UUID) (*ComputeInstance, error)
	SaveInstanceStats(ctx context.Context, id uuid.UUID, rdbInstance *ComputeInstanceRds) error
	GetInstanceStats(ctx context.Context, id uuid.UUID) ([]*ComputeInstanceRds, error)
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
	p2pClient    *goipfsp2p.P2pClient
	log          *log.Helper
}

func NewComputeInstanceUsercase(
	specRepo ComputeSpecRepo,
	instanceRepo ComputeInstanceRepo,
	imageRepo ComputeImageRepo,
	agentRepo AgentRepo,
	p2pClient *goipfsp2p.P2pClient,
	logger log.Logger) *ComputeInstanceUsercase {
	return &ComputeInstanceUsercase{
		specRepo:     specRepo,
		instanceRepo: instanceRepo,
		imageRepo:    imageRepo,
		agentRepo:    agentRepo,
		p2pClient:    p2pClient,
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

	// 选择一个agent节点进行通信
	agent, err := uc.agentRepo.FindOneActiveAgent(ctx, computeSpec.Core, computeSpec.Memory)
	if err != nil {
		return nil, err
	}

	instance := &ComputeInstance{
		Owner:          claim.UserID,
		Name:           cic.Name,
		Core:           computeSpec.Core,
		Memory:         computeSpec.Memory,
		Port:           fmt.Sprintf("%d", computeImage.Port),
		Image:          fmt.Sprintf("%s:%s", computeImage.Image, computeImage.Tag),
		Command:        computeImage.Command,
		ExpirationTime: time.Now().AddDate(0, int(cic.Duration), 0),
		PeerID:         agent.PeerId,
		Status:         InstanceStatusStarting,
	}

	err = uc.instanceRepo.Create(ctx, instance)

	go uc.CreateInstanceOnAgent(agent.PeerId, instance)

	return instance, err
}

func (uc *ComputeInstanceUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	go func() {
		instance, err := uc.Get(ctx, id)
		if err != nil {
			return
		}

		if instance.ContainerID == "" || instance.PeerID == "" {
			return
		}

		vmClient, cleanup, err := uc.getVmClient(instance.PeerID)
		if err != nil {
			return
		}
		defer cleanup()

		_, err = vmClient.DeleteVm(ctx, &clientcomputev1.DeleteVmRequest{
			Id: instance.ContainerID,
		})

		if err != nil {
			return
		}
	}()
	return uc.instanceRepo.Delete(ctx, id)
}

func (uc *ComputeInstanceUsercase) CreateInstanceOnAgent(peerId string, instance *ComputeInstance) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*20)

	vmClient, cleanup, err := uc.getVmClient(peerId)
	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return
	}
	defer cleanup()

	reply, err := vmClient.CreateVm(ctx, &clientcomputev1.CreateVmRequest{
		Image:   instance.Image,
		Port:    instance.Port,
		Command: strings.Fields(instance.Command),
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

func (uc *ComputeInstanceUsercase) getVmClient(peerId string) (clientcomputev1.VmHTTPClient, func(), error) {
	ip, port, err := uc.p2pClient.ForwardWithRandomPort(peerId)
	if err != nil {
		return nil, nil, err
	}

	time.Sleep(time.Second * 2)

	client, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(fmt.Sprintf("%s:%s", ip, port)),
		transhttp.WithTimeout(time.Second*10),
	)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return nil, nil, err
	}

	vmClient := clientcomputev1.NewVmHTTPClient(client)
	return vmClient, func() {
		_ = client.Close()
	}, nil
}

func (uc *ComputeInstanceUsercase) ListComputeInstance(ctx context.Context, owner string) ([]*ComputeInstance, error) {
	list, err := uc.instanceRepo.List(ctx, owner)
	for _, ins := range list {
		ins.Stats, _ = uc.instanceRepo.GetInstanceStats(ctx, ins.ID)
	}
	return list, err
}

func (uc *ComputeInstanceUsercase) Get(ctx context.Context, id uuid.UUID) (*ComputeInstance, error) {
	return uc.instanceRepo.Get(ctx, id)
}

func (uc *ComputeInstanceUsercase) Start(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	if instance.ContainerID == "" || instance.PeerID == "" {
		return fmt.Errorf("instance is not avaliable")
	}

	vmClient, cleanup, err := uc.getVmClient(instance.PeerID)
	if err != nil {
		return err
	}
	defer cleanup()

	_, err = vmClient.StartVm(ctx, &clientcomputev1.GetVmRequest{
		Id: instance.ContainerID,
	})

	if err != nil {
		return err
	}

	instance.Status = InstanceStatusRunning

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return err
	}

	return nil

}

func (uc *ComputeInstanceUsercase) Stop(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	if instance.ContainerID == "" || instance.PeerID == "" {
		return fmt.Errorf("instance is not avaliable")
	}

	vmClient, cleanup, err := uc.getVmClient(instance.PeerID)
	if err != nil {
		return err
	}
	defer cleanup()

	_, err = vmClient.StopVm(ctx, &clientcomputev1.GetVmRequest{
		Id: instance.ContainerID,
	})

	if err != nil {
		return err
	}

	instance.Status = InstanceStatusTerminal

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return err
	}

	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (uc *ComputeInstanceUsercase) Terminal(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	instanceId, err := uuid.Parse(r.Form.Get("instanceId"))
	if err != nil {
		return
	}
	instance, err := uc.Get(context.Background(), instanceId)
	if err != nil {
		return
	}

	if instance.PeerID == "" {
		return
	}

	ip, port, err := uc.p2pClient.ForwardWithRandomPort(instance.PeerID)
	if err != nil {
		return
	}

	// websocket握手
	// 建立与目标WebSocket服务器的连接
	targetConn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/v1/vm/%s/terminal?container=%s&workdir=/bin", ip, port, instanceId, instance.ContainerID), nil)
	if err != nil {
		uc.log.Error(err)
		return
	}
	defer targetConn.Close()

	// 升级客户端连接为WebSocket
	clientConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		uc.log.Error(err)
		return
	}
	defer clientConn.Close()

	// 开始在两个WebSocket连接之间传递消息
	go func() {
		for {
			msgType, msg, err := clientConn.ReadMessage()
			if err != nil {
				return
			}
			if err := targetConn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	}()

	for {
		msgType, msg, err := targetConn.ReadMessage()
		if err != nil {
			return
		}
		if err := clientConn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func (uc *ComputeInstanceUsercase) SyncContainerStats() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	list, err := uc.instanceRepo.ListAll(ctx)
	if err != nil {
		return
	}

	for _, instance := range list {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		uc.syncInstanceStats(ctx, instance)
	}
}

func (uc *ComputeInstanceUsercase) syncInstanceStats(ctx context.Context, instance *ComputeInstance) {
	if instance.ContainerID == "" || instance.PeerID == "" {
		return
	}

	vmClient, cleanup, err := uc.getVmClient(instance.PeerID)
	if err != nil {
		return
	}
	defer cleanup()

	vm, err := vmClient.GetVm(ctx, &clientcomputev1.GetVmRequest{
		Id: instance.ContainerID,
	})
	if err != nil {
		return
	}

	instanceRdb := &ComputeInstanceRds{
		ID:          instance.ID.String(),
		CpuUsage:    vm.CpuUsage,
		MemoryUsage: vm.MemoryUsage,
		StatsTime:   time.Now(),
	}

	_ = uc.instanceRepo.SaveInstanceStats(ctx, instance.ID, instanceRdb)

}
