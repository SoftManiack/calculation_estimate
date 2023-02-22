package handler

import (
	service "calculation-estimate/pkg/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Authorization, User-Agent, Accept, Origin, Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	inquirer := router.Group("/v1")
	{
		inquirer.GET("/inquirer", h.GetInquirer)
		inquirer.POST("/inquirer", h.GetEstimate)
	}

	return router
}
