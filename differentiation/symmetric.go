package differentiation

import "github.com/rocas777/kairos"

// Symmetric provides methods for calculating the first derivative.
// The algorithm is based on the symmetric definition of the derivative as a limit.
// 'H' is used as an approximation of infinitesimals. The derivative is computed as the slope of the function between points 'x - H' and 'x + H'.
// For accurate results, 'H' should be close to 0. Note that the closer 'H' is to 0, the closer the slope estimations are to the real value. However, this also increases the impact of errors introduced by floating-point calculations.
// It is necessary to strike a balance between the errors introduced by calculations and the errors associated with the fact that 'H' is not truly an infinitesimal.
// The symmetric approach usually provides better results as it is not one-sided, like the [Simple] definition, and the center of the interval is the point 'x', unlike the regular definition where the center of the interval is the point 'x + H/2'.
//
// If 'H' is not specified, it defaults to 0.1. If 'H' is less than 0, a panic is raised.
//
// If 'Order' is not specified, it defaults to 1.
//
// [definition]: https://en.wikipedia.org/wiki/Symmetric_derivative
type Symmetric struct {
	H float64
}

// NewSymmetric creates and returns a pointer to a new Symmetric instance with the specified value of 'h'.
//
// If 'h' is less than 0, a panic is raised.
func NewSymmetric(h float64) *Symmetric {
	return &Symmetric{H: h}
}

// LocalDerivative calculates the first order derivative of the function 'f' at the point 'x' using the Symmetric method.
// It returns the calculated derivative value.
func (s *Symmetric) LocalDerivative(f func(x float64) float64, x float64) float64 {
	s.handleInput()
	return (f(x+s.H) - f(x-s.H)) / (s.H * 2)
}

func (s *Symmetric) handleInput() {
	if s.H == 0 {
		s.H = 0.1
	} else if s.H <= 0 {
		panic("Symmetric struct value of H should be higher than 0")
	}
}

// RangeDerivative calculates the first order derivative of the function 'f' over the specified range [a, b] using the [Symmetric] method.
// It divides the range into 'samples' points and returns a slice of [kairos.Pair] representing the points and their corresponding derivative values.
// This function is useful, for example, in drawing the line of the first derivative function of 'f'.
func (s *Symmetric) RangeDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
	s.handleInput()
	if samples < 2 {
		samples = 2
	}
	out := make([]kairos.Pair, samples)
	sampleH := (b - a) / float64(samples-1)
	for i := 0; i < int(samples); i++ {
		x := a + float64(i)*sampleH
		out[i] = kairos.Pair{X: x, Y: s.LocalDerivative(f, x)}
	}
	return out
}
