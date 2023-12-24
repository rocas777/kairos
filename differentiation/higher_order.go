package differentiation

type higherOrder struct {
	h float64
	f func(x float64) float64
}

func NewHigherOrder(h float64, f func(x float64) float64) *higherOrder {
	return &higherOrder{h, f}
}

func DefaultHigherOrder(f func(x float64) float64) *higherOrder {
	return NewHigherOrder(0.1, f)
}

func (s *higherOrder) LocalDerivative(x float64, order uint) float64 {
	if order == 1 {
		return NewSymmetric(s.h, s.f).LocalDerivative(x)
	}
	return (s.LocalDerivative(x+s.h, order-1) - s.LocalDerivative(x-s.h, order-1)) / (s.h * 2)
}
