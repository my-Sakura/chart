package chart

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeScatterChart() {
	scatter := charts.NewScatter()

	scatter.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "scatter",
	}), charts.WithTooltipOpts(opts.Tooltip{
		Show:      true,
		Trigger:   "axis",
		TriggerOn: "mousemove",
		Formatter: "{a0}",
	}), charts.WithTooltipOpts(opts.Tooltip{
		Show:      true,
		Trigger:   "axis",
		TriggerOn: "mousemove",
		Formatter: "{a1}",
	}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}))
	scatter.SetXAxis([]string{"8:00", "10:00", "12:00", "14:00", "16:00", "18:00"}).
		AddSeries("A", generateScatterItems()).
		AddSeries("B", generateScatterItems())

	f, _ := os.Create("example/scatter.html")
	scatter.Render(f)
}

func generateScatterItems() []opts.ScatterData {
	items := make([]opts.ScatterData, 0)

	for i := 0; i < 6; i++ {
		items = append(items, opts.ScatterData{Name: "A", Value: rand.Intn(50)})
	}

	return items
}
