package integration

import (
	"github.com/rocas777/kairos"
	"math"
)

type SimpsonAdaptive struct {
	Epsilon float64
	simpson *Simpson_1_3
	cycles  uint
}

func NewSimpsonAdaptive(epsilon float64) *SimpsonAdaptive {
	return &SimpsonAdaptive{Epsilon: epsilon}
}

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
		panic("SimpsonAdaptive struct value of Epsilon should be higher that 0")
	}
	if s.simpson == nil {
		s.simpson = &Simpson_1_3{}
	}
}
