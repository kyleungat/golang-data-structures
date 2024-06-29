package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests1 = []struct {
		input    int // input
		list     []int
		expected []int // expected result
	}{
		{1, []int{}, []int{1}},
		{1, []int{1}, []int{1, 1}},
		{1, []int{1, 2}, []int{1, 2, 1}},
	}

	t.Run("TestAdd without index", func(t *testing.T) {
		for tid, tt := range tests1 {
			target := new(LinkedList[int])
			for _, v := range tt.list {
				target.Add(v)
			}
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
	})

	var tests2 = []struct {
		input    int // input
		index    int
		list     []int
		expected []int // expected result
	}{
		{1, -1, []int{}, []int{}},
		{1, 0, []int{}, []int{1}},
		{1, 1, []int{}, []int{}},
		{1, -1, []int{1}, []int{1}},
		{1, 0, []int{1}, []int{1, 1}},
		{1, 1, []int{1}, []int{1}},
		{1, -1, []int{1, 2}, []int{1, 2}},
		{1, 0, []int{1, 2}, []int{1, 1, 2}},
		{1, 1, []int{1, 2}, []int{1, 2, 1}},
		{1, 2, []int{1, 2}, []int{1, 2}},
	}

	t.Run("TestAdd with one index", func(t *testing.T) {
		for tid, tt := range tests2 {
			target := new(LinkedList[int])
			for _, v := range tt.list {
				target.Add(v)
			}
			target.Add(tt.input, tt.index)
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
		input    int // input
		indexes  []int
		list     []int
		expected []int // expected result
	}{
		{1, []int{0, 1}, []int{}, []int{}},
		{1, []int{0, 1}, []int{1}, []int{1}},
		{1, []int{0, 1}, []int{1, 2}, []int{1, 2}},
	}
	t.Run("TestAdd with more than one index", func(t *testing.T) {
		for tid, tt := range tests3 {
			target := new(LinkedList[int])
			for _, v := range tt.list {
				target.Add(v)
			}
			target.Add(tt.input, tt.indexes...)
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

func BenchmarkAdd1(b *testing.B)      { benchmarkAdd(1, b) }
func BenchmarkAdd10(b *testing.B)     { benchmarkAdd(10, b) }
func BenchmarkAdd100(b *testing.B)    { benchmarkAdd(100, b) }
func BenchmarkAdd1000(b *testing.B)   { benchmarkAdd(1000, b) }
func BenchmarkAdd10000(b *testing.B)  { benchmarkAdd(10000, b) }
func BenchmarkAdd100000(b *testing.B) { benchmarkAdd(100000, b) }

func benchmarkAdd(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("no index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Add(1)
		}
	})
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Add(1, 0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Add(1, size/2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Add(1, size-1)
		}
	})
}

func BenchmarkAddAll1(b *testing.B)      { benchmarkAddAll(1, b) }
func BenchmarkAddAll10(b *testing.B)     { benchmarkAddAll(10, b) }
func BenchmarkAddAll100(b *testing.B)    { benchmarkAddAll(100, b) }
func BenchmarkAddAll1000(b *testing.B)   { benchmarkAddAll(1000, b) }
func BenchmarkAddAll10000(b *testing.B)  { benchmarkAddAll(10000, b) }
func BenchmarkAddAll100000(b *testing.B) { benchmarkAddAll(100000, b) }

func benchmarkAddAll(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	add := []int{}
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
		add = append(add, i)
	}
	b.ResetTimer()
	b.Run("no index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.AddAll(add)
		}
	})
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.AddAll(add, 0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.AddAll(add, size/2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.AddAll(add, size-1)
		}
	})
}

func BenchmarkAddFirst1(b *testing.B)      { benchmarkAddFirst(1, b) }
func BenchmarkAddFirst10(b *testing.B)     { benchmarkAddFirst(10, b) }
func BenchmarkAddFirst100(b *testing.B)    { benchmarkAddFirst(100, b) }
func BenchmarkAddFirst1000(b *testing.B)   { benchmarkAddFirst(1000, b) }
func BenchmarkAddFirst10000(b *testing.B)  { benchmarkAddFirst(10000, b) }
func BenchmarkAddFirst100000(b *testing.B) { benchmarkAddFirst(100000, b) }

func benchmarkAddFirst(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.AddFirst(1)
	}
}

