package main

import "zipcode/presentation/api"

func main() {
	apiApp := api.NewApiConroller()
	apiApp.Init()
}
