package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersCollection ...[]int) []int {
	var sumCollection []int
	for _, numbers := range numbersCollection {
		sumCollection = append(sumCollection, Sum(numbers))
	}

	return sumCollection
}

func SumAllTails(numbersCollection ...[]int) []int {
	var sumCollection []int
	for _, numbers := range numbersCollection {
		if len(numbers) == 0 {
			sumCollection = append(sumCollection, Sum(numbers))
		} else {
			tail := numbers[1:]
			sumCollection = append(sumCollection, Sum(tail))

		}
	}

	return sumCollection
}
