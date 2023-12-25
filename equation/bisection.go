// Package equation provides utilities for solving equations and finding the roots of functions.
// It offers multiple root-finding methods, including the Bisection method, False Position method,
// Newton-Raphson method, and Secant method. Users can choose the most suitable method for their
// specific functions and interval constraints to efficiently locate zeros of the given function.
//   - [Bisection]
//   - [FalsePosition]
//   - [NewtonRaphson]
//   - [Secant]
//
// Note: These methods assume the provided function is continuous on the considered interval.
package equation

import "math"

// Bisection provides a method to find the zero of a function using the [bisection] method on an interval [a, b].
// The method can be limited by CycleLimit, which restricts the number of cycles to prevent the algorithm from running indefinitely.
// A solution is considered definitive once the difference of the interval [a, b] is below Epsilon.
//
// The bisection method is the slowest of the bunch; however, it is the safest as it always converges.
//
// If 'Epsilon' is not specified, it defaults to 0.01. If 'Epsilon' is less than 0, a panic is raised.
//
// If 'CycleLimit' is not specified, it defaults to 100. If 'CycleLimit' is less than 0, a panic is raised.
//
// [bisection]: https://en.wikipedia.org/wiki/Bisection_method
type Bisection struct {
	Epsilon    float64
	cycles     uint
	CycleLimit uint
}

// NewBisection creates and returns a pointer to a new Bisection instance with the specified values of 'epsilon' and 'cycleLimit'.
//
// If epsilon is below 0, a panic is raised.
func NewBisection(epsilon float64, cycleLimit uint) *Bisection {
	return &Bisection{Epsilon: epsilon, CycleLimit: cycleLimit}
}

func (s *Bisection) Cycles() uint {
	return s.cycles
}

// Zero finds the zero of the function 'f' using the [Bisection] method on the interval [a, b].
// It iteratively narrows down the interval until the solution is found within the specified precision ('Epsilon') or until the maximum number of cycles ('CycleLimit') is reached.
// The result is returned as a float64. If no zero is found within the given constraints, it returns math.NaN().
//
// Note: The function 'f' must have a zero on the interval [a, b], and it must be continuous on that interval.
func (s *Bisection) Zero(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()
	var c float64
	for s.cycles < s.CycleLimit {
		c = (a + b) / 2
		yc := f(c)
		if yc == 0 || (b-a)/2 < s.Epsilon {
			return c
		}
		s.cycles++
		if yc*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return math.NaN()
}

func (s *Bisection) handleInput() {
	if s.CycleLimit == 0 {
		s.CycleLimit = 100
	} else if s.CycleLimit < 0 {
		panic("Bisection struct value of CycleLimit should be higher than 0")
	}
	if s.Epsilon == 0 {
		s.Epsilon = 0.01
	} else if s.Epsilon < 0 {
		panic("Bisection struct value of Epsilon should be higher than 0")
	}
}
