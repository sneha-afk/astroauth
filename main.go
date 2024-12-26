package main

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sneha-afk/astroauth/routes"
	"github.com/sneha-afk/astroauth/store"
	"github.com/sneha-afk/astroauth/utils"
)

func envOrDefault(key string, def string) string {
	attempt := os.Getenv(key)
	if attempt == "" {
		return def
	}
	return attempt
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize DB + load schema
	if err := store.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer store.CloseDB()

	if err := store.ExecuteSQLFile("./store/db_schema.sql"); err != nil {
		log.Fatal(err)
	}

	// Load in keys
	privKeyPath := envOrDefault("PRIV_KEY_LOC", "private_key.pem")
	pubKeyPath := envOrDefault("PUBLIC_KEY_LOC", "public_key.pem")
	authKeyType := envOrDefault("AUTH_KEY_TYPE", "RSA")
	if err := utils.LoadKeys(privKeyPath, pubKeyPath, authKeyType); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	routes.RegisterRoutes(router)

	// Signal handler recommended by Gin
	port := envOrDefault("PORT", "8080")
	endless.ListenAndServe(":"+port, router)
}
