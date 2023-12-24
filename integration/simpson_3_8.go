package integration

type Simpson_3_8 struct {
	N      uint
	cycles uint
}

func NewSimpson_3_8(n uint) *Simpson_3_8 {
	return &Simpson_3_8{N: n * 3}
}

func (s *Simpson_3_8) Cycles() uint {
	return s.cycles
}

func (s *Simpson_3_8) DefiniteIntegral(f func(x float64) float64, a, b float64) float64 {
	s.handleInput()

	s.cycles = 0
	h := (b - a) / float64(s.N)
	partialOut := f(a) + f(b)
	for i := 1; i < int(s.N/3+1); i++ {
		partialOut += 3 * f(a+float64(3*i-2)*h)
		partialOut += 3 * f(a+float64(3*i-1)*h)
		s.cycles++
	}
	for i := 1; i < int(s.N/3); i++ {
		partialOut += 2 * f(a+float64(3*i)*h)
		s.cycles++
	}
	return partialOut * h * 3 / 8
}

func (s *Simpson_3_8) handleInput() {
	if s.N == 0 {
		s.N = 3
	} else if s.N < 3 {
		panic("Simpson_3_8 struct value of N should be higher that 0")
	}
}
