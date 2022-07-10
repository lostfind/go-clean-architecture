//go:build wireinject
// +build wireinject

package main

import (
	"zipcode/app/restapi"
	"zipcode/data/repositories"
	"zipcode/domain/usecase"
	"zipcode/infrastructure/server"
	"zipcode/presentation/api"

	"github.com/google/wire"
)

func InitializeRestApi() server.Server {
	wire.Build(restapi.NewApiServer, api.NewApiController, usecase.NewAddressFinder, repositories.NewFileRepository)

	return nil
}
