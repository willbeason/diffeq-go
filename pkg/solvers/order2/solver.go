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

	for i := 0; i < n; i++ {
		y, yp = s.Solve(eq, t, y, yp0, h)
		t0 += h
	}

	return y, yp
}
