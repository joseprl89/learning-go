package main

func NotDivisibleByAny(primes []int, number int) bool {
	return false
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
