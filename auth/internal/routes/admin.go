package routes

import (
	"go/auth-service/internal/controllers"
	"go/auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminManager(approachingRoute *gin.Engine) {
	adminRoute := approachingRoute.Group("/admin")
	{
		adminRoute.Use(middleware.Authentification())
		adminRoute.POST("/promote/:id", controllers.PromoteAdmin())
		adminRoute.DELETE("/demote/:id", controllers.DeleteAdmin())
	}
}
