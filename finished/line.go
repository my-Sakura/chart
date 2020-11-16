package main

import (
	"context"
	"os"
	"strconv"
	"time"

	api "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func main() {
	ListImage()
}

func ListImage() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	for {
		images, err := cli.ImageList(context.Background(), api.ImageListOptions{All: true})
		if err != nil {
			panic(err)
		}

		MakeLineChart(len(images))
	}
}

func MakeLineChart(count int) {
	line := charts.NewLine()

	title := opts.Title{Title: "本地镜像数量"}
	legend := opts.Legend{Show: true}
	xaxis := opts.XAxis{Name: "时间"}
	yaxis := opts.YAxis{Name: "镜像数量"}
	toolTip := opts.Tooltip{Show: true, Trigger: "axis", TriggerOn: "mousemove", Formatter: "{c}"}

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(title),
		charts.WithLegendOpts(legend),
		charts.WithXAxisOpts(xaxis),
		charts.WithYAxisOpts(yaxis),
		charts.WithTooltipOpts(toolTip),
	)

	a := strconv.Itoa(time.Now().Hour())
	b := strconv.Itoa(time.Now().Hour() + 1)
	c := strconv.Itoa(time.Now().Hour() + 2)
	d := strconv.Itoa(time.Now().Hour() + 3)
	e := strconv.Itoa(time.Now().Hour() + 4)
	f := strconv.Itoa(time.Now().Hour() + 5)
	line.SetXAxis([]string{a, b, c, d, e, f}).
		AddSeries("A", generateLineItems(count), charts.WithLineStyleOpts(opts.LineStyle{Color: "green"}),
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	file, _ := os.Create("example/line.html")
	line.Render(file)
}

func generateLineItems(count int) []opts.LineData {
	items := make([]opts.LineData, 0)

	for i := 0; i < 6; i++ {
		items = append(items, opts.LineData{Value: count})
	}

	return items
}
