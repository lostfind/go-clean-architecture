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

func NewApiConroller() *ApiConroller {
	return &ApiConroller{
		useCase: usecase.NewAddressFinder(repositories.AddressFile{}),
	}
}

func (c ApiConroller) Init() {
	repositories.LoadData()
	c.serverInit()
}

func (c ApiConroller) serverInit() {
	http.HandleFunc("/zipcodes", c.GetAddress)
	http.ListenAndServe(":8080", nil)
}

func (c ApiConroller) GetAddress(w http.ResponseWriter, req *http.Request) {
	zipcode := req.URL.Query().Get("zipcode")
	zipCodeModel, err := c.useCase.GetAddress(zipcode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := Json(viewmodel.AddressByZipCode(zipCodeModel))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
