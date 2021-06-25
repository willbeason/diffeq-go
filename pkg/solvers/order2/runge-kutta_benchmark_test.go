package order2

import (
	"testing"

	"github.com/willbeason/diffeq-go/pkg/equations"
)

func BenchmarkRungeKutta_Solve(b *testing.B) {
	var eq equations.SecondOrder = func(t, y, yp float64) float64 {
		return y
	}

	for i := 0; i < b.N; i++ {
		Solve(RK4, eq, 0.0, 1.0, 1.0, 1.0, 100)
	}
}
