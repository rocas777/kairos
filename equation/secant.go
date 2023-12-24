package equation

import "math"

type secant struct {
	epsilon    float64
	cycles     uint
	cycleLimit uint
	f          func(x float64) float64
}

func (s *secant) Cycles() uint {
	return s.cycles
}

func NewSecant(epsilon float64, cycleLimit uint, f func(x float64) float64) *secant {
	return &secant{epsilon, 0, cycleLimit, f}
}

func DefaultSecant(f func(x float64) float64) *secant {
	return NewSecant(0.001, 100, f)
}

func (s *secant) Result(a, b float64) float64 {
	x2 := 20.0
	x0 := a
	x1 := b
	for s.cycles = 0; s.cycles < s.cycleLimit; s.cycles++ {
		x2 = (x0*s.f(x1) - x1*s.f(x0)) / (s.f(x1) - s.f(x0))
		x0 = x1
		x1 = x2
		g := s.f(x2)
		if math.Abs(g) < s.epsilon {
			return x2
		}
	}
	return math.NaN()
}
