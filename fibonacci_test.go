package main

import "testing"

func TestFibonacci(t *testing.T) {
	expected := []int{1, 2, 3, 5, 8, 13, 21, 34}
	fibonacci := FibonacciSequenceUpTo(40)
	if !IsEqual(fibonacci, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", fibonacci, expected)
	}
}

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("IsEven says 2 is not even.")
	}
}

func TestFilter(t *testing.T) {
	expected := []int{2, 4, 6}
	filtered := Filter(IsEven, []int{1, 2, 3, 4, 5, 6})
	if !IsEqual(filtered, expected) {
		t.Errorf("Incorrect filtered, got: %v, want: %v.", filtered, expected)
	}
}
