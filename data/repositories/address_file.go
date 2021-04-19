package repositories

import (
	"zipcode/domain/model"
	"zipcode/domain/repository"
	"zipcode/domain/usecase"
)

var addressDatas map[string]addressData

func init() {
	addressDatas = make(map[string]addressData)

	addressDatas["4630062"] = addressData{
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
	addressDatas["4630063"] = addressData{
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

func NewFileRepository() repository.ZipCodeRepository {
	return new(addressFile)
}

type addressFile struct {
}

func (f addressFile) GetAddressForZipCode(zipcode string) (model.Address, error) {
	addr, ok := addressDatas[zipcode]
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

type addressData struct {
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
