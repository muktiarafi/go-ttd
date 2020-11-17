package main

import "math"

type Rectangle struct {
	Width  float64
	Heigth float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Heigth)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Heigth
}
