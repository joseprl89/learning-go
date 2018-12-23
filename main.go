package main

func FindMultiplesOf(x int, limit int) int {
  for i := 1; i < limit; i++ {
    if i % x == 0 {
      return i
    }
  }
  return 2
}
