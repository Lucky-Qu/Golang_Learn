package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type location struct {
	Name string  `json:"name"`
	X    float64 `json:"Longitude"`
	Y    float64 `json:"Latitude"`
}

type gps struct {
	start location
	end   location
}

func (g gps) distance() float64 {
	return math.Sqrt((g.start.X-g.end.X)*(g.start.X-g.end.X) + (g.start.Y-g.end.Y)*(g.start.Y-g.end.Y))
}

func (l location) description() {
	x, _ := json.MarshalIndent(l, "", " ")
	y, _ := json.MarshalIndent(l, "", " ")
	fmt.Println(string(x), string(y))
}

func main() {
	var l = gps{start: location{Name: "Home", X: 0.0, Y: 0.0}, end: location{Name: "School", X: 3, Y: 4}}
	fmt.Println(l.distance())
}
