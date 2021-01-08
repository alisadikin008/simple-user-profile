package controller

import (
	userModel "simple-user-profile/entities/user/model"
	userResponse "simple-user-profile/entities/user/response"

	_ "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Register -()
func Register(Context *gin.Context) {
	var userObject userModel.User
	Context.Bind(&userObject)
	//check whether email exist
	_, exist := userObject.OfEmail(userObject.Email)
	if exist == nil {
		Context.JSON(200, userResponse.DisplayResponse(userObject.Email, "Email "+userObject.Email+" Already Exist", "Error"))
		return
	}

	//check whether username exist
	_, none := userObject.OfUsername(userObject.Username)
	if none == nil {
		Context.JSON(200, userResponse.DisplayResponse(userObject.Username, "Username "+userObject.Username+" Already Exist", "Error"))
		return
	}

	data := userObject.PostData(userObject)
	Context.JSON(200, userResponse.DisplayResponse(data, "Successfully Save the Data"))
}
