package main

import "testing"

func TestMultiples(t *testing.T) {
	expected := []int{2, 4}
	multiples := FindMultiplesOf(2, 4)
	if !IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestArraySum(t *testing.T) {
	expected := 6
	sum := SumOf([]int{2, 4})
	if sum != expected {
		t.Errorf("Incorrect sum of array, got: %d, want: %d.", sum, expected)
	}
}

func TestSumOfMultiples(t *testing.T) {
	expected := 6
	sum := SumOfMultiples(2, 4)
	if sum != expected {
		t.Errorf("Incorrect sum of multiples, got: %d, want: %d.", sum, expected)
	}
}

func IsEqual(lhs []int, rhs []int) bool {
	if (lhs == nil) != (rhs == nil) {
		return false
	}

	if len(lhs) != len(rhs) {
		return false
	}

	for i := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}
