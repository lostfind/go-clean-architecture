package repositories

import (
	"zipcode/domain/model"
	"zipcode/domain/usecase"
)

type AddressFile struct {
}

func (z AddressFile) GetAddressForZipCode(zipcode string) (model.Address, error) {
	addr, ok := AddressDatas[zipcode]
	if !ok {
		return model.Address{}, usecase.ErrNotFound
	}

	address := model.Address{
		ZipCode:      addr.zipCode,
		Prefecture:   model.Prefecture{ID: addr.prefID, Name: addr.prefName, NameKana: addr.prefNameKana},
		City:         model.City{ID: addr.cityID, Name: addr.cityName, NameKana: addr.cityNameKana},
		TownName:     addr.townName,
		TownNameKana: addr.townNameKana,
	}

	return address, nil
}

type AddressData struct {
	zipCode      string
	prefID       int64
	prefName     string
	prefNameKana string
	cityID       int64
	cityName     string
	cityNameKana string
	townName     string
	townNameKana string
}

var AddressDatas map[string]AddressData

func LoadData() {
	AddressDatas = make(map[string]AddressData)

	AddressDatas["4630062"] = AddressData{
		zipCode:      "4630062",
		prefID:       23,
		prefName:     "愛知県",
		prefNameKana: "アイチケン",
		cityID:       23113,
		cityName:     "名古屋市守山区",
		cityNameKana: "ナゴヤシモリヤマク",
		townName:     "長栄",
		townNameKana: "チョウエイ",
	}
	AddressDatas["4630063"] = AddressData{
		zipCode:      "4630063",
		prefID:       23,
		prefName:     "愛知県",
		prefNameKana: "アイチケン",
		cityID:       23113,
		cityName:     "名古屋市守山区",
		cityNameKana: "ナゴヤシモリヤマク",
		townName:     "八反",
		townNameKana: "ハッタン",
	}
}
