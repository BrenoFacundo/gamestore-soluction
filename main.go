package main

import (
	"github.com/brenofacundo/gamestore-soluction/routes"
)

func main() {
	server := routes.Routes()
	server.Run(":8000")
}