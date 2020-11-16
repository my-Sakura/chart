package chart

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func MakeLineChart() {
	line := charts.NewLine()

	title := opts.Title{Title: "本地镜像数量"}
	legend := opts.Legend{Show: true}
	xaxis := opts.XAxis{Name: "时间"}
	toolTip := opts.Tooltip{Show: true, Trigger: "axis", TriggerOn: "mousemove", Formatter: "{c}"}

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(title),
		charts.WithLegendOpts(legend),
		charts.WithXAxisOpts(xaxis),
		charts.WithTooltipOpts(toolTip),
	)

	line.SetXAxis([]string{"1", "2", "3", "4", "5"}).
		AddSeries("A", generateLineItems(), charts.WithLineStyleOpts(opts.LineStyle{Color: "green"}),
			charts.WithLineChartOpts(opts.LineChart{Smooth: true})).
		AddSeries("B", generateLineItems(), charts.WithLineStyleOpts(opts.LineStyle{Color: "blue"}),
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, _ := os.Create("example/line.html")
	line.Render(f)
}

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)

	for i := 0; i < 5; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(10)})
	}
	return items
}
