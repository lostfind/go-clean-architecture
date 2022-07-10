package main

import (
	"zipcode/app/restapi"
	"zipcode/data/repositories"
	"zipcode/domain/usecase"

	"zipcode/presentation/api"
)

var apiController api.ApiController

func main() {
	inject()
	server := restapi.NewApiServer(apiController)
	server.Router()
	server.Run()
}

func inject() {
	repository := repositories.NewFileRepository()
	useCase := usecase.NewAddressFinder(repository)
	apiController = api.NewApiController(useCase)
}
