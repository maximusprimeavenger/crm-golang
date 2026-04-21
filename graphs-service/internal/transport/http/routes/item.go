package routes

import "github.com/gin-gonic/gin"

func RouteManager(e *gin.Engine) {
	e.GET("/line-chart/:item_id")
}
