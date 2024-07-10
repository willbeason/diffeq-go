package order1_test

import (
	"testing"

	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/solvers/order1"
)

func BenchmarkEuler_Solve(b *testing.B) {
	var eq equations.FirstOrder = func(t, y float64) float64 {
		return y
	}

	euler := order1.Euler{}
	for i := 0; i < b.N; i++ {
		order1.Solve(euler, eq, 0.0, 1.0, 1.0, 100)
	}
}

func BenchmarkRungeKuttaEuler_Solve(b *testing.B) {
	var eq equations.FirstOrder = func(t, y float64) float64 {
		return y
	}

	euler := order1.NewRungeKutta(order1.Step(1.0, 0.0))
	rk4 := order1.NewRungeKuttaSolver(euler)
	for i := 0; i < b.N; i++ {
		order1.Solve(rk4, eq, 0.0, 1.0, 1.0, 100)
	}
}

func BenchmarkRungeKutta_Solve(b *testing.B) {
	var eq equations.FirstOrder = func(t, y float64) float64 {
		return y
	}

	rk4 := order1.NewRungeKuttaSolver(order1.RK4())
	for i := 0; i < b.N; i++ {
		order1.Solve(rk4, eq, 0.0, 1.0, 1.0, 100)
	}
}

func BenchmarkRungeKuttaInline_Solve(b *testing.B) {
	var eq equations.FirstOrder = func(t, y float64) float64 {
		return y
	}

	rk4 := order1.RungeKuttaInline{}
	for i := 0; i < b.N; i++ {
		order1.Solve(rk4, eq, 0.0, 1.0, 1.0, 100)
	}
}
