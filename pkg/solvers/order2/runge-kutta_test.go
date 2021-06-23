package order2

import (
	"testing"

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
