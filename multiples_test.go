package main

import "testing"

func TestMultiples(t *testing.T) {
	expected := []int{2, 4}
	multiples := FindMultiplesOf(5, 2)
	if !IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestMultiplesOnEdge(t *testing.T) {
	expected := []int{2}
	multiples := FindMultiplesOf(4, 2)
	if !IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestMultiplesOfMultipleNumbers(t *testing.T) {
	expected := []int{2, 3}
	multiples := FindMultiplesOf(4, 2, 3)
	if !IsEqual(multiples, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
	}
}

func TestMultiplesOfMultipleNumbersReturnsUnique(t *testing.T) {
	expected := []int{3, 5, 6, 9}
	multiples := FindMultiplesOf(10, 3, 5)
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
	expected := 23
	sum := SumOfMultiples(10, 3, 5)
	if sum != expected {
		t.Errorf("Incorrect sum of multiples, got: %d, want: %d.", sum, expected)
	}
}

func TestMergeArrays(t *testing.T) {
	expected := []int{1, 2, 3}
	result := MergeArrays([]int{1, 2}, []int{2, 3})

	if !IsEqual(result, expected) {
		t.Errorf("Incorrect merged arrays, got: %v, want: %v.", result, expected)
	}
}