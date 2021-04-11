package repository

import "zipcode/domain/model"

type ZipCodeRepository interface {
	GetAddressForZipCode(zipcode string) (model.Address, error)
}
