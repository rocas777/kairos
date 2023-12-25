package equation

import "math"

// NewtonRaphson provides a method to find the zero of a function using the [Newton-Raphson] method.
// The method iteratively refines the estimate of the zero based on the function's local behavior.
// A solution is considered definitive once the change in the estimate is below Epsilon or the maximum number of cycles (CycleLimit) is reached.
//
// The Newton-Raphson method is generally faster than the bisection method but requires the function to be differentiable.
//
// If 'Epsilon' is not specified, it defaults to 0.01. If 'Epsilon' is less than or equal to 0, a panic is raised.
//
// If 'CycleLimit' is not specified, it defaults to 100. If 'CycleLimit' is less than or equal to 0, a panic is raised.
//
// [Newton-Raphson]: https://en.wikipedia.org/wiki/Newton%27s_method
type NewtonRaphson struct {
	Epsilon    float64
	cycles     uint
	CycleLimit uint
}

// NewNewtonRaphson creates and returns a pointer to a new NewtonRaphson instance with the specified values of 'epsilon' and 'cycleLimit'.
//
// If epsilon is below 0, a panic is raised.
func NewNewtonRaphson(epsilon float64, cycleLimit uint) *NewtonRaphson {
	return &NewtonRaphson{Epsilon: epsilon, CycleLimit: cycleLimit}
}

func (s *NewtonRaphson) Cycles() uint {
	return s.cycles
}

// Zero finds the zero of the function 'f' using the [Newton-Raphson] method.
// It iteratively refines the estimate of the zero based on the function's local behavior using the derivative function 'dxF'.
// A solution is considered definitive once the change in the estimate is below 'Epsilon' or the maximum number of cycles ('CycleLimit') is reached.
// The initial estimate is provided by the parameter 'a'.
// If no zero is found within the given constraints, it returns math.NaN().
func (s *NewtonRaphson) Zero(f func(x float64) float64, dxF func(x float64) float64, a float64) float64 {
	s.handleInput()
	x := a
	for s.cycles = 0; s.cycles < s.CycleLimit; s.cycles++ {
		if f(x) < s.Epsilon {
			return x
		}
		x = x - f(x)/dxF(x)
	}
	return math.NaN()
}

func (s *NewtonRaphson) handleInput() {
	if s.CycleLimit == 0 {
		s.CycleLimit = 1
	} else if s.CycleLimit < 0 {
		panic("NewtonRaphson struct value of CycleLimit should be higher than 0")
	}
	if s.Epsilon == 0 {
		s.Epsilon = 0.01
	} else if s.Epsilon < 0 {
		panic("NewtonRaphson struct value of Epsilon should be higher than 0")
	}
}
