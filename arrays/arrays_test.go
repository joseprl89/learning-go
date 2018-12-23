package arrays

import "testing"

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

func TestArraySum(t *testing.T) {
	expected := 6
	sum := SumOf([]int{2, 4})
	if sum != expected {
		t.Errorf("Incorrect sum of array, got: %d, want: %d.", sum, expected)
	}
}

func TestMergeArrays(t *testing.T) {
	expected := []int{1, 2, 3}
	result := MergeArrays([]int{1, 2}, []int{2, 3})

	if !IsEqual(result, expected) {
		t.Errorf("Incorrect merged arrays, got: %v, want: %v.", result, expected)
	}
}
