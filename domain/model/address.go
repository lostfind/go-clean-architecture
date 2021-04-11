package model

type Address struct {
	ZipCode      string
	Prefecture   Prefecture
	City         City
	TownName     string
	TownNameKana string
}
