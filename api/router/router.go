package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-boilerplate/api/handlers"
	"go-boilerplate/api/middleware"
	"go-boilerplate/infrastructure/ent"
	"os"
)

func InitRoutes(r *gin.Engine, client *ent.Client) {
	authHandler := handlers.NewAuthHandler(client, os.Getenv("JWT_SECRET"))
	userHandler := handlers.NewUserHandler(client)
	authMiddleware := middleware.NewAuthMiddleware(client)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api := r.Group("/api")
	api.POST("/auth/login", authHandler.Login)
	api.POST("/auth/register", authHandler.Register)

	protected := api.Group("/user")
	protected.Use(authMiddleware.Handle())
	{
		protected.GET("/me", userHandler.GetSelf)
	}
}
