package equation

import "math"

type Bisection struct {
	Epsilon    float64
	cycles     uint
	CycleLimit uint
}

func NewBisection(epsilon float64, cycleLimit uint) *Bisection {
	return &Bisection{Epsilon: epsilon, CycleLimit: cycleLimit}
}

func (s *Bisection) Cycles() uint {
	return s.cycles
}

func (s *Bisection) Zero(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()
	var c float64
	for s.cycles < s.CycleLimit {
		c = (a + b) / 2
		yc := f(c)
		if yc == 0 || (b-a)/2 < s.Epsilon {
			return c
		}
		s.cycles++
		if yc*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return math.NaN()
}

func (s *Bisection) handleInput() {
	if s.CycleLimit == 0 {
		s.CycleLimit = 1
	} else if s.CycleLimit < 0 {
		panic("Bisection struct value of CycleLimit should be higher that 0")
	}
	if s.Epsilon == 0 {
		s.Epsilon = 0.01
	} else if s.Epsilon < 0 {
		panic("Bisection struct value of Epsilon should be higher that 0")
	}
}
