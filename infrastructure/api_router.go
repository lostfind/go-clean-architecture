package infrastructure

import (
	"zipcode/presentation/api"

	"github.com/gin-gonic/gin"
)

func GinServer() *gin.Engine {
	apiController := api.NewAPIController()
	r := gin.Default()
	r.GET("/zipcodes/:zipcode", func(c *gin.Context) { apiController.GetAddress(c) })

	return r
}
