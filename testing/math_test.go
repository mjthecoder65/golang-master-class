package testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 5)
	expected := 7

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}

}
