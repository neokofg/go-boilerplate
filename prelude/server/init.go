package server

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/api/middleware"
	"go-boilerplate/api/router"
	"go-boilerplate/infrastructure/ent"
	"go.uber.org/zap"
)

func InitGin(sugar *zap.SugaredLogger, client *ent.Client) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(sugar))

	router.InitRoutes(r, client)

	r.Run(":8080")
}
