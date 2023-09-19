package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
)

type ComputePowerService struct {
	pb.UnimplementedComputePowerServer

	uc  *biz.ScriptUseCase
	log *log.Helper
}

func NewComputePowerService(uc *biz.ScriptUseCase, logger log.Logger) *ComputePowerService {
	return &ComputePowerService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *ComputePowerService) UploadScriptFile(ctx context.Context, req *pb.UploadScriptFileRequest) (*pb.UploadScriptFileReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}
	script := biz.Script{
		UserId:        token.UserID,
		ScriptName:    req.Name,
		ScriptContent: string(req.Body),
	}
	createScript, err := s.uc.CreateScript(ctx, &script)
	if err != nil {
		return nil, err
	}
	return &pb.UploadScriptFileReply{
		Id:            createScript.ID,
		TaskNumber:    createScript.TaskNumber,
		ScriptName:    createScript.ScriptName,
		ScriptContent: createScript.ScriptContent,
		ExecuteState:  createScript.ExecuteState,
	}, nil
}
func (s *ComputePowerService) GetScriptList(ctx context.Context, req *pb.GetScriptListRequest) (*pb.GetScriptListReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}
	data, total, err := s.uc.GetScriptPage(ctx, token.UserID, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	var pointerList []*pb.UploadScriptFileReply
	for _, script := range data {
		var uploadScriptFileReply pb.UploadScriptFileReply
		copier.Copy(uploadScriptFileReply, *script)
		pointerList = append(pointerList, &uploadScriptFileReply)
	}
	return &pb.GetScriptListReply{List: pointerList, Total: total, Page: req.GetPage(), Size: req.GetSize()}, nil
}
func (s *ComputePowerService) RunPythonPackage(ctx context.Context, req *pb.RunPythonPackageServerRequest) (*pb.RunPythonPackageServerReply, error) {
	script, err := s.uc.RunPythonPackage(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.RunPythonPackageServerReply{
		Id:            script.ID,
		TaskNumber:    script.TaskNumber,
		ScriptName:    script.ScriptName,
		ScriptContent: script.ScriptContent,
		ExecuteState:  script.ExecuteState,
		ExecuteResult: script.ExecuteResult,
	}, nil
}

func (s *ComputePowerService) CancelExecPythonPackage(ctx context.Context, req *pb.CancelExecPythonPackageRequest) (*pb.CancelExecPythonPackageReply, error) {
	script, err := s.uc.CancelExecPythonPackage(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.CancelExecPythonPackageReply{
		Id:            script.ID,
		TaskNumber:    script.TaskNumber,
		ScriptName:    script.ScriptName,
		ScriptContent: script.ScriptContent,
		ExecuteState:  script.ExecuteState,
		ExecuteResult: script.ExecuteResult,
	}, nil
}
func (s *ComputePowerService) GetScriptInfo(ctx context.Context, req *pb.GetScriptInfoRequest) (*pb.GetScriptInfoReply, error) {
	script, err := s.uc.GetScriptInfo(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetScriptInfoReply{
		Id:            script.ID,
		TaskNumber:    script.TaskNumber,
		ScriptName:    script.ScriptName,
		ScriptContent: script.ScriptContent,
		ExecuteState:  script.ExecuteState,
		ExecuteResult: script.ExecuteResult,
	}, nil
}
func (s *ComputePowerService) DownloadScriptExecuteResult(ctx context.Context, req *pb.DownloadScriptExecuteResultRequest) (*pb.DownloadScriptExecuteResultReply, error) {
	script, err := s.uc.GetScriptInfo(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.DownloadScriptExecuteResultReply{
		Body: []byte(script.ScriptContent),
		Name: script.ScriptName,
	}, nil
}
