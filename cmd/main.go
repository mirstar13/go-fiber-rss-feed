package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/routes"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initDatabase() {
	godotenv.Load("key.env")
	db, err := sql.Open("postgres", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal(err)
	}
	config.DBQueris = config.New(db)
}

func main() {
	initDatabase()
	app := fiber.New()
	routes.RegisterFeedRoutes(app)

	go utils.StartScraping(10, time.Minute)

	log.Fatal(app.Listen(8080))
}
