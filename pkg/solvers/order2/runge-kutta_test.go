package order2

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/willbeason/diffeq-go/pkg/equations"
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
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(Euler{}, eq)
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
		NewTestCase(0.0, 1.0, 1.0, 1.0, 2.721, 2.751),
		NewTestCase(0.0, 1.0, 1.0, 0.5, 1.648, 1.650),
		NewTestCase(0.0, 1.0, 1.0, 0.1, 1.105, 1.105),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(RK4, eq)
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
		NewTestCase(0.0, 0.0, 1.0, math.Pi, 1.161, -2.853),
		NewTestCase(0.0, 0.0, 1.0, math.Pi/2, 1.024, 0.068),
		NewTestCase(0.0, 0.0, 1.0, math.Pi/10, 0.309, 0.951),
		// Cosine
		NewTestCase(0.0, 1.0, 0.0, math.Pi, -2.379, 1.177),
		NewTestCase(0.0, 1.0, 0.0, math.Pi/2, -0.0191, -1.0435),
		NewTestCase(0.0, 1.0, 0.0, math.Pi/10, 0.951, -0.309),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(RK4, eq)
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
		NewTestCase(0.0, 1.0, 1.0, 1.0, 2.724, 2.754),
		NewTestCase(0.0, 1.0, 1.0, 0.5, 1.648, 1.650),
		NewTestCase(0.0, 1.0, 1.0, 0.1, 1.105, 1.105),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(RK38, eq)
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
		NewTestCase(0.0, 1.0, 1.0, 1.0, 2.719, 2.751),
		NewTestCase(0.0, 1.0, 1.0, 0.5, 1.648, 1.650),
		NewTestCase(0.0, 1.0, 1.0, 0.1, 1.105, 1.105),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(Ralston, eq)
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
		{n: 1, wantDiff: 0.009948},
		{n: 10, wantDiff: 0.000002084},
		{n: 100, wantDiff: 0.0000000002246},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			t.Parallel()

			got, _ := Solve(RK4, eq, 0.0, 1.0, 1.0, 1.0, tc.n)
			if diff := cmp.Diff(tc.wantDiff, math.Abs(math.E-got), cmpopts.EquateApprox(0.001, 0.0)); diff != "" {
				t.Error("y:", diff)
			}
		})
	}
}
