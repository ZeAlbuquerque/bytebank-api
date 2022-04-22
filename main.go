package main

import (
	"bytebank-api/database"
	"bytebank-api/routes"
)

func main() {
	database.Conecta()
	routes.HandleRequests()
}
