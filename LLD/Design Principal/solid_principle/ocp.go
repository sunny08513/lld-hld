package solidprinciple

import "math"

// Suppose we have a simple function that calculates the area of a rectangle:
type Rectangle struct {
	Width  float64
	Height float64
}

func Area1(rectangle *Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

/*If we need to add support for calculating the area of a circle,
the current implementation would require modification:*/

type Circle struct {
	Radius float64
}

// Area2 after modification of area1
func Area2(shape interface{}) float64 {
	switch s := shape.(type) {
	case *Rectangle:
		return s.Width * s.Height
	case *Circle:
		return math.Pi * math.Pow(s.Radius, 2)
	default:
		return 0
	}
}

/*To follow the Open/Closed Principle,
we can define an interface and implement it for each shape:*/

type Shape interface {
	Area() float64
}

// Implement area for circle
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Implement area for Rectangle
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
