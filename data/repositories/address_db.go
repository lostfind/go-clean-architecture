package repositories

import (
	"database/sql"
	"zipcode/domain/model"
	"zipcode/domain/repository"
	"zipcode/domain/usecase"
)

func NewDbRepository(sqlDb *sql.DB) repository.ZipCodeRepository {
	return &addressDb{db: sqlDb}
}

type addressDb struct {
	db *sql.DB
}

func (r addressDb) GetAddressForZipCode(zipcode string) (model.Address, error) {
	address := model.Address{}
	sql := `SELECT zipcode
	, prefecture_id
	, prefecutre_name
	, prefecture_name_kana
	, city_id
	, city_name
	, city_name_kana
	, town_name
	, town_name_kana
	FROM address WHERE zipcode = ?;`

	rows, err := r.db.Query(sql, zipcode)
	if err != nil {
		return address, err
	}

	if rows.Next() {
		err := rows.Scan(&address.ZipCode,
			&address.Prefecture.ID,
			&address.Prefecture.Name,
			&address.Prefecture.NameKana,
			&address.City.ID,
			&address.City.Name,
			&address.City.NameKana,
			&address.TownName,
			&address.TownNameKana,
		)
		if err != nil {
			return address, err
		}
	} else {
		return address, usecase.ErrNotFound
	}

	return address, nil
}
