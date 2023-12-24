package equation

import "math"

type newtonRaphson struct {
	epsilon    float64
	cycles     uint
	cycleLimit uint
	f          func(x float64) float64
}

func (s *newtonRaphson) Cycles() uint {
	return s.cycles
}

func NewNewtonRaphson(epsilon float64, cycleLimit uint, f func(x float64) float64) *newtonRaphson {
	return &newtonRaphson{epsilon, 0, cycleLimit, f}
}

func DefaultNewtonRaphson(f func(x float64) float64) *newtonRaphson {
	return NewNewtonRaphson(0.001, 100, f)
}

func (s *newtonRaphson) Result(f1 func(x float64) float64, a float64) float64 {
	x := a
	for s.cycles = 0; s.cycles < s.cycleLimit; s.cycles++ {
		if s.f(x) < s.epsilon {
			return x
		}
		x = x - s.f(x)/f1(x)
	}
	return math.NaN()
}
