package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
)

type ComputePowerService struct {
	pb.UnimplementedComputePowerServer

	uc  *biz.ScriptUseCase
	log *log.Helper
}

func NewComputePowerService(uc *biz.ScriptUseCase, logger log.Logger) (*ComputePowerService, error) {
	return &ComputePowerService{
		uc:  uc,
		log: log.NewHelper(logger),
	}, nil
}

func (s *ComputePowerService) UploadScriptFile(ctx context.Context, req *pb.UploadScriptFileRequest) (*pb.UploadScriptFileReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}

	//opts := []options.UnixfsAddOption{
	//	options.Unixfs.Hash(18),
	//
	//	options.Unixfs.Inline(false),
	//	options.Unixfs.InlineLimit(32),
	//
	//	options.Unixfs.Chunker("size-262144"),
	//
	//	options.Unixfs.Pin(true),
	//	options.Unixfs.HashOnly(false),
	//	options.Unixfs.FsCache(false),
	//	options.Unixfs.Nocopy(false),
	//
	//	options.Unixfs.Progress(true),
	//	options.Unixfs.Silent(false),
	//}
	//
	//fileNode := files.NewBytesFile(req.Body)
	//pathAdded, err := s.ipfsApi.Unixfs().Add(ctx, fileNode, opts...)

	script := biz.Script{
		UserId:        token.UserID,
		ScriptName:    req.Name,
		ScriptContent: string(req.Body),
		FileAddress:   "pathAdded.Cid().String()",
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
	if err != nil {
		return nil, err
	}
	return &pb.DownloadScriptExecuteResultReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.DownloadScriptExecuteResultReply_Data{
			Body: []byte(script.ExecuteResult),
			Name: script.ScriptName,
		},
	}, nil
}
