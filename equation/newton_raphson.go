package equation

import "math"

type NewtonRaphson struct {
	Epsilon    float64
	cycles     uint
	CycleLimit uint
}

func NewNewtonRaphson(epsilon float64, cycleLimit uint) *NewtonRaphson {
	return &NewtonRaphson{Epsilon: epsilon, CycleLimit: cycleLimit}
}

func (s *NewtonRaphson) Cycles() uint {
	return s.cycles
}

func (s *NewtonRaphson) Result(f func(x float64) float64, dxF func(x float64) float64, a float64) float64 {
	s.handleInput()
	x := a
	for s.cycles = 0; s.cycles < s.CycleLimit; s.cycles++ {
		if f(x) < s.Epsilon {
			return x
		}
		x = x - f(x)/dxF(x)
	}
	return math.NaN()
}

func (s *NewtonRaphson) handleInput() {
	if s.CycleLimit == 0 {
		s.CycleLimit = 1
	} else if s.CycleLimit < 0 {
		panic("NewtonRaphson struct value of CycleLimit should be higher that 0")
	}
	if s.Epsilon == 0 {
		s.Epsilon = 0.01
	} else if s.Epsilon < 0 {
		panic("NewtonRaphson struct value of Epsilon should be higher that 0")
	}
}
