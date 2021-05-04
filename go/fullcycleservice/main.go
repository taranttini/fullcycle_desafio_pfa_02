package main

import (
	"fullcycleservice/database"
	"fullcycleservice/module"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	module.SetupRoutes(apiBasePath)
	log.Printf("inicia servidor")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
