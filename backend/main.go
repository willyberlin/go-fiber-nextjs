package main

import (
	"fiber-rooms/app"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fiberApp := fiber.New()
	app.SetupApp(fiberApp)

	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(fiberApp.Listen(":" + port))
}
