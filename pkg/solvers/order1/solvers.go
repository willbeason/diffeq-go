package order1

// RK4 is the traditional 4-step Runge-Kutta method.
var RK4 = NewRungeKutta(
	Step(1.0/6.0, 0.0),
	Step(1.0/3.0, 0.5, 0.5),
	Step(1.0/3.0, 0.5, 0.0, 0.5),
	Step(1.0/6.0, 1.0, 0.0, 0.0, 1.0),
)

// RK38 is the 3/8-rule Runge-Kutta method.
var RK38 = NewRungeKutta(
	Step(1.0/8.0, 0.0),
	Step(3.0/8.0, 1.0/3.0, 1.0/3.0),
	Step(3.0/8.0, 2.0/3.0, -1.0/3.0, 1.0),
	Step(1.0/8.0, 1.0, 1.0, -1.0, 1.0),
)

// Ralston is Ralston's fourth-order Runge-Kutta method which has minimum
// truncation error.
var Ralston = NewRungeKutta(
	Step(0.17476028, 0.0),
	Step(-0.55148066, 0.4, 0.4),
	Step(1.20553560, 0.45573725, 0.29697761, 0.15875964),
	Step(0.17118478, 1.0, 0.21810040, -3.05096516, 3.83286476),
)
