package equations

// FirstOrder represents a differential equation dependent on time and the
// current value of the function.
//
// For example, y' = y would be written func(t, y float64) { return y }
type FirstOrder func(t, y float64) float64
