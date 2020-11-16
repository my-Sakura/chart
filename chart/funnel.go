package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeFunnelChart() {
	funnel := charts.NewFunnel()

	funnel.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "funnel"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Trigger:   "item",
			TriggerOn: "mousemove",
			Formatter: "{c}",
		}),
	)

	funnel.AddSeries("a", generateFunnelItems())

	f, _ := os.Create("example/funnel.html")
	funnel.Render(f)
}

func generateFunnelItems() []opts.FunnelData {
	items := make([]opts.FunnelData, 0)

	items = append(items, opts.FunnelData{Name: "hahaha", Value: "10"})

	return items
}
