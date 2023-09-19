package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"
	"time"

	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/script"

	"github.com/go-kratos/kratos/v2/log"
)

type scriptRepo struct {
	data *Data
	log  *log.Helper
}

// NewScriptRepo .
func NewScriptRepo(data *Data, logger log.Logger) biz.ScriptRepo {
	return &scriptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *scriptRepo) Save(ctx context.Context, g *biz.Script) (*biz.Script, error) {
	//先查用户最新的脚本序号
	taskNumber := 1
	first, err := s.data.db.Script.Query().Select(script.FieldTaskNumber).Where(script.UserID(g.UserId)).Order(script.ByTaskNumber(sql.OrderDesc())).First(ctx)
	if first == nil {
		taskNumber = 1
	}
	save, err := s.data.db.Script.Create().SetUserID(g.UserId).SetTaskNumber(int32(taskNumber)).
		SetScriptName(g.ScriptName).SetFileAddress(g.FileAddress).SetScriptContent(g.ScriptContent).
		SetExecuteResult(g.ExecuteResult).SetExecuteState(consts.UnExecuted).SetCreateTime(time.Now()).SetUpdateTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toBiz(save, 0), nil
}

func (s *scriptRepo) Update(ctx context.Context, g *biz.Script) (*biz.Script, error) {
	_, err := s.data.db.Script.Query().Where(script.ID(g.ID)).First(ctx)
	if err != nil {
		return nil, err
	}
	update, err := s.data.db.Script.UpdateOneID(g.ID).SetExecuteState(g.ExecuteState).SetExecuteResult(g.ExecuteResult).
		SetUpdateTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return s.toBiz(update, 0), nil
}

func (s *scriptRepo) FindByID(ctx context.Context, id int32) (*biz.Script, error) {
	first, err := s.data.db.Script.Query().Where(script.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return s.toBiz(first, 0), nil
}

func (s *scriptRepo) PageByUserID(ctx context.Context, userId string, page int32, size int32) ([]*biz.Script, int32, error) {
	count, err := s.data.db.Script.Query().Select(script.FieldID).Where(script.UserID(userId)).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	var offset int32
	if page > 0 {
		offset = (page - 1) * size
	} else {
		offset = page * size
	}
	scripts, err := s.data.db.Script.Query().Where(script.UserID(userId)).Order(script.ByCreateTime(sql.OrderDesc())).Offset(int(offset)).Limit(int(size)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return lo.Map(scripts, s.toBiz), int32(count), nil
}

func (s *scriptRepo) toBiz(p *ent.Script, _ int) *biz.Script {
	if p == nil {
		return nil
	}
	return &biz.Script{
		ID:            p.ID,
		UserId:        p.UserID,
		TaskNumber:    p.TaskNumber,
		ScriptName:    p.ScriptName,
		FileAddress:   p.FileAddress,
		ScriptContent: p.ScriptContent,
		ExecuteState:  p.ExecuteState,
		ExecuteResult: p.ExecuteResult,
		CreateTime:    p.CreateTime,
		UpdateTime:    p.UpdateTime,
	}
}
