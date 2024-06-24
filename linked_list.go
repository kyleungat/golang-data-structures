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

	for i, current := 0, list.firstNode; current.next != nil; i, current = i+1, current.next {
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
// void	addLast(E e)
// Appends the specified element to the end of this list.
// void	clear()
// Removes all of the elements from this list.
// Object	clone()
// Returns a shallow copy of this LinkedList.
// boolean	contains(Object o)
// Returns true if this list contains the specified element.
// Iterator<E>	descendingIterator()
// Returns an iterator over the elements in this deque in reverse sequential order.
// E	element()
// Retrieves, but does not remove, the head (first element) of this list.
// E	get(int index)
// Returns the element at the specified position in this list.
// E	getFirst()
// Returns the first element in this list.
// E	getLast()
// Returns the last element in this list.
// int	indexOf(Object o)
// Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
// int	lastIndexOf(Object o)
// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
// ListIterator<E>	listIterator(int index)
// Returns a list-iterator of the elements in this list (in proper sequence), starting at the specified position in the list.
// boolean	offer(E e)
// Adds the specified element as the tail (last element) of this list.
// boolean	offerFirst(E e)
// Inserts the specified element at the front of this list.
// boolean	offerLast(E e)
// Inserts the specified element at the end of this list.
// E	peek()
// Retrieves, but does not remove, the head (first element) of this list.
// E	peekFirst()
// Retrieves, but does not remove, the first element of this list, or returns null if this list is empty.
// E	peekLast()
// Retrieves, but does not remove, the last element of this list, or returns null if this list is empty.
// E	poll()
// Retrieves and removes the head (first element) of this list.
// E	pollFirst()
// Retrieves and removes the first element of this list, or returns null if this list is empty.
// E	pollLast()
// Retrieves and removes the last element of this list, or returns null if this list is empty.
// E	pop()
// Pops an element from the stack represented by this list.
// void	push(E e)
// Pushes an element onto the stack represented by this list.
// E	remove()
// Retrieves and removes the head (first element) of this list.
// E	remove(int index)
// Removes the element at the specified position in this list.
// boolean	remove(Object o)
// Removes the first occurrence of the specified element from this list, if it is present.
// E	removeFirst()
// Removes and returns the first element from this list.
// boolean	removeFirstOccurrence(Object o)
// Removes the first occurrence of the specified element in this list (when traversing the list from head to tail).
// E	removeLast()
// Removes and returns the last element from this list.
// boolean	removeLastOccurrence(Object o)
// Removes the last occurrence of the specified element in this list (when traversing the list from head to tail).
// E	set(int index, E element)
// Replaces the element at the specified position in this list with the specified element.
// int	size()
// Returns the number of elements in this list.
// Spliterator<E>	spliterator()
// Creates a late-binding and fail-fast Spliterator over the elements in this list.
// Object[]	toArray()
// Returns an array containing all of the elements in this list in proper sequence (from first to last element).
// <T> T[]	toArray(T[] a)
// Returns an array containing all of the elements in this list in proper sequence (from first to last element); the runtime type of the returned array is that of the specified array.
