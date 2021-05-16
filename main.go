package main

import (
	"zipcode/data/repositories"
	"zipcode/domain/usecase"
	"zipcode/infrastructure/database"
	"zipcode/infrastructure/server"

	"zipcode/presentation/api"
)

var apiController api.ApiController

func main() {
	inject()
	apiServer := server.NewApiServer(apiController)
	apiServer.Router()
	apiServer.Run()
}

func inject() {
	db := database.NewMySqlDB()
	repository := repositories.NewDbRepository(db)
	useCase := usecase.NewAddressFinder(repository)
	apiController = api.NewApiController(useCase)
}
