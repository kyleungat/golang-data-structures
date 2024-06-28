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
	var tests1 = []struct {
		input    []int // input
		list     []int
		expected []int // expected result
	}{
		{nil, []int{}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{nil, []int{1}, []int{1}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{1}, []int{1, 1}},
		{[]int{1, 2}, []int{1}, []int{1, 1, 2}},
		{nil, []int{1, 2}, []int{1, 2}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1, 2}, []int{1, 2, 1}},
		{[]int{1, 2}, []int{1, 2}, []int{1, 2, 1, 2}},
	}

	t.Run("TestAddAll without index", func(t *testing.T) {
		for tid, tt := range tests1 {
			target := new(LinkedList[int])
			for _, v := range tt.list {
				target.Add(v)
			}
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
	})

	var tests2 = []struct {
		input    []int // input
		index    int
		list     []int
		expected []int // expected result
	}{
		{nil, -1, []int{}, []int{}},
		{nil, 0, []int{}, []int{}},
		{nil, 1, []int{}, []int{}},
		{[]int{}, -1, []int{}, []int{}},
		{[]int{}, 0, []int{}, []int{}},
		{[]int{}, 1, []int{}, []int{}},
		{[]int{1}, 0, []int{}, []int{1}},
		{[]int{1}, 1, []int{}, []int{}},
		{[]int{1}, -1, []int{}, []int{}},
		{[]int{1, 2}, -1, []int{}, []int{}},
		{[]int{1, 2}, 0, []int{}, []int{1, 2}},
		{[]int{1, 2}, 1, []int{}, []int{}},
		{nil, -1, []int{1}, []int{1}},
		{nil, 0, []int{1}, []int{1}},
		{nil, 1, []int{1}, []int{1}},
		{[]int{}, -1, []int{1}, []int{1}},
		{[]int{}, 0, []int{1}, []int{1}},
		{[]int{}, 1, []int{1}, []int{1}},
		{[]int{1}, -1, []int{1}, []int{1}},
		{[]int{1}, 0, []int{1}, []int{1, 1}},
		{[]int{1}, 1, []int{1}, []int{1}},
		{[]int{1, 2}, -1, []int{1}, []int{1}},
		{[]int{1, 2}, 0, []int{1}, []int{1, 1, 2}},
		{[]int{1, 2}, 1, []int{1}, []int{1}},
		{nil, -1, []int{1, 2}, []int{1, 2}},
		{nil, 0, []int{1, 2}, []int{1, 2}},
		{nil, 1, []int{1, 2}, []int{1, 2}},
		{[]int{}, -1, []int{1, 2}, []int{1, 2}},
		{[]int{}, 0, []int{1, 2}, []int{1, 2}},
		{[]int{}, 1, []int{1, 2}, []int{1, 2}},
		{[]int{1}, -1, []int{1, 2}, []int{1, 2}},
		{[]int{1}, 0, []int{1, 2}, []int{1, 1, 2}},
		{[]int{1}, 1, []int{1, 2}, []int{1, 2, 1}},
		{[]int{1}, 2, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, -1, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, 0, []int{1, 2}, []int{1, 1, 2, 2}},
		{[]int{1, 2}, 1, []int{1, 2}, []int{1, 2, 1, 2}},
		{[]int{1, 2}, 2, []int{1, 2}, []int{1, 2}},
	}

	t.Run("TestAddAll with one index", func(t *testing.T) {
		for tid, tt := range tests2 {
			target := new(LinkedList[int])
			for _, v := range tt.list {
				target.Add(v)
			}
			target.AddAll(tt.input, tt.index)
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
	})

	var tests3 = []struct {
		input    []int // input
		indexes  []int
		list     []int
		expected []int // expected result
	}{
		{nil, []int{0, 1}, []int{}, []int{}},
		{[]int{1}, []int{0, 1}, []int{}, []int{}},
		{[]int{1}, []int{0, 2}, []int{}, []int{}},
		{[]int{1}, []int{0, 1}, []int{1}, []int{1}},
		{[]int{1}, []int{0, 1, 2}, []int{1}, []int{1}},
		{[]int{1}, []int{2, 1, 2}, []int{1}, []int{1}},
		{[]int{1, 2}, []int{0, 1}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{0, 1, 2}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{2, 1, 2}, []int{1, 2}, []int{1, 2}},
	}
	t.Run("TestAddAll with more than one index", func(t *testing.T) {
		for tid, tt := range tests3 {
			target := new(LinkedList[int])
			for _, v := range tt.list {
				target.Add(v)
			}
			target.AddAll(tt.input, tt.indexes...)
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
	})
}

func TestAddFirst(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected []int // expected result
	}{
		{-1, []int{}, []int{-1}},
		{0, []int{}, []int{0}},
		{-1, []int{1}, []int{-1, 1}},
		{0, []int{1}, []int{0, 1}},
		{0, []int{1, 2}, []int{0, 1, 2}},
		{2, []int{1, 2}, []int{2, 1, 2}},
		{2, []int{1, 2, 3}, []int{2, 1, 2, 3}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		target.AddFirst(tt.input)
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

func TestAddLast(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected []int // expected result
	}{
		{-1, []int{}, []int{-1}},
		{0, []int{1}, []int{1, 0}},
		{-1, []int{1, 1}, []int{1, 1, -1}},
		{-1, []int{1, 2}, []int{1, 2, -1}},
		{-1, []int{1, 2, 3}, []int{1, 2, 3, -1}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		target.AddLast(tt.input)
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

func TestClear(t *testing.T) {
	var tests = []struct {
		input []int // input
	}{
		{[]int{}},
		{nil},
		{[]int{1, 2, 3}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.input {
			target.Add(v)
		}
		target.Clear()
		result := target.ToArray()
		if len(result) != 0 {
			t.Fatalf("test case %v not cleared", tid)
		}
	}
}

func TestClone(t *testing.T) {
	var tests = []struct {
		input    []int // input
		expected []int // expected result
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.input {
			target.Add(v)
		}
		clone := target.Clone()
		result := clone.ToArray()
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

func TestContains(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected bool // expected result
	}{
		{1, []int{1}, true},
		{1, []int{1, 2, 3}, true},
		{1, []int{2, 1, 3}, true},
		{1, []int{3, 2, 1}, true},
		{1, []int{1, 1, 1}, true},
		{1, []int{}, false},
		{1, []int{2}, false},
		{1, []int{2, 3, 4}, false},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.Contains(tt.input)
		if result != tt.expected {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, result, tt.expected)
		}
	}
}

func TestGet(t *testing.T) {
	var tests = []struct {
		input int // input
		list  []int
		exist bool
		value int // expected result
	}{

		{-1, []int{1}, false, 0},
		{0, []int{1}, true, 1},
		{1, []int{1}, false, 0},
		{-1, []int{1, 2, 3}, false, 0},
		{0, []int{1, 2, 3}, true, 1},
		{1, []int{1, 2, 3}, true, 2},
		{2, []int{1, 2, 3}, true, 3},
		{3, []int{1, 2, 3}, false, 0},
		{-1, []int{}, false, 0},
		{0, []int{}, false, 0},
		{1, []int{}, false, 0},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.Get(tt.input)
		if result == nil && tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && !tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && *result != tt.value {
			t.Fatalf("test case %v not equal, actual", tid)
		}
	}
}

func TestGetFirst(t *testing.T) {
	var tests = []struct {
		list  []int
		exist bool
		value int // expected result
	}{

		{[]int{}, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2}, true, 1},
		{[]int{1, 2, 3}, true, 1},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.GetFirst()
		if result == nil && tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && !tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && *result != tt.value {
			t.Fatalf("test case %v not equal, actual", tid)
		}
	}
}

func TestGetLast(t *testing.T) {
	var tests = []struct {
		list  []int
		exist bool
		value int // expected result
	}{

		{[]int{}, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2}, true, 2},
		{[]int{1, 2, 3}, true, 3},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.GetLast()
		if result == nil && tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && !tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && *result != tt.value {
			t.Fatalf("test case %v not equal, actual", tid)
		}
	}
}

func TestIndexOf(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected int // expected result
	}{
		{1, []int{}, -1},
		{1, []int{1}, 0},
		{1, []int{1, 2, 3}, 0},
		{1, []int{2, 1, 3}, 1},
		{1, []int{3, 2, 1}, 2},
		{1, []int{1, 1, 1}, 0},
		{1, []int{2}, -1},
		{1, []int{2, 3, 4}, -1},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.IndexOf(tt.input)
		if result != tt.expected {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, result, tt.expected)
		}
	}
}

func TestLastIndexOf(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected int // expected result
	}{
		{1, []int{}, -1},
		{1, []int{1}, 0},
		{1, []int{1, 2, 3}, 0},
		{1, []int{2, 1, 3}, 1},
		{1, []int{3, 2, 1}, 2},
		{1, []int{1, 1, 1}, 2},
		{1, []int{1, 1, 1, 2}, 2},
		{1, []int{1, 1, 2, 1}, 3},
		{1, []int{2}, -1},
		{1, []int{2, 3, 4}, -1},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.LastIndexOf(tt.input)
		if result != tt.expected {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, result, tt.expected)
		}
	}
}

func TestPeekFirst(t *testing.T) {
	var tests = []struct {
		list  []int
		exist bool
		value int // expected result
	}{

		{[]int{}, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2}, true, 1},
		{[]int{1, 2, 3}, true, 1},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.PeekFirst()
		if result == nil && tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && !tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && *result != tt.value {
			t.Fatalf("test case %v not equal, actual", tid)
		}
	}
}

