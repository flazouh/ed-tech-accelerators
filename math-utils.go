// Package mathutils provides basic mathematical utility functions.
// This package includes arithmetic operations, comparison functions,
// and mathematical utilities with proper error handling.
package main

import (
	"errors"
	"fmt"
)

// ErrDivisionByZero is returned when attempting to divide by zero.
var ErrDivisionByZero = errors.New("division by zero")

// ErrNegativeFactorial is returned when calculating factorial of a negative number.
var ErrNegativeFactorial = errors.New("factorial of negative number is undefined")

// Add performs addition of two float64 numbers.
// Returns the sum of a and b.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction of two float64 numbers.
// Returns the difference of a and b (a - b).
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication of two float64 numbers.
// Returns the product of a and b.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division of two float64 numbers.
// Returns the quotient of a and b, or an error if b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Max returns the larger of two float64 numbers.
// Returns a if a >= b, otherwise returns b.
func Max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

// Min returns the smaller of two float64 numbers.
// Returns a if a <= b, otherwise returns b.
func Min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

// Abs returns the absolute value of a float64 number.
// Returns the non-negative value of x.
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("Math Utils Package - Basic Mathematical Operations")
	
	// Example usage of arithmetic functions
	fmt.Printf("Add(10, 5) = %.2f\n", Add(10, 5))
	fmt.Printf("Subtract(10, 5) = %.2f\n", Subtract(10, 5))
	fmt.Printf("Multiply(10, 5) = %.2f\n", Multiply(10, 5))
	
	result, err := Divide(10, 5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Divide(10, 5) = %.2f\n", result)
	}
	
	// Example of division by zero error
	_, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// Example usage of utility functions
	fmt.Printf("Max(10, 5) = %.2f\n", Max(10, 5))
	fmt.Printf("Min(10, 5) = %.2f\n", Min(10, 5))
	fmt.Printf("Abs(-7.5) = %.2f\n", Abs(-7.5))
	fmt.Printf("Abs(3.2) = %.2f\n", Abs(3.2))
}