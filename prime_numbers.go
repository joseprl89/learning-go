package main

import "math"

func NotDivisibleByAny(primes []int, number int) bool {
	limit := int(math.Sqrt(float64(number))) + 1
	for i := range primes {
		prime := primes[i]

		if prime > limit {
			return true
		}

		if number%prime == 0 {
			return false
		}
	}
	return true
}

func PrimesUpTo(limit int) []int {
	var primes []int

	for i := 1; i < limit; i++ {
		if NotDivisibleByAny(primes, i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func HighestPrimeNumberFor(number int) int {
	return 3
}
