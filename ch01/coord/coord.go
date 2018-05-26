// Package coord defines a Coordinate type and calculates distances
// between Cartesian and Polar Coordinates. This demonstrates using an
// interface to represent a discriminated union.
package coord

import (
	"math"
)

// Coordinate is a coordinate expressed in either Cartesian or Polar form.
type Coordinate interface {
	coord()
}

// Cartesian is a point expressed in Cartesian coordinates.
type Cartesian struct {
	X, Y float64
}

func (Cartesian) coord() {}

// Polar is a point expressed in Polar coordinates.
type Polar struct {
	Radius, Angle float64
}

func (Polar) coord() {}

func square(x float64) float64 {
	return x * x
}

// Distance returns the distance between Coordinates a and b.
func Distance(a, b Coordinate) float64 {
	switch a := a.(type) {
	case Cartesian:
		switch b := b.(type) {
		case Cartesian:
			return math.Sqrt(square(a.X-b.X) + square(a.Y-b.Y))
		case Polar:
			return math.Sqrt(square(a.X-b.Radius*math.Cos(b.Angle)) +
				square(a.Y-b.Radius*math.Sin(b.Angle)))
		}
	case Polar:
		switch b := b.(type) {
		case Cartesian:
			return math.Sqrt(square(a.Radius*math.Cos(a.Angle)-b.X) +
				square(a.Radius*math.Sin(a.Angle)-b.Y))
		case Polar:
			return math.Sqrt(square(a.Radius) + square(b.Radius) -
				2*a.Radius*b.Radius*math.Cos(a.Angle-b.Angle))
		}
	}
	panic("unknown Coordinate type")
}
