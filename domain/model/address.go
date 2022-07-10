package model

type Address struct {
	ZipCode      string
	TownName     string
	TownNameKana string
	Prefecture
	City
}
