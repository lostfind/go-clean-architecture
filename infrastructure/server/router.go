package server

import (
	"zipcode/presentation/api"

	"github.com/gin-gonic/gin"
)

type apiServer struct {
	controller api.ApiController
	engine     *gin.Engine
}

func NewApiServer(apiController api.ApiController) Server {
	return apiServer{
		controller: apiController,
		engine:     gin.Default(),
	}
}

func (s apiServer) Run(addr ...string) (err error) {
	s.Router()
	return s.engine.Run(addr...)
}

func (s apiServer) Router() {
	s.engine.GET("/zipcodes/:zipcode", func(c *gin.Context) { s.controller.GetAddress(c) })
	s.engine.GET("/query")
}
