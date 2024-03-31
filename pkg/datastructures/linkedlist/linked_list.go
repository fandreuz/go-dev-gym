package linkedlist

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	next  *node[T]
	value T
	head  bool
}

func New[T any](value T) *node[T] {
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

func (handle *node[T]) Length() (int64, error) {
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

func (handle *node[T]) At(index int64) (*node[T], error) {
	if !handle.head {
		return nil, errors.New("handle is not head")
	}

	var current *node[T] = handle
	for currentIndex := int64(0); currentIndex < index; currentIndex++ {
		if !current.HasNext() {
			msg := fmt.Sprintf("list too short for the given index: %v", index)
			return nil, errors.New(msg)
		}
		current = current.Next()
	}
	return current, nil
}

func (handle *node[T]) RemoveNext() {
	handle.next = nil
}
