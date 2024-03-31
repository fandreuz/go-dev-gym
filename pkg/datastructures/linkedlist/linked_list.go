package linkedlist

import "errors"

type node[T any] struct {
	next  *node[T]
	value T
}

func NewLinkedList[T any](value T) *node[T] {
	return &node[T]{nil, value}
}

func (handle *node[T]) HasNext() bool {
	return handle.next != nil
}

func (handle *node[T]) Next() *node[T] {
	return handle.next
}

func (handle *node[T]) Append(values ...T) (*node[T], error) {
	if handle.HasNext() {
		return nil, errors.New("given handle has already a next")
	}
	for _, value := range values {
		handle.next = &node[T]{nil, value}
		handle = handle.next
	}
	return handle, nil
}

func (handle *node[T]) Last() *node[T] {
	var current *node[T] = handle
	for current.HasNext() {
		current = current.Next()
	}
	return current
}

func (handle *node[T]) Lenght() int64 {
	var length int64 = 1

	var current *node[T] = handle
	for current.HasNext() {
		current = current.Next()
		length++
	}
	return length
}
