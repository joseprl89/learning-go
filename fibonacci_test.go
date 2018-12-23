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
