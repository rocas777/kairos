package integration

import "github.com/rocas777/kairos"

// Trapezoid provides a method to calculate the definite integral of a given function using the [Trapezoidal Rule].
// This rule approximates the integral by dividing the interval into N trapezoids and summing their areas.
// The higher the value of N, the more accurate the integration will be; however, this will lead to a more time-expensive method.
//
// If N is not specified, it defaults to 10. If the value is below 0, a panic will be raised.
//
// [Trapezoidal Rule]: https://en.wikipedia.org/wiki/Trapezoidal_rule
type Trapezoid struct {
	N      uint
	cycles uint
}

// NewTrapezoid creates and returns a pointer to a new [Trapezoid] instance with the specified value of 'n'.
//
// If 'n' is below 0, a panic is raised.
func NewTrapezoid(n uint) *Trapezoid {
	return &Trapezoid{N: n}
}

func (t *Trapezoid) Cycles() uint {
	return t.cycles
}

// DefiniteIntegral calculates the definite integral of the given function 'f' using the Trapezoidal Rule.
// It divides the interval [a, b] into N subintervals and performs the Trapezoidal Rule on each subinterval.
// The result is the sum of the areas under the function's curve within each subinterval.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a float64 representing the calculated definite integral.
func (t *Trapezoid) DefiniteIntegral(f func(x float64) float64, a, b float64) float64 {
	t.cycles = 0
	h := (b - a) / float64(t.N)
	out := 0.0
	for i := 0; i < int(t.N); i++ {
		t.cycles++
		out += (f(a+float64(i)*h) + f(a+float64(i+1)*h)) / 2.0 * h
	}
	return out
}

// AntiDerivative calculates the approximate antiderivative of the given function 'f' using the Trapezoidal Rule.
// It samples the antiderivative at 'samples' points within the interval [a, b].
// The antiderivative is computed by performing the Trapezoidal Rule integration from 0 to each sampled point 'x'.
//
// If 'samples' is less than 2, it defaults to 2.
//
// The function 'f' represents the integrand, and 'a' and 'b' define the integration interval.
// The result is returned as a slice of [kairos.Pair] representing the sampled points and their corresponding antiderivative values.
func (t *Trapezoid) AntiDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
	t.handleInput()
	if samples < 2 {
		samples = 2
	}

	y := 0.0
	out := make([]kairos.Pair, samples)
	sampleH := (b - a) / float64(samples-1)
	for i := 0; i < int(samples); i++ {
		x := a + float64(i)*sampleH
		y = t.DefiniteIntegral(f, 0, x)
		out[i] = kairos.Pair{X: x, Y: y}
	}
	return out
}

func (t *Trapezoid) handleInput() {
	if t.N == 0 {
		t.N = 10
	} else if t.N < 0 {
		panic("Trapezoid struct value of N should be higher than 0")
	}

}
