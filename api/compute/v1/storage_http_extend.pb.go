package v1

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

func Storage_UploadFile_Extend_HTTP_Handler(srv StorageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadFileRequest

		req := ctx.Request()
		fmt.Println(req.Header.Get("x-md-globa-owner"))
		parentId := req.FormValue("parentId")
		file, handler, err := req.FormFile("file")
		if err != nil {
			return err
		}

		in.Name = handler.Filename
		in.Body, err = io.ReadAll(file)
		in.ParentId = &parentId
		defer file.Close()

		http.SetOperation(ctx, OperationStorageUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFile(ctx, req.(*UploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*File)
		return ctx.Result(200, reply)
	}
}

func Storage_DownloadFile_Extend_HTTP_Handler(srv StorageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageDownload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Download(ctx, req.(*DownloadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadReply)
		disposition := fmt.Sprintf("attachment; filename=%s", reply.Name)
		ctx.Response().Header().Set("Content-Type", "application/octet-stream")
		ctx.Response().Header().Set("Content-Disposition", disposition)
		ctx.Response().Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		_, err = ctx.Response().Write(reply.Body)
		return err
	}
}
