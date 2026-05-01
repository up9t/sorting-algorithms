package main

func SelectionSort(numbers []int) []int {
	n := len(numbers)

	for i := 0; i < n; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[i] {
				minIndex = j
			}
		}

		if minIndex != i {
			numbers[minIndex], numbers[i] = numbers[i], numbers[minIndex]
		}
	}

	return numbers
}
