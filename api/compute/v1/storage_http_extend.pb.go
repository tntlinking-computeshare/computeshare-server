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

		owner := req.Header.Get("x-md-globa-owner")

		in.Name = handler.Filename
		in.Body, err = io.ReadAll(file)
		in.Owner = owner
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
