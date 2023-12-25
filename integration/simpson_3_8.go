package integration

import "github.com/rocas777/kairos"

type Simpson_3_8 struct {
	N      uint
	cycles uint
}

func NewSimpson_3_8(n uint) *Simpson_3_8 {
	return &Simpson_3_8{N: n * 3}
}

func (s *Simpson_3_8) Cycles() uint {
	return s.cycles
}

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
		panic("Simpson_3_8 struct value of N should be higher that 0")
	}
}
