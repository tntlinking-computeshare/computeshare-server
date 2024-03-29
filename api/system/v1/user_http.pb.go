// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/system/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserCreateUser = "/api.server.system.v1.User/CreateUser"
const OperationUserGetUser = "/api.server.system.v1.User/GetUser"
const OperationUserListUser = "/api.server.system.v1.User/ListUser"
const OperationUserLogin = "/api.server.system.v1.User/Login"
const OperationUserLoginWithClient = "/api.server.system.v1.User/LoginWithClient"
const OperationUserLoginWithValidateCode = "/api.server.system.v1.User/LoginWithValidateCode"
const OperationUserSendValidateCode = "/api.server.system.v1.User/SendValidateCode"
const OperationUserUpdateUser = "/api.server.system.v1.User/UpdateUser"
const OperationUserUpdateUserPassword = "/api.server.system.v1.User/UpdateUserPassword"
const OperationUserUpdateUserTelephone = "/api.server.system.v1.User/UpdateUserTelephone"
const OperationUserVerifyCode = "/api.server.system.v1.User/VerifyCode"

type UserHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	// Login Login
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// {{import "tables.md"}}
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	LoginWithClient(context.Context, *LoginWithClientRequest) (*LoginReply, error)
	LoginWithValidateCode(context.Context, *LoginWithValidateCodeRequest) (*LoginReply, error)
	SendValidateCode(context.Context, *SendValidateCodeRequest) (*SendValidateCodeReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	UpdateUserPassword(context.Context, *UpdateUserPasswordRequest) (*UpdateUserPasswordReply, error)
	UpdateUserTelephone(context.Context, *UpdateUserTelephoneRequest) (*UpdateUserTelephoneReply, error)
	VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user", _User_CreateUser0_HTTP_Handler(srv))
	r.PUT("/v1/user", _User_UpdateUser0_HTTP_Handler(srv))
	r.PUT("/v1/user/password", _User_UpdateUserPassword0_HTTP_Handler(srv))
	r.PUT("/v1/user/telephone", _User_UpdateUserTelephone0_HTTP_Handler(srv))
	r.GET("/v1/user", _User_GetUser0_HTTP_Handler(srv))
	r.GET("/v1/system/user", _User_ListUser0_HTTP_Handler(srv))
	r.POST("/v1/user/login", _User_Login0_HTTP_Handler(srv))
	r.POST("/v1/user/loginWithClient", _User_LoginWithClient0_HTTP_Handler(srv))
	r.POST("/v1/user/login_by_vc", _User_LoginWithValidateCode0_HTTP_Handler(srv))
	r.POST("/v1/sms/send", _User_SendValidateCode0_HTTP_Handler(srv))
	r.POST("/v1/sms/code/verify", _User_VerifyCode0_HTTP_Handler(srv))
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUserPassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUserPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserPassword(ctx, req.(*UpdateUserPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserPasswordReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUserTelephone0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserTelephoneRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUserTelephone)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserTelephone(ctx, req.(*UpdateUserTelephoneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserTelephoneReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_LoginWithClient0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginWithClientRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLoginWithClient)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginWithClient(ctx, req.(*LoginWithClientRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_LoginWithValidateCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginWithValidateCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLoginWithValidateCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginWithValidateCode(ctx, req.(*LoginWithValidateCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_SendValidateCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendValidateCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSendValidateCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendValidateCode(ctx, req.(*SendValidateCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendValidateCodeReply)
		return ctx.Result(200, reply)
	}
}

func _User_VerifyCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyCode(ctx, req.(*VerifyCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerifyCodeReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginWithClient(ctx context.Context, req *LoginWithClientRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	LoginWithValidateCode(ctx context.Context, req *LoginWithValidateCodeRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	SendValidateCode(ctx context.Context, req *SendValidateCodeRequest, opts ...http.CallOption) (rsp *SendValidateCodeReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
	UpdateUserPassword(ctx context.Context, req *UpdateUserPasswordRequest, opts ...http.CallOption) (rsp *UpdateUserPasswordReply, err error)
	UpdateUserTelephone(ctx context.Context, req *UpdateUserTelephoneRequest, opts ...http.CallOption) (rsp *UpdateUserTelephoneReply, err error)
	VerifyCode(ctx context.Context, req *VerifyCodeRequest, opts ...http.CallOption) (rsp *VerifyCodeReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/v1/system/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) LoginWithClient(ctx context.Context, in *LoginWithClientRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/loginWithClient"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLoginWithClient))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) LoginWithValidateCode(ctx context.Context, in *LoginWithValidateCodeRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login_by_vc"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLoginWithValidateCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SendValidateCode(ctx context.Context, in *SendValidateCodeRequest, opts ...http.CallOption) (*SendValidateCodeReply, error) {
	var out SendValidateCodeReply
	pattern := "/v1/sms/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSendValidateCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUserPassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...http.CallOption) (*UpdateUserPasswordReply, error) {
	var out UpdateUserPasswordReply
	pattern := "/v1/user/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUserTelephone(ctx context.Context, in *UpdateUserTelephoneRequest, opts ...http.CallOption) (*UpdateUserTelephoneReply, error) {
	var out UpdateUserTelephoneReply
	pattern := "/v1/user/telephone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUserTelephone))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...http.CallOption) (*VerifyCodeReply, error) {
	var out VerifyCodeReply
	pattern := "/v1/sms/code/verify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
