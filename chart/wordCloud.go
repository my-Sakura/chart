package chart

import (
	"context"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Events() {
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	events, _ := cli.Events(context.Background(), types.EventsOptions{})

	for {
		messages := <-events
		data := messages.Action + " " + messages.Type
		MakeWordCloudChart(data)
	}
}

func MakeWordCloudChart(data string) {
	wordCloud := charts.NewWordCloud()

	title := opts.Title{
		Title: "docker Events",
	}

	wordCloud.SetGlobalOptions(charts.WithTitleOpts(title))
	wordCloud.AddSeries("events", a(data), charts.WithWorldCloudChartOpts(opts.WordCloudChart{Shape: "diamond"}))

	f, _ := os.Create("example/wordCloud.html")
	wordCloud.Render(f)
}

func a(data string) []opts.WordCloudData {
	items := make([]opts.WordCloudData, 0)

	items = append(items, opts.WordCloudData{Name: "a", Value: data})

	return items
}
