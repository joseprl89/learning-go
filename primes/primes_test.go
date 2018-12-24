package primes

import "testing"

func TestDivisibleByAny(t *testing.T) {
	expectedToPass := []int{3, 5}

	if DivisibleByAny(expectedToPass, 2) {
		t.Errorf("Incorrect result when calculating 2 not divisible by any of %v.", expectedToPass)
	}
}

func TestDivisibleByAnyWithDivisibleNumber(t *testing.T) {
	expectedToPass := []int{3, 5}

	if !DivisibleByAny(expectedToPass, 9) {
		t.Errorf("Incorrect result when calculating 9 not divisible by any of %v.", expectedToPass)
	}
}
