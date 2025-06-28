package main

import (
	"math"
	"testing"
)

// TestAdd tests the Add function with various inputs
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 8.0},
		{"negative numbers", -5.0, -3.0, -8.0},
		{"mixed signs", 5.0, -3.0, 2.0},
		{"zero values", 0.0, 5.0, 5.0},
		{"decimal numbers", 2.5, 3.7, 6.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSubtract tests the Subtract function with various inputs
func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 8.0, 3.0, 5.0},
		{"negative numbers", -5.0, -3.0, -2.0},
		{"mixed signs", 5.0, -3.0, 8.0},
		{"zero values", 5.0, 0.0, 5.0},
		{"decimal numbers", 5.7, 2.5, 3.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMultiply tests the Multiply function with various inputs
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 15.0},
		{"negative numbers", -5.0, -3.0, 15.0},
		{"mixed signs", 5.0, -3.0, -15.0},
		{"zero multiplication", 5.0, 0.0, 0.0},
		{"decimal numbers", 2.5, 4.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestDivide tests the Divide function with various inputs
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		shouldError bool
	}{
		{"positive numbers", 15.0, 3.0, 5.0, false},
		{"negative numbers", -15.0, -3.0, 5.0, false},
		{"mixed signs", 15.0, -3.0, -5.0, false},
		{"decimal numbers", 7.5, 2.5, 3.0, false},
		{"division by zero", 5.0, 0.0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			
			if tt.shouldError {
				if err == nil {
					t.Errorf("Divide(%v, %v) should return an error", tt.a, tt.b)
				}
				if err != ErrDivisionByZero {
					t.Errorf("Divide(%v, %v) should return ErrDivisionByZero, got %v", tt.a, tt.b, err)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%v, %v) returned unexpected error: %v", tt.a, tt.b, err)
				}
				if math.Abs(result-tt.expected) > 1e-9 {
					t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

// TestMax tests the Max function with various inputs
func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"a greater than b", 10.0, 5.0, 10.0},
		{"b greater than a", 5.0, 10.0, 10.0},
		{"equal values", 7.0, 7.0, 7.0},
		{"negative numbers", -5.0, -10.0, -5.0},
		{"mixed signs", -5.0, 3.0, 3.0},
		{"decimal numbers", 3.14, 2.71, 3.14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMin tests the Min function with various inputs
func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"a smaller than b", 5.0, 10.0, 5.0},
		{"b smaller than a", 10.0, 5.0, 5.0},
		{"equal values", 7.0, 7.0, 7.0},
		{"negative numbers", -5.0, -10.0, -10.0},
		{"mixed signs", -5.0, 3.0, -5.0},
		{"decimal numbers", 3.14, 2.71, 2.71},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Min(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestAbs tests the Abs function with various inputs
func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"positive number", 5.0, 5.0},
		{"negative number", -5.0, 5.0},
		{"zero", 0.0, 0.0},
		{"positive decimal", 3.14, 3.14},
		{"negative decimal", -3.14, 3.14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.input)
			if result != tt.expected {
				t.Errorf("Abs(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestFactorial tests the Factorial function with various inputs
func TestFactorial(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    int64
		shouldError bool
	}{
		{"factorial of 0", 0, 1, false},
		{"factorial of 1", 1, 1, false},
		{"factorial of 5", 5, 120, false},
		{"factorial of 10", 10, 3628800, false},
		{"negative number", -1, 0, true},
		{"large negative", -10, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Factorial(tt.input)
			
			if tt.shouldError {
				if err == nil {
					t.Errorf("Factorial(%v) should return an error", tt.input)
				}
				if err != ErrNegativeFactorial {
					t.Errorf("Factorial(%v) should return ErrNegativeFactorial, got %v", tt.input, err)
				}
			} else {
				if err != nil {
					t.Errorf("Factorial(%v) returned unexpected error: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("Factorial(%v) = %v; want %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

// TestFactorialLargeValues tests factorial with larger values to verify no overflow for reasonable inputs
func TestFactorialLargeValues(t *testing.T) {
	// Test some larger values that should still work with int64
	tests := []struct {
		input    int
		expected int64
	}{
		{12, 479001600},
		{15, 1307674368000},
	}

	for _, tt := range tests {
		t.Run("large factorial", func(t *testing.T) {
			result, err := Factorial(tt.input)
			if err != nil {
				t.Errorf("Factorial(%v) returned unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Factorial(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}