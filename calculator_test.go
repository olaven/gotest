package calculator

import (
	"testing"
)

type Container interface{}

type Test struct {
	equals func(a, b Container) Test
}

func test(t *testing.T) Test {

	equals := func(a, b Container) Test {

		if a != b {
			t.Error("Expected", a, "to equal", b)
		}

		return test(t)
	}

	return Test{
		equals,
	}
}
func TestAdd(t *testing.T) {

	test(t).
		equals(6, Add(2, 4)).
		equals(2, Add(4, -2)).
		equals(-3, Add(-1, -2))
}

func TestSubtract(t *testing.T) {

	test(t).
		equals(10, Subtract(20, 10)).
		equals(-5, Subtract(5, 10))
}
