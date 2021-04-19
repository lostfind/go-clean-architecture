package api

import (
	"log"
	"zipcode/domain/usecase"
	"zipcode/presentation/viewmodel"
)

type ApiController interface {
	GetAddress(ctx Context) error
}

type apiController struct {
	useCase usecase.AddressFinder
}

func NewApiController(uc usecase.AddressFinder) ApiController {
	return &apiController{
		useCase: uc,
	}
}

func (c apiController) GetAddress(ctx Context) error {
	zipcode := ctx.Param("zipcode")
	address, err := c.useCase.GetAddress(zipcode)
	if err != nil {
		log.Print(err)
		return err
	}

	JsonResponse(ctx, viewmodel.NewAddressViewModel(address))

	return nil
}
