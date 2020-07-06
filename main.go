package main

import (
	"ironchip.net/gin-example/server"
)

func main() {
	address := "localhost"
	port := "8010"

	err := server.HTTPServer(address, port).Init()
	if err != nil {
		panic(err)
	}
}
