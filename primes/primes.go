package primes

import "math"

func DivisibleByAny(primes []int, number int) bool {
	for i := range primes {
		prime := primes[i]
		if number % prime == 0 {
			return true
		}
	}
	return false
}

func NotDivisibleByAny(primes []int, number int) bool {
	for i := range primes {
		prime := primes[i]
		if number % prime == 0 {
			return false
		}
	}
	return true
}

func IsPrime(number int) bool {
	maximumDivisor := int(math.Sqrt(float64(number))) + 1

	for i := maximumDivisor ; i >= 2 ; i-- {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func MaximumPrimeDivisor(number int) int {
	return 3
}
