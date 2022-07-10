package main

func main() {
	server := InitializeRestApi()
	server.Router()
	server.Run()
}
