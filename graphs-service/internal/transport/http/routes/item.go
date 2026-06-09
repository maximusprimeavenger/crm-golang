package routes

import (
	"graphs-service/internal/application/aggregator"
	linechart "graphs-service/internal/application/use_cases/line_chart"

	"github.com/gin-gonic/gin"
)

func RouteManager(e *gin.Engine, agg *aggregator.Aggregator) {
	e.GET("/line-chart/:item_id", linechart.DrawLineChartByItem(agg))
	e.GET("/line-chart", linechart.DrawLineChart(agg))
}
