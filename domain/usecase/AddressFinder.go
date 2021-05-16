package usecase

import (
	"zipcode/domain/model"
	"zipcode/domain/repository"
)

type AddressFinder struct {
	zipCodeRepository repository.ZipCodeRepository
}

func NewAddressFinder(repo repository.ZipCodeRepository) AddressFinder {
	return AddressFinder{zipCodeRepository: repo}
}

func (a AddressFinder) GetAddress(zipcode string) (model.Address, error) {
	zipCode, err := a.zipCodeRepository.GetAddressForZipCode(zipcode)
	return zipCode, err
}

func (u AddressFinder) GetPrefectures() ([]model.Prefecture, error) {
	prefs, err := u.zipCodeRepository.GetPrefectures()
	return prefs, err
}
