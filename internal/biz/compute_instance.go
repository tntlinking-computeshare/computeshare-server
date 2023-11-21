package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	queue "github.com/mohaijiang/computeshare-server/api/queue/v1"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"net/http"
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
	UpdateStatus(ctx context.Context, id uuid.UUID, status consts.InstanceStatus) error
	SetInstanceExpiration(ctx context.Context) error
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
	taskRepo     TaskRepo
	p2pClient    *P2pClient
	log          *log.Helper
}

func NewComputeInstanceUsercase(
	specRepo ComputeSpecRepo,
	instanceRepo ComputeInstanceRepo,
	imageRepo ComputeImageRepo,
	agentRepo AgentRepo,
	taskRepo TaskRepo,
	logger log.Logger) *ComputeInstanceUsercase {
	return &ComputeInstanceUsercase{
		specRepo:     specRepo,
		instanceRepo: instanceRepo,
		imageRepo:    imageRepo,
		agentRepo:    agentRepo,
		taskRepo:     taskRepo,
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
		AgentId:        agent.ID.String(),
		Status:         consts.InstanceStatusCreating,
	}

	err = uc.instanceRepo.Create(ctx, instance)
	if err != nil {
		return nil, err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_CREATE, func() (string, string) {
		return cic.PublicKey, cic.Password
	})
	if err != nil {
		return nil, err
	}

	return instance, err
}

func (uc *ComputeInstanceUsercase) SendTaskQueue(ctx context.Context, instance *ComputeInstance, cmd queue.TaskCmd, publicKeyAndPassword func() (string, string)) error {

	publicKey := ""
	password := ""

	if publicKeyAndPassword != nil {
		publicKey, password = publicKeyAndPassword()
	}

	taskParam := queue.ComputeInstanceTaskParamVO{
		Id:         uuid.NewString(),
		Name:       instance.Name,
		Cpu:        instance.GetCore(),
		Memory:     instance.GetMemory(),
		Image:      instance.Image,
		PublicKey:  publicKey,
		Password:   password,
		InstanceId: instance.ID.String(),
	}
	paramData, err := json.Marshal(taskParam)
	if err != nil {
		return err
	}
	param := string(paramData)
	task := Task{
		AgentID:    instance.AgentId,
		Cmd:        cmd,
		Params:     &param,
		Status:     queue.TaskStatus_CREATED,
		CreateTime: time.Now(),
	}
	err = uc.taskRepo.CreateTask(ctx, &task)
	return err
}

func (uc *ComputeInstanceUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_DELETE, nil)
	if err != nil {
		return err
	}

	return uc.instanceRepo.UpdateStatus(ctx, instance.ID, consts.InstanceStatusDeleting)
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

	instance.Status = consts.InstanceStatusRunning

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_START, nil)
	if err != nil {
		return err
	}

	return nil

}

func (uc *ComputeInstanceUsercase) Stop(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	instance.Status = consts.InstanceStatusClosing

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_SHUTDOWN, nil)
	if err != nil {
		return err
	}

	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Terminal Deprecate
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

	if instance.AgentId == "" {
		return
	}

	ip, port, err := uc.p2pClient.ForwardWithRandomPort(instance.AgentId)
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

// SyncContainerOverdue 同步资源实例的过期状态
func (uc *ComputeInstanceUsercase) SyncContainerOverdue() {
	ctx := context.Background()
	_ = uc.instanceRepo.SetInstanceExpiration(ctx)
}
