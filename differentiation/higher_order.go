// Package differentiation implements 2 + 1 methods for local differentiation and range differentiation
package differentiation

import "github.com/rocas777/kairos"

type HigherOrder struct {
	H     float64
	Order uint
}

func NewHigherOrder(h float64, order uint) *HigherOrder {
	return &HigherOrder{H: h, Order: order}
}

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
		panic("HigherOrder struct value of H should be higher that 0")
	}
	if s.Order == 0 {
		s.Order = 1
	}
}

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
