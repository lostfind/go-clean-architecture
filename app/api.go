package main

import (
	"zipcode/data/repositories"
	"zipcode/infrastructure"
)

func main() {
	repositories.LoadData()
	infrastructure.GinServer().Run()
}
