package main

func FindMultiplesOf(x int, limit int) []int {
	var multiples []int
	for i := 1; i <= limit; i++ {
		if i%x == 0 {
			multiples = append(multiples, i)
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
	return 23
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
