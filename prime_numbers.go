package main

func NotDivisibleByAny(primes []int, number int) bool {
	for i := range primes {
		if number%primes[i] == 0 {
			return false
		}
	}
	return true
}

func PrimesUpTo(limit int) []int {
	var primes []int

	for i := 0; i < limit; i++ {
		if NotDivisibleByAny(primes, i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func HighestPrimeNumberFor(number int) int {
	return 3
}
