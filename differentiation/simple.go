package differentiation

import "github.com/rocas777/kairos"

type Simple struct {
	H float64
}

func NewSimple(h float64) *Simple {
	return &Simple{H: h}
}

func (s *Simple) LocalDerivative(f func(x float64) float64, x float64) float64 {
	s.handleInput()
	return (f(x+s.H) - f(x)) / s.H
}

func (s *Simple) handleInput() {
	if s.H == 0 {
		s.H = 0.1
	} else if s.H <= 0 {
		panic("Simple struct value of H should be higher that 0")
	}
}

func (s *Simple) RangeDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
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
