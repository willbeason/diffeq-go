package order2

import (
	"github.com/willbeason/diffeq-go/pkg/equations"
)

type Euler struct {}

var _ Solver = Euler{}

func (e Euler) Solve(eq equations.SecondOrder, t0, y0, yp0, h float64) (float64, float64) {
	k := eq(t0, y0, yp0)

	ypf := yp0 + h*k
	yf := y0 + (yp0 + ypf)*h/2

	return yf, ypf
}
