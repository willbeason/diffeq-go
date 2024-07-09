package order2

import "github.com/willbeason/diffeq-go/pkg/equations"

type RungeKuttaSolver struct {
	rk  RungeKutta
	k   []float64
	yps []float64
}

var _ Solver = (*RungeKuttaSolver)(nil)

func NewRungeKuttaSolver(rk RungeKutta) *RungeKuttaSolver {
	return &RungeKuttaSolver{
		rk:  rk,
		k:   make([]float64, len(rk.Steps)),
		yps: make([]float64, len(rk.Steps)),
	}
}

func (rks *RungeKuttaSolver) Solve(eq equations.SecondOrder, t, y, yp, h float64) (float64, float64) {
	for i, step := range rks.rk.Steps {
		ypp := 0.0
		for j, w := range step.Coefficients {
			ypp += w * rks.k[j]
		}

		ypi := yp + h*ypp
		rks.yps[i] = ypi

		yi := y + step.Node*(h/6.0)*(h*step.Node*rks.k[0]+4*yp+2*ypi)
		rks.k[i] = eq(t+h*step.Node, yi, ypi)
	}

	ypp := 0.0
	for i, step := range rks.rk.Steps {
		ypp += step.Weight * rks.k[i]
	}

	ypf := yp + h*ypp

	ypt := 0.0
	for i, step := range rks.rk.Steps {
		ypt += step.Weight * rks.yps[i]
	}
	yf := y + h*ypt

	return yf, ypf
}
