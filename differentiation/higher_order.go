// Package differentiation provides utilities for calculating derivatives of functions.
// It supports the calculation of the first derivative using two algorithms: Simple (based on the regular definition)
// and Symmetric (based on the symmetric definition). Additionally, it offers the HigherOrder method to calculate
// arbitrary order derivatives. Users can choose the method that best fits their accuracy and efficiency requirements.
//   - 1st order derivative based on the regular derivative definition [Simple]
//   - 1st order derivative based on the symmetric derivative definition [Symmetric]
//   - nth order derivative based on the symmetric derivative definition [HigherOrder]
package differentiation

import "github.com/rocas777/kairos"

// HigherOrder contains methods for calculating the nth-order derivative, where the order is specified by the 'Order' field.
// The algorithm used is the [Symmetric] algorithm. It achieves higher-order derivatives by recursively applying the first-order derivative.
// For more information about the algorithm and the use of the 'H' field, refer to [Symmetric].
//
// If 'H' is not specified, it defaults to 0.1. If 'H' is less than 0, a panic is raised.
//
// If 'Order' is not specified, it defaults to 1.
type HigherOrder struct {
	H     float64
	Order uint
}

// NewHigherOrder creates and returns a pointer to the [HigherOrder] structure.
//
// If 'h' is not specified, it defaults to 0.1. If 'h' is less than 0, a panic is raised.
//
// If 'Order' is not specified, it defaults to 1.
func NewHigherOrder(h float64, order uint) *HigherOrder {
	if h < 0 {
		panic("HigherOrder struct value of H should be higher than 0")
	}
	return &HigherOrder{H: h, Order: order}
}

// LocalDerivative calculates the nth order derivative of the function 'f' at the point 'x' using the [Symmetric] method.
// It returns the calculated derivative value.
func (s *HigherOrder) LocalDerivative(f func(x float64) float64, x float64) float64 {
	s.handleInput()
	if s.Order == 1 {
		return NewSymmetric(s.H).LocalDerivative(f, x)
	}
	return (NewHigherOrder(s.H, s.Order-1).LocalDerivative(f, x+s.H) - NewHigherOrder(s.H, s.Order-1).LocalDerivative(f, x-s.H)) / (s.H * 2)
}

func (s *HigherOrder) handleInput() {
	if s.H == 0 {
		s.H = 0.1
	} else if s.H <= 0 {
		panic("HigherOrder struct value of H should be higher than 0")
	}
	if s.Order == 0 {
		s.Order = 1
	}
}

// RangeDerivative calculates the nth order derivative of the function 'f' over the specified range [a, b] using the [Symmetric] method.
// It divides the range into 'samples' points and returns a slice of [kairos.Pair] representing the points and their corresponding derivative values.
// This function is useful, for example, in drawing the line of the nth order derivative function of 'f'.
func (s *HigherOrder) RangeDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
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
