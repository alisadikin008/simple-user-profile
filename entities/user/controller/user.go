package controller

/*
	created by Ali Sadikin
	this is user controller that invoked by user route in main app

	## function list
	GetUsers()
	GetOne()
	PostUsers()
	PutOne()
	DeleteOne()
*/

import (
	"strconv"

	userModel "simple-user-profile/entities/user/model"
	userResponse "simple-user-profile/entities/user/response"

	"github.com/gin-gonic/gin"
)

var user userModel.User

// GetUsers -()
func GetUsers(Context *gin.Context) {
	queryParams := Context.Request.URL.Query()
	page, _ := strconv.Atoi(Context.Query("page"))
	limit, _ := strconv.Atoi(Context.Query("limit"))
	data, result := user.OfUsers(queryParams, page, limit)
	if result == nil {
		Context.JSON(200, userResponse.DisplayResponse(data, "Data Not Found"))
		return
	}

	Context.JSON(200, userResponse.DisplayResponse(data, "Data Collected"))
}

// GetOne -()
func GetOne(Context *gin.Context) {
	id, _ := strconv.Atoi(Context.Param("id"))
	data, err := user.OfID(id)
	if err != nil {
		Context.JSON(200, userResponse.DisplayResponse(err, "Data Not Found"))
		return
	}

	Context.JSON(200, userResponse.DisplayResponse(data, "Data Collected"))
}

// PostUsers -()
func PostUsers(Context *gin.Context) {
	Context.Bind(&user)
	//check whether email exist
	_, exist := user.OfEmail(user.Email)
	if exist == nil {
		Context.JSON(200, userResponse.DisplayResponse(user.Email, "Email "+user.Email+" Already Exist", "Error"))
		return
	}

	//check whether username exist
	_, none := user.OfUsername(user.Username)
	if none == nil {
		Context.JSON(200, userResponse.DisplayResponse(user.Username, "Username "+user.Username+" Already Exist", "Error"))
		return
	}

	data := user.PostData(user)
	Context.JSON(200, userResponse.DisplayResponse(data, "Successfully Save the Data"))
}

// PutOne -()
func PutOne(Context *gin.Context) {
	Context.Bind(&user)
	id, _ := strconv.Atoi(Context.Param("id"))
	_, err := user.OfID(id)
	if err != nil {
		Context.JSON(200, userResponse.DisplayResponse(err, "User Not Found"))
		return
	}

	data := user.PutOne(id, user)
	response := userResponse.DisplayResponse(data, "Successfully Saved the Data")
	Context.JSON(200, response)

}
