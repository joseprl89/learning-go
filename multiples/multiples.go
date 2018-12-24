package multiples

import "euler/arrays"

func FindMultiplesOf(limit int, x ...int) []int {
	var multiples []int

	for i := 1; i < limit; i++ {
		for j := range x {
			var divisor = x[j]
			if i%divisor == 0 {
				multiples = append(multiples, i)
				break
			}
		}
	}

	return multiples
}

func SumOfMultiples(limit int, multiples ...int) int {
	return arrays.SumOf(FindMultiplesOf(limit, multiples...))
}
