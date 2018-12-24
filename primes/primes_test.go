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

	if !IsPrime(5) {
		t.Errorf("Incorrect result when calculating 5 being prime.")
	}
}

func TestMaximumDivisor(t *testing.T) {
	if MaximumPrimeDivisor(9) != 3 {
		t.Errorf("Incorrect result when calculating maximum prime divisor of 9 not being 3.")
	}

	primeDivisorOfLargeNumber := MaximumPrimeDivisor(13195)
	expected := 29
	if primeDivisorOfLargeNumber != expected {
		t.Errorf("Incorrect result when calculating maximum prime divisor of 13195. Expected %d, got %d.", expected, primeDivisorOfLargeNumber)
	}
}
