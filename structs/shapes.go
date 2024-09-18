package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter += 2 * (rectangle.Width + rectangle.Height)
	return
}

func Area(rectangle Rectangle) (area float64) {
	area = rectangle.Width * rectangle.Height
	return
}
