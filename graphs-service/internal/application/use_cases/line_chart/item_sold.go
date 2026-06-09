package linechart

import (
	"fmt"
	"graphs-service/internal/application/aggregator"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func DrawLineChart(agg *aggregator.Aggregator) gin.HandlerFunc {
	return func(c *gin.Context) {
		line := charts.NewLine()

		analytics := agg.GetItemAnalytics()

		xAxisSet := make(map[string]struct{})

		for _, item := range analytics {
			for day := range item.SalesByDay {
				xAxisSet[day] = struct{}{}
			}
		}

		xAxis := make([]string, 0, len(xAxisSet))
		for day := range xAxisSet {
			xAxis = append(xAxis, day)
		}
		sort.Strings(xAxis)

		for itemID, item := range analytics {
			data := make([]opts.LineData, len(xAxis))

			for i, day := range xAxis {
				data[i] = opts.LineData{
					Value: item.SalesByDay[day],
				}
			}

			line.AddSeries(
				fmt.Sprintf("Item %d", itemID),
				data,
			)
		}

		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{
				Theme: types.ThemeWesteros,
			}),
			charts.WithTitleOpts(opts.Title{
				Title:    "Sales by Item",
				Subtitle: "Line chart of sales per day",
			}),
		)

		line.SetXAxis(xAxis).
			SetSeriesOptions(charts.WithLineChartOpts(
				opts.LineChart{Smooth: opts.Bool(true)},
			))

		c.Header("Content-Type", "text/html; charset=utf-8")

		if err := line.Render(c.Writer); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
	}
}

func DrawLineChartByItem(agg *aggregator.Aggregator) gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Param("item_id")

		id, err := strconv.ParseUint(idStr, 10, 0)
		if !agg.Ready(uint(id)) {
			c.JSON(503, gin.H{"error": "aggregator is not ready"})
			return
		}
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid item_id"})
			return
		}

		salesByDay, revenueByDay, itemName := agg.GetSalesAndRevenueByID(uint(id))

		if itemName == "" {
			c.JSON(404, gin.H{"error": "item not found"})
			return
		}

		salesData, revenueData, xAxis := getSalesRevenueXaxis(salesByDay, revenueByDay)

		line := charts.NewLine()

		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{
				Theme: types.ThemeWesteros,
			}),
			charts.WithTitleOpts(opts.Title{
				Title:    "Sales by Day",
				Subtitle: fmt.Sprintf("Line chart of sales per day for item: %s", itemName),
			}),
		)

		line.SetXAxis(xAxis).
			AddSeries("Sales", salesData).
			AddSeries("Revenue", revenueData).
			SetSeriesOptions(
				charts.WithLineChartOpts(opts.LineChart{
					Smooth: opts.Bool(true),
				}),
			)

		c.Header("Content-Type", "text/html; charset=utf-8")

		if err := line.Render(c.Writer); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}

func getSalesRevenueXaxis(salesByDay map[string]int, revenueByDay map[string]float64) ([]opts.LineData, []opts.LineData, []string) {
	xAxis := []string{}
	salesData := []opts.LineData{}
	revenueData := []opts.LineData{}
	for date := range salesByDay {
		xAxis = append(xAxis, date)
	}
	sort.Strings(xAxis)
	for _, date := range xAxis {
		salesData = append(salesData, opts.LineData{Value: salesByDay[date]})
		revenueData = append(revenueData, opts.LineData{Value: revenueByDay[date]})
	}
	return salesData, revenueData, xAxis
}
