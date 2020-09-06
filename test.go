package calculator

import "testing"

type container interface{}

//Assertions defines possible assertion methods
type Assertions struct {
	equals    func(a, b container) Assertions
	notEquals func(a, b container) Assertions
}

func equals(t *testing.T) func(a, b container) Assertions {

	return func(a, b container) Assertions {

		if a != b {
			t.Error("Expected", a, "to equal", b)
		}

		return Test(t)
	}
}

func notEquals(t *testing.T) func(a, b container) Assertions {

	return func(a, b container) Assertions {

		if a == b {
			t.Error("Expected", a, "_not_ to equal", b)
		}

		return Test(t)
	}
}

//Test exposes assertion functions
func Test(t *testing.T) Assertions {

	return Assertions{
		equals:    equals(t),
		notEquals: notEquals(t),
	}
}
