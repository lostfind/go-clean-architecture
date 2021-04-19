package api

import (
	"zipcode/domain/usecase"
	"zipcode/presentation/viewmodel"
)

type ApiController interface {
	GetAddress(ctx Context)
}

type apiController struct {
	useCase usecase.AddressFinder
}

func NewApiController(uc usecase.AddressFinder) ApiController {
	return &apiController{
		useCase: uc,
	}
}

func (c apiController) GetAddress(ctx Context) {
	zipcode := ctx.Param("zipcode")
	address, err := c.useCase.GetAddress(zipcode)
	if err != nil {
		ErrResponse(ctx, err.Error())
		return
	}

	JsonResponse(ctx, viewmodel.NewAddressViewModel(address))
}
