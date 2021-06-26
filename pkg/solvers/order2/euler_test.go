package order2

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/willbeason/diffeq-go/pkg/equations"
)

type TestCase struct {
	t0, y0, yp0, h float64
	wantY, wantYP  float64
}

func (tc TestCase) Name() string {
	return fmt.Sprintf("t0 = %.2f, y0 = %.2f, h = %.2f", tc.t0, tc.y0, tc.h)
}

func NewTestCase(t0, y0, yp0, h, wantY, wantYP float64) TestCase {
	return TestCase{t0: t0, y0: y0, yp0: yp0, h: h, wantY: wantY, wantYP: wantYP}
}

var ErrSolverFail = errors.New("solver returned unexpected result")

func (tc TestCase) Run(solver Solver, eq equations.SecondOrder) error {
	gotY, gotYP := solver.Solve(eq, tc.t0, tc.y0, tc.yp0, tc.h)

	if math.Abs(gotY-tc.wantY) > 0.001 {
		return fmt.Errorf("%w: got y=%f, want y=%f", ErrSolverFail, gotY, tc.wantY)
	}
	if math.Abs(gotYP-tc.wantYP) > 0.001 {
		return fmt.Errorf("%w: got yp=%f, want yp=%f", ErrSolverFail, gotYP, tc.wantYP)
	}

	return nil
}

func TestEuler_Solve(t *testing.T) {
	t.Parallel()

	eq := func(t, y, yp float64) float64 { return y }

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
