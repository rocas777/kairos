package integration

import "math"

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
