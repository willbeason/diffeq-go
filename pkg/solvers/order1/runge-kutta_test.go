package order1_test

import (
	"testing"

	"github.com/willbeason/diffeq-go/pkg/solvers/order1"
)

func TestRungeKutta_Euler(t *testing.T) {
	t.Parallel()

	eq := func(t, y float64) float64 { return y }

	solver := order1.NewRungeKutta(order1.Step(1.0, 0.0))

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.0),
		NewTestCase(0.0, 1.0, 0.5, 1.5),
		NewTestCase(0.0, 1.0, 0.1, 1.1),
	}

	for _, tc := range tcs {
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
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			rk4 := order1.RK4()
			err := tc.Run(rk4, eq)
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
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			rk38 := order1.RK38()
			err := tc.Run(rk38, eq)
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
			ralston := order1.Ralston()
			err := tc.Run(ralston, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
