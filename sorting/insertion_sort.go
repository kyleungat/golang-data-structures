package sorting

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](input []T) []T {
	for i := 0; i+1 < len(input); i++ {
		for j := i + 1; j-1 >= 0; j-- {
			if input[j] < input[j-1] { // comparison
				// swaping
				tmp := input[j-1]
				input[j-1] = input[j]
				input[j] = tmp
			} else {
				break
			}
		}
	}
	return input
}
