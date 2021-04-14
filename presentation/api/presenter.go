package api

import (
	"net/http"
)

func JsonResponse(ctx Context, data interface{}) error {

	ctx.JSON(http.StatusOK, data)
	return nil
}
