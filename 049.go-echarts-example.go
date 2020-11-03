package main

import (
	"io"
	"math/rand"
	"os"

	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		mapBase(),
		mapShowLabel(),
		mapVisualMap(),
		mapRegion(),
		mapTheme(),
	)

	f, err := os.Create("map.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))

	fs := http.FileServer(http.Dir("./"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}

var (
	baseMapData = []opts.MapData{
		{"北京", float64(rand.Intn(150))},
		{"上海", float64(rand.Intn(150))},
		{"广东", float64(rand.Intn(150))},
		{"辽宁", float64(rand.Intn(150))},
		{"山东", float64(rand.Intn(150))},
		{"山西", float64(rand.Intn(150))},
		{"陕西", float64(rand.Intn(150))},
		{"新疆", float64(rand.Intn(150))},
		{"内蒙古", float64(rand.Intn(150))},
	}

	guangdongMapData = map[string]float64{
		"深圳市": float64(rand.Intn(150)),
		"广州市": float64(rand.Intn(150)),
		"湛江市": float64(rand.Intn(150)),
		"汕头市": float64(rand.Intn(150)),
		"东莞市": float64(rand.Intn(150)),
		"佛山市": float64(rand.Intn(150)),
		"云浮市": float64(rand.Intn(150)),
		"肇庆市": float64(rand.Intn(150)),
		"梅州市": float64(rand.Intn(150)),
	}
)

func generateMapData(data map[string]float64) (items []opts.MapData) {
	items = make([]opts.MapData, 0)
	for k, v := range data {
		items = append(items, opts.MapData{Name: k, Value: v})
	}
	return
}

func mapBase() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Map-example",
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Map-show-label",
		}),
	)

	mc.AddSeries("map", baseMapData).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return mc
}

func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Map-VisualMap",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapRegion() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("广东")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Map-religion-Guangdong",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}},
		}),
	)

	mc.AddSeries("map", generateMapData(guangdongMapData))
	return mc
}

func mapTheme() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "macarons",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Map-theme",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        150,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

type MapExamples struct{}

func (MapExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		mapBase(),
		mapShowLabel(),
		mapVisualMap(),
		mapRegion(),
		mapTheme(),
	)

	f, err := os.Create("examples/html/map.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
