package main

import "testing"

func TestSum(t *testing.T) {
    expected := []int { 2 }
    multiples := FindMultiplesOf(2, 3)
    if multiples[0] != expected[0] {
       t.Errorf("Incorrect multiples, got: %v, want: %v.", multiples, expected)
    }
}
