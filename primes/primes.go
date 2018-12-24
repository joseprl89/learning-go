package primes

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
	return false
}
