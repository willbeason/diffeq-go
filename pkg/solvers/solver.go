package solvers

import (
	"github.com/willbeason/diffeq-go/pkg/equations"
)

type Solver interface {
	Solve(order equations.FirstOrder, t, y, h float64) float64
}

// Solve uses Solver to determine the value of y for the provided first-order differential equation
// at time tf, given initial condition y(t0) = y0 and the desired number of steps.
func Solve(s Solver, eq equations.FirstOrder, t0, y0, tf float64, n int) float64 {
	h := (tf - t0) / float64(n)

	t := t0
	y := y0

	for i := 0; i < n; i++ {
		y = s.Solve(eq, t, y, h)
		t0 += h
	}

	return y
}
