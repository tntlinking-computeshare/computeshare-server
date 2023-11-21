package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/script"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/scriptexecutionrecord"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"
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
	//先查用户最新的脚本序号
	taskNumber := 1
	first, err := s.data.db.ScriptExecutionRecord.Query().Select(script.FieldTaskNumber).Where(scriptexecutionrecord.UserID(scriptExecutionRecord.UserID)).Order(scriptexecutionrecord.ByTaskNumber(sql.OrderDesc())).First(ctx)
	if first == nil {
		taskNumber = 1
	} else {
		taskNumber = int(first.TaskNumber + 1)
	}
	save, err := s.data.db.ScriptExecutionRecord.Create().SetFkScriptID(scriptExecutionRecord.FkScriptID).SetScriptContent(scriptExecutionRecord.ScriptContent).
		SetUserID(scriptExecutionRecord.UserID).SetFileAddress(scriptExecutionRecord.FileAddress).SetExecuteState(consts.Executing).SetTaskNumber(int32(taskNumber)).
		SetScriptName(scriptExecutionRecord.ScriptName).SetExecuteResult(scriptExecutionRecord.ExecuteResult).SetCreateTime(time.Now()).SetUpdateTime(time.Now()).Save(ctx)
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

func (s *scriptExecutionRecordRepo) PageByUserId(ctx context.Context, userId string, page int32, size int32) ([]*biz.ScriptExecutionRecord, int32, error) {
	count, err := s.data.db.ScriptExecutionRecord.Query().Select(scriptexecutionrecord.FieldID).Where(scriptexecutionrecord.UserID(userId)).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	var offset int32
	if page > 0 {
		offset = (page - 1) * size
	} else {
		offset = page * size
	}
	scripts, err := s.data.db.ScriptExecutionRecord.Query().Where(scriptexecutionrecord.UserID(userId)).Order(scriptexecutionrecord.ByCreateTime(sql.OrderDesc())).Offset(int(offset)).Limit(int(size)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return lo.Map(scripts, s.toBiz), int32(count), nil
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
		TaskNumber:    p.TaskNumber,
		FileAddress:   p.FileAddress,
		ScriptName:    p.ScriptName,
		ExecuteState:  p.ExecuteState,
		ExecuteResult: p.ExecuteResult,
		CreateTime:    p.CreateTime,
		UpdateTime:    p.UpdateTime,
	}
}
