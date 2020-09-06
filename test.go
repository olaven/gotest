package calculator

import "testing"

type container interface{}

//Assertions defines possible assertion methods
type Assertions struct {
	equals func(a, b container) Assertions
}

//Test exposes assertion functions
func Test(t *testing.T) Assertions {

	equals := func(a, b container) Assertions {

		if a != b {
			t.Error("Expected", a, "to equal", b)
		}

		return Test(t)
	}

	return Assertions{
		equals,
	}
}
