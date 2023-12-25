package equation_test

import (
	"github.com/rocas777/kairos/equation"
	"math"
	"testing"
)

func smooth(x float64) float64 {
	return math.Pow(math.E, -math.Pow(x, 2)) - 0.3
}

func oscillatory(x float64) float64 {
	return math.Sin(x)
}

func exponential(x float64) float64 {
	return math.Pow(math.E, x) - 4
}

func singularity(x float64) float64 {
	return 1/x - 0.2
}

func dxSmooth(x float64) float64 {
	return -2 * x * math.Pow(math.E, -math.Pow(x, 2))
}

func dxOscillatory(x float64) float64 {
	return math.Cos(x)
}

func dxExponential(x float64) float64 {
	return math.Pow(math.E, x)
}

func dxSingularity(x float64) float64 {
	return -1 / (x * x)
}

func check(got float64, t *testing.T) {
	if math.Abs(got) > 0.1 || math.IsNaN(got) {
		t.Fatalf("Got: %f, wanted: 0", got)
	}
}

func TestBisection(t *testing.T) {
	a := 0.0
	b := 10.0

	tests := []struct {
		name string
		f    func(x float64) float64
		a    float64
		b    float64
	}{
		{"smooth", smooth, a, b},
		{"oscillatory", oscillatory, a, 4},
		{"exponential", exponential, a, b},
		{"singularity", singularity, a + 1, b},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(test.f(equation.NewBisection(0.001, 100).Zero(test.f, test.a, test.b)), t)
		})
	}
}

func TestFalsePosition(t *testing.T) {
	a := 0.0
	b := 10.0
	tests := []struct {
		name string
		f    func(x float64) float64
		a    float64
		b    float64
	}{
		{"smooth", smooth, a, b},
		{"oscillatory", oscillatory, a, 4},
		{"exponential", exponential, a, b},
		{"singularity", singularity, a + 1, b},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(test.f(equation.NewFalsePosition(0.001, 100).Zero(test.f, test.a, test.b)), t)
		})
	}
}

func TestSecant(t *testing.T) {
	a := 1.0
	tests := []struct {
		name string
		f    func(x float64) float64
		a    float64
		b    float64
	}{
		{"smooth", smooth, a, 3},
		{"oscillatory", oscillatory, a, 4},
		{"exponential", exponential, a, 10},
		{"singularity", singularity, a + 1, 10},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(test.f(equation.NewSecant(0.001, 100).Zero(test.f, test.a, test.b)), t)
		})
	}
}

func TestNewton(t *testing.T) {
	tests := []struct {
		name string
		f    func(x float64) float64
		dxF  func(x float64) float64
		a    float64
	}{
		{"smooth", smooth, dxSmooth, 1},
		{"oscillatory", oscillatory, dxOscillatory, 3},
		{"exponential", exponential, dxExponential, 1.5},
		{"singularity", singularity, dxSingularity, 0.5},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(test.f(equation.NewNewtonRaphson(0.001, 100).Zero(test.f, test.dxF, test.a)), t)
		})
	}
}
