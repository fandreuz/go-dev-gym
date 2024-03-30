package datastructures

import "errors"

type node[T any] struct {
	next  *node[T]
	value T
}

type LLHandle[T any] node[T]

func NewLinkedList[T any](value T) *LLHandle[T] {
	return &LLHandle[T]{nil, value}
}

func (handle *LLHandle[T]) HasNext() bool {
	return handle.next != nil
}

func (handle *LLHandle[T]) Next() *LLHandle[T] {
	return (*LLHandle[T])(handle.next)
}

func (handle *LLHandle[T]) Append(value T) (*LLHandle[T], error) {
	if handle.HasNext() {
		return nil, errors.New("given handle has already a next")
	}
	handle.next = &node[T]{nil, value}
	return (*LLHandle[T])(handle.next), nil
}

func (handle *LLHandle[T]) Last() *LLHandle[T] {
	var current *LLHandle[T] = handle
	for current.HasNext() {
		current = current.Next()
	}
	return current
}
