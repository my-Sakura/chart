package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeGaugeChart() {
	gauge := charts.NewGauge()

	gauge.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "gauge",
	}))

	gauge.AddSeries("a", generateGaugeItems())

	f, _ := os.Create("example/gauge.html")
	gauge.Render(f)
}

func generateGaugeItems() []opts.GaugeData {
	items := make([]opts.GaugeData, 0)

	items = append(items, opts.GaugeData{Name: "b", Value: "10"})

	return items
}
