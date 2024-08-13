package structs

import "testing"

// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(Rectangle{10.0, 10.0})
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		CheckAreaHelper(t, rectangle, 72.0)

	})

	t.Run("Circles", func(t *testing.T) {
		circle := Cirlce{10}
		CheckAreaHelper(t, circle, 314.1592653589793)
	})
}

func CheckAreaHelper(t *testing.T, shape Shape, expected float64) {
	t.Helper()

	got := shape.Area()
	want := expected

	if got != want {
		t.Errorf("got %g want %g", want, got)
	}
}
