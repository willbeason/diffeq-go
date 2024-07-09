package order1

import "github.com/willbeason/diffeq-go/pkg/equations"

type RungeKuttaSolver struct {
	rk RungeKutta
	k  []float64
}

var _ Solver = (*RungeKuttaSolver)(nil)

func NewRungeKuttaSolver(rk RungeKutta) *RungeKuttaSolver {
	return &RungeKuttaSolver{
		rk: rk,
		k:  make([]float64, len(rk.Steps)),
	}
}

func (rks *RungeKuttaSolver) Solve(eq equations.FirstOrder, t, y, h float64) float64 {
	for i, step := range rks.rk.Steps {
		ypi := 0.0
		for j, w := range step.Coefficients {
			ypi += w * rks.k[j]
		}

		yi := y + h*ypi
		rks.k[i] = eq(t+h*step.Node, yi)
	}

	yp := 0.0
	for i, step := range rks.rk.Steps {
		yp += step.Weight * rks.k[i]
	}

	return y + h*yp
}
