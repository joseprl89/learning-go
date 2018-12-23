package main

import "testing"

func isEqual(lhs []int, rhs []int) bool {
  if (lhs == nil) != (rhs == nil) {
    return false
  }

  if len(lhs) != len(rhs) {
    return false
  }
  return true
}

func TestMultiples(t *testing.T) {
    expected := []int { 2, 4 }
    multiples := FindMultiplesOf(2, 4)
    if !isEqual(multiples, expected) {
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
