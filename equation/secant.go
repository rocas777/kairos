package equation

import "math"

type Secant struct {
	Epsilon    float64
	cycles     uint
	CycleLimit uint
}

func NewSecant(epsilon float64, cycleLimit uint) *Secant {
	return &Secant{Epsilon: epsilon, CycleLimit: cycleLimit}
}

func (s *Secant) Cycles() uint {
	return s.cycles
}

func (s *Secant) Result(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()
	x2 := 20.0
	x0 := a
	x1 := b
	for s.cycles = 0; s.cycles < s.CycleLimit; s.cycles++ {
		x2 = (x0*f(x1) - x1*f(x0)) / (f(x1) - f(x0))
		x0 = x1
		x1 = x2
		g := f(x2)
		if math.Abs(g) < s.Epsilon {
			return x2
		}
	}
	return math.NaN()
}

func (s *Secant) handleInput() {
	if s.CycleLimit == 0 {
		s.CycleLimit = 1
	} else if s.CycleLimit < 0 {
		panic("Secant struct value of CycleLimit should be higher that 0")
	}
	if s.Epsilon == 0 {
		s.Epsilon = 0.01
	} else if s.Epsilon < 0 {
		panic("Secant struct value of Epsilon should be higher that 0")
	}
}
