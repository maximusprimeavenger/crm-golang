package routes

import (
	controllers "go/auth-service/internal/controllers"
	"go/auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthintificateRoute(approachingRoute *gin.Engine) {
	authRoutes := approachingRoute.Group("/users")
	{
		authRoutes.POST("/login", controllers.Login())
		authRoutes.POST("/signup", controllers.SignUp())
	}
	appRoutes := approachingRoute.Group("/validate-token")
	{
		appRoutes.POST("", middleware.AdminRoute())
		appRoutes.POST("/id-taking", middleware.TakeIdFromToken())
	}
}
