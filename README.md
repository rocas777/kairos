![Automation Badge](https://github.com/rocas777/kairos/actions/workflows/go.yml/badge.svg)

# Kairos - numerical calculus

The Kairos project is a versatile library developed in [Go](https://go.dev/) that provides comprehensive tools for numerical computations in the domains of integration, equation solving, and differentiation. Each subpackage within Kairos is designed to offer specialized functionality, allowing users to perform mathematical operations with ease and precision.

**The key features of Kairos:**

- Differentiation
- Integration
- Equation Solver


# Index

1. [Kairos - Numerical Calculus](#kairos---numerical-calculus)
    1. [Getting started](#getting-started)
2. [Kairos: Differentiation Package](#kairos-differentiation-package)
    1. [Simple Derivative](#simple-derivative)
        1. [Local Derivative](#local-derivative)
        2. [Range Derivative](#range-derivative)
    2. [Symmetric Derivative](#symmetric-derivative)
        1. [Local Derivative](#local-derivative-1)
        2. [Range Derivative](#range-derivative-1)
    3. [Higher Order Derivative](#higher-order-derivative)
        1. [Local Derivative](#local-derivative-2)
        2. [Range Derivative](#range-derivative-2)
3. [Kairos: Equation Solver Package](#kairos-equation-solver-package)
    1. [Bisection](#bisection)
    2. [False Position (Regula False)](#falseposition)
    3. [NewtonRaphson](#newtonraphson)
    4. [Secant](#secant)
4. [Kairos: Integration Package](#kairos-integration-package-)
    1. [Trapezoidal Rule](#trapezoid-rule)
        1. [Definite Integral](#definite-integral)
        2. [Anti-Derivative](#anti-derivative)
    2. [Simpson's 1/3 Rule](#simpson-13-rule)
        1. [Definite Integral](#definite-integral-1)
        2. [Anti-Derivative](#anti-derivative-1)
    3. [Simpson's 3/8 Rule](#simpson-38-rule)
        1. [Definite Integral](#definite-integral-2)
        2. [Anti-Derivative](#anti-derivative-2)
    4. [Adaptative Simpson Rule](#adaptive-simpson-integration)
        1. [Definite Integral](#definite-integral-3)
        2. [Anti-Derivative](#anti-derivative-3)
5.  [Documentation Reference](#documentation-reference)


## Getting started

### Getting Kairos

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/rocas777/kairos"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `kairos` package:

```sh
$ go get -u github.com/rocas777/kairos
```

# Kairos: Differentiation Package

The `differentiation` package in the Kairos library equips Go developers with utilities for calculating derivatives of functions. Whether you need a first-order derivative or an arbitrary nth-order derivative, Kairos has you covered.

## Overview

- **First Order Derivatives:**
    - **[Simple Algorithm](#simple-derivative):** Based on the regular derivative definition.
    - **[Symmetric Algorithm](#symmetric-derivative):** Based on the symmetric derivative definition.

- **Arbitrary Order Derivatives:**
    - **[HigherOrder Method](#higher-order-derivative):** Utilizes the symmetric algorithm recursively to calculate nth-order derivatives.

Users can choose the method that best suits their accuracy and efficiency requirements.

## Simple Derivative

The `Simple` struct provides methods for calculating the first derivative based on the regular definition. It uses the limit concept to approximate infinitesimals with 'H'. The derivative is computed as the slope of the function between points 'x' and 'x + H'.

### Local derivative
```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/differentiation"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Simple instance with default H value (0.1)
	simple := differentiation.NewSimple(0.1)

	// Calculate the first order derivative at the point x = 2
	result := simple.LocalDerivative(f, 2)
	fmt.Println("First Derivative at x = 2:", result)
}
```

### Range Derivative

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/differentiation"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Simple instance with default H value (0.1)
	simple := differentiation.NewSimple(0.1)

	// Calculate the first order derivative over the range [0, 2] with 5 samples
	rangeDerivative := simple.RangeDerivative(f, 0, 2, 5)
	fmt.Println("First Derivative over the Range [0, 2]:", rangeDerivative)
}
```


## Symmetric Derivative

The `Symmetric` struct provides methods for calculating the first derivative based on the symmetric definition. It uses 'H' as an approximation of infinitesimals and computes the derivative as the slope between points 'x - H' and 'x + H'.
### Local derivative
### Usage

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/differentiation"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Symmetric instance with default H value (0.1)
	symmetric := differentiation.NewSymmetric(0.1)

	// Calculate the first order derivative at the point x = 2
	result := symmetric.LocalDerivative(f, 2)
	fmt.Println("First Derivative at x = 2:", result)
}
```

### Range Derivative
### Usage

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/differentiation"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Symmetric instance with default H value (0.1)
	symmetric := differentiation.NewSymmetric(0.1)

	// Calculate the first order derivative over the range [0, 2] with 5 samples
	rangeDerivative := symmetric.RangeDerivative(f, 0, 2, 5)
	fmt.Println("First Derivative over the Range [0, 2]:", rangeDerivative)
}
```

## Higher Order Derivative

The `HigherOrder` struct contains methods for calculating nth-order derivatives. It utilizes the symmetric algorithm recursively to achieve higher-order derivatives.


### Local derivative
```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/differentiation"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new HigherOrder instance with default H value (0.1) and Order (2)
	higherOrder := differentiation.NewHigherOrder(0.1, 2)

	// Calculate the first order derivative at the point x = 2
	result := higherOrder.LocalDerivative(f, 2)
	fmt.Println("First Derivative at x = 2:", result)
}

```

### Range Derivative

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/differentiation"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new HigherOrder instance with default H value (0.1) and Order (2)
	higherOrder := differentiation.NewHigherOrder(0.1, 2)

	// Calculate the first order derivative over the range [0, 2] with 5 samples
	rangeDerivative := higherOrder.RangeDerivative(f, 0, 2, 5)
	fmt.Println("First Derivative over the Range [0, 2]:", rangeDerivative)
}
```






















































# Kairos: Equation Solver Package

The `equation` package in the Kairos library provides utilities for solving equations and finding the roots of functions. It offers multiple root-finding methods, including the Bisection method, False Position method, Newton-Raphson method, and Secant method. Users can choose the most suitable method for their specific functions and interval constraints to efficiently locate zeros of the given function.
## Overview

- [Bisection](#bisection)
- [FalsePosition](#falseposition)
- [NewtonRaphson](#newtonraphson)
- [Secant](#secant)


**Note:** These methods assume the provided function is continuous on the considered interval.
## Bisection

The `Bisection` struct provides a method to find the zero of a function using the [Bisection](https://en.wikipedia.org/wiki/Bisection_method) method on an interval [a, b]. The method can be limited by CycleLimit, which restricts the number of cycles to prevent the algorithm from running indefinitely. A solution is considered definitive once the difference of the interval [a, b] is below Epsilon.

### Usage

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/equation"
)

func main() {
	// Example function: f(x) = x^2 - 4
	f := func(x float64) float64 {
		return x*x - 4
	}

	// Create a new Bisection instance with default Epsilon (0.01) and CycleLimit (100)
	bisection := equation.NewBisection(0.01, 100)

	// Find the zero of the function on the interval [1, 3]
	result := bisection.Zero(f, 1, 3)
	fmt.Println("Zero of the function:", result)
}
```


## FalsePosition

The `FalsePosition` struct provides a method to find the zero of a function using the [False Position](https://en.wikipedia.org/wiki/Regula_falsi) method on an interval [a, b]. The method iteratively refines the estimate of the zero based on linear interpolation.

### Usage

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/equation"
)

func main() {
	// Example function: f(x) = x^2 - 4
	f := func(x float64) float64 {
		return x*x - 4
	}

	// Create a new FalsePosition instance with default Epsilon (0.01) and CycleLimit (100)
	falsePosition := equation.NewFalsePosition(0.01, 100)

	// Find the zero of the function on the interval [1, 3]
	result := falsePosition.Zero(f, 1, 3)
	fmt.Println("Zero of the function:", result)
}
```

## NewtonRaphson

The `NewtonRaphson` struct provides a method to find the zero of a function using the [Newton-Raphson](https://en.wikipedia.org/wiki/Newton%27s_method) method. The method iteratively refines the estimate of the zero based on the function's local behavior.

### Usage

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/equation"
)

func main() {
	// Example function: f(x) = x^2 - 4
	f := func(x float64) float64 {
		return x*x - 4
	}

	// Example derivative function: f'(x) = 2*x
	dxF := func(x float64) float64 {
		return 2 * x
	}

	// Create a new NewtonRaphson instance with default Epsilon (0.01) and CycleLimit (100)
	newtonRaphson := equation.NewNewtonRaphson(0.01, 100)

	// Find the zero of the function using the derivative on the initial estimate 3
	result := newtonRaphson.Zero(f, dxF, 3)
	fmt.Println("Zero of the function:", result)
}
```

## Secant

The Secant struct provides a method to find the zero of a function using the [Secant](https://en.wikipedia.org/wiki/Secant_method) method. The method iteratively refines the estimate of the zero based on a secant line between two points.

### Usage

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/equation"
)

func main() {
	// Example function: f(x) = x^2 - 4
	f := func(x float64) float64 {
		return x*x - 4
	}

	// Create a new Secant instance with default Epsilon (0.01) and CycleLimit (100)
	secant := equation.NewSecant(0.01, 100)

	// Find the zero of the function on the interval [1, 3]
	result := secant.Zero(f, 1, 3)
	fmt.Println("Zero of the function:", result)
}

```





















































































# Kairos: Integration Package 

The `integration` package in the Kairos library provides utilities for numerical integration of functions. It includes several methods for calculating definite integrals, such as the Trapezoidal Rule, Simpson's 1/3 Rule, Simpson's 3/8 Rule, and adaptive Simpson integration. Users can choose the appropriate method based on the precision and efficiency requirements of their mathematical analysis.

## Overview

- [Trapezoid rule](#trapezoid-rule)
- [Simpson 1/3 Rule](#simpson-13-rule)
- [Simpson 3/8 Rule](#falseposition)
- [Simpson 1/3 adaptative](#newtonraphson)



## Trapezoid Rule

The `Trapezoid` struct provides a method to calculate the definite integral of a given function using the [Trapezoidal](https://en.wikipedia.org/wiki/Trapezoidal_rule) Rule. This rule approximates the integral by dividing the interval into N trapezoids and summing their areas. The higher the value of N, the more accurate the integration will be; however, this will lead to a more time-expensive method.
### Definite Integral
```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Trapezoid instance with default N (10)
	trapezoid := integration.NewTrapezoid(10)

	// Calculate the definite integral of the function over the interval [0, 1]
	result := trapezoid.DefiniteIntegral(f, 0, 1)
	fmt.Println("Definite Integral:", result)
}
```

### Anti-derivative

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Trapezoid instance with default N (10)
	trapezoid := integration.NewTrapezoid(10)

	// Calculate the approximate antiderivative of the function over the interval [0, 1] with 5 samples
	result := trapezoid.AntiDerivative(f, 0, 1, 5)
	fmt.Println("Approximate Anti-Derivative:", result)
}
```






## Simpson 1/3 Rule

The `Simpson_1_3` struct provides a method to calculate the definite integral of a given function using the [Simpson](https://en.wikipedia.org/wiki/Simpson%27s_rule#Composite_Simpson's_1/3_rule) algorithm, specifically the 1/3 composite rule. This works by dividing the interval into N pieces and making a polynomial interpolation between two successive points to calculate the area below the interpolation. The higher the value of N, the more accurate the integration will be; however, this will lead to a more time-expensive method.

### Definite Integral
```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Simpson_1_3 instance with default N (2)
	simpson13 := integration.NewSimpson_1_3(2)

	// Calculate the definite integral of the function over the interval [0, 1]
	result := simpson13.DefiniteIntegral(f, 0, 1)
	fmt.Println("Definite Integral:", result)
}
```

### Anti-derivative

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Simpson_1_3 instance with default N (2)
	simpson13 := integration.NewSimpson_1_3(2)

	// Calculate the approximate antiderivative of the function over the interval [0, 1] with 5 samples
	result := simpson13.AntiDerivative(f, 0, 1, 5)
	fmt.Println("Approximate Anti-Derivative:", result)
}
```


## Simpson 3/8 Rule

The `Simpson_3_8` struct provides a method to calculate the definite integral of a given function using the [Simpson algorithm](https://en.wikipedia.org/wiki/Simpson%27s_rule#Composite_Simpson's_3/8_rule), specifically the 3/8 composite rule. This works by dividing the interval into N pieces and making a polynomial interpolation using three successive points to calculate the area below the interpolation. The higher the value of N, the more accurate the integration will be; however, this will lead to a more time-expensive method.

### Definite Integral
```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Simpson_3_8 instance with default N (3)
	simpson38 := integration.NewSimpson_3_8(3)

	// Calculate the definite integral of the function over the interval [0, 1]
	result := simpson38.DefiniteIntegral(f, 0, 1)
	fmt.Println("Definite Integral:", result)
}
```

### Anti-derivative

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new Simpson_3_8 instance with default N (3)
	simpson38 := integration.NewSimpson_3_8(3)

	// Calculate the approximate antiderivative of the function over the interval [0, 1] with 5 samples
	result := simpson38.AntiDerivative(f, 0, 1, 5)
	fmt.Println("Approximate Anti-Derivative:", result)
}
```





## Adaptive Simpson Integration

The `SimpsonAdaptive` struct provides a method to calculate the definite integral of a given function using adaptive Simpson integration. It combines the simplicity of the [Simpson 1/3 Rule](https://en.wikipedia.org/wiki/Simpson%27s_rule#Composite_Simpson's_1/3_rule) with adaptivity to improve accuracy. The algorithm automatically adjusts the number of intervals based on the function's behavior, comparing the results by dividing intervals into more subintervals using the epsilon criterion.
### Definite Integral
```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new SimpsonAdaptive instance with default Epsilon (0.1)
	simpsonAdaptive := integration.NewSimpsonAdaptive(0.1)

	// Calculate the definite integral of the function over the interval [0, 1]
	result := simpsonAdaptive.DefiniteIntegral(f, 0, 1)
	fmt.Println("Definite Integral:", result)
}
```

### Anti-derivative

```go
package main

import (
	"fmt"
	"github.com/rocas777/kairos/integration"
)

func main() {
	// Example function: f(x) = x^2
	f := func(x float64) float64 {
		return x * x
	}

	// Create a new SimpsonAdaptive instance with default Epsilon (0.1)
	simpsonAdaptive := integration.NewSimpsonAdaptive(0.1)

	// Calculate the approximate antiderivative of the function over the interval [0, 1] with 5 samples
	result := simpsonAdaptive.AntiDerivative(f, 0, 1, 5)
	fmt.Println("Approximate Anti-Derivative:", result)
}
```

# Documentation Reference

For detailed documentation and examples, please refer to the official documentation on [pkg.go.dev](https://pkg.go.dev/github.com/rocas777/kairos).


