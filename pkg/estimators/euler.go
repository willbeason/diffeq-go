package estimators

import (
	"github.com/willbeason/diffeq-go/pkg/equations"
)

type Euler struct {}

var _ Solver = Euler{}

func (e Euler) Solve(eq equations.FirstOrder, t, y, h float64) float64 {
	k := eq(t, y)
	return y + h*k
}
