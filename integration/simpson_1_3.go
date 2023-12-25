package integration

import "github.com/rocas777/kairos"

type Simpson_1_3 struct {
	N      uint
	cycles uint
}

func NewSimpson_1_3(n uint) *Simpson_1_3 {
	return &Simpson_1_3{N: n * 2}
}

func (s *Simpson_1_3) Cycles() uint {
	return s.cycles
}

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
		panic("Simpson_1_3 struct value of N should be higher that 0")
	}
}
