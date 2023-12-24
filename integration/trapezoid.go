package integration

type Trapezoid struct {
	N      uint
	cycles uint
}

func NewTrapezoid(n uint) *Trapezoid {
	return &Trapezoid{N: n}
}

func (t *Trapezoid) Cycles() uint {
	return t.cycles
}

func (t *Trapezoid) DefiniteIntegral(f func(x float64) float64, a, b float64) float64 {
	t.cycles = 0
	h := (b - a) / float64(t.N)
	out := 0.0
	for i := 0; i < int(t.N); i++ {
		t.cycles++
		out += (f(a+float64(i)*h) + f(a+float64(i+1)*h)) / 2.0 * h
	}
	return out
}
