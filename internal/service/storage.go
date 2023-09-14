package service

import (
	pb "computeshare-server/api/compute/v1"
	"computeshare-server/internal/biz"
	"computeshare-server/internal/global"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	iface "github.com/ipfs/boxo/coreiface"
	"github.com/ipfs/boxo/coreiface/options"
	"github.com/ipfs/boxo/coreiface/path"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/samber/lo"
	"io"
	"time"
)

type StorageService struct {
	pb.UnimplementedStorageServer
	uc       *biz.Storagecase
	ipfsNode *core.IpfsNode
	ipfsapi  iface.CoreAPI
	log      *log.Helper
}

func NewStorageService(uc *biz.Storagecase, ipfsNode *core.IpfsNode, logger log.Logger) (*StorageService, error) {
	api, err := coreapi.NewCoreAPI(ipfsNode, options.Api.FetchBlocks(true))
	if err != nil {
		return nil, err
	}
	return &StorageService{
		uc:       uc,
		ipfsNode: ipfsNode,
		ipfsapi:  api,
		log:      log.NewHelper(logger),
	}, err
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
			LastModify: item.LastModify.Unix(),
			Cid:        &item.Cid,
		}
	})
	return &pb.ListReply{
		Result: list,
	}, err
}
func (s *StorageService) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.File, error) {
	token, ok := global.FromContext(ctx)
	if ok == false {
		return nil, errors.New("cannot get user ID")
	}

	opts := []options.UnixfsAddOption{
		options.Unixfs.Hash(18),

		options.Unixfs.Inline(false),
		options.Unixfs.InlineLimit(32),

		options.Unixfs.Chunker("size-262144"),

		options.Unixfs.Pin(true),
		options.Unixfs.HashOnly(false),
		options.Unixfs.FsCache(false),
		options.Unixfs.Nocopy(false),

		options.Unixfs.Progress(true),
		options.Unixfs.Silent(false),
	}

	fileNode := files.NewBytesFile(req.Body)
	pathAdded, err := s.ipfsapi.Unixfs().Add(ctx, fileNode, opts...)

	if err != nil {
		return nil, err
	}

	size, err := fileNode.Size()
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
		Cid:        pathAdded.Cid().String(),
	}
	err = s.uc.Create(ctx, storage)
	return &pb.File{
		Id:         storage.ID.String(),
		Name:       storage.Name,
		Cid:        &storage.Cid,
		LastModify: storage.LastModify.Unix(),
		Type:       pb.FileType(storage.Type),
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
	f, err := s.ipfsapi.Unixfs().Get(ctx, path.New(cid))
	var file files.File
	switch f := f.(type) {
	case files.File:
		file = f
	case files.Directory:
		return nil, iface.ErrIsDir
	default:
		return nil, iface.ErrNotSupported
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return &pb.DownloadReply{
		Name: storage.Name,
		Body: data,
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

	return &pb.DeleteReply{}, nil
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
		Id: storage.ID.String(),
	}, err
}
