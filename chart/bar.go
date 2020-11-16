package chart

import (
	"os"

	"test/docker"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeBarChart() {
	bar := charts.NewBar()

	title := opts.Title{
		Title:  "本地镜像大小",
		Bottom: "0",
		Right:  "425",
	}

	toolTip := opts.Tooltip{
		Show:      true,
		Trigger:   "axis",
		TriggerOn: "mousemove",
		Formatter: "{c}",
	}

	xaxis := opts.XAxis{Name: "imageID"}
	yaxis := opts.YAxis{Name: "ImageSize(单位 b)"}
	bar.SetGlobalOptions(charts.WithTitleOpts(title), charts.WithTooltipOpts(toolTip),
		charts.WithXAxisOpts(xaxis), charts.WithYAxisOpts(yaxis))

	bar.SetXAxis(docker.ImageSlice).
		AddSeries("LocalImages", generateBarItems())

	f, _ := os.Create("example/bar.html")
	bar.Render(f)
}

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)

	for _, imageSize := range docker.LocalImages {
		items = append(items, opts.BarData{Value: imageSize})
	}

	return items
}
