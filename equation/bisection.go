package equation

import "math"

type bisection struct {
	epsilon    float64
	cycles     uint
	cycleLimit uint
	f          func(x float64) float64
}

func (s *bisection) Cycles() uint {
	return s.cycles
}

func NewBisection(epsilon float64, cycleLimit uint, f func(x float64) float64) *bisection {
	return &bisection{epsilon, 0, cycleLimit, f}
}

func DefaultBisection(f func(x float64) float64) *bisection {
	return NewBisection(0.001, 100, f)
}

func (s *bisection) Zero(a, b float64) float64 {
	c := -20.0
	for s.cycles < s.cycleLimit {
		c = (a + b) / 2
		yc := s.f(c)
		if yc == 0 || (b-a)/2 < s.epsilon {
			return c
		}
		s.cycles++
		if yc*s.f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return math.NaN()
}
