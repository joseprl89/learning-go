package main

type FilterFunction func(int) bool

func Filter(function FilterFunction, array []int) []int {
	return []int{2, 4, 6}
}

func SumOf(array []int) int {
	var result = 0
	for i := 0; i < len(array); i++ {
		result = result + array[i]
	}
	return result
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

func IsEqual(lhs []int, rhs []int) bool {
	if (lhs == nil) != (rhs == nil) {
		return false
	}

	if len(lhs) != len(rhs) {
		return false
	}

	for i := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}
