package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Cirlce struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Cirlce) Area() float64 {
	return math.Pi * c.radius * c.radius
}
