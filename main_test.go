package main

import "testing"

func TestMultiples(t *testing.T) {
    expected := []int { 2, 4 }
    multiples := FindMultiplesOf(2, 4)
    if multiples[0] != expected[0] {
       t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
    }
}

func TestArraySum(t *testing.T) {
    expected := 6
    sum := sumOf([]int { 2, 4 })
    if sum != expected {
       t.Errorf("Incorrect sum of array, got: %v, want: %v.", sum, expected)
    }
}
