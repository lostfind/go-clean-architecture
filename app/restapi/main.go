package main

import (
	"zipcode/data/repositories"
	"zipcode/domain/usecase"
	"zipcode/infrastructure"
	"zipcode/infrastructure/database"

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
	apiServer := infrastructure.ApiServer(apiController)
	apiServer.Run()
}
