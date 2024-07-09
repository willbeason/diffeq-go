package order1_test

import (
	"testing"

	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/solvers/order1"
)

func BenchmarkRungeKutta_Solve(b *testing.B) {
	var eq equations.FirstOrder = func(t, y float64) float64 {
		return y
	}

	for i := 0; i < b.N; i++ {
		order1.Solve(order1.RK4, eq, 0.0, 1.0, 1.0, 100)
	}
}
