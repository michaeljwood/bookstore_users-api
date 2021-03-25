package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/michaeljwood/bookstore_users-api/app"
)

func main() {
	app.StartApplication()
}
