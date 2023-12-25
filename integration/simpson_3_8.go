package integration

import "github.com/rocas777/kairos"

// Simpson_3_8 provides a method to calculate the definite integral of a given function using the [Simpson] algorithm, specifically the 3/8 composite rule.
// This works by dividing the interval into N pieces and making a polynomial interpolation using three successive points to calculate the area below the interpolation.
// The higher the value of N, the more accurate the integration will be; however, this will lead to a more time-expensive method.
//
// If N is not specified, it defaults to 3. If the value is below 3, a panic will be raised.
//
// [Simpson]: https://en.wikipedia.org/wiki/Simpson%27s_rule#Composite_Simpson's_3/8_rule
type Simpson_3_8 struct {
	N      uint
	cycles uint
}

// NewSimpson_3_8 creates and returns a pointer to a new [Simpson_3_8] instance with the specified value of 'n'.
//
// If 'n' is below 3, a panic is raised.
func NewSimpson_3_8(n uint) *Simpson_3_8 {
	return &Simpson_3_8{N: n * 3}
}

func (s *Simpson_3_8) Cycles() uint {
	return s.cycles
}

// DefiniteIntegral calculates the definite integral of the given function 'f' using the Simpson 3/8 rule.
// It divides the interval [a, b] into N subintervals and performs Simpson's 3/8 rule on each subinterval.
// The result is the sum of the areas under the function's curve within each subinterval.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a float64 representing the calculated definite integral.
func (s *Simpson_3_8) DefiniteIntegral(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()

	s.cycles = 0
	h := (b - a) / float64(s.N)
	partialOut := f(a) + f(b)
	for i := 1; i < int(s.N/3+1); i++ {
		partialOut += 3 * f(a+float64(3*i-2)*h)
		partialOut += 3 * f(a+float64(3*i-1)*h)
		s.cycles++
	}
	for i := 1; i < int(s.N/3); i++ {
		partialOut += 2 * f(a+float64(3*i)*h)
		s.cycles++
	}
	return partialOut * h * 3 / 8
}

// AntiDerivative calculates the approximate antiderivative of the given function 'f' using the Simpson 3/8 rule.
// It samples the antiderivative at 'samples' points within the interval [a, b].
// The antiderivative is computed by performing the Simpson 3/8 rule integration from 0 to each sampled point 'x'.
//
// If 'samples' is less than 2, it defaults to 2.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a slice of [kairos.Pair] representing the sampled points and their corresponding antiderivative values.
func (s *Simpson_3_8) AntiDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
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

func (s *Simpson_3_8) handleInput() {
	if s.N == 0 {
		s.N = 3
	} else if s.N < 3 {
		panic("Simpson_3_8 struct value of N should be higher than 0")
	}
}
