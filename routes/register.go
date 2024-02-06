package routes

import (
	"net/http"
	"restwithjwt/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err})
		return
	}
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled registration successfully"})
}
