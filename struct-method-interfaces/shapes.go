package main

type Rectangle struct {
	Width  float64
	Heigth float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Heigth)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Heigth
}
