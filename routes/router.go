package routes

import (
	"github.com/Frientia/gin-firebase-backend/handlers"
	"github.com/Frientia/gin-firebase-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Init handlers
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()

	// API v1 group
	v1 := r.Group("/v1")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "gin-firebase-backend"})
		})

		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/verify-token", authHandler.VerifyToken)
		}

		// Protected routes (require Backend JWT)
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			products := protected.Group("/products")
			{
				products.GET("", productHandler.GetAll)
				products.GET("/:id", productHandler.GetByID)

				// Admin only
				adminProducts := products.Group("")
				adminProducts.Use(middleware.AdminOnly())
				{
					adminProducts.POST("", productHandler.Create)
					adminProducts.PUT("/:id", productHandler.Update)
					adminProducts.DELETE("/:id", productHandler.Delete)
				}
			}
		}
	}

	return r
}
