package estimators

import (
	"testing"
)

func TestRungeKutta_Euler(t *testing.T) {
	eq := func(t, y float64) float64 { return y }

	solver := NewRungeKutta(Step(1.0, 0.0))

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.0),
		NewTestCase(0.0, 1.0, 0.5, 1.5),
		NewTestCase(0.0, 1.0, 0.1, 1.1),
	}

	for _, tc := range tcs {
		tc.Run(t, solver, eq)
	}
}

func TestRungeKutta_RK4(t *testing.T) {
	eq := func(t, y float64) float64 { return y }

	solver := NewRungeKutta(
		Step(1.0/6.0, 0.0),
		Step(1.0/3.0, 0.5, 0.5),
		Step(1.0/3.0, 0.5, 0.0, 0.5),
		Step(1.0/6.0, 1.0, 0.0, 0.0, 1.0),
	)

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.708),
		NewTestCase(0.0, 1.0, 0.5, 1.648),
		NewTestCase(0.0, 1.0, 0.1, 1.105),
	}

	for _, tc := range tcs {
		tc.Run(t, solver, eq)
	}
}

func TestRungeKutta_RK38(t *testing.T) {
	eq := func(t, y float64) float64 { return y }

	solver := NewRungeKutta(
		Step(1.0/8.0, 0.0),
		Step(3.0/8.0, 1.0/3.0, 1.0/3.0),
		Step(3.0/8.0, 2.0/3.0, -1.0/3.0, 1.0),
		Step(1.0/8.0, 1.0, 1.0, -1.0, 1.0),
	)

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.708),
		NewTestCase(0.0, 1.0, 0.5, 1.648),
		NewTestCase(0.0, 1.0, 0.1, 1.105),
	}

	for _, tc := range tcs {
		tc.Run(t, solver, eq)
	}
}
