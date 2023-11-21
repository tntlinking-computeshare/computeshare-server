package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	clientcomputev1 "github.com/mohaijiang/computeshare-client/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Script is a Script model.
type Script struct {
	ID            int32     `json:"id,omitempty"`
	UserId        string    `json:"userId,omitempty"`
	TaskNumber    int32     `json:"taskNumber,omitempty"`
	ScriptName    string    `json:"scriptName,omitempty"`
	FileAddress   string    `json:"fileAddress,omitempty"`
	ScriptContent string    `json:"scriptContent,omitempty"`
	CreateTime    time.Time `json:"createTime,omitempty"`
	UpdateTime    time.Time `json:"updateTime,omitempty"`
}

// ScriptRepo is a Script repo.
type ScriptRepo interface {
	Save(context.Context, *Script) (*Script, error)
	Update(context.Context, *Script) (*Script, error)
	FindByID(context.Context, int32) (*Script, error)
	PageByUserID(context.Context, string, int32, int32) ([]*Script, int32, error)
}

// ScriptExecutionRecord is a ScriptExecutionRecord model.
type ScriptExecutionRecord struct {
	ID            int32     `json:"id,omitempty"`
	UserID        string    `json:"user_id,omitempty"`
	FkScriptID    int32     `json:"fk_script_id,omitempty"`
	TaskNumber    int32     `json:"taskNumber,omitempty"`
	ScriptContent string    `json:"fk_script_content,omitempty"`
	ScriptName    string    `json:"scriptName,omitempty"`
	FileAddress   string    `json:"fileAddress,omitempty"`
	ExecuteState  int32     `json:"execute_state,omitempty"`
	ExecuteResult string    `json:"execute_result,omitempty"`
	CreateTime    time.Time `json:"create_time,omitempty"`
	UpdateTime    time.Time `json:"update_time,omitempty"`
}

// ScriptExecutionRecordRepo is a ScriptExecutionRecord repo.
type ScriptExecutionRecordRepo interface {
	Save(context.Context, *ScriptExecutionRecord) (*ScriptExecutionRecord, error)
	Update(context.Context, *ScriptExecutionRecord) (*ScriptExecutionRecord, error)
	FindByID(context.Context, int32) (*ScriptExecutionRecord, error)
	PageByUserId(context.Context, string, int32, int32) ([]*ScriptExecutionRecord, int32, error)
	FindLatestByUserIdAndScript(context.Context, string, int32) (*ScriptExecutionRecord, error)
}

// ScriptUseCase is a Script UseCase.
type ScriptUseCase struct {
	repo                      ScriptRepo
	scriptExecutionRecordRepo ScriptExecutionRecordRepo
	agentRepo                 AgentRepo
	p2pClient                 *P2pClient
	log                       *log.Helper
}

// NewScriptUseCase new a Script UseCase.
func NewScriptUseCase(repo ScriptRepo, scriptExecutionRecordRepo ScriptExecutionRecordRepo, agentRepo AgentRepo, logger log.Logger) *ScriptUseCase {
	return &ScriptUseCase{repo: repo, scriptExecutionRecordRepo: scriptExecutionRecordRepo, agentRepo: agentRepo, log: log.NewHelper(logger)}
}

// CreateScript creates a Script, and returns the new Script.
func (uc *ScriptUseCase) CreateScript(ctx context.Context, s *Script) (*Script, error) {
	uc.log.WithContext(ctx).Infof("CreateScript: %v", s.ScriptContent)
	return uc.repo.Save(ctx, s)
}

func (uc *ScriptUseCase) GetScriptExecutionRecordPage(ctx context.Context, userId string, page, size int32) ([]*ScriptExecutionRecord, int32, error) {
	uc.log.WithContext(ctx).Infof("GetScriptExecutionRecordPage %s %d %d", userId, page, size)
	return uc.scriptExecutionRecordRepo.PageByUserId(ctx, userId, page, size)
}

