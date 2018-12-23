package primes

func NotDivisibleByAny(primes []int, number int) bool {
	for i := range primes {
		prime := primes[i]

		if number%prime == 0 {
			return false
		}
	}

	return true
}