func TestPeekLast(t *testing.T) {
	var tests = []struct {
		list  []int
		exist bool
		value int // expected result
	}{

		{[]int{}, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2}, true, 2},
		{[]int{1, 2, 3}, true, 3},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.PeekLast()
		if result == nil && tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && !tt.exist {
			t.Fatalf("test case %v not equal, actual", tid)
		}
		if result != nil && *result != tt.value {
			t.Fatalf("test case %v not equal, actual", tid)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected []int // expected result
	}{
		{-1, []int{}, []int{}},
		{0, []int{}, []int{}},
		{1, []int{}, []int{}},
		{-1, []int{1}, []int{1}},
		{0, []int{1}, []int{}},
		{1, []int{1}, []int{1}},
		{-1, []int{1, 2, 3}, []int{1, 2, 3}},
		{0, []int{1, 2, 3}, []int{2, 3}},
		{1, []int{1, 2, 3}, []int{1, 3}},
		{3, []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		target.Remove(tt.input)
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

func TestRemoveValue(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected []int // expected result
	}{
		{-1, []int{}, []int{}},
		{0, []int{}, []int{}},
		{1, []int{}, []int{}},
		{-1, []int{1}, []int{1}},
		{0, []int{1}, []int{1}},
		{1, []int{1}, []int{}},
		{-1, []int{1, 2, 3}, []int{1, 2, 3}},
		{1, []int{1, 2, 3}, []int{2, 3}},
		{2, []int{1, 2, 3}, []int{1, 3}},
		{3, []int{1, 2, 3}, []int{1, 2}},
		{2, []int{1, 2, 3, 2}, []int{1, 3, 2}},
		{2, []int{1, 2, 2, 2, 3, 2}, []int{1, 2, 2, 3, 2}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		target.RemoveValue(tt.input)
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

func TestRemoveFirst(t *testing.T) {
	var tests = []struct {
		list     []int
		expected []int // expected result
		exists   bool
		value    int
	}{
		{[]int{}, []int{}, false, 0},
		{[]int{1}, []int{}, true, 1},
		{[]int{1, 1}, []int{1}, true, 1},
		{[]int{1, 2}, []int{2}, true, 1},
		{[]int{1, 2, 3}, []int{2, 3}, true, 1},
		{[]int{2, 2, 2}, []int{2, 2}, true, 2},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		value := target.RemoveFirst()
		result := target.ToArray()
		if len(result) != len(tt.expected) {
			t.Fatalf("test case %v not equal length, actual: %v; expected: %v", tid, len(result), len(tt.expected))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.expected[i] {
				t.Fatalf("test case %v not equal at element %v, actual: %v; expected: %v", tid, i, result[i], tt.expected[i])
			}
		}
		if value == nil && tt.exists {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, value, tt.exists)
		}
		if value != nil && !tt.exists {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, value, tt.exists)
			if value != nil && (*value != tt.value) {
				t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, *value, tt.value)
			}
		}
	}
}

func TestRemoveFirstOccurrence(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected []int // expected result
	}{
		{1, []int{}, []int{}},
		{-1, []int{1}, []int{1}},
		{1, []int{1}, []int{}},
		{-1, []int{1, 1}, []int{1, 1}},
		{1, []int{1, 1}, []int{1}},
		{-1, []int{1, 2, 3}, []int{1, 2, 3}},
		{1, []int{1, 2, 3}, []int{2, 3}},
		{2, []int{1, 2, 3}, []int{1, 3}},
		{3, []int{1, 2, 3}, []int{1, 2}},
		{-1, []int{1, 2, 1}, []int{1, 2, 1}},
		{1, []int{1, 2, 1}, []int{2, 1}},
		{-1, []int{1, 2, 1, 2, 1}, []int{1, 2, 1, 2, 1}},
		{1, []int{1, 2, 1, 2, 1}, []int{2, 1, 2, 1}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		target.RemoveFirstOccurrence(tt.input)
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

func TestRemoveLast(t *testing.T) {
	var tests = []struct {
		list     []int
		expected []int // expected result
		exists   bool
		value    int
	}{
		{[]int{}, []int{}, false, 0},
		{[]int{1}, []int{}, true, 1},
		{[]int{1, 1}, []int{1}, true, 1},
		{[]int{1, 2}, []int{1}, true, 2},
		{[]int{1, 2, 3}, []int{1, 2}, true, 3},
		{[]int{2, 2, 2}, []int{2, 2}, true, 2},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		value := target.RemoveLast()
		result := target.ToArray()
		if len(result) != len(tt.expected) {
			t.Fatalf("test case %v not equal length, actual: %v; expected: %v", tid, len(result), len(tt.expected))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.expected[i] {
				t.Fatalf("test case %v not equal at element %v, actual: %v; expected: %v", tid, i, result[i], tt.expected[i])
			}
		}
		if value == nil && tt.exists {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, value, tt.exists)
		}
		if value != nil && !tt.exists {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, value, tt.exists)
			if value != nil && (*value != tt.value) {
				t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, *value, tt.value)
			}
		}
	}
}

func TestRemoveLastOccurrence(t *testing.T) {
	var tests = []struct {
		input    int // input
		list     []int
		expected []int // expected result
	}{
		{1, []int{}, []int{}},
		{-1, []int{1}, []int{1}},
		{1, []int{1}, []int{}},
		{-1, []int{1, 1}, []int{1, 1}},
		{1, []int{1, 1}, []int{1}},
		{-1, []int{1, 2, 3}, []int{1, 2, 3}},
		{1, []int{1, 2, 3}, []int{2, 3}},
		{2, []int{1, 2, 3}, []int{1, 3}},
		{3, []int{1, 2, 3}, []int{1, 2}},
		{-1, []int{1, 2, 1}, []int{1, 2, 1}},
		{1, []int{1, 2, 1}, []int{1, 2}},
		{-1, []int{1, 2, 1, 2, 1}, []int{1, 2, 1, 2, 1}},
		{1, []int{1, 2, 1, 2, 1}, []int{1, 2, 1, 2}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		target.RemoveLastOccurrence(tt.input)
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

func TestSet(t *testing.T) {
	var tests = []struct {
		index    int // input
		value    int
		list     []int
		expected []int // expected result
	}{
		{-1, 1, []int{}, []int{}},
		{0, 1, []int{}, []int{}},
		{1, 1, []int{}, []int{}},
		{-1, 1, []int{1}, []int{1}},
		{0, 1, []int{1}, []int{1}},
		{0, 2, []int{1}, []int{2}},
		{1, 1, []int{1}, []int{1}},
		{-1, 1, []int{1, 2, 3}, []int{1, 2, 3}},
		{0, 1, []int{1, 2, 3}, []int{1, 2, 3}},
		{0, 2, []int{1, 2, 3}, []int{2, 2, 3}},
		{1, 1, []int{1, 2, 3}, []int{1, 1, 3}},
		{1, 2, []int{1, 2, 3}, []int{1, 2, 3}},
		{2, 1, []int{1, 2, 3}, []int{1, 2, 1}},
		{2, 3, []int{1, 2, 3}, []int{1, 2, 3}},
		{3, 3, []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		target.Set(tt.index, tt.value)
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

func TestSize(t *testing.T) {
	var tests = []struct {
		list     []int
		expected int // expected result
	}{

		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 1}, 2},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 1}, 3},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
		result := target.Size()
		if result != tt.expected {
			t.Fatalf("test case %v not equal, actual: %v; expected: %v", tid, result, tt.expected)
		}
	}
}

func TestToArray(t *testing.T) {
	var tests = []struct {
		list     []int
		expected []int // expected result
	}{

		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
	}

	for tid, tt := range tests {
		target := new(LinkedList[int])
		for _, v := range tt.list {
			target.Add(v)
		}
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
