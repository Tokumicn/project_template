//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Tokumicn/kratos_tlp/internal/biz"
	"github.com/Tokumicn/kratos_tlp/internal/conf"
	"github.com/Tokumicn/kratos_tlp/internal/data"
	"github.com/Tokumicn/kratos_tlp/internal/server"
	"github.com/Tokumicn/kratos_tlp/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
