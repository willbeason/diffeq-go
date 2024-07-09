package main

import (
	"fmt"
	"math"

	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/solvers/order1"
	"github.com/willbeason/diffeq-go/pkg/solvers/order2"
)

func main() {
	order2rk4 := order2.NewRungeKuttaSolver(order2.RK4())

	var eq equations.SecondOrder = func(t, y, yp float64) float64 {
		return (y + yp) / 2
	}

	for _, i := range []int{1, 10, 100, 1000} {
		yf, ypf := order2.Solve(order2rk4, eq, 0.0, 1.0, 1.0, 1.0, i)
		fmt.Println("dy", yf-math.E)
		fmt.Println("dy'", ypf-math.E)
		fmt.Println()
	}

	order1rk4 := order1.RK4()

	var eq1 equations.FirstOrder = func(t, y float64) float64 {
		return y
	}

	fmt.Println()
	fmt.Println()

	for _, i := range []int{1, 10, 100, 1000} {
		yf := order1.Solve(order1rk4, eq1, 0.0, 1.0, 1.0, i)
		fmt.Println("dy", yf-math.E)
		fmt.Println()
	}
}
