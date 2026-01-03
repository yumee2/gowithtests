package arrays

func Sum(numbers []int) int {
	var sum int

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers))

	for index, val := range numbers {
		sumOfArray := Sum(val)
		sums[index] = sumOfArray
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
