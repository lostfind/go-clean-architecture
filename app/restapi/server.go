package restapi

import (
	"zipcode/app/graphql/graph"
	"zipcode/app/graphql/graph/generated"
	"zipcode/infrastructure/server"
	"zipcode/presentation/api"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	controller api.ApiController
	engine     *gin.Engine
}

func NewApiServer(apiController api.ApiController) server.Server {
	return apiServer{
		controller: apiController,
		engine:     gin.Default(),
	}
}

func (s apiServer) Run(addr ...string) (err error) {
	return s.engine.Run(addr...)
}

func (s apiServer) Router() {
	s.engine.GET("/zipcodes/:zipcode", func(c *gin.Context) { s.controller.GetAddress(c) })
	s.engine.POST("/query", graphqlHandler())
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
