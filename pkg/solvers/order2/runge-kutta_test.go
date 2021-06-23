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
