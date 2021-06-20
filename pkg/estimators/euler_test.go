package estimators

import (
	"fmt"
	"math"
	"testing"

	"github.com/willbeason/diffeq-go/pkg/equations"
)

type TestCase struct {
	t0, y0, h, want float64
}

func NewTestCase(t0, y0, h, want float64) TestCase {
	return TestCase{t0: t0, y0: y0, h: h, want: want}
}

func (tc TestCase) Run(t *testing.T, solver Solver, eq equations.FirstOrder) {
	t.Run(fmt.Sprintf("t0 = %.2f, y0 = %.2f, h = %.2f", tc.t0, tc.y0, tc.h), func(t *testing.T) {
		got := solver.Solve(eq, tc.t0, tc.y0, tc.h)

		if math.Abs(got-tc.want) > 0.001 {
			t.Errorf("got %f, want %f", got, tc.want)
		}
	})

}

func TestEuler_Solve(t *testing.T) {
	eq := func(t, y float64) float64 { return y }

	tcs := []TestCase{
		NewTestCase(0.0, 1.0, 1.0, 2.0),
		NewTestCase(0.0, 1.0, 0.5, 1.5),
		NewTestCase(0.0, 1.0, 0.1, 1.1),
	}

	for _, tc := range tcs {
		tc.Run(t, Euler{}, eq)
	}
}
