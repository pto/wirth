package coord

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	cases := []struct {
		a, b Coordinate
		d    float64
	}{
		{Cartesian{X: 1, Y: 0},
			Cartesian{X: 0, Y: 1}, math.Sqrt(2)},
		{Cartesian{X: 1, Y: 1},
			Polar{Radius: math.Sqrt(2), Angle: math.Pi / 4}, 0},
		{Polar{Radius: math.Sqrt(2), Angle: math.Pi / 4},
			Cartesian{X: 1, Y: 1}, 0},
		{Polar{Radius: 5, Angle: math.Pi / 2},
			Polar{Radius: 1, Angle: math.Pi / 2}, 4},
		{Cartesian{X: 5, Y: 0},
			Polar{Radius: 2, Angle: 0}, 3},
		{Polar{Radius: 4, Angle: math.Pi / 2},
			Polar{Radius: 4, Angle: 0}, math.Sqrt(32)},
		{Polar{Radius: 1, Angle: 0},
			Cartesian{X: 0.5, Y: 0.5}, math.Sqrt(2) / 2},
		{Polar{Radius: math.Sqrt(0.5), Angle: math.Pi / 4},
			Cartesian{X: 1, Y: 0}, math.Sqrt(2) / 2},
	}
	for _, c := range cases {
		result := Distance(c.a, c.b)
		if math.Abs(result-c.d) > 1e-15 {
			t.Errorf("Distance(%v, %v) is %v, want %v",
				c.a, c.b, result, c.d)
		}
	}
}
