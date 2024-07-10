package order2

import "github.com/willbeason/diffeq-go/pkg/equations"

type RungeKuttaInline struct{}

var _ Solver = RungeKuttaInline{}

func (r RungeKuttaInline) Solve(eq equations.SecondOrder, t, y0, yp0, h float64) (float64, float64) {
	k1 := eq(t+h/2.0, y0, yp0)
	yp1 := yp0 + k1*h/2.0
	y1 := y0 + (h/12.0)*(h/2.0*k1+4.0*yp0+2.0*yp1)

	k2 := eq(t+h/2.0, y1, yp1)
	yp2 := yp0 + k2*h/2.0
	y2 := y0 + (h/12.0)*(h/2.0*k1+4.0*yp0+2.0*yp2)

	k3 := eq(t+h/2.0, y2, yp2)
	yp3 := yp0 + k3*h
	y3 := y0 + (h/6.0)*(h*k1+4.0*yp0+2.0*yp3)

	k4 := eq(t+h, y3, yp3)

	ypf := yp0 + (h/6.0)*(k1+2*k2+2*k3+k4)
	yf := y0 + (h/6.0)*(yp0+2*yp1+2*yp2+yp3)

	return yf, ypf
}
