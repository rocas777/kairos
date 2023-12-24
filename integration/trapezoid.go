package integration

type trapezoid struct {
	N      uint
	cycles uint
	f      func(x float64) float64
}

func (t *trapezoid) Cycles() uint {
	return t.cycles
}

func NewTrapezoid(N uint, f func(x float64) float64) *trapezoid {
	return &trapezoid{N, 0, f}
}

func DefaultTrapezoid(f func(x float64) float64) *trapezoid {
	return NewTrapezoid(10, f)
}

func (t *trapezoid) DefiniteIntegral(a, b float64) float64 {
	t.cycles = 0
	h := (b - a) / float64(t.N)
	out := 0.0
	for i := 0; i < int(t.N); i++ {
		t.cycles++
		out += (t.f(a+float64(i)*h) + t.f(a+float64(i+1)*h)) / 2.0 * h
	}
	return out
}
