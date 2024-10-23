package main

import (
	"context"
	"log"

	"github.com/tvllyit/t38c"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
)

func main() {
	tile38, err := t38c.New(t38c.Config{
		Address: "localhost:9851",
		Debug:   true,
	})
	if err != nil {
		log.Fatalf("unable to connect to tile38 %v", err)
	}
	defer tile38.Close()
	g := geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{33.5123, -112.2693})
	feature := &geojson.Feature{Geometry: g}
	feature.Properties = map[string]interface{}{
		"speed": 55,
		"name":  "Carol",
		"age":   "23",
	}
	g2 := geom.NewPoint(geom.XY).MustSetCoords([]float64{33.553653, -112.112222})
	f2 := &geojson.Feature{Geometry: g2}
	f2.Properties = map[string]interface{}{
		"speed": 40,
		"name":  "Andy",
		"age":   "25",
	}
	tile38.Keys.Set("fleet", "carol").Feature(feature).Do(context.Background())
	tile38.Keys.Set("fleet", "andy").Feature(f2).Do(context.Background())

	// references: https://tile38.com/topics/filter-expressions
	tile38.Search.Scan("fleet").RawQuery("properties.age == 25 && properties.speed > 50").Do(context.Background())
}
