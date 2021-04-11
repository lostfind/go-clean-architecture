package viewmodel

import "zipcode/domain/model"

type AddressViewModel struct {
	PrefCode      int64       `json:"prefCode"`
	CityCode      int64       `json:"cityCode"`
	ZipCode       string      `json:"zipCode"`
	PrefName      string      `json:"prefName"`
	CityName      string      `json:"cityName"`
	TownName      string      `json:"townName"`
	FullWidthKana AddressKana `json:"fullWidthKana"`
}

type AddressKana struct {
	PrefName string `json:"prefKana"`
	CityName string `json:"cityKana"`
	TownName string `json:"townKana"`
}

func AddressByZipCode(address model.Address) AddressViewModel {
	viewModel := AddressViewModel{
		ZipCode:  address.ZipCode,
		PrefCode: address.Prefecture.ID,
		PrefName: address.Prefecture.Name,
		CityName: address.City.Name,
		TownName: address.TownName,
		FullWidthKana: AddressKana{
			PrefName: address.Prefecture.NameKana,
			CityName: address.City.NameKana,
			TownName: address.TownNameKana,
		},
	}

	return viewModel
}
