package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {

	Test(t).
		notEquals(Add(2, 2), Add(2, 1)).
		equals(6, Add(2, 4)).
		equals(2, Add(4, -2)).
		equals(-3, Add(-1, -2))
}

func TestSubtract(t *testing.T) {

	Test(t).
		equals(10, Subtract(20, 10)).
		equals(-5, Subtract(5, 10))
}

func TestMultiply(t *testing.T) {

	Test(t).
		equals(4, Multiply(2, 2)).
		equals(56, Multiply(7, 8)).
		equals(-20, Multiply(-10, 2)).
		equals(-20, Multiply(2, -10))
}

func TestDivide(t *testing.T) {

	Test(t).
		equals(float64(2), Divide(2, 1)).
		equals(float64(12), Divide(24, 2)).
		equals(2.5, Divide(10, 4)).
		equals(2.5, Divide(5, 2))
}
