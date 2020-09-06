package gotest

import "testing"

//TODO: figure out how to check if test failed. I.e. "pass on failure"
/* func TestEqualAssertion(t *testing.T) {

	//should pass
	mockT := new(testing.T)
	_, equal := Test(mockT).equals(2, 2)

	if !equal {
		t.Error("2 should equal 2")
	}
}

func TestEqualityFailure(t *testing.T) {

	mockT := new(testing.T)
	_, equal := Test(mockT).equals(2, 1)

	if equal {
		t.Error("2 should not be equal to 1")
	}
} */

func TestUsage(t *testing.T) {

	Test(t).
		equals(2, 2).
		equals("dog", "dog").
		notEquals(2, 1).
		notEquals("first", "second")
}
