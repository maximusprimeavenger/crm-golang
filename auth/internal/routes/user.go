package routes

import (
	"go/auth-service/internal/controllers"
	"go/auth-service/internal/helpers"
	"go/auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserManager(aproachingRoute *gin.Engine) {
	userRoutes := aproachingRoute.Group("/users")
	{
		userRoutes.Use(middleware.Authentification())
		userRoutes.GET("", controllers.GetAll())
		userRoutes.GET("/:id", controllers.GetUser())
	}
	userNotify := aproachingRoute.Group("/name-taking")
	{
		userNotify.POST("", helpers.TakeName())
	}
}
