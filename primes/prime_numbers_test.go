package primes

import "testing"

func TestNotDivisibleByAny(t *testing.T) {
	expectedToPass := []int{1, 3, 5}

	if NotDivisibleByAny(expectedToPass, 2) {
		t.Errorf("Incorrect result when calculating 2 not divisible by any of %v.", expectedToPass)
	}
}