func (uc *ScriptUseCase) GetScriptExecutionRecordInfo(ctx context.Context, id int32) (*ScriptExecutionRecord, error) {
	uc.log.WithContext(ctx).Infof("GetScriptExecutionRecordInfo is %d", id)
	return uc.scriptExecutionRecordRepo.FindByID(ctx, id)
}

func (uc *ScriptUseCase) RunPythonPackage(ctx context.Context, id int32, userId string) (*ScriptExecutionRecord, error) {
	script, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	record := ScriptExecutionRecord{
		UserID:        userId,
		FkScriptID:    script.ID,
		ScriptContent: script.ScriptContent,
		FileAddress:   script.FileAddress,
		ScriptName:    script.ScriptName,
		ExecuteState:  consts.Executing,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	save, err := uc.scriptExecutionRecordRepo.Save(ctx, &record)
	if err != nil {
		return nil, err
	}
	// 选择一个agent节点进行通信
	agent, err := uc.agentRepo.FindOneActiveAgent(ctx, "", "")
	if err != nil {
		return nil, err
	}
	go uc.RunPythonPackageOnAgent(agent.PeerId, save)

	return save, nil
}

func (uc *ScriptUseCase) RunPythonPackageOnAgent(peerId string, record *ScriptExecutionRecord) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*20)

	computePowerClient, cleanup, err := uc.getComputePowerHTTPClient(peerId)
	if err != nil {
		uc.log.Error("创建ComputePowerHTTPClient链接失败")
		uc.log.Error(err)
		return
	}
	defer cleanup()

	rsp, runPythonPackageErr := computePowerClient.RunPythonPackage(ctx, &clientcomputev1.RunPythonPackageClientRequest{Cid: record.FileAddress})
	scriptExecutionRecord, err := uc.scriptExecutionRecordRepo.FindByID(ctx, record.ID)
	if err != nil {
		uc.log.Error("computePowerClient RunPythonPackage FindByID fail")
		uc.log.Error(err)
		return
	}
	if scriptExecutionRecord.ExecuteState == consts.Executing {
		executeState := consts.Completed
		if runPythonPackageErr != nil {
			uc.log.Error("computePowerClient RunPythonPackage fail")
			uc.log.Error(err)
			executeState = consts.ExecutionFailed
			record.ExecuteResult = runPythonPackageErr.Error()
		} else {
			if rsp == nil {
				record.ExecuteResult = ""
			} else {
				record.ExecuteResult = rsp.ExecuteResult
			}
		}
		record.ExecuteState = int32(executeState)
		_, err = uc.scriptExecutionRecordRepo.Update(ctx, record)
		if err != nil {
			uc.log.Error("客户端执行py完成，向db保存scriptExecutionRecord失败")
			uc.log.Error(err)
			return
		}
	} else if scriptExecutionRecord.ExecuteState == consts.Canceled {
		uc.log.Info("本次执行任务已经取消，不能写入执行结果")
		return
	} else {
		uc.log.Info("本次执行任务状态不符合写入执行结果的条件")
		return
	}

}

func (uc *ScriptUseCase) getComputePowerHTTPClient(peerId string) (clientcomputev1.ComputePowerClientHTTPClient, func(), error) {
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
		uc.log.Error("创建ComputePowerClient链接失败")
		uc.log.Error(err)
		return nil, nil, err
	}

	vmClient := clientcomputev1.NewComputePowerClientHTTPClient(client)
	return vmClient, func() {
		_ = client.Close()
	}, nil
}

func (uc *ScriptUseCase) CancelExecPythonPackage(ctx context.Context, scriptId int32) (*ScriptExecutionRecord, error) {
	executionRecord, err := uc.scriptExecutionRecordRepo.FindByID(ctx, scriptId)
	if err != nil {
		return nil, err
	}
	executionRecord.ExecuteState = consts.Canceled
	return uc.scriptExecutionRecordRepo.Update(ctx, executionRecord)
}