func BenchmarkAddLast1(b *testing.B)      { benchmarkAddLast(1, b) }
func BenchmarkAddLast10(b *testing.B)     { benchmarkAddLast(10, b) }
func BenchmarkAddLast100(b *testing.B)    { benchmarkAddLast(100, b) }
func BenchmarkAddLast1000(b *testing.B)   { benchmarkAddLast(1000, b) }
func BenchmarkAddLast10000(b *testing.B)  { benchmarkAddLast(10000, b) }
func BenchmarkAddLast100000(b *testing.B) { benchmarkAddLast(100000, b) }

func benchmarkAddLast(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.AddLast(1)
	}
}

func BenchmarkClear1(b *testing.B)      { benchmarkClear(1, b) }
func BenchmarkClear10(b *testing.B)     { benchmarkClear(10, b) }
func BenchmarkClear100(b *testing.B)    { benchmarkClear(100, b) }
func BenchmarkClear1000(b *testing.B)   { benchmarkClear(1000, b) }
func BenchmarkClear10000(b *testing.B)  { benchmarkClear(10000, b) }
func BenchmarkClear100000(b *testing.B) { benchmarkClear(100000, b) }

func benchmarkClear(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.Clear()
	}
}

func BenchmarkClone1(b *testing.B)      { benchmarkClone(1, b) }
func BenchmarkClone10(b *testing.B)     { benchmarkClone(10, b) }
func BenchmarkClone100(b *testing.B)    { benchmarkClone(100, b) }
func BenchmarkClone1000(b *testing.B)   { benchmarkClone(1000, b) }
func BenchmarkClone10000(b *testing.B)  { benchmarkClone(10000, b) }
func BenchmarkClone100000(b *testing.B) { benchmarkClone(100000, b) }

func benchmarkClone(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.Clone()
	}
}

func BenchmarkContains1(b *testing.B)      { benchmarkContains(1, b) }
func BenchmarkContains10(b *testing.B)     { benchmarkContains(10, b) }
func BenchmarkContains100(b *testing.B)    { benchmarkContains(100, b) }
func BenchmarkContains1000(b *testing.B)   { benchmarkContains(1000, b) }
func BenchmarkContains10000(b *testing.B)  { benchmarkContains(10000, b) }
func BenchmarkContains100000(b *testing.B) { benchmarkContains(100000, b) }

func benchmarkContains(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Contains(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Contains(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Contains(size - 1)
		}
	})
}

func BenchmarkGet1(b *testing.B)      { benchmarkGet(1, b) }
func BenchmarkGet10(b *testing.B)     { benchmarkGet(10, b) }
func BenchmarkGet100(b *testing.B)    { benchmarkGet(100, b) }
func BenchmarkGet1000(b *testing.B)   { benchmarkGet(1000, b) }
func BenchmarkGet10000(b *testing.B)  { benchmarkGet(10000, b) }
func BenchmarkGet100000(b *testing.B) { benchmarkGet(100000, b) }

func benchmarkGet(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Get(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Get(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Get(size - 1)
		}
	})
}

func BenchmarkGetFirst1(b *testing.B)      { benchmarkGetFirst(1, b) }
func BenchmarkGetFirst10(b *testing.B)     { benchmarkGetFirst(10, b) }
func BenchmarkGetFirst100(b *testing.B)    { benchmarkGetFirst(100, b) }
func BenchmarkGetFirst1000(b *testing.B)   { benchmarkGetFirst(1000, b) }
func BenchmarkGetFirst10000(b *testing.B)  { benchmarkGetFirst(10000, b) }
func BenchmarkGetFirst100000(b *testing.B) { benchmarkGetFirst(100000, b) }

func benchmarkGetFirst(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.GetFirst()
	}
}

func BenchmarkGetLast1(b *testing.B)      { benchmarkGetLast(1, b) }
func BenchmarkGetLast10(b *testing.B)     { benchmarkGetLast(10, b) }
func BenchmarkGetLast100(b *testing.B)    { benchmarkGetLast(100, b) }
func BenchmarkGetLast1000(b *testing.B)   { benchmarkGetLast(1000, b) }
func BenchmarkGetLast10000(b *testing.B)  { benchmarkGetLast(10000, b) }
func BenchmarkGetLast100000(b *testing.B) { benchmarkGetLast(100000, b) }

func benchmarkGetLast(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.GetLast()
	}
}

func BenchmarkIndexOf1(b *testing.B)      { benchmarkIndexOf(1, b) }
func BenchmarkIndexOf10(b *testing.B)     { benchmarkIndexOf(10, b) }
func BenchmarkIndexOf100(b *testing.B)    { benchmarkIndexOf(100, b) }
func BenchmarkIndexOf1000(b *testing.B)   { benchmarkIndexOf(1000, b) }
func BenchmarkIndexOf10000(b *testing.B)  { benchmarkIndexOf(10000, b) }
func BenchmarkIndexOf100000(b *testing.B) { benchmarkIndexOf(100000, b) }

