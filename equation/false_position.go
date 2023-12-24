package equation

import "math"

type falsePosition struct {
	epsilon    float64
	cycles     uint
	cycleLimit uint
	f          func(x float64) float64
}

func (s *falsePosition) Cycles() uint {
	return s.cycles
}

func NewFalsePosition(epsilon float64, cycleLimit uint, f func(x float64) float64) *falsePosition {
	return &falsePosition{epsilon, 0, cycleLimit, f}
}

func DefaultFalsePosition(f func(x float64) float64) *falsePosition {
	return NewFalsePosition(0.001, 100, f)
}

func (s *falsePosition) Result(a, b float64) float64 {
	var c float64
	var fc float64
	side := 0

	fa := s.f(a)
	fb := s.f(b)

	for s.cycles = 0; s.cycles < s.cycleLimit; s.cycles++ {
		c = (fa*b - fb*a) / (fa - fb)
		if math.Abs(b-a) < s.epsilon*math.Abs(b+a) {
			return c
		}
		fc = s.f(c)
		if fc*fb > 0 {
			b = c
			fb = fc
			if side == -1 {
				fa /= 2
			}
			side = -1
		} else if fa*fc > 0 {
			a = c
			fa = fc
			if side == 1 {
				fb /= 2
			}
			side = 1
		} else {
			return c
		}
	}
	return math.NaN()
}
