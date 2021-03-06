package geometry

import (
	"math"
	"fmt"
)

type Angle struct {
	Rads float64
}

// Returns the equivalent value between -math.Pi and +math.Pi
func normalizeRads(rads float64) float64 {
	r := math.Mod(rads, math.Pi * 2.0)

	if r > math.Pi {
		r -= math.Pi * 2.0
	}

	return r
}

func (a *Angle) Radians() float64 {
	return normalizeRads(a.Rads)
}

func (a *Angle) Degrees() float64 {
	return 360.0 * (a.Radians() / (2.0 * math.Pi))
}

func (a *Angle) Perpendicular() *Angle {
	return &Angle{
		Rads: normalizeRads(a.Rads + (math.Pi / 2.0)),
	}
}

func (a *Angle) Opposite() *Angle {
	return &Angle{
		Rads: normalizeRads(a.Rads + math.Pi),
	}
}

func (a *Angle) Neg() *Angle {
	return &Angle{
		Rads: -a.Rads,
	}
}

func (a *Angle) Add(other *Angle) *Angle {
	return &Angle{
		Rads: a.Rads + other.Rads,
	}
}

func (a *Angle) Subtract(other *Angle) *Angle {
	return &Angle{
		Rads: a.Rads - other.Rads,
	}
}

func (a *Angle) Divide(x float64) *Angle {
	return &Angle{
		Rads: a.Rads / x,
	}
}

func (a *Angle) Multiply(x float64) *Angle {
	return &Angle{
		Rads: a.Rads * x,
	}
}

func (a *Angle) Sin() float64 {
	return math.Sin(a.Radians())
}

func (a *Angle) Cos() float64 {
	return math.Cos(a.Radians())
}

func (a *Angle) Tan() float64 {
	return math.Tan(a.Radians())
}

func (a *Angle) Asin() float64 {
	return math.Asin(a.Radians())
}

func (a *Angle) Acos() float64 {
	return math.Acos(a.Radians())
}

func (a *Angle) Atan() float64 {
	return math.Atan(a.Radians())
}

func (a *Angle) String() string {
	measure := fmt.Sprintf("%.2f", a.Rads)

	switch a.Rads {
	case math.Pi:
		measure = "π"
	case math.Pi * 1.5:
		measure = "3π/2"
	case math.Pi * 2.0:
		measure = "2π"
	case math.Pi / 2.0:
		measure = "π/2"
	case math.Pi / 4.0:
		measure = "π/4"
	case math.Pi / 8.0:
		measure = "π/8"
	case math.Pi / 16.0:
		measure = "π/16"
	}

	return fmt.Sprintf("%s radians", measure)
}

// Angles are "equivalent" if they are the same direction even if the true Radians are different.
// For example, 0 rads is equivalent to 2π rads is equivalent to -2π rads.
func (a *Angle) Equivalent(o *Angle) bool {
	return a.Radians() == o.Radians()
}
