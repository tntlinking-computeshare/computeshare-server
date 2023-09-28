package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	shell "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-files"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
)

type ComputePowerService struct {
	pb.UnimplementedComputePowerServer

	ipfsShell *shell.Shell
	uc        *biz.ScriptUseCase
	log       *log.Helper
}

func NewComputePowerService(uc *biz.ScriptUseCase, ipfsShell *shell.Shell, logger log.Logger) (*ComputePowerService, error) {
	return &ComputePowerService{
		uc:        uc,
		ipfsShell: ipfsShell,
		log:       log.NewHelper(logger),
	}, nil
}

func (s *ComputePowerService) UploadScriptFile(ctx context.Context, req *pb.UploadScriptFileRequest) (*pb.UploadScriptFileReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}

	file := files.NewBytesFile(req.Body)
	pathAdded, err := s.ipfsShell.Add(file, shell.Pin(true))
	if err != nil {
		s.log.Error("ipfs add failed err is", err)
		return nil, errors.New("ipfs add failed")
	}
	script := biz.Script{
		UserId:        token.UserID,
		ScriptName:    req.Name,
		ScriptContent: string(req.Body),
		FileAddress:   pathAdded,
	}
	createScript, err := s.uc.CreateScript(ctx, &script)
	if err != nil {
		return nil, err
	}
	return &pb.UploadScriptFileReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.ScriptReply{
			Id:            createScript.ID,
			TaskNumber:    createScript.TaskNumber,
			ScriptName:    createScript.ScriptName,
			ScriptContent: createScript.ScriptContent,
		},
	}, nil
}
func (s *ComputePowerService) GetScriptExecutionRecordList(ctx context.Context, req *pb.GetScriptExecutionRecordListRequest) (*pb.GetScriptListReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}
	data, total, err := s.uc.GetScriptExecutionRecordPage(ctx, token.UserID, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	var pointerList []*pb.ScriptReply
	for _, script := range data {
		var scriptReply pb.ScriptReply
		scriptReply.Id = script.ID
		scriptReply.TaskNumber = script.TaskNumber
		scriptReply.ScriptName = script.ScriptName
		scriptReply.ScriptContent = script.ScriptContent
		scriptReply.ExecuteState = script.ExecuteState
		scriptReply.ExecuteResult = script.ExecuteResult
		pointerList = append(pointerList, &scriptReply)
	}
	return &pb.GetScriptListReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.GetScriptListReply_Data{
			List:  pointerList,
			Total: total,
			Page:  req.GetPage(),
			Size:  req.GetSize(),
		},
	}, nil
}
func (s *ComputePowerService) RunPythonPackage(ctx context.Context, req *pb.RunPythonPackageServerRequest) (*pb.RunPythonPackageServerReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}
	script, err := s.uc.RunPythonPackage(ctx, req.GetId(), token.UserID)
	if err != nil {
		return nil, err
	}
	return &pb.RunPythonPackageServerReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.ScriptReply{
			Id:            script.ID,
			TaskNumber:    script.TaskNumber,
			ScriptName:    script.ScriptName,
			ScriptContent: script.ScriptContent,
			ExecuteState:  script.ExecuteState,
			ExecuteResult: script.ExecuteResult,
		},
	}, nil
}

func (s *ComputePowerService) CancelExecPythonPackage(ctx context.Context, req *pb.CancelExecPythonPackageRequest) (*pb.CancelExecPythonPackageReply, error) {
	script, err := s.uc.CancelExecPythonPackage(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.CancelExecPythonPackageReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.ScriptReply{
			Id:            script.ID,
			TaskNumber:    script.TaskNumber,
			ScriptName:    script.ScriptName,
			ScriptContent: script.ScriptContent,
			ExecuteState:  script.ExecuteState,
			ExecuteResult: script.ExecuteResult,
		},
	}, nil
}
func (s *ComputePowerService) GetScriptExecutionRecordInfo(ctx context.Context, req *pb.GetScriptExecutionRecordInfoRequest) (*pb.GetScriptInfoReply, error) {
	script, err := s.uc.GetScriptExecutionRecordInfo(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetScriptInfoReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.ScriptReply{
			Id:            script.ID,
			TaskNumber:    script.TaskNumber,
			ScriptName:    script.ScriptName,
			ScriptContent: script.ScriptContent,
			ExecuteState:  script.ExecuteState,
			ExecuteResult: script.ExecuteResult,
		},
	}, nil
}

func (s *ComputePowerService) DownloadScriptExecuteResult(ctx context.Context, req *pb.DownloadScriptExecuteResultRequest) (*pb.DownloadScriptExecuteResultReply, error) {
	script, err := s.uc.GetScriptExecutionRecordInfo(ctx, req.GetId())
	name := string(script.TaskNumber) + ".log"
	if err != nil {
		return nil, err
	}
	return &pb.DownloadScriptExecuteResultReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.DownloadScriptExecuteResultReply_Data{
			Body: []byte(script.ExecuteResult),
			Name: name,
		},
	}, nil
}
