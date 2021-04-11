package api

import (
	"net/http"
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

func (c ApiConroller) GetAddress(w http.ResponseWriter, req *http.Request) error {
	zipcode := req.URL.Query().Get("zipcode")
	address, err := c.useCase.GetAddress(zipcode)
	if err != nil {
		return err
	}

	return JsonResponse(w, viewmodel.NewAddressViewModel(address))
}
