package integration_test

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
	"math"
	"testing"
)

func smooth(x float64) float64 {
	return math.Pow(math.E, -math.Pow(x, 2))
}

func smoothSol() float64 {
	return math.Sqrt(math.Pi) / 2
}

func oscillatory(x float64) float64 {
	return math.Sin(x)
}

func oscillatorySol() float64 {
	return math.Cos(0) - math.Cos(10)
}

func exponential(x float64) float64 {
	return math.Pow(math.E, x)
}

func exponentialSol() float64 {
	return math.Pow(math.E, 10) - math.Pow(math.E, 0)
}

func simple(x float64) float64 {
	return 13
}

func simpleSol() float64 {
	return 130
}

func singularity(x float64) float64 {
	return 1 / x
}

func singularitySol() float64 {
	return math.Ln10
}

func check(got, real float64, t *testing.T) {
	if math.Abs((got-real)/real) > 0.1 {
		t.Fatalf("Got: %f, wanted: %f -> %f, %f, %f", got, real, math.Abs(got-real), real, math.Abs(got-real)/real)
	}
}

func TestTrapezoid(t *testing.T) {
	a := 0.0
	b := 10.0
	tests := []struct {
		name string
		f    func(x float64) float64
		a    float64
		b    float64
		sol  func() float64
	}{
		{"simple", simple, a, b, simpleSol},
		{"smooth", smooth, a, b, smoothSol},
		{"oscillatory", oscillatory, a, b, oscillatorySol},
		{"exponential", exponential, a, b, exponentialSol},
		{"singularity", singularity, a + 1, b, singularitySol},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(integration.NewTrapezoid(10).DefiniteIntegral(test.f, test.a, test.b), test.sol(), t)
		})
	}
}

func TestSimpson_1_3(t *testing.T) {
	a := 0.0
	b := 10.0
	tests := []struct {
		name string
		f    func(x float64) float64
		a    float64
		b    float64
		sol  func() float64
	}{
		{"simple", simple, a, b, simpleSol},
		{"smooth", smooth, a, b, smoothSol},
		{"oscillatory", oscillatory, a, b, oscillatorySol},
		{"exponential", exponential, a, b, exponentialSol},
		{"singularity", singularity, a + 1, b, singularitySol},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(integration.NewSimpson_1_3(10).DefiniteIntegral(test.f, test.a, test.b), test.sol(), t)
		})
	}

	for _, v := range integration.NewSimpson_1_3(10).AntiDerivative(smooth, -10, 10, 200) {
		fmt.Printf("%f %f\n", v.X, v.Y)
	}
}

func TestSimpson_3_8(t *testing.T) {
	a := 0.0
	b := 10.0
	tests := []struct {
		name string
		f    func(x float64) float64
		a    float64
		b    float64
		sol  func() float64
	}{
		{"simple", simple, a, b, simpleSol},
		{"smooth", smooth, a, b, smoothSol},
		{"oscillatory", oscillatory, a, b, oscillatorySol},
		{"exponential", exponential, a, b, exponentialSol},
		{"singularity", singularity, a + 1, b, singularitySol},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(integration.NewSimpson_3_8(10).DefiniteIntegral(test.f, test.a, test.b), test.sol(), t)
		})
	}
}

func TestSimpson_adaptive(t *testing.T) {
	a := 0.0
	b := 10.0

	tests := []struct {
		name string
		f    func(x float64) float64
		a    float64
		b    float64
		sol  func() float64
	}{
		{"simple", simple, a, b, simpleSol},
		{"smooth", smooth, a, b, smoothSol},
		{"oscillatory", oscillatory, a, b, oscillatorySol},
		{"exponential", exponential, a, b, exponentialSol},
		{"singularity", singularity, a + 1, b, singularitySol},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			check(integration.NewSimpsonAdaptive(0.0001).DefiniteIntegral(test.f, test.a, test.b), test.sol(), t)
		})
	}
}
