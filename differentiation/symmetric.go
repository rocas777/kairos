package differentiation

type symmetric struct {
	h float64
	f func(x float64) float64
}

func NewSymmetric(h float64, f func(x float64) float64) *symmetric {
	return &symmetric{h: h, f: f}
}

func DefaultSymmetric(f func(x float64) float64) *symmetric {
	return NewSymmetric(0.1, f)
}

func (s *symmetric) LocalDerivative(x float64) float64 {
	return (s.f(x+s.h) - s.f(x-s.h)) / (s.h * 2)
}
