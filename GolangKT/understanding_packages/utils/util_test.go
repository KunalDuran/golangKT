package utils

import "testing"

func TestAdd(t *testing.T) {

	a, b := 1, 3
	expectedResult := 4

	result := Add(a, b)

	if result != expectedResult {
		t.Error("Unexpected result")
	}

}
