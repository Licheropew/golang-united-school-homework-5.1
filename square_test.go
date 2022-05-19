package square

import (
	"testing"
)

func TestSquare(t *testing.T) {
	p := Point{x: 0, y: 0}
	s := Square{start: p, a: 4}
	t.Run("Area", func(t *testing.T) {
		funcResult := s.Area()
		result := s.a * s.a
		if funcResult != result {
			t.Errorf("Real result not equal expected result. %d != %d\n", funcResult, result)
		}
	})
	t.Run("Perimeter", func(t *testing.T) {
		funcResult := s.Perimeter()
		result := 4 * s.a
		if funcResult != result {
			t.Errorf("Real result not equal expected result. %d != %d\n", funcResult, result)
		}
	})
	t.Run("End Point", func(t *testing.T) {
		funcResult := s.End()
		result := Point{x: s.start.x + int(s.a), y: s.start.y + int(s.a)}
		if funcResult != result {
			t.Errorf("Real result not equal expected result. %d != %d\n", funcResult, result)
		}
	})
}
