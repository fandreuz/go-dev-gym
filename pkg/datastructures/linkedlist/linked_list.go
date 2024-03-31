package linkedlist

import "errors"

type node[T any] struct {
	next  *node[T]
	value T
	head  bool
}

func NewLinkedList[T any](value T) *node[T] {
	return &node[T]{nil, value, true}
}

func (handle *node[T]) HasNext() bool {
	return handle.next != nil
}

func (handle *node[T]) Next() *node[T] {
	return handle.next
}

func (handle *node[T]) Append(values ...T) *node[T] {
	handle = handle.Last()

	for _, value := range values {
		handle.next = &node[T]{nil, value, false}
		handle = handle.next
	}
	return handle
}

func (handle *node[T]) Last() *node[T] {
	var current *node[T] = handle
	for current.HasNext() {
		current = current.Next()
	}
	return current
}

func (handle *node[T]) Lenght() (int64, error) {
	if !handle.head {
		return -1, errors.New("handle is not head")
	}

	var length int64 = 1
	var current *node[T] = handle
	for current.HasNext() {
		current = current.Next()
		length++
	}
	return length, nil
}
