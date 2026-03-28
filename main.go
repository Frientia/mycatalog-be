package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yourusername/gin-firebase-backend/config"
	"github.com/yourusername/gin-firebase-backend/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
	}

	// Initialize Firebase Admin SDK
	config.InitFirebase()

	// Initialize database
	config.InitDatabase()

	// Setup Gin router
	router := routes.SetupRouter()

	// Get port from environment
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di http://localhost:%s", port)
	log.Printf("Health check: http://localhost:%s/v1/health", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}