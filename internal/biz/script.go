package biz

import (
	"context"
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
	ExecuteState  int32     `json:"executeState,omitempty"`
	ExecuteResult string    `json:"executeResult,omitempty"`
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
	ScriptContent string    `json:"fk_script_content,omitempty"`
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
	FindLatestByUserIdAndScript(context.Context, string, int32) (*ScriptExecutionRecord, error)
}

// ScriptUseCase is a Script UseCase.
type ScriptUseCase struct {
	repo                      ScriptRepo
	scriptExecutionRecordRepo ScriptExecutionRecordRepo
	log                       *log.Helper
}

// NewScriptUseCase new a Script UseCase.
func NewScriptUseCase(repo ScriptRepo, scriptExecutionRecordRepo ScriptExecutionRecordRepo, logger log.Logger) *ScriptUseCase {
	return &ScriptUseCase{repo: repo, scriptExecutionRecordRepo: scriptExecutionRecordRepo, log: log.NewHelper(logger)}
}

// CreateScript creates a Script, and returns the new Script.
func (uc *ScriptUseCase) CreateScript(ctx context.Context, s *Script) (*Script, error) {
	uc.log.WithContext(ctx).Infof("CreateScript: %v", s.ScriptContent)
	return uc.repo.Save(ctx, s)
}

func (uc *ScriptUseCase) GetScriptPage(ctx context.Context, userId string, page, size int32) ([]*Script, int32, error) {
	uc.log.WithContext(ctx).Infof("GetScriptPage %s %d %d", userId, page, size)
	return uc.repo.PageByUserID(ctx, userId, page, size)
}

func (uc *ScriptUseCase) GetScriptInfo(ctx context.Context, id int32) (*Script, error) {
	uc.log.WithContext(ctx).Infof("GetScriptInfo is %d", id)
	return uc.repo.FindByID(ctx, id)
}

func (uc *ScriptUseCase) RunPythonPackage(ctx context.Context, id int32) (*Script, error) {
	script, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	record := ScriptExecutionRecord{
		UserID:        script.UserId,
		FkScriptID:    script.ID,
		ScriptContent: script.ScriptContent,
		ExecuteState:  consts.Executing,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	script.ExecuteState = consts.Executing
	_, err = uc.scriptExecutionRecordRepo.Save(ctx, &record)
	if err != nil {
		return nil, err
	}
	uc.log.WithContext(ctx).Infof("GetScriptInfo is %d", id)
	return uc.repo.Update(ctx, script)
}

func (uc *ScriptUseCase) CancelExecPythonPackage(ctx context.Context, id int32) (*Script, error) {
	script, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	executionRecord, err := uc.scriptExecutionRecordRepo.FindLatestByUserIdAndScript(ctx, script.UserId, script.ID)
	if err != nil {
		return nil, err
	}
	executionRecord.ExecuteState = consts.Canceled
	script.ExecuteState = consts.Canceled
	uc.log.WithContext(ctx).Infof("GetScriptInfo is %d", id)
	_, err = uc.scriptExecutionRecordRepo.Update(ctx, executionRecord)
	if err != nil {
		return nil, err
	}
	return uc.repo.Update(ctx, script)
}
