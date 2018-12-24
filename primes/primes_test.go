package primes

import "testing"

func TestDivisibleByAny(t *testing.T) {
	expectedToPass := []int{3, 5}

	if DivisibleByAny(expectedToPass, 2) {
		t.Errorf("Incorrect result when calculating 2 divisible by any of %v.", expectedToPass)
	}
}

func TestDivisibleByAnyWithDivisibleNumber(t *testing.T) {
	expectedToPass := []int{3, 5}

	if !DivisibleByAny(expectedToPass, 9) {
		t.Errorf("Incorrect result when calculating 9 divisible by any of %v.", expectedToPass)
	}
}

func TestNotDivisibleByAny(t *testing.T) {
	expectedToPass := []int{3, 5}

	if !NotDivisibleByAny(expectedToPass, 2) {
		t.Errorf("Incorrect result when calculating 2 not divisible by any of %v.", expectedToPass)
	}
}

func TestNotDivisibleByAnyWithDivisibleNumber(t *testing.T) {
	expectedToPass := []int{3, 5}

	if NotDivisibleByAny(expectedToPass, 9) {
		t.Errorf("Incorrect result when calculating 9 not divisible by any of %v.", expectedToPass)
	}
}

func TestIsPrimeNumber(t *testing.T) {
	if IsPrime(9) {
		t.Errorf("Incorrect result when calculating 9 being prime.")
	}
}
