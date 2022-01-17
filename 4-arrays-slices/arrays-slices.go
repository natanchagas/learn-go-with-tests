package main

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumAll(arr ...[]int) []int {
	var result []int

	for _, numbers := range arr {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(arr ...[]int) []int {
	var result []int

	for _, numbers := range arr {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(numbers[1:]))
		}
	}

	return result
}
