package infrastructure

import (
	"net/http"
	"zipcode/domain/usecase"
	"zipcode/presentation/api"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		status := http.StatusInternalServerError

		switch err {
		case usecase.ErrNotFound:
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
	}
}

func ApiServerInit() {
	apiController := api.NewAPIController()

	http.Handle("/zipcodes", appHandler(apiController.GetAddress))
	http.ListenAndServe(":8080", nil)
}
