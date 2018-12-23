package main

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

func SumOf(array []int) int {
	var result = 0
	for i := 0; i < len(array); i++ {
		result = result + array[i]
	}
	return result
}

func SumOfMultiples(limit int, multiples ...int) int {
	return SumOf(FindMultiplesOf(limit, multiples...))
}

func contains(array []int, value int) bool {
	for i := range array {
		if array[i] == value {
			return true
		}
	}

	return false
}

func MergeArrays(lhs []int, rhs []int) []int {
	for i := range rhs {
		if !contains(lhs, rhs[i]) {
			lhs = append(lhs, rhs[i])
		}
	}
	return lhs
}
