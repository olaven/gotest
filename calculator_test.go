package calculator

import (
	"testing"
)

type Container interface{}

/* func equals(test *testing.T, a, b Container) {

	if a != b {
		test.Error("Expected to be the same", a, b)
	}
} */

type ITest interface {
	equality(a, b Container)
}
type Test struct {
}

func (Test) equality(a, b Container) {
	if a != b {
		test.Error("Expected to be the same", a, b)
	}
}

func equalsWrapper(test *testing.T) func(a, b Container) {

	return func(a, b Container) {

		if a != b {
			test.Error("Expected to be the same", a, b)
		}
	}
}

func test(test *testing.T) ITest {

}

func TestAdd(test *testing.T) {

	equals(test, 6, Add(2, 4))
}

func TestAddNegative(test *testing.T) {

	equals(test, 2, Add(4, -2))
}

func TestAddTwoNegative(test *testing.T) {

	equals(test, -3, Add(-1, -2))
}

func TestSubtract(test *testing.T) {

	equals(test, 10, Subtract(20, 10))
}

func TestSubtractNegative(test *testing.T) {

	equals(test, -5, Subtract(5, 10))
}
