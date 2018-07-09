package main

import (
	"github.com/daniel/basic_microservice/article/routes"
)

var (
	SERVER_PORT = ":8080"
)
func main() {

	e := routes.InitRoutes()

	e.Logger.Fatal(e.Start(SERVER_PORT))
}
