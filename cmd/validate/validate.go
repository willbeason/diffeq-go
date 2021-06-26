package main

import (
	"fmt"
	"math"

	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/solvers/order1"
	"github.com/willbeason/diffeq-go/pkg/solvers/order2"
)

func main() {
	s := order2.RK4

	var eq equations.SecondOrder = func(t, y, yp float64) float64 {
		return (y + yp) / 2
	}

	for _, i := range []int{1, 10, 100, 1000} {
		yf, ypf := order2.Solve(s, eq, 0.0, 1.0, 1.0, 1.0, i)
		fmt.Println("dy", yf-math.E)
		fmt.Println("dy'", ypf-math.E)
		fmt.Println()
	}

	s1 := order1.RK4

	var eq1 equations.FirstOrder = func(t, y float64) float64 {
		return y
	}

	fmt.Println()
	fmt.Println()

	for _, i := range []int{1, 10, 100, 1000} {
		yf := order1.Solve(s1, eq1, 0.0, 1.0, 1.0, i)
		fmt.Println("dy", yf-math.E)
		fmt.Println()
	}
}
