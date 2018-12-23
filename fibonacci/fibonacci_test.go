package fibonacci

import "testing"
import "arrays"

func TestFibonacci(t *testing.T) {
	expected := []int{1, 2, 3, 5, 8, 13, 21, 34}
	fibonacci := FibonacciSequenceUpTo(40)
	if !arrays.IsEqual(fibonacci, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", fibonacci, expected)
	}
}

func TestEvenFibonacciSequence(t *testing.T) {
	expected := []int{2, 8, 34}
	fibonacci := EvenFibonacciSequenceUpTo(40)
	if !arrays.IsEqual(fibonacci, expected) {
		t.Errorf("Incorrect multiples, got: %v, want: %v.", fibonacci, expected)
	}
}

func TestSumEvenFibonacciSequence(t *testing.T) {
	expected := 44
	result := arrays.SumOf(EvenFibonacciSequenceUpTo(40))
	if result != expected {
		t.Errorf("Incorrect sum of fibonacci, got: %d, want: %d.", result, expected)
	}
}
