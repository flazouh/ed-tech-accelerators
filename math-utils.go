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

func main() {
	fmt.Println("Math Utils Package - Basic Mathematical Operations")
}