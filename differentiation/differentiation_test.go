package differentiation_test

import (
	"math"
	"numericalgo/pkg/differentiation"
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

func polynomial(x float64) float64 {
	return x * x * x
}

func dxPolynomial(x float64) float64 {
	return 3 * x * x
}
func dx2Polynomial(x float64) float64 {
	return 6 * x
}
func dx3Polynomial(x float64) float64 {
	return 6
}
func check(got, real float64, t *testing.T) {
	if math.Abs((got-real)/real) > 0.01 {
		t.Fatalf("Got: %f, wanted: %f -> %f, %f, %f", got, real, math.Abs(got-real), real, math.Abs(got-real)/real)
	}
}

func TestSimple(t *testing.T) {
	h := 0.001
	x := 3.0

	tests := []struct {
		name string
		f    func(x float64) float64
		dxF  func(x float64) float64
		h    float64
		x    float64
	}{
		{"smooth", smooth, dxSmooth, h, x},
		{"oscillatory", oscillatory, dxOscillatory, h, x},
		{"exponential", exponential, dxExponential, h, x},
		{"singularity", singularity, dxSingularity, h, x},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := differentiation.NewSimple(test.h, test.f)
			check(m.LocalDerivative(test.x), test.dxF(test.x), t)
		})
	}
}

func TestSymmetric(t *testing.T) {
	h := 0.001
	x := 3.0

	tests := []struct {
		name string
		f    func(x float64) float64
		dxF  func(x float64) float64
		h    float64
		x    float64
	}{
		{"smooth", smooth, dxSmooth, h, x},
		{"oscillatory", oscillatory, dxOscillatory, h, x},
		{"exponential", exponential, dxExponential, h, x},
		{"singularity", singularity, dxSingularity, h, x},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := differentiation.NewSymmetric(test.h, test.f)
			check(m.LocalDerivative(test.x), test.dxF(test.x), t)
		})
	}
}

func TestHigherOrder(t *testing.T) {
	h := 0.001
	x := 3.0

	tests := []struct {
		name  string
		f     func(x float64) float64
		dxF   func(x float64) float64
		h     float64
		x     float64
		order uint
	}{
		{"smooth", smooth, dxSmooth, h, x, 1},
		{"oscillatory", oscillatory, dxOscillatory, h, x, 1},
		{"exponential1", exponential, dxExponential, h, x, 1},
		{"exponential2", exponential, dxExponential, h, x, 2},
		{"exponential3", exponential, dxExponential, h, x, 3},
		{"exponential4", exponential, dxExponential, h, x, 4},
		{"singularity", singularity, dxSingularity, h, x, 1},
		{"polynomial1", polynomial, dxPolynomial, h, x, 1},
		{"polynomial2", polynomial, dx2Polynomial, h, x, 2},
		{"polynomial3", polynomial, dx3Polynomial, h, x, 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := differentiation.NewHigherOrder(test.h, test.f)
			check(m.LocalDerivative(test.x, test.order), test.dxF(test.x), t)
		})
	}
}
