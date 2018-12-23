package main

import "testing"

func TestHighestPrimeNumber(t *testing.T) {
	expected := 3
	sum := HighestPrimeNumberFor(6)
	if sum != expected {
		t.Errorf("Incorrect sum of array, got: %d, want: %d.", sum, expected)
	}
}

func TestNotDivisibleByAny(t *testing.T) {
	expectedToPass := []int{1, 3, 5}

	if NotDivisibleByAny(expectedToPass, 2) {
		t.Errorf("Incorrect result when calculating 2 not divisible by any of %v.", expectedToPass)
	}
}
