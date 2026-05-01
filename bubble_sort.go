package main

func BubbleSort(numbers []int) []int {
	n := len(numbers)

	for i := 0; i < n; i++ {
		swap := false

		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swap = true
			}
		}

		if !swap {
			break
		}
	}

	return numbers
}
