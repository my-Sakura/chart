package chart

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func MakeSurfaceChart() {
	surface := charts.NewSurface3D()

	surface.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "surface",
	}), charts.WithXAxis3DOpts(opts.XAxis3D{
		Show: true,
		Name: "我是 X",
		Type: "",
		Min:  -3,
		Max:  3,
	}), charts.WithYAxis3DOpts(opts.YAxis3D{
		Show: true,
		Name: "我是 Y",
		Type: "value",
		Min:  -3,
		Max:  3,
	}), charts.WithZAxis3DOpts(opts.ZAxis3D{
		Show: true,
		Name: "我是 Z",
		Type: "value",
		Min:  -3,
		Max:  3,
	}))

	surface.AddSeries("surface3D", generateSurfaceItems())

	f, _ := os.Create("example/surface.html")
	surface.Render(f)
}

func generateSurfaceItems() []opts.Chart3DData {
	items := make([]opts.Chart3DData, 0)

	for i := 0; i < 100; i++ {
		items = append(items, opts.Chart3DData{Name: "test", Value: []interface{}{rand.Intn(3), rand.Intn(3), rand.Intn(3)}})
	}

	return items
}
