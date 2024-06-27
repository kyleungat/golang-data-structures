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
	for tid, tt := range tests {
		target.Add(tt.input)
		result := target.ToArray()
		if len(result) != len(tt.expected) {
			t.Fatalf("test case %v not equal length, actual: %v; expected: %v", tid, len(result), len(tt.expected))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.expected[i] {
				t.Fatalf("test case %v not equal at element %v, actual: %v; expected: %v", tid, i, result[i], tt.expected[i])
			}
		}
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		input    []int // input
		expected []int // expected result
	}{
		{nil, []int{}},
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1, 2, 1}},
		{[]int{1, 2}, []int{1, 2, 1, 1, 2}},
		{nil, []int{1, 2, 1, 1, 2}},
	}

	target := new(LinkedList[int])
	for tid, tt := range tests {
		target.AddAll(tt.input)
		result := target.ToArray()
		if len(result) != len(tt.expected) {
			t.Fatalf("test case %v not equal length, actual: %v; expected: %v", tid, len(result), len(tt.expected))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.expected[i] {
				t.Fatalf("test case %v not equal at element %v, actual: %v; expected: %v", tid, i, result[i], tt.expected[i])
			}
		}
	}
}
