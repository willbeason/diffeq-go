package order2

import (
	"github.com/willbeason/diffeq-go/pkg/equations"
)

type Solver interface {
	Solve(eq equations.SecondOrder, t, y, yp, h float64) (float64, float64)
}

func Solve(s Solver, eq equations.SecondOrder, t0, y0, yp0, tf float64, n int) (float64, float64) {
	h := (tf - t0) / float64(n)

	t := t0
	y := y0
	yp := yp0

	for range n {
		y, yp = s.Solve(eq, t, y, yp, h)
		t += h
	}

	return y, yp
}
