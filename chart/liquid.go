package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeLiquidChart() {
	liquid := charts.NewLiquid()

	liquid.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:  "liquid",
		Bottom: "50",
		Right:  "420",
	}), charts.WithLegendOpts(opts.Legend{
		Show:       true,
		ItemWidth:  35,
		ItemHeight: 20,
		Right:      "44%",
		Top:        "5%",
		TextStyle:  &opts.TextStyle{FontSize: 20, FontStyle: "normal"},
	}), charts.WithToolboxOpts(opts.Toolbox{
		Show: true,
		Feature: &opts.ToolBoxFeature{SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{Show: true, Type: "png", Name: "liquid", Title: "download"},
			DataView: &opts.ToolBoxFeatureDataView{Show: true, Title: "view", Lang: []string{"data view", "turn off", "refresh"}, BackgroundColor: "green"},
		}}))

	liquid.AddSeries("liquid", generateLiquidItems(), charts.WithLabelOpts(opts.Label{Show: true, Position: "inside", Formatter: "{c}"}), charts.WithLiquidChartOpts(opts.LiquidChart{IsShowOutline: true, IsWaveAnimation: false}))
	f, _ := os.Create("example/liquid.html")
	liquid.Render(f)
}

func generateLiquidItems() []opts.LiquidData {
	items := make([]opts.LiquidData, 0)

	items = append(items, opts.LiquidData{Name: "liquid", Value: "30%"})

	return items
}
