package order1

import "github.com/willbeason/diffeq-go/pkg/equations"

type RungeKuttaInline struct{}

var _ Solver = RungeKuttaInline{}

func (r RungeKuttaInline) Solve(eq equations.FirstOrder, t, y, h float64) float64 {
	// Step 1
	k1 := eq(t, y)

	// Step 2
	k2 := eq(t+h/2.0, y+(h/2.0)*k1)

	// Step 3
	k3 := eq(t+h/2.0, y+(h/2.0)*k2)

	// Step 4
	k4 := eq(t+h, y+h*k3)

	return y + (h/6.0)*(k1+2*k2+2*k3+k4)
}
