package main

import "testing"

func TestSum(t *testing.T) {
    total := FindMultiplesOf(2, 3)
    if total[0] != 2 {
       t.Errorf("Incorrect multiple, got: %d, want: 2.", total)
    }
}
