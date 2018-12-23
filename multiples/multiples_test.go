package multiples

import "testing"
import "arrays"

func TestMultiples(t *testing.T) {
	expected := []int{2, 4}
	multiples := FindMultiplesOf(5, 2)
	if !arrays.IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestMultiplesOnEdge(t *testing.T) {
	expected := []int{2}
	multiples := FindMultiplesOf(4, 2)
	if !arrays.IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestMultiplesOfMultipleNumbers(t *testing.T) {
	expected := []int{2, 3}
	multiples := FindMultiplesOf(4, 2, 3)
	if !arrays.IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestMultiplesOfMultipleNumbersReturnsUnique(t *testing.T) {
	expected := []int{3, 5, 6, 9}
	multiples := FindMultiplesOf(10, 3, 5)
	if !arrays.IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestSumOfMultiples(t *testing.T) {
	expected := 23
	sum := SumOfMultiples(10, 3, 5)
	if sum != expected {
		t.Errorf("Incorrect sum of multiples, got: %d, want: %d.", sum, expected)
	}
}
