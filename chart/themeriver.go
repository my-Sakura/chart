package chart

import (
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeThemeRiverChart() {
	themeRiver := charts.NewThemeRiver()

	themeRiver.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "themeriver",
	}), charts.WithLegendOpts(opts.Legend{
		Show: true,
	}), charts.WithSingleAxisOpts(opts.SingleAxis{
		Type: "time",
	}), charts.WithTooltipOpts(opts.Tooltip{
		Show:      true,
		Trigger:   "item",
		TriggerOn: "mousemove",
		Formatter: "{c}",
	}),
	)

	themeRiver.AddSeries("a", generateThemeRiverItemsOne()).
		AddSeries("b", generateThemeRiverItemsTwo())

	f, _ := os.Create("example/themeriver.html")
	themeRiver.Render(f)
}

func generateThemeRiverItemsOne() []opts.ThemeRiverData {
	items := make([]opts.ThemeRiverData, 0)

	rand.Seed(time.Now().Unix())

	items = append(items, opts.ThemeRiverData{Date: "abc", Name: "hhh", Value: 10.1})

	return items
}

func generateThemeRiverItemsTwo() []opts.ThemeRiverData {
	items := make([]opts.ThemeRiverData, 0)

	rand.Seed(time.Now().Unix())

	items = append(items, opts.ThemeRiverData{Date: "abc", Name: "heihei", Value: 10.2})

	return items
}
