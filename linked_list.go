package main

type LinkedList[T comparable] struct {
	firstNode *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

// Add(value T)
// Appends the specified element to the end of this list.
// Inserts the specified element at the specified position in this list if the index is smaller than the length of the list.
func (list *LinkedList[T]) Add(value T, index ...int) {
	insert := -1
	if len(index) == 1 {
		insert = index[0]
	}

	for i, current := 0, list.firstNode; current != nil; i, current = i+1, current.next {
		if current.next == nil {
			current.next = &node[T]{
				value: value,
			}
		} else if i == insert {
			current.next = &node[T]{
				value: value,
				next:  current.next,
			}
			break
		}
	}

}

// boolean	addAll(Collection<? extends E> c)
// Appends all of the elements in the specified collection to the end of this list, in the order that they are returned by the specified collection's iterator.

// boolean	addAll(int index, Collection<? extends E> c)
// Inserts all of the elements in the specified collection into this list, starting at the specified position.

// void	addFirst(E e)
// Inserts the specified element at the beginning of this list.
func (list *LinkedList[T]) AddFirst(value T) {
	list.firstNode = &node[T]{
		value: value,
		next:  list.firstNode,
	}
}

// void	addLast(E e)
// Appends the specified element to the end of this list.
func (list *LinkedList[T]) AddLast(value T) {
	for current := list.firstNode; current != nil; current = current.next {
		if current.next == nil {
			current.next = &node[T]{
				value: value,
			}
		}
	}
}

// void	clear()
// Removes all of the elements from this list.

// Object	clone()
// Returns a shallow copy of this LinkedList.

// boolean	contains(Object o)
// Returns true if this list contains the specified element.
func (list *LinkedList[T]) Ccontains(value T) bool {
	var exist bool
	for current := list.firstNode; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}
	return exist
}

// Iterator<E>	descendingIterator()
// Returns an iterator over the elements in this deque in reverse sequential order.

// E	get(int index)
// Returns the element at the specified position in this list.
func (list *LinkedList[T]) Get(index int) *T {
	if index < 0 {
		return nil
	}
	for i, current := 0, list.firstNode; current != nil; i, current = i+1, current.next {
		if i == index {
			return &current.value
		}
	}
	return nil
}

// E	getFirst()
// Returns the first element in this list.
func (list *LinkedList[T]) GetFirst(index int) *T {
	if list.firstNode != nil {
		return &list.firstNode.value
	}
	return nil
}

// E	getLast()
// Returns the last element in this list.
func (list *LinkedList[T]) GetLast(index int) *T {
	for current := list.firstNode; current != nil; current = current.next {
		if current.next == nil {
			return &current.value
		}
	}
	return nil
}

// int	indexOf(Object o)
// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (list *LinkedList[T]) IndexOf(value T) int {
	for i, current := 0, list.firstNode; current != nil; i, current = i+1, current.next {
		if current.value == value {
			return i
		}
	}
	return -1
}

// int	lastIndexOf(Object o)
// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (list *LinkedList[T]) LastIndexOf(value T) int {
	index := -1
	for i, current := 0, list.firstNode; current != nil; i, current = i+1, current.next {
		if current.value == value {
			index = i
		}
	}
	return index
}

// ListIterator<E>	listIterator(int index)
// Returns a list-iterator of the elements in this list (in proper sequence), starting at the specified position in the list.

// E	peek()
// Retrieves, but does not remove, the head (first element) of this list.

// E	peekFirst()
// Retrieves, but does not remove, the first element of this list, or returns null if this list is empty.

// E	peekLast()
// Retrieves, but does not remove, the last element of this list, or returns null if this list is empty.

// E	remove(int index)
// Removes the element at the specified position in this list.
func (list *LinkedList[T]) Remove(index int) {
	if index == 0 && list.firstNode != nil {
		list.firstNode = list.firstNode.next
	}
	if index > 0 && list.firstNode != nil {
		for i, current := 0, list.firstNode.next; current != nil; i, current = i+1, current.next {
			if i+1 == index && current.next != nil {
				current.next = current.next.next
			}
		}
	}
}

// boolean	remove(Object o)
// Removes the first occurrence of the specified element from this list, if it is present.

// E	removeFirst()
// Removes and returns the first element from this list.
func (list *LinkedList[T]) RemoveFirst(value T) *T {
	var result *T
	if list.firstNode == nil {
		return result
	}
	result = &list.firstNode.value
	list.firstNode = list.firstNode.next
	return result
}

// boolean	removeFirstOccurrence(Object o)
// Removes the first occurrence of the specified element in this list (when traversing the list from head to tail).

// E	removeLast()
// Removes and returns the last element from this list.
func (list *LinkedList[T]) RemoveLast(value T) *T {
	var result *T
	if list.firstNode == nil {
		return result
	}
	for current := list.firstNode; current != nil; current = current.next {
		if current.next == nil {
			result = &list.firstNode.value
			list.firstNode = nil
			return result
		}
	}
	return result
}

// boolean	removeLastOccurrence(Object o)
// Removes the last occurrence of the specified element in this list (when traversing the list from head to tail).

// E	set(int index, E element)
// Replaces the element at the specified position in this list with the specified element.
func (list *LinkedList[T]) Set(index int, value T) {
	for i, current := 0, list.firstNode; current != nil; i, current = i+1, current.next {
		if i == index {
			current.value = value
		}
	}
}

// int	size()
// Returns the number of elements in this list.
func (list *LinkedList[T]) Size(value T) int {
	for i, current := 0, list.firstNode; current != nil; i, current = i+1, current.next {
		if current.next == nil {
			return i
		}
	}
	return 0
}

// Spliterator<E>	spliterator()
// Creates a late-binding and fail-fast Spliterator over the elements in this list.

// Object[]	toArray()
// Returns an array containing all of the elements in this list in proper sequence (from first to last element).
func (list *LinkedList[T]) ToArray(value T) []T {
	var arr []T
	for current := list.firstNode; current != nil; current = current.next {
		arr = append(arr, current.value)
	}
	return arr
}
