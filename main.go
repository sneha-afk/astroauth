package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sneha-afk/astroauth/routes"
	"github.com/sneha-afk/astroauth/store"
)

const PORT = "8080"

func main() {
	if err := store.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer store.CloseDB()

	if err := store.ExecuteSQLFile("./store/db_schema.sql"); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	routes.RegisterRoutes(router)

	// Signal handler recommended by Gin
	endless.ListenAndServe(":"+PORT, router)
}
