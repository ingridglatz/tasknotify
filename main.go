package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"tasknotify/config"
	migrate "tasknotify/internal"
	"tasknotify/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	config.ConnectDB()
	migrate.AutoMigrate()

	env := os.Getenv("ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor rodando em http://localhost:" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
