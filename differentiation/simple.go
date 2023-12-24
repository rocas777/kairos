package differentiation

type simple struct {
	h float64
	f func(x float64) float64
}

func NewSimple(h float64, f func(x float64) float64) *simple {
	return &simple{h, f}
}

func DefaultSimple(f func(x float64) float64) *simple {
	return NewSimple(0.1, f)
}

func (s *simple) LocalDerivative(x float64) float64 {
	return (s.f(x+s.h) - s.f(x)) / s.h
}
