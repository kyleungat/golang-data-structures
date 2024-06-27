package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		input    int   // input
		expected []int // expected result
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 2, 3}},
	}

	target := new(LinkedList[int])
	for _, tt := range tests {
		target.Add(tt.input)
		result := target.ToArray()
		if len(result) != len(tt.expected) {
			t.Fatalf("not equal length, actual: %v; expected: %v", len(result), len(tt.expected))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.expected[i] {
				t.Fatalf("not equal at element %v, actual: %v; expected: %v", i, result[i], tt.expected[i])
			}
		}
	}
}
