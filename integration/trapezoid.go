package integration

import "github.com/rocas777/kairos"

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

func (t *Trapezoid) AntiDerivative(f func(x float64) float64, a, b float64, samples uint) []kairos.Pair {
	t.handleInput()
	if samples < 2 {
		samples = 2
	}

	y := 0.0
	out := make([]kairos.Pair, samples)
	sampleH := (b - a) / float64(samples-1)
	for i := 0; i < int(samples); i++ {
		x := a + float64(i)*sampleH
		y = t.DefiniteIntegral(f, 0, x)
		out[i] = kairos.Pair{X: x, Y: y}
	}
	return out
}

func (t *Trapezoid) handleInput() {
	if t.N == 0 {
		t.N = 10
	} else if t.N < 0 {
		panic("Trapezoid struct value of N should be higher that 0")
	}

}
