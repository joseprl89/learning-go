package main

import "testing"

func TestHighestPrimeNumber(t *testing.T) {
	expected := 3
	sum := HighestPrimeNumberFor(6)
	if sum != expected {
		t.Errorf("Incorrect sum of array, got: %d, want: %d.", sum, expected)
	}

}
