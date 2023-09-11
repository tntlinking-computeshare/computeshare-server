package service

import (
	pb "computeshare-server/api/compute/v1"
	"computeshare-server/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	iface "github.com/ipfs/boxo/coreiface"
	"github.com/ipfs/boxo/coreiface/options"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/samber/lo"
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
	result, err := s.uc.List(ctx, req.GetOwner(), req.GetParentId())
	if err != nil {
		return nil, err
	}
	files := lo.Map(result, func(item *biz.Storage, index int) *pb.File {
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
		Result: files,
	}, err
}
func (s *StorageService) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.File, error) {

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
		Owner:      req.GetOwner(),
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
	return &pb.DownloadReply{}, nil
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

	storage := &biz.Storage{
		Owner:      req.GetOwner(),
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
