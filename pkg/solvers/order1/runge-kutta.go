package order1

import (
	"github.com/willbeason/diffeq-go/pkg/equations"
)

// RungeKutta is an explicit Runge-Kutta solver for first-order differential
// equations.
type RungeKutta struct {
	// Steps are the steps of the Runge-Kutta solver, in order.
	Steps []RungeKuttaStep
}

var _ Solver = RungeKutta{}

func NewRungeKutta(steps ...RungeKuttaStep) RungeKutta {
	return RungeKutta{Steps: steps}
}

func (rk RungeKutta) Solve(eq equations.FirstOrder, t, y, h float64) float64 {
	k := make([]float64, len(rk.Steps))

	for i, step := range rk.Steps {
		ypi := 0.0
		for j, w := range step.Coefficients {
			ypi += w * k[j]
		}

		yi := y + h*ypi
		k[i] = eq(t+h*step.Node, yi)
	}

	yp := 0.0
	for i, step := range rk.Steps {
		yp += step.Weight * k[i]
	}

	return y + h*yp
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
