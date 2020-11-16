package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeGeoChart() {
	geo := charts.NewGeo()

	geo.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:  "geo",
		Bottom: "0",
		Right:  "40%",
	}), charts.WithTooltipOpts(opts.Tooltip{
		Show:      true,
		Trigger:   "item",
		TriggerOn: "mousemove",
		Formatter: "{c}",
	}), charts.WithLegendOpts(opts.Legend{
		Show: true,
	}))

	geo.AddSeries("a", "b", generateGeoItems())

	f, _ := os.Create("example/geo.html")
	geo.Render(f)
}

func generateGeoItems() []opts.GeoData {
	items := make([]opts.GeoData, 0)

	items = append(items, opts.GeoData{Name: "a", Value: "hh"})

	return items
}
