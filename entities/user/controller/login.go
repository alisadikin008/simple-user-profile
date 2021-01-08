package controller

import (
	sha "crypto/sha256"
	"encoding/hex"

	userModel "simple-user-profile/entities/user/model"
	userResponse "simple-user-profile/entities/user/response"

	_ "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type loginObject struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type blankObject struct {
}

// Login -()
func Login(Context *gin.Context) {
	var login loginObject
	var userObject userModel.User
	var blank blankObject
	Context.Bind(&login)
	//check whether email exist
	user, err := userObject.OfUsername(login.Username)
	if err != nil {
		Context.JSON(200, userResponse.DisplayResponse(err, "Username "+login.Username+" Not Found", "Error"))
		return
	}

	// check whether password match
	hash := sha.Sum256([]byte(login.Password))
	encodedPassword := hex.EncodeToString(hash[:])
	if encodedPassword != user.Password {
		Context.JSON(200, userResponse.DisplayResponse(blank, "Password Not Match", "Error"))
		return
	}

	Context.JSON(200, userResponse.DisplayResponse(user, "Login Success"))
}
