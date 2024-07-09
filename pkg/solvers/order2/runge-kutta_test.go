package order2_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/solvers/order2"
)

func TestRungeKutta_Euler(t *testing.T) {
	t.Parallel()

	var eq equations.SecondOrder = func(t, y, yp float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 1.0, 2.5, 2.0),
		NewTestCase(0.0, 1.0, 1.0, 0.5, 1.625, 1.5),
		NewTestCase(0.0, 1.0, 1.0, 0.1, 1.105, 1.1),
	}

	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(order2.Euler{}, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_RK4(t *testing.T) {
	t.Parallel()

	var eq equations.SecondOrder = func(t, y, yp float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 1.0, 2.717, 2.721),
		NewTestCase(0.0, 1.0, 1.0, 0.5, 1.648, 1.648),
		NewTestCase(0.0, 1.0, 1.0, 0.1, 1.105, 1.105),
	}

	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(order2.RK4, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_RK4_Sine(t *testing.T) {
	t.Parallel()

	var eq equations.SecondOrder = func(t, y, yp float64) float64 { return -y }

	tcs := []TestCase{
		// Sine
		NewTestCase(0.0, 0.0, 1.0, math.Pi, 0.09901, -2.101),
		NewTestCase(0.0, 0.0, 1.0, math.Pi/2, 0.9912, -0.0148),
		NewTestCase(0.0, 0.0, 1.0, math.Pi/10, 0.309, 0.951),
		// Cosine
		NewTestCase(0.0, 1.0, 0.0, math.Pi, -1.545, 0.5862),
		NewTestCase(0.0, 1.0, 0.0, math.Pi/2, -0.006110, -1.010),
		NewTestCase(0.0, 1.0, 0.0, math.Pi/10, 0.951, -0.309),
	}

	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(order2.RK4, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_RK38(t *testing.T) {
	t.Parallel()

	var eq equations.SecondOrder = func(t, y, yp float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 1.0, 2.719, 2.723),
		NewTestCase(0.0, 1.0, 1.0, 0.5, 1.648, 1.648),
		NewTestCase(0.0, 1.0, 1.0, 0.1, 1.105, 1.105),
	}

	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(order2.RK38, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_Ralston(t *testing.T) {
	t.Parallel()

	var eq equations.SecondOrder = func(t, y, yp float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 1.0, 2.715, 2.721),
		NewTestCase(0.0, 1.0, 1.0, 0.5, 1.648, 1.648),
		NewTestCase(0.0, 1.0, 1.0, 0.1, 1.105, 1.105),
	}

	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(order2.Ralston, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRungeKutta_Solve(t *testing.T) {
	t.Parallel()

	var eq equations.SecondOrder = func(t, y, yp float64) float64 { return yp }

	tcs := []struct {
		n        int
		wantDiff float64
	}{
		{n: 1, wantDiff: 9.948e-3},
		{n: 10, wantDiff: 2.084e-6},
		{n: 100, wantDiff: 2.246e-10},
	}

	for _, tc := range tcs {
		t.Run(strconv.Itoa(tc.n), func(t *testing.T) {
			t.Parallel()

			got, _ := order2.Solve(order2.RK4, eq, 0.0, 1.0, 1.0, 1.0, tc.n)
			if diff := cmp.Diff(tc.wantDiff, math.Abs(math.E-got), cmpopts.EquateApprox(0.001, 0.0)); diff != "" {
				t.Error("y:", diff)
			}
		})
	}
}

func TestRungeKutta_Solve_Sine(t *testing.T) {
	t.Parallel()

	var eq equations.SecondOrder = func(t, y, yp float64) float64 { return -y }

	tcs := []struct {
		n              int
		wantDiffSine   float64
		wantDiffCosine float64
	}{
		{n: 1, wantDiffSine: 8.757e-3, wantDiffCosine: 6.110e-3},
		{n: 10, wantDiffSine: 1.099e-6, wantDiffCosine: 3.130e-7},
		{n: 100, wantDiffSine: 1.0612e-10, wantDiffCosine: 3.293e-11},
	}

	for _, tc := range tcs {
		t.Run(strconv.Itoa(tc.n), func(t *testing.T) {
			t.Parallel()

			// Sine
			got, _ := order2.Solve(order2.RK4, eq, 0.0, 0.0, 1.0, math.Pi/2, tc.n)
			if diff := cmp.Diff(tc.wantDiffSine, math.Abs(1.0-got), cmpopts.EquateApprox(0.001, 0.0)); diff != "" {
				t.Error("y:", diff)
			}

			// Cosine
			got, _ = order2.Solve(order2.RK4, eq, 0.0, 1.0, 0.0, math.Pi/2, tc.n)
			if diff := cmp.Diff(tc.wantDiffCosine, math.Abs(got), cmpopts.EquateApprox(0.001, 0.0)); diff != "" {
				t.Error("y:", diff)
			}
		})
	}
}
