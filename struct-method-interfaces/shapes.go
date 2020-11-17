package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Heigth float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Heigth
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Heigth)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Heigth
}
