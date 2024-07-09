package order2_test

import (
	"testing"

	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/solvers/order2"
)

func BenchmarkRungeKutta_Solve(b *testing.B) {
	var eq equations.SecondOrder = func(t, y, yp float64) float64 {
		return y
	}

	rks4 := order2.NewRungeKuttaSolver(order2.RK4())
	for i := 0; i < b.N; i++ {
		order2.Solve(rks4, eq, 0.0, 1.0, 1.0, 1.0, 100)
	}
}
