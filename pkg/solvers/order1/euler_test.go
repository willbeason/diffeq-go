package order1_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/solvers/order1"
)

type TestCase struct {
	t0, y0, h, want float64
}

func (tc TestCase) Name() string {
	return fmt.Sprintf("t0 = %.2f, y0 = %.2f, h = %.2f", tc.t0, tc.y0, tc.h)
}

func NewTestCase(t0, y0, h, want float64) TestCase {
	return TestCase{t0: t0, y0: y0, h: h, want: want}
}

var ErrSolverFail = errors.New("solver returned unexpected result")

func (tc TestCase) Run(solver order1.Solver, eq equations.FirstOrder) error {
	got := solver.Solve(eq, tc.t0, tc.y0, tc.h)

	if math.Abs(got-tc.want) > 0.001 {
		return fmt.Errorf("%w: got %f, want %f", ErrSolverFail, got, tc.want)
	}

	return nil
}

func TestEuler_Solve(t *testing.T) {
	t.Parallel()

	eq := func(t, y float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.0),
		NewTestCase(0.0, 1.0, 0.5, 1.5),
		NewTestCase(0.0, 1.0, 0.1, 1.1),
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.Name(), func(t *testing.T) {
			t.Parallel()
			err := tc.Run(order1.Euler{}, eq)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
