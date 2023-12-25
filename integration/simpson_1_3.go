// Package integration provides utilities for numerical integration of functions.
// It includes several methods for calculating definite integrals, such as the Trapezoidal Rule,
// Simpson's 1/3 Rule, Simpson's 3/8 Rule, and adaptive Simpson integration.
// Users can choose the appropriate method based on the precision and efficiency requirements
// of their mathematical analysis.
//   - [Trapezoid]
//   - [Simpson_1_3]
//   - [Simpson_1_3]
//   - [SimpsonAdaptive]
package integration

import "github.com/rocas777/kairos"

// Simpson_1_3 provides a method to calculate the definite integral of a given function using the [Simpson] algorithm, specifically the 1/3 composite rule.
// This works by dividing the interval into N pieces and making a polynomial interpolation between two successive points to calculate the area below the interpolation.
// The higher the value of N, the more accurate the integration will be; however, this will lead to a more time-expensive method.
//
// If N is not specified, it defaults to 2. If the value is below 2, a panic will be raised.
//
// [Simpson]: https://en.wikipedia.org/wiki/Simpson%27s_rule#Composite_Simpson's_1/3_rule
type Simpson_1_3 struct {
	N      uint
	cycles uint
}

// NewSimpson_1_3 creates and returns a pointer to a new [Simpson_1_3] instance with the specified value of 'n'.
//
// If 'n' is below 2, a panic is raised.
func NewSimpson_1_3(n uint) *Simpson_1_3 {
	return &Simpson_1_3{N: n * 2}
}

func (s *Simpson_1_3) Cycles() uint {
	return s.cycles
}

// DefiniteIntegral calculates the definite integral of the given function 'f' using the Simpson 1/3 rule.
// It divides the interval [a, b] into N subintervals and performs Simpson's 1/3 rule on each subinterval.
// The result is the sum of the areas under the function's curve within each subinterval.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a float64 representing the calculated definite integral.
func (s *Simpson_1_3) DefiniteIntegral(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()

	s.cycles = 0
	h := (b - a) / float64(s.N)
	partialOut := f(a) + f(b)
	for i := 1; i < int(s.N/2+1); i++ {
		s.cycles++
		partialOut += 4 * f(a+float64(2*i-1)*h)
	}
	for i := 1; i < int(s.N/2); i++ {
		s.cycles++
		partialOut += 2 * f(a+float64(2*i)*h)
	}
	return partialOut * h / 3
}

// AntiDerivative calculates the approximate antiderivative of the given function 'f' using the Simpson 1/3 rule.
// It samples the antiderivative at 'samples' points within the interval [a, b].
// The antiderivative is computed by performing the Simpson 1/3 rule integration from 0 to each sampled point 'x'.
//
// If 'samples' is less than 2, it defaults to 2.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a slice of [kairos.Pair] representing the sampled points and their corresponding antiderivative values.
func (s *Simpson_1_3) AntiDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
	s.handleInput()
	if samples < 2 {
		samples = 2
	}

	y := 0.0
	out := make([]kairos.Pair, samples)
	sampleH := (b - a) / float64(samples-1)
	for i := 0; i < int(samples); i++ {
		x := a + float64(i)*sampleH
		y = s.DefiniteIntegral(f, 0, x)
		out[i] = kairos.Pair{X: x, Y: y}
	}
	return out
}

func (s *Simpson_1_3) handleInput() {
	if s.N == 0 {
		s.N = 2
	} else if s.N < 2 {
		panic("Simpson_1_3 struct value of N should be higher than 0")
	}
}
