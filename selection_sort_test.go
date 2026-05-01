package main

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"unsorted all", []int{3, 2, 1}, []int{1, 2, 3}},
		{"unsorted partial", []int{1, 3, 2}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SelectionSort(test.input)
			if !slices.Equal(result, test.want) {
				t.Errorf("want %v got %v", test.want, result)
			}
		})
	}
}
