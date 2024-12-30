package main

import (
	"backend/database"
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Conecta ao banco de dados
	database.Connect()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8888"))
}