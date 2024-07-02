package sorting

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](input []T) []T {
	selected := 0
	for selected < len(input) {
		min := input[selected]
		selectIndex := selected
		for i := selected + 1; i < len(input); i++ {
			// comparison
			if input[i] < min {
				min = input[i]
				selectIndex = i
			}
		}
		// swaping
		if selectIndex != selected {
			input[selectIndex] = input[selected]
			input[selected] = min
		}

		selected++
	}
	return input
}
