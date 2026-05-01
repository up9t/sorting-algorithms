package main

import (
	"fmt"
	"slices"
)

func main() {
	n := []int{2, 45, 1}

	fmt.Println(SelectionSort(slices.Clone(n)))
	fmt.Println(BubbleSort(slices.Clone(n)))
}
