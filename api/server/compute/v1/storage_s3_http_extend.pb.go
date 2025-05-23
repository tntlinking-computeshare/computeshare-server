package v1

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

func Storage_S3_UploadFile_Extend_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in S3StorageUploadFileRequest

		req := ctx.Request()
		fmt.Println(req.Header.Get("x-md-globa-owner"))
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		prefix := req.FormValue("prefix")
		file, fileHeader, err := req.FormFile("file")
		if err != nil {
			return err
		}
		in.Prefix = prefix
		in.FileName = fileHeader.Filename
		in.Body, err = io.ReadAll(file)
		defer file.Close()

		http.SetOperation(ctx, OperationStorageS3S3StorageUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.S3StorageUploadFile(ctx, req.(*S3StorageUploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*S3StorageUploadFileReply)
		return ctx.Result(200, reply)
	}
}

func Storage_S3_DownloadFile_Extend_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in S3StorageDownloadRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3S3StorageDownload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.S3StorageDownload(ctx, req.(*S3StorageDownloadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*S3StorageDownloadReply)
		disposition := fmt.Sprintf("attachment; filename=%s", reply.Data.Name)
		ctx.Response().Header().Set("Content-Type", "application/octet-stream")
		ctx.Response().Header().Set("Content-Disposition", disposition)
		ctx.Response().Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		_, err = ctx.Response().Write(reply.Data.Body)
		return err
	}
}
