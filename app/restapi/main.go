package main

import (
	"zipcode/data/repositories"
	"zipcode/domain/usecase"
	"zipcode/infrastructure/database"
	"zipcode/infrastructure/server"

	"zipcode/presentation/api"
)

var apiController api.ApiController

func init() {
	db := database.NewMySqlDB()
	repository := repositories.NewDbRepository(db)
	useCase := usecase.NewAddressFinder(repository)
	apiController = api.NewApiController(useCase)
}

func main() {
	apiServer := server.NewApiServer(apiController)
	apiServer.Run()
}
