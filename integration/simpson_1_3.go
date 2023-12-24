package integration

type simpson_1_3 struct {
	n      uint
	cycles uint
	f      func(x float64) float64
}

func (s *simpson_1_3) Cycles() uint {
	return s.cycles
}

func NewSimpson_1_3(N uint, f func(x float64) float64) *simpson_1_3 {
	return &simpson_1_3{N * 2, 0, f}
}

func DefaultSimpson_1_3(f func(x float64) float64) *simpson_1_3 {
	return NewSimpson_1_3(10, f)
}

func (s *simpson_1_3) DefiniteIntegral(a, b float64) float64 {
	s.cycles = 0
	h := (b - a) / float64(s.n)
	partialOut := s.f(a) + s.f(b)
	for i := 1; i < int(s.n/2+1); i++ {
		s.cycles++
		partialOut += 4 * s.f(a+float64(2*i-1)*h)
	}
	for i := 1; i < int(s.n/2); i++ {
		s.cycles++
		partialOut += 2 * s.f(a+float64(2*i)*h)
	}
	return partialOut * h / 3
}
