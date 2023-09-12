//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"computeshare-server/internal/biz"
	"computeshare-server/internal/conf"
	"computeshare-server/internal/data"
	"computeshare-server/internal/server"
	"computeshare-server/internal/service"
	"computeshare-server/third_party/p2p"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, p2p.ProviderSet, newApp))
}
