package main

import "testing"

func TestSum(t *testing.T) {
    expected := []int { 2 }
    total := FindMultiplesOf(2, 3)
    if total[0] != expected[0] {
       t.Errorf("Incorrect multiple, got: %v, want: %v.", total, expected)
    }
}
