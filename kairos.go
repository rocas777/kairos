// Package kairos provides utilities for mathematical computations and analyses related to calculus and equations.
// It consists of three subpackages: integration, equation, and differentiation.
//
// # Integration Package:
//
// The integration package offers methods to calculate definite integrals using different techniques.
// It includes the Trapezoidal Rule, Simpson's 1/3 Rule, Simpson's 3/8 Rule, and adaptive Simpson integration.
//
// # Equation Package:
//
// The equation package provides methods to find the zero of a given function using various root-finding algorithms.
// The supported methods are Bisection, FalsePosition, NewtonRaphson, and Secant.
//
// # Differentiation Package:
//
// The differentiation package offers methods to calculate derivatives of functions.
// It supports the calculation of the first derivative using two algorithms: Simple (based on the regular definition)
// and Symmetric (based on the symmetric definition). Additionally, it provides the ability to calculate arbitrary
// order derivatives using the HigherOrder method.
//
// The [kairos] package aims to assist users in performing mathematical computations with a focus on calculus and equation solving.
// It provides flexibility in choosing different methods depending on the specific requirements of the user's mathematical analysis.
//
// Note: Each subpackage within kairos has its own set of functions and structures, and users are encouraged to refer to the
// documentation of each subpackage for detailed information on usage and functionality.
package kairos

type Pair struct {
	X float64
	Y float64
}
