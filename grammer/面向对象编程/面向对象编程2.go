package main

import (
	"fmt"
	"math"
)

// 公共方法
type Shaper interface {
	Area() float64
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

// ----------------------------------------------------
type Circle1 struct {
	radius float64
}

func (c Circle1) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// ------------------------------------
func ComputeArea(shaper Shaper) float64 {
	return shaper.Area()
}

func main() {
	s := Square{5}
	c := Circle1{4}
	fmt.Println(ComputeArea(s)) // 输出 25
	fmt.Println(ComputeArea(c)) // 输出 50.26548245743669
}
