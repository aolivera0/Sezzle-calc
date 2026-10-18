package calculator

import "testing"

func TestAdd(t *testing.T){
	result := Add(3,2)

	expected := float64(5)

	if result != expected {
		t.Errorf("Add(3,2) = %v; want %v", result, expected)
	}
}

func TestSubtract(t *testing.T){
	result := Subtract(3,5)

	expected := float64(-2)

	if result != expected {
		t.Errorf("Subtract(3,5) = %v; want %v", result, expected)
	}
}

func TestMultiply(t *testing.T){
	result := Multiply(5,-11)

	expected := float64(-55)

	if result != expected {
		t.Errorf("Multiply(5,-11) = %v; want %v", result, expected)
	}
}

func TestDivide(t *testing.T){
	result, err := Divide(4, 2)

	if err != nil {
		t.Fatalf("Divide(4,2) returned an unexpected error: %v", err)
	}

	expected := float64(2)

	if result != expected {
		t.Errorf("Divide(4,2) = %v; want %v", result, expected)
	}
}

func TestDivideByZero(t *testing.T){
	_, err := Divide(3,0)

	if err == nil {
		t.Fatalf("Divide by 0 should not be allowed")
	}
}