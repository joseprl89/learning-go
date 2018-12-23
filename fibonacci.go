package main

func FibonacciSequenceUpTo(limit int) []int {
	current := 1
	next := 2
	var result []int

	for current < limit {
		result = append(result, current)

		var previous = current
		current = next
		next = previous + current
	}

	return result
}

func EvenFibonacciSequenceUpTo(limit int) []int {
	current := 1
	next := 2
	var result []int

	for current < limit {
		if IsEven(current) {
			result = append(result, current)
		}

		var previous = current
		current = next
		next = previous + current
	}

	return result
}
