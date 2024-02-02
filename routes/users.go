package routes

import (
	"net/http"
	"restwithjwt/models"
	"restwithjwt/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		print(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not login user"})
		return
	}

	token, err := utils.GenerateToken(user.EMAIL, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
