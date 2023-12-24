package differentiation

import "github.com/rocas777/kairos"

type Symmetric struct {
	H float64
}

func NewSymmetric(h float64) *Symmetric {
	return &Symmetric{H: h}
}

func (s *Symmetric) LocalDerivative(f func(x float64) float64, x float64) float64 {
	s.handleInput()
	return (f(x+s.H) - f(x-s.H)) / (s.H * 2)
}

func (s *Symmetric) handleInput() {
	if s.H == 0 {
		s.H = 0.1
	} else if s.H <= 0 {
		panic("Symmetric struct value of H should be higher that 0")
	}
}

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
