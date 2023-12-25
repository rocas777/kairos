package equation

import "math"

// FalsePosition provides a method to find the zero of a function using the [False Position] method on an interval [a, b].
// The method iteratively refines the estimate of the zero based on linear interpolation.
// A solution is considered definitive once the change in the estimate is below Epsilon or the maximum number of cycles (CycleLimit) is reached.
//
// The False Position method is generally faster than the bisection method but may fail if the function changes sign rapidly.
//
// If 'Epsilon' is not specified, it defaults to 0.01. If 'Epsilon' is less than or equal to 0, a panic is raised.
//
// If 'CycleLimit' is not specified, it defaults to 100. If 'CycleLimit' is less than or equal to 0, a panic is raised.
//
// [False Position]: https://en.wikipedia.org/wiki/Regula_falsi
type FalsePosition struct {
	Epsilon    float64
	cycles     uint
	CycleLimit uint
}

// NewFalsePosition creates and returns a pointer to a new [FalsePosition] instance with the specified values of 'epsilon' and 'cycleLimit'.
//
// If epsilon is below 0, a panic is raised.
func NewFalsePosition(epsilon float64, cycleLimit uint) *FalsePosition {
	return &FalsePosition{Epsilon: epsilon, CycleLimit: cycleLimit}
}

func (s *FalsePosition) Cycles() uint {
	return s.cycles
}

// Zero finds the zero of the function 'f' using the [FalsePosition] method on the interval [a, b].
// It iteratively narrows down the interval until the solution is found within the specified precision ('Epsilon') or until the maximum number of cycles ('CycleLimit') is reached.
// The result is returned as a float64. If no zero is found within the given constraints, it returns math.NaN().
//
// Note: The function 'f' must have a zero on the interval [a, b], and it must be continuous on that interval.
func (s *FalsePosition) Zero(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()
	side := 0

	fa := f(a)
	fb := f(b)

	for s.cycles = 0; s.cycles < s.CycleLimit; s.cycles++ {
		c := (fa*b - fb*a) / (fa - fb)
		if math.Abs(b-a) < s.Epsilon*math.Abs(b+a) {
			return c
		}
		fc := f(c)
		if fc*fb > 0 {
			b = c
			fb = fc
			if side == -1 {
				fa /= 2
			}
			side = -1
		} else if fa*fc > 0 {
			a = c
			fa = fc
			if side == 1 {
				fb /= 2
			}
			side = 1
		} else {
			return c
		}
	}
	return math.NaN()
}

func (s *FalsePosition) handleInput() {
	if s.CycleLimit == 0 {
		s.CycleLimit = 1
	} else if s.CycleLimit < 0 {
		panic("FalsePosition struct value of CycleLimit should be higher than 0")
	}
	if s.Epsilon == 0 {
		s.Epsilon = 0.01
	} else if s.Epsilon < 0 {
		panic("FalsePosition struct value of Epsilon should be higher than 0")
	}
}
