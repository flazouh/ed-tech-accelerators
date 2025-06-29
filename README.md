# Math Utils - Comprehensive Math Utilities Library

A comprehensive math utilities library implemented in Rust, providing:

## Features

### Basic Arithmetic
- Addition, subtraction, multiplication, division
- Safe division with error handling

### Advanced Mathematics  
- Power calculations
- Square root with error handling
- Factorial computation
- Greatest Common Divisor (GCD)
- Least Common Multiple (LCM)

### Statistics
- Mean (average) calculation
- Median calculation
- Mode (most frequent value)
- Standard deviation

### Geometry
- Circle area and circumference
- Rectangle area and perimeter  
- Triangle area using Heron's formula
- Input validation for all geometric functions

## Usage

The main implementation is in `math-utils.rust` with modular organization:

```rust
use math_utils::*;

// Basic arithmetic
let sum = arithmetic::add(5.0, 3.0);
let quotient = arithmetic::divide(10.0, 2.0).unwrap();

// Advanced functions
let power = advanced::power(2.0, 3.0);
let factorial = advanced::factorial(5);

// Statistics
let data = vec![1.0, 2.0, 3.0, 4.0, 5.0];
let mean = statistics::mean(&data).unwrap();

// Geometry
let area = geometry::circle_area(5.0).unwrap();
```

## Testing

Run tests with:
```bash
cargo test
```

## Requirements

- Rust 1.70.0 or later
- Cargo package manager