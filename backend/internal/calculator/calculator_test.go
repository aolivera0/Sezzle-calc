package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 2)

	expected := float64(5)

	if result != expected {
		t.Errorf("Add(3,2) = %v; want %v", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(3, 5)

	expected := float64(-2)

	if result != expected {
		t.Errorf("Subtract(3,5) = %v; want %v", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, -11)

	expected := float64(-55)

	if result != expected {
		t.Errorf("Multiply(5,-11) = %v; want %v", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(4, 2)

	if err != nil {
		t.Fatalf("Divide(4,2) returned an unexpected error: %v", err)
	}

	expected := float64(2)

	if result != expected {
		t.Errorf("Divide(4,2) = %v; want %v", result, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(3, 0)

	if err == nil {
		t.Fatalf("Divide by 0 should not be allowed")
	}
}

func TestPower(t *testing.T) {
	result := Power(3, 2)

	expected := float64(9)

	if result != expected {
		t.Errorf(`Power(a,b)=%v, want %v`, result, expected)
	}
}

func TestSqrt(t *testing.T) {
	result, err := Sqrt(9)

	if err != nil {
		t.Fatalf("Sqrt(9) returned an unexpected error: %v", err)
	}

	expected := float64(3)

	if result != expected {
		t.Errorf("Sqrt(9) = %v; want %v", result, expected)
	}
}

func TestSqrtOfNegativeNumber(t *testing.T) {
	_, err := Sqrt(-2)

	if err == nil {
		t.Fatalf("Sqrt of -2 should not be allowed")
	}
}

func TestPercentage(t *testing.T) {
	result := Percentage(10, 50)

	expected := float64(5)

	if result != expected {
		t.Errorf(`Percentage(10,50)=%v, want %v`, result, expected)
	}
}
