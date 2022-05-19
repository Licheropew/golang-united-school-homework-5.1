package square

import (
	"testing"
)

func TestSquare(t *testing.T) {
	p := Point{x: 0, y: 0}
	s := Square{start: p, a: 4}
	var perimeterResult uint = 16
	var areaResult uint = 16
	endPointResult := Point{x: 4, y: 4}
	t.Run("Area", func(t *testing.T) {
		funcResult := s.Area()
		if funcResult != areaResult {
			t.Errorf("square area is not correct. Got: %d. Expect: %d.\n", funcResult, areaResult)
		}
	})
	t.Run("Perimeter", func(t *testing.T) {
		funcResult := s.Perimeter()

		if funcResult != perimeterResult {
			t.Errorf("square perimeter is not correct. Got: %d. Expect: %d.\n", funcResult, perimeterResult)
		}
	})
	t.Run("End Point", func(t *testing.T) {
		funcResult := s.End()

		if funcResult != endPointResult {
			t.Errorf("square end point is not correct. Got: %d. Expect: %d.\n", funcResult, endPointResult)
		}
	})
}
