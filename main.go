package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

type square struct {
	d float64
}

func (s *square) area() float64 {
	return s.d * s.d
}

type geo interface {
	area() float64
}

func whatYourArea(g geo) {
	fmt.Println("Your area is ", g.area())
}

func main() {
	c := circle{r: 7}
	s := square{d: 10}

	whatYourArea(&c)
	whatYourArea(&s)
}
