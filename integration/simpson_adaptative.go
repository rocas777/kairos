package integration

import (
	"github.com/rocas777/kairos"
	"math"
)

// SimpsonAdaptive provides a method to calculate the definite integral of a given function using adaptive Simpson integration.
// It combines the simplicity of the [Simpson 1/3 Rule] with adaptivity to improve accuracy.
// The algorithm automatically adjusts the number of intervals based on the function's behavior, comparing the results by dividing intervals into more subintervals using the epsilon criterion.
//
// If 'Epsilon' is not specified, it defaults to 0.001. If 'Epsilon' is less than or equal to 0, a panic will be raised.
//
// [Simpson 1/3 Rule]: https://en.wikipedia.org/wiki/Simpson%27s_rule#Composite_Simpson's_1/3_rule
type SimpsonAdaptive struct {
	Epsilon float64
	simpson *Simpson_1_3
	cycles  uint
}

// NewSimpsonAdaptive creates and returns a pointer to a new [SimpsonAdaptive] instance with the specified value of 'n'.
//
// If 'n' is below 0, a panic is raised.
func NewSimpsonAdaptive(epsilon float64) *SimpsonAdaptive {
	return &SimpsonAdaptive{Epsilon: epsilon}
}

// DefiniteIntegral calculates the definite integral of the given function 'f' using adaptive Simpson integration.
// It combines the results of two consecutive applications of the Simpson 1/3 rule on subintervals [(a, (a+b)/2)] and [((a+b)/2, b)].
// The adaptive nature of the algorithm checks the difference between the combined result and the result from the Simpson 1/3 rule on the entire interval [a, b].
// If the difference is within an acceptable range specified by 'Epsilon', it returns the combined result; otherwise, it recursively applies adaptive Simpson integration on each half of the interval.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a float64 representing the calculated definite integral.
func (s *SimpsonAdaptive) DefiniteIntegral(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()

	s_a_m := s.simpson.DefiniteIntegral(f, a, (b+a)/2.0)
	s_m_b := s.simpson.DefiniteIntegral(f, (b+a)/2.0, b)
	s_a_b := s.simpson.DefiniteIntegral(f, a, b)
	s.cycles++
	if math.Abs(s_a_m+s_m_b-s_a_b) < 15*s.Epsilon {
		return s_a_m + s_m_b
	}
	return s.DefiniteIntegral(f, a, (b+a)/2.0) + s.DefiniteIntegral(f, (b+a)/2.0, b)
}

// AntiDerivative calculates the approximate antiderivative of the given function 'f' using adaptive Simpson integration.
// It samples the antiderivative at 'samples' points within the interval [a, b].
// The antiderivative is computed by performing adaptive Simpson integration from 0 to each sampled point 'x'.
//
// If 'samples' is less than 2, it defaults to 2.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a slice of [kairos.Pair] representing the sampled points and their corresponding antiderivative values.
func (s *SimpsonAdaptive) AntiDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
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

func (s *SimpsonAdaptive) Cycles() uint {
	return s.cycles
}

func (s *SimpsonAdaptive) handleInput() {
	if s.Epsilon == 0 {
		s.Epsilon = 0.1
	} else if s.Epsilon < 0 {
		panic("SimpsonAdaptive struct value of Epsilon should be higher than 0")
	}
	if s.simpson == nil {
		s.simpson = &Simpson_1_3{}
	}
}
