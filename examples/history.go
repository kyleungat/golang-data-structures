package examples

import (
	"golang-data-structures/linked_list"
)

// An example usage of linked list
// Undo, Redo functions

type System struct {
	history     linked_list.LinkedList[int]
	undoHistory linked_list.LinkedList[int]
}

func (system *System) Do(input int) {
	system.history.Add(input)
}

func (system *System) Undo() {
	value := system.history.RemoveLast()
	if value != nil {
		system.undoHistory.Add(*value)
	}
}

func (system *System) Redo() {
	value := system.undoHistory.RemoveLast()
	if value != nil {
		system.Do(*value)
	}
}

func (system *System) ShowHistory() []int {
	return system.history.ToArray()
}
