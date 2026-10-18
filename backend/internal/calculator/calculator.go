package calculator

import (
	"errors"
	"math"
)

func Add(a float64, b float64) float64 {
	return a + b
}

func Subtract(a float64, b float64) float64 {
	return a - b
}

func Multiply(a float64, b float64) float64 {
	return a * b
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by 0 is not allowed")
	}

	return a / b, nil
}

func Power(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return a, errors.New("Square root of negative number is not allowed")
	}

	return math.Sqrt(a), nil
}

func Percentage(a float64, b float64) float64 {
	return a * b / 100
}
