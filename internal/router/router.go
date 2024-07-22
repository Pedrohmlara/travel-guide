package router

import (
	"demo-travel-guide/internal/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	InitializeRoutes(router)

	err := router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic("Error on initializing application.")
	}
}

func InitializeRoutes(router *gin.Engine) {
	routes := router.Group("/api/")
	{
		inineraryRoutes := routes.Group("/itinerary/")
		{
			inineraryRoutes.POST("/make", controller.MakeItinerary)
		}
	}
}
