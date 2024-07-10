package order1

// RungeKutta is an explicit Runge-Kutta solver for first-order differential
// equations.
type RungeKutta struct {
	// Steps are the steps of the Runge-Kutta solver, in order.
	Steps []RungeKuttaStep
}

func NewRungeKutta(steps ...RungeKuttaStep) RungeKutta {
	return RungeKutta{Steps: steps}
}

// RungeKuttaStep is a step in a Runge-Kutta solver.
type RungeKuttaStep struct {
	// Weight is the weight given to this step's calculated slope in the final
	// calculation.
	Weight float64

	// Node is the fraction of the estimated time step forward to be estimated.
	Node float64

	// Coefficients are the weights of previous steps to include.
	Coefficients []float64
}

func Step(b, c float64, coefficients ...float64) RungeKuttaStep {
	return RungeKuttaStep{
		Weight:       b,
		Node:         c,
		Coefficients: coefficients,
	}
}
