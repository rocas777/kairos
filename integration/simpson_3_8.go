package integration

type simpson_3_8 struct {
	n      uint
	cycles uint
	f      func(x float64) float64
}

func (s *simpson_3_8) Cycles() uint {
	return s.cycles
}

func NewSimpson_3_8(N uint, f func(x float64) float64) *simpson_3_8 {
	return &simpson_3_8{N * 3, 0, f}
}

func DefaultSimpson_3_8(f func(x float64) float64) *simpson_3_8 {
	return NewSimpson_3_8(9, f)
}

func (s *simpson_3_8) DefiniteIntegral(a, b float64) float64 {
	s.cycles = 0
	h := (b - a) / float64(s.n)
	partial_out := s.f(a) + s.f(b)
	for i := 1; i < int(s.n/3+1); i++ {
		partial_out += 3 * s.f(a+float64(3*i-2)*h)
		partial_out += 3 * s.f(a+float64(3*i-1)*h)
		s.cycles++
	}
	for i := 1; i < int(s.n/3); i++ {
		partial_out += 2 * s.f(a+float64(3*i)*h)
		s.cycles++
	}
	return partial_out * h * 3 / 8
}
