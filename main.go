package main

func FindMultiplesOf(x int, limit int) int {
  var multiples []int
  for i := 1; i < limit; i++ {
    if i % x == 0 {
      multiples = append(multiples, i)
      return multiples[0]
    }
  }
  return 2
}
