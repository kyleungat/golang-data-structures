package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{2, 1, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{6, 2, 5, 4, 3},
			[]int{2, 3, 4, 5, 6},
		},
		{
			[]int{-3, -10, 0, 0, 5, 1},
			[]int{-10, -3, 0, 0, 1, 5},
		},
	}

	for index, tt := range tests {
		result := BubbleSort(tt.input)
		if len(result) != len(tt.expected) {
			t.Fatalf("test case %v: result %v does not equal to expected %v", index, result, tt.expected)
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.expected[i] {
				t.Fatalf("test case %v: elements[%v] result %v does not equal to expected %v", index, i, result[i], tt.expected[i])
			}
		}
	}

}
