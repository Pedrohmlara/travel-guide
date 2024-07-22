package controller

import (
	"demo-travel-guide/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeItinerary(context *gin.Context) {
	var itineraryRequest model.ItineraryRequest
	if err := context.ShouldBindJSON(&itineraryRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data."})
	}

}
