package v1

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

func Compute_Power_UploadSceipt_Extend_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadScriptFileRequest

		req := ctx.Request()
		fmt.Println(req.Header.Get("x-md-globa-owner"))
		file, handler, err := req.FormFile("file")
		if err != nil {
			return err
		}

		in.Name = handler.Filename
		in.Body, err = io.ReadAll(file)
		defer file.Close()

		http.SetOperation(ctx, OperationComputePowerUploadScriptFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadScriptFile(ctx, req.(*UploadScriptFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadScriptFileReply)
		return ctx.Result(200, reply)
	}
}

func Compute_Powere_DownloadResult_Extend_HTTP_Handler(srv ComputePowerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadScriptExecuteResultRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationComputePowerDownloadScriptExecuteResult)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DownloadScriptExecuteResult(ctx, req.(*DownloadScriptExecuteResultRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadScriptExecuteResultReply)
		disposition := fmt.Sprintf("attachment; filename=%s", reply.Name)
		ctx.Response().Header().Set("Content-Type", "application/octet-stream")
		ctx.Response().Header().Set("Content-Disposition", disposition)
		ctx.Response().Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		_, err = ctx.Response().Write(reply.Body)
		return err
	}
}
