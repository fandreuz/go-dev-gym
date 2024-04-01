package doublylinkedlist

import (
	"fmt"
)

type node[T any] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

type explorer[T any] func(*node[T]) (currentHandle *node[T])

func New[T any](value T) *node[T] {
	return &node[T]{nil, nil, value}
}

func (handle *node[T]) IsFirst() bool {
	return handle.prev == nil
}

func (handle *node[T]) IsLast() bool {
	return handle.next == nil
}

func (handle *node[T]) HasNext() bool {
	return !handle.IsLast()
}

func (handle *node[T]) Next() *node[T] {
	return handle.next
}

func (handle *node[T]) Previous() *node[T] {
	return handle.prev
}

func (handle *node[T]) Append(values ...T) *node[T] {
	handle = handle.Last()

	for _, value := range values {
		handle.next = &node[T]{handle, nil, value}
		handle = handle.next
	}
	return handle
}

func (handle *node[T]) loop(explorerFunc explorer[T]) *node[T] {
	previousHandle := handle
	currentHandle := handle
	for currentHandle != nil {
		previousHandle = currentHandle
		currentHandle = explorerFunc(currentHandle)
	}
	return previousHandle
}

func (handle *node[T]) First() *node[T] {
	return handle.loop(func(currentHandle *node[T]) *node[T] {
		return currentHandle.prev
	})
}

func (handle *node[T]) Last() *node[T] {
	return handle.loop(func(currentHandle *node[T]) *node[T] {
		return currentHandle.next
	})
}

func (handle *node[T]) Length() uint64 {
	length := uint64(0)
	handle.loop(func(currentHandle *node[T]) *node[T] {
		length++
		return currentHandle.next
	})
	return length
}

func (handle *node[T]) Shift(of uint64) (*node[T], error) {
	currentShift := uint64(0)
	handleAt := handle.loop(func(currentHandle *node[T]) *node[T] {
		if currentShift == of {
			return nil
		}
		if currentHandle.next != nil {
			currentShift++
		}
		return currentHandle.next
	})

	if currentShift < of {
		return nil, fmt.Errorf("shifted of %v and could not procede", currentShift-1)
	}
	return handleAt, nil
}

func (handle *node[T]) Remove() *node[T] {
	if handle.prev != nil {
		handle.prev.next = handle.next
	}
	if handle.next != nil {
		handle.next.prev = handle.prev
	}
	return handle.prev
}

func (handle *node[T]) RemoveFirst() *node[T] {
	return handle.First().Remove()
}

func (handle *node[T]) RemoveLast() *node[T] {
	return handle.Last().Remove()
}

func (handle *node[T]) InsertAfter(values ...T) {
	originalNext := handle.next

	currentParent := handle
	for _, value := range values {
		newNode := &node[T]{currentParent, nil, value}
		currentParent.next = newNode
		newNode.prev = currentParent

		currentParent = newNode
	}
	if originalNext != nil {
		currentParent.next = originalNext
		originalNext.prev = currentParent
	}
}

func (handle *node[T]) InsertBefore(values ...T) {
	if handle.IsFirst() {
		if len(values) == 0 {
			return
		}
		temp := New(values[0])
		temp.Append(values[1:]...)

		temp.prev, handle.prev = handle.prev, temp.Last()
		temp.Last().next = handle
	} else {
		handle.Previous().InsertAfter(values...)
	}
}

func (handle *node[T]) AsArray() *[]T {
	l := handle.Length()
	arr := make([]T, l)
	for i := uint64(0); i < l; i++ {
		arr[i] = handle.value
		handle, _ = handle.Shift(1)
	}
	return &arr
}
