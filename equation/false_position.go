package equation

import "math"

type FalsePosition struct {
	Epsilon    float64
	cycles     uint
	CycleLimit uint
}

func NewFalsePosition(epsilon float64, cycleLimit uint) *FalsePosition {
	return &FalsePosition{Epsilon: epsilon, CycleLimit: cycleLimit}
}

func (s *FalsePosition) Cycles() uint {
	return s.cycles
}

func (s *FalsePosition) Result(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()
	side := 0

	fa := f(a)
	fb := f(b)

	for s.cycles = 0; s.cycles < s.CycleLimit; s.cycles++ {
		c := (fa*b - fb*a) / (fa - fb)
		if math.Abs(b-a) < s.Epsilon*math.Abs(b+a) {
			return c
		}
		fc := f(c)
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

func (s *FalsePosition) handleInput() {
	if s.CycleLimit == 0 {
		s.CycleLimit = 1
	} else if s.CycleLimit < 0 {
		panic("FalsePosition struct value of CycleLimit should be higher that 0")
	}
	if s.Epsilon == 0 {
		s.Epsilon = 0.01
	} else if s.Epsilon < 0 {
		panic("FalsePosition struct value of Epsilon should be higher that 0")
	}
}
