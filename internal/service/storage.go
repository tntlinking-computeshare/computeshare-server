package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/samber/lo"
	"time"
)

type StorageService struct {
	pb.UnimplementedStorageServer
	uc *biz.Storagecase
	//ipfsNode *core.IpfsNode
	//ipfsapi  coreiface.CoreAPI
	log *log.Helper
}

func NewStorageService(uc *biz.Storagecase, logger log.Logger) (*StorageService, error) {
	//api, err := coreapi.NewCoreAPI(ipfsNode, options.Api.FetchBlocks(true))
	//if err != nil {
	//	return nil, err
	//}
	return &StorageService{
		uc: uc,
		//ipfsNode: ipfsNode,
		//ipfsapi:  api,
		log: log.NewHelper(logger),
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
	//pathAdded, err := s.ipfsapi.Unixfs().Add(ctx, fileNode, opts...)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//size, err := fileNode.Size()
	//if err != nil {
	//	return nil, err
	//}

	storage := &biz.Storage{
		Owner: token.UserID,
		Type:  int32(pb.FileType_FILE),
		//Size:       int32(size),
		Size:       int32(len(req.Body)),
		Name:       req.GetName(),
		ParentID:   req.GetParentId(),
		LastModify: time.Now(),
		//Cid:        pathAdded.Cid().String(),
		Cid: "pathAdded.Cid().String()",
	}

	s.log.Info("uploaded: ", req.GetName())

	err := s.uc.Create(ctx, storage)
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
	//f, err := s.ipfsapi.Unixfs().Get(ctx, path.New(cid))
	//var file files.File
	//switch f := f.(type) {
	//case files.File:
	//	file = f
	//case files.Directory:
	//	return nil, coreiface.ErrIsDir
	//default:
	//	return nil, coreiface.ErrNotSupported
	//}

	//data, err := io.ReadAll(file)
	data := []byte{}
	if err != nil {
		return nil, err
	}
	return &pb.DownloadReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.DownloadReply_Data{
			Name: storage.Name,
			Body: data,
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
