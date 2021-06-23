package models

import (
	"math"
)

type DuffingOscillator struct {
	// Delta is the drag force.
	Delta float64

	// Alpha is the spring constant.
	Alpha float64

	// Beta is the nonlinear force response.
	Beta float64

	// Gamma is the magnitude of the driving force.
	Gamma float64

	// Frequency is the frequency of the driving force.
	Frequency float64
}

func (o DuffingOscillator) Acceleration(t, y, yp float64) float64 {
	return -o.Delta*yp - o.Alpha*y - o.Beta*y*y*y + o.Gamma*math.Cos(o.Frequency*t)
}