func benchmarkIndexOf(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.IndexOf(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.IndexOf(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.IndexOf(size - 1)
		}
	})
}

func BenchmarkLastIndexOf1(b *testing.B)      { benchmarkLastIndexOf(1, b) }
func BenchmarkLastIndexOf10(b *testing.B)     { benchmarkLastIndexOf(10, b) }
func BenchmarkLastIndexOf100(b *testing.B)    { benchmarkLastIndexOf(100, b) }
func BenchmarkLastIndexOf1000(b *testing.B)   { benchmarkLastIndexOf(1000, b) }
func BenchmarkLastIndexOf10000(b *testing.B)  { benchmarkLastIndexOf(10000, b) }
func BenchmarkLastIndexOf100000(b *testing.B) { benchmarkLastIndexOf(100000, b) }

func benchmarkLastIndexOf(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.LastIndexOf(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.LastIndexOf(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.LastIndexOf(size - 1)
		}
	})
}

func BenchmarkPeekFirst1(b *testing.B)      { benchmarkPeekFirst(1, b) }
func BenchmarkPeekFirst10(b *testing.B)     { benchmarkPeekFirst(10, b) }
func BenchmarkPeekFirst100(b *testing.B)    { benchmarkPeekFirst(100, b) }
func BenchmarkPeekFirst1000(b *testing.B)   { benchmarkPeekFirst(1000, b) }
func BenchmarkPeekFirst10000(b *testing.B)  { benchmarkPeekFirst(10000, b) }
func BenchmarkPeekFirst100000(b *testing.B) { benchmarkPeekFirst(100000, b) }

func benchmarkPeekFirst(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.PeekFirst()
	}
}

func BenchmarkPeekLast1(b *testing.B)      { benchmarkPeekLast(1, b) }
func BenchmarkPeekLast10(b *testing.B)     { benchmarkPeekLast(10, b) }
func BenchmarkPeekLast100(b *testing.B)    { benchmarkPeekLast(100, b) }
func BenchmarkPeekLast1000(b *testing.B)   { benchmarkPeekLast(1000, b) }
func BenchmarkPeekLast10000(b *testing.B)  { benchmarkPeekLast(10000, b) }
func BenchmarkPeekLast100000(b *testing.B) { benchmarkPeekLast(100000, b) }

func benchmarkPeekLast(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.PeekLast()
	}
}

func BenchmarkRemove1(b *testing.B)      { benchmarkRemove(1, b) }
func BenchmarkRemove10(b *testing.B)     { benchmarkRemove(10, b) }
func BenchmarkRemove100(b *testing.B)    { benchmarkRemove(100, b) }
func BenchmarkRemove1000(b *testing.B)   { benchmarkRemove(1000, b) }
func BenchmarkRemove10000(b *testing.B)  { benchmarkRemove(10000, b) }
func BenchmarkRemove100000(b *testing.B) { benchmarkRemove(100000, b) }

func benchmarkRemove(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Remove(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Remove(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Remove(size - 1)
		}
	})
}

func BenchmarkRemoveValue1(b *testing.B)      { benchmarkRemoveValue(1, b) }
func BenchmarkRemoveValue10(b *testing.B)     { benchmarkRemoveValue(10, b) }
func BenchmarkRemoveValue100(b *testing.B)    { benchmarkRemoveValue(100, b) }
func BenchmarkRemoveValue1000(b *testing.B)   { benchmarkRemoveValue(1000, b) }
func BenchmarkRemoveValue10000(b *testing.B)  { benchmarkRemoveValue(10000, b) }
func BenchmarkRemoveValue100000(b *testing.B) { benchmarkRemoveValue(100000, b) }

func benchmarkRemoveValue(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveValue(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveValue(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveValue(size - 1)
		}
	})
}

func BenchmarkRemoveFirst1(b *testing.B)      { benchmarkRemoveFirst(1, b) }
func BenchmarkRemoveFirst10(b *testing.B)     { benchmarkRemoveFirst(10, b) }
func BenchmarkRemoveFirst100(b *testing.B)    { benchmarkRemoveFirst(100, b) }
func BenchmarkRemoveFirst1000(b *testing.B)   { benchmarkRemoveFirst(1000, b) }
func BenchmarkRemoveFirst10000(b *testing.B)  { benchmarkRemoveFirst(10000, b) }
func BenchmarkRemoveFirst100000(b *testing.B) { benchmarkRemoveFirst(100000, b) }

func benchmarkRemoveFirst(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.RemoveFirst()
	}
}

