package solvers_test

import (
	"testing"

	"github.com/willbeason/diffeq-go/pkg/solvers"
)

func TestRungeKutta_Euler(t *testing.T) {
	t.Parallel()

	eq := func(t, y float64) float64 { return y }

	solver := solvers.NewRungeKutta(solvers.Step(1.0, 0.0))

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.0),
		NewTestCase(0.0, 1.0, 0.5, 1.5),
		NewTestCase(0.0, 1.0, 0.1, 1.1),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()

			err := tc.Run(solver, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_RK4(t *testing.T) {
	t.Parallel()

	eq := func(t, y float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.708),
		NewTestCase(0.0, 1.0, 0.5, 1.648),
		NewTestCase(0.0, 1.0, 0.1, 1.105),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(solvers.RK4, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_RK38(t *testing.T) {
	t.Parallel()

	eq := func(t, y float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.708),
		NewTestCase(0.0, 1.0, 0.5, 1.648),
		NewTestCase(0.0, 1.0, 0.1, 1.105),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(solvers.RK38, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_Ralston(t *testing.T) {
	t.Parallel()

	eq := func(t, y float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.708),
		NewTestCase(0.0, 1.0, 0.5, 1.648),
		NewTestCase(0.0, 1.0, 0.1, 1.105),
	}

	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(solvers.Ralston, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
