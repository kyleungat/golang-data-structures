package sorting

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](input []T) []T {
	sortedIndex := len(input) - 1
	for sortedIndex > 0 {
		for i := 0; i < sortedIndex; i++ {
			if input[i] > input[i+1] { // comparison
				// swaping
				tmp := input[i+1]
				input[i+1] = input[i]
				input[i] = tmp
			}
		}
		sortedIndex--
	}
	return input
}
