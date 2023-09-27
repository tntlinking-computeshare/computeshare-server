package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	shell "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-files"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/samber/lo"
	"time"
)

type StorageService struct {
	pb.UnimplementedStorageServer

	ipfsShell *shell.Shell
	uc        *biz.Storagecase
	log       *log.Helper
}

func NewStorageService(uc *biz.Storagecase, ipfsShell *shell.Shell, logger log.Logger) (*StorageService, error) {
	return &StorageService{
		uc:        uc,
		ipfsShell: ipfsShell,
		log:       log.NewHelper(logger),
	}, nil
}

func (s *StorageService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}
	result, err := s.uc.List(ctx, token.UserID, req.GetParentId())
	if err != nil {
		return nil, err
	}
	list := lo.Map(result, func(item *biz.Storage, index int) *pb.File {
		return &pb.File{
			Id:         item.ID.String(),
			Type:       pb.FileType(item.Type),
			Name:       item.Name,
			Size:       item.Size,
			LastModify: item.LastModify.UnixMilli(),
			Cid:        &item.Cid,
		}
	})
	return &pb.ListReply{
		Code:    200,
		Message: SUCCESS,
		Data:    list,
	}, err
}
func (s *StorageService) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileReply, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}

	file := files.NewBytesFile(req.Body)
	pathAdded, err := s.ipfsShell.Add(file, shell.OnlyHash(false), shell.Pin(true), shell.Progress(true))
	if err != nil {
		s.log.Error("ipfs add failed err is", err)
		return nil, errors.New("ipfs add failed")
	}
	size, err := file.Size()
	if err != nil {
		return nil, err
	}

	storage := &biz.Storage{
		Owner:      token.UserID,
		Type:       int32(pb.FileType_FILE),
		Size:       int32(size),
		Name:       req.GetName(),
		ParentID:   req.GetParentId(),
		LastModify: time.Now(),
		Cid:        pathAdded,
	}

	s.log.Info("uploaded: ", req.GetName())

	err = s.uc.Create(ctx, storage)
	return &pb.UploadFileReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.File{
			Id:         storage.ID.String(),
			Name:       storage.Name,
			Cid:        &storage.Cid,
			LastModify: storage.LastModify.UnixMilli(),
			Type:       pb.FileType(storage.Type),
			Size:       storage.Size,
		},
	}, err
}
func (s *StorageService) Download(ctx context.Context, req *pb.DownloadRequest) (*pb.DownloadReply, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	storage, err := s.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	cid := storage.Cid
	if cid == "" {
		return nil, errors.New("download file error")
	}
	ipfsReadCloser, err := s.ipfsShell.Cat(cid)
	if err != nil {
		s.log.Error("ipfs get failed err is", err)
		return nil, err
	}
	ipfsDataBuffer := new(bytes.Buffer)
	_, err = ipfsDataBuffer.ReadFrom(ipfsReadCloser)
	if err != nil {
		s.log.Error("ipfsReadCloser to ipfsDataBuffer failed err is", err)
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &pb.DownloadReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.DownloadReply_Data{
			Name: storage.Name,
			Body: ipfsDataBuffer.Bytes(),
		},
	}, err
}
func (s *StorageService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	for _, id := range req.GetIds() {
		fileId, err := uuid.Parse(id)
		if err != nil {
			continue
		}
		err = s.uc.Delete(ctx, fileId)
		if err != nil {
			return nil, err
		}
	}

	return &pb.DeleteReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}

func (s *StorageService) CreateDir(ctx context.Context, req *pb.CreateDirRequest) (*pb.CreateDirReply, error) {
	// 	判断该目录有无重复名字的文件夹
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}
	storage := &biz.Storage{
		Owner:      token.UserID,
		Type:       int32(pb.FileType_DIR),
		Name:       req.GetName(),
		ParentID:   req.GetParentId(),
		Cid:        "",
		LastModify: time.Now(),
	}
	err := s.uc.Create(ctx, storage)
	return &pb.CreateDirReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CreateDirReply_Data{
			Id: storage.ID.String(),
		},
	}, err
}
