package equations

// FirstOrder represents a differential equation dependent on time and the
// current value of the function.
//
// For example, y' = y would be written func(t, y float64) { return y }.
type FirstOrder func(t, y float64) float64

// SecondOrder represents a differential equation dependent on time, the current
// value of the function, and the function's derivative.
type SecondOrder func(t, y, yp float64) float64
