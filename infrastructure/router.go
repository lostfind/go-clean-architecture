package infrastructure

import (
	"zipcode/presentation/api"

	"github.com/gin-gonic/gin"
)

func ApiServer(controller api.ApiController) *gin.Engine {
	r := gin.Default()
	r.GET("/zipcodes/:zipcode", func(c *gin.Context) { controller.GetAddress(c) })

	return r
}
