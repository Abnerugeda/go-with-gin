package main

import (
	"github.com/abnerugeda/go-with-gin/database"
	"github.com/abnerugeda/go-with-gin/routes"
)

func main() {
	database.ConnectDB()

	routes.HandleRequests()
}
