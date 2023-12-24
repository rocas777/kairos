package integration

import "math"

type simpsonAdaptive struct {
	epsilon float64
	simpson *simpson_1_3
	cycles  uint
	f       func(x float64) float64
}

func NewSimpson_adaptive(epsilon float64, f func(x float64) float64) *simpsonAdaptive {
	return &simpsonAdaptive{epsilon, NewSimpson_1_3(1, f), 0, f}
}

func DefaultSimpsonAdaptive(f func(x float64) float64) *simpsonAdaptive {
	return NewSimpson_adaptive(0.10, f)
}

func (s *simpsonAdaptive) DefiniteIntegral(a, b float64) float64 {
	s.cycles = 0
	return s._recursiveSimpson(a, b)
}

func (s *simpsonAdaptive) _recursiveSimpson(a, b float64) float64 {
	s_a_m := s.simpson.DefiniteIntegral(a, (b+a)/2.0)
	s_m_b := s.simpson.DefiniteIntegral((b+a)/2.0, b)
	s_a_b := s.simpson.DefiniteIntegral(a, b)
	s.cycles++
	if math.Abs(s_a_m+s_m_b-s_a_b) < 15*s.epsilon {
		return s_a_m + s_m_b
	}
	return s._recursiveSimpson(a, (b+a)/2.0) + s._recursiveSimpson((b+a)/2.0, b)
}

func (s *simpsonAdaptive) Cycles() uint {
	return s.cycles
}
