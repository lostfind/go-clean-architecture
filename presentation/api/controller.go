package api

import (
	"zipcode/data/repositories"
	"zipcode/domain/usecase"
	"zipcode/presentation/viewmodel"
)

type ApiConroller struct {
	useCase usecase.AddressFinder
}

func NewAPIController() *ApiConroller {
	return &ApiConroller{
		useCase: usecase.NewAddressFinder(repositories.AddressFile{}),
	}
}

func (c ApiConroller) GetAddress(ctx Context) error {
	zipcode := ctx.Param("zipcode")
	address, err := c.useCase.GetAddress(zipcode)
	if err != nil {
		return err
	}

	return JsonResponse(ctx, viewmodel.NewAddressViewModel(address))
}
