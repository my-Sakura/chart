package chart

import (
	"os"
	"test/docker"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakePieChart() {
	pie := charts.NewPie()

	title := opts.Title{
		Title:  "运行容器镜像扇形图",
		Bottom: "0",
		Right:  "40%",
	}

	toolTip := opts.Tooltip{
		Show:      true,
		Trigger:   "item",
		TriggerOn: "mousemove",
		Formatter: "precent: {d}",
	}

	legend := opts.Legend{
		Show: true,
	}

	pie.SetGlobalOptions(charts.WithTitleOpts(title), charts.WithTooltipOpts(toolTip), charts.WithLegendOpts(legend))

	pie.AddSeries("dockerImage", generatePieItems())

	f, _ := os.Create("example/pie.html")
	pie.Render(f)
}

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)

	for image, num := range docker.Images {
		items = append(items, opts.PieData{Name: image, Value: num, Label: &opts.Label{Position: "top", Formatter: "{c}"}})
	}

	return items
}