func BenchmarkRemoveFirstOccurrence1(b *testing.B)      { benchmarkRemoveFirstOccurrence(1, b) }
func BenchmarkRemoveFirstOccurrence10(b *testing.B)     { benchmarkRemoveFirstOccurrence(10, b) }
func BenchmarkRemoveFirstOccurrence100(b *testing.B)    { benchmarkRemoveFirstOccurrence(100, b) }
func BenchmarkRemoveFirstOccurrence1000(b *testing.B)   { benchmarkRemoveFirstOccurrence(1000, b) }
func BenchmarkRemoveFirstOccurrence10000(b *testing.B)  { benchmarkRemoveFirstOccurrence(10000, b) }
func BenchmarkRemoveFirstOccurrence100000(b *testing.B) { benchmarkRemoveFirstOccurrence(100000, b) }

func benchmarkRemoveFirstOccurrence(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveFirstOccurrence(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveFirstOccurrence(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveFirstOccurrence(size - 1)
		}
	})
}

func BenchmarkRemoveLast1(b *testing.B)      { benchmarkRemoveLast(1, b) }
func BenchmarkRemoveLast10(b *testing.B)     { benchmarkRemoveLast(10, b) }
func BenchmarkRemoveLast100(b *testing.B)    { benchmarkRemoveLast(100, b) }
func BenchmarkRemoveLast1000(b *testing.B)   { benchmarkRemoveLast(1000, b) }
func BenchmarkRemoveLast10000(b *testing.B)  { benchmarkRemoveLast(10000, b) }
func BenchmarkRemoveLast100000(b *testing.B) { benchmarkRemoveLast(100000, b) }

func benchmarkRemoveLast(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.RemoveLast()
	}
}

func BenchmarkRemoveLastOccurrence1(b *testing.B)      { benchmarkRemoveLastOccurrence(1, b) }
func BenchmarkRemoveLastOccurrence10(b *testing.B)     { benchmarkRemoveLastOccurrence(10, b) }
func BenchmarkRemoveLastOccurrence100(b *testing.B)    { benchmarkRemoveLastOccurrence(100, b) }
func BenchmarkRemoveLastOccurrence1000(b *testing.B)   { benchmarkRemoveLastOccurrence(1000, b) }
func BenchmarkRemoveLastOccurrence10000(b *testing.B)  { benchmarkRemoveLastOccurrence(10000, b) }
func BenchmarkRemoveLastOccurrence100000(b *testing.B) { benchmarkRemoveLastOccurrence(100000, b) }

func benchmarkRemoveLastOccurrence(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveLastOccurrence(0)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveLastOccurrence(size / 2)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.RemoveLastOccurrence(size - 1)
		}
	})
}

func BenchmarkSet10(b *testing.B)     { benchmarkSet(10, b) }
func BenchmarkSet1000(b *testing.B)   { benchmarkSet(1000, b) }
func BenchmarkSet10000(b *testing.B)  { benchmarkSet(10000, b) }
func BenchmarkSet100000(b *testing.B) { benchmarkSet(100000, b) }

func benchmarkSet(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	b.Run("first", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Set(0, 1)
		}
	})
	b.Run("middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Set(size/2, 1)
		}
	})
	b.Run("last", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			linkedlist.Set(size-1, 1)
		}
	})
}

func BenchmarkSize1(b *testing.B)      { benchmarkSize(1, b) }
func BenchmarkSize10(b *testing.B)     { benchmarkSize(10, b) }
func BenchmarkSize100(b *testing.B)    { benchmarkSize(100, b) }
func BenchmarkSize1000(b *testing.B)   { benchmarkSize(1000, b) }
func BenchmarkSize10000(b *testing.B)  { benchmarkSize(10000, b) }
func BenchmarkSize100000(b *testing.B) { benchmarkSize(100000, b) }

func benchmarkSize(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.Size()
	}
}

func BenchmarkToArray1(b *testing.B)      { benchmarkToArray(1, b) }
func BenchmarkToArray10(b *testing.B)     { benchmarkToArray(10, b) }
func BenchmarkToArray100(b *testing.B)    { benchmarkToArray(100, b) }
func BenchmarkToArray1000(b *testing.B)   { benchmarkToArray(1000, b) }
func BenchmarkToArray10000(b *testing.B)  { benchmarkToArray(10000, b) }
func BenchmarkToArray100000(b *testing.B) { benchmarkToArray(100000, b) }

func benchmarkToArray(size int, b *testing.B) {
	linkedlist := new(LinkedList[int])
	for i := 0; i < size; i++ {
		linkedlist.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.ToArray()
	}
}
