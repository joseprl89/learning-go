package main

func FindMultiplesOf(x int, limit int) []int {
  var multiples []int
  for i := 1; i <= limit; i++ {
    if i % x == 0 {
      multiples = append(multiples, i)
    }
  }
  return multiples
}

func sumOf(array []int) int {
  var result = 0
  for i:= 0; i < len(array); i++ {
    result = result + array[i]
  }
  return result
}
