//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	"github.com/tvllyit/t38c"
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

	// To set a field when setting an object:
	if err := tile38.Keys.Set("fleet", "truck1").
		Point(33.5123, -112.2693).
		Field("speed", 90).
		Field("age", 21).
		Do(context.TODO()); err != nil {
		log.Fatalf("unable to set field %v", err)
	}

	// To set a field when an object already exists:
	if err := tile38.Keys.FSet("fleet", "truck1").Field("speed", 90).Do(context.TODO()); err != nil {
		log.Fatalf("faield to fset field %v", err)
	}
}
