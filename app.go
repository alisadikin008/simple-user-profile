package main

import (
	config "simple-user-profile/config"
	userController "simple-user-profile/entities/user/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(appMode())
	generalVersion := router.Group("/api/v1")
	{
		generalVersion.POST("/registration", userController.Register)
		generalVersion.POST("/login", userController.Login)

		generalVersion.Use(tokenAuthMiddleware())
		users := generalVersion.Group("/users")
		{
			users.GET("/", userController.GetUsers)
			users.GET("/:id", userController.GetOne)
			users.POST("/", userController.PostUsers)
			//users.PUT("/:id", userController.PutOne)
		}
	}

	return router
}

// respondWithError -()
func respondWithError(Context *gin.Context, code int, message interface{}) {
	Context.AbortWithStatusJSON(code, gin.H{"error": message})
}

// appMode -()
func appMode() gin.HandlerFunc {
	return func(Context *gin.Context) {
		mode := config.LoadConfiguration()
		if mode != "production" {
			respondWithError(Context, 401, "Under Development")
			return
		}

		Context.Next()
	}

}

// tokenAuthMiddleware -()
func tokenAuthMiddleware() gin.HandlerFunc {
	return func(Context *gin.Context) {
		appKey := config.GetCredential("key")
		appToken := config.GetCredential("token")

		headerKey := Context.GetHeader("App-Key")
		headerToken := Context.GetHeader("X-Lemo-Token")
		if headerKey != appKey {
			respondWithError(Context, 401, "Invalid API Key")
			return
		}
		if headerToken != appToken {
			respondWithError(Context, 401, "Invalid API Token")
			return
		}

		Context.Next()
	}
}
