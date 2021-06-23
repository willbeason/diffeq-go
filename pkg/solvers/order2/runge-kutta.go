package order2

import (
	"github.com/willbeason/diffeq-go/pkg/equations"
)

type RungeKutta struct {
	Steps []RungeKuttaStep
}

var _ Solver = RungeKutta{}

func NewRungeKutta(steps ...RungeKuttaStep) RungeKutta {
	return RungeKutta{Steps: steps}
}

func (rk RungeKutta) Solve(eq equations.SecondOrder, t, y, yp, h float64) (float64, float64) {
	k := make([]float64, len(rk.Steps))
	yps := make([]float64, len(rk.Steps))

	for i, step := range rk.Steps {
		ypp := 0.0
		for j, w := range step.Coefficients {
			ypp += w * k[j]
		}

		ypi := yp + h*ypp
		yps[i] = ypi

		yi := y + h*step.Node*(yp+ypi)/2
		k[i] = eq(t+h*step.Node, yi, ypi)
	}

	ypp := 0.0
	for i, step := range rk.Steps {
		ypp += step.Weight * k[i]
	}

	ypt := 0.0
	for i, step := range rk.Steps {
		ypt += step.Weight * yps[i]
	}

	ypf := yp + h*ypp
	yf := y + h*ypt

	return yf, ypf
}

// RungeKuttaStep is a step in a Runge-Kutta solver.
type RungeKuttaStep struct {
	// Weight is the weight given to this step's calculated slope in the final
	// calculation.
	Weight float64

	// Node is the fraction of a the estimated time step forward to be estimated.
	Node float64

	// Coefficients are the weights of previous steps to include.
	Coefficients []float64
}

func Step(b, c float64, coefficients ...float64) RungeKuttaStep {
	sumCoefficients := 0.0
	for _, w := range coefficients {
		sumCoefficients += w
	}

	return RungeKuttaStep{
		Weight:       b,
		Node:         c,
		Coefficients: coefficients,
	}
}
