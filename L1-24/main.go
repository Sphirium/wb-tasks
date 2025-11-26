package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	distance := p1.Distance(*p2)
	fmt.Println(distance)
}
