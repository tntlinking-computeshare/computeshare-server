package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/scriptexecutionrecord"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type scriptExecutionRecordRepo struct {
	data *Data
	log  *log.Helper
}

// NewScriptExecutionRecordRepo .
func NewScriptExecutionRecordRepo(data *Data, logger log.Logger) biz.ScriptExecutionRecordRepo {
	return &scriptExecutionRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *scriptExecutionRecordRepo) Save(ctx context.Context, scriptExecutionRecord *biz.ScriptExecutionRecord) (*biz.ScriptExecutionRecord, error) {
	save, err := s.data.db.ScriptExecutionRecord.Create().SetFkScriptID(scriptExecutionRecord.FkScriptID).SetScriptContent(scriptExecutionRecord.ScriptContent).
		SetExecuteState(consts.Executing).SetCreateTime(time.Now()).SetUpdateTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toBiz(save, 0), nil
}

func (s *scriptExecutionRecordRepo) Update(ctx context.Context, scriptExecutionRecord *biz.ScriptExecutionRecord) (*biz.ScriptExecutionRecord, error) {
	save, err := s.data.db.ScriptExecutionRecord.UpdateOneID(scriptExecutionRecord.ID).SetExecuteResult(scriptExecutionRecord.ExecuteResult).
		SetExecuteState(scriptExecutionRecord.ExecuteState).SetUpdateTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toBiz(save, 0), nil
}

func (s *scriptExecutionRecordRepo) FindByID(ctx context.Context, id int32) (*biz.ScriptExecutionRecord, error) {
	first, err := s.data.db.ScriptExecutionRecord.Query().Where(scriptexecutionrecord.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toBiz(first, 0), nil
}

func (s *scriptExecutionRecordRepo) FindLatestByUserIdAndScript(ctx context.Context, userId string, scriptId int32) (*biz.ScriptExecutionRecord, error) {
	first, err := s.data.db.ScriptExecutionRecord.Query().Where(scriptexecutionrecord.UserIDEQ(userId), scriptexecutionrecord.FkScriptIDEQ(scriptId)).
		Order(scriptexecutionrecord.ByCreateTime(sql.OrderDesc())).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toBiz(first, 0), nil
}

func (s *scriptExecutionRecordRepo) toBiz(p *ent.ScriptExecutionRecord, _ int) *biz.ScriptExecutionRecord {
	if p == nil {
		return nil
	}
	return &biz.ScriptExecutionRecord{
		ID:            p.ID,
		UserID:        p.UserID,
		FkScriptID:    p.FkScriptID,
		ScriptContent: p.ScriptContent,
		ExecuteState:  p.ExecuteState,
		ExecuteResult: p.ExecuteResult,
		CreateTime:    p.CreateTime,
		UpdateTime:    p.UpdateTime,
	}
}
