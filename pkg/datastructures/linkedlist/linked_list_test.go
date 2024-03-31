package linkedlist

import "testing"

func TestNew(t *testing.T) {
	handle := NewLinkedList(10)

	if handle.value != 10 {
		t.Error()
	}
	if handle.next != nil {
		t.Error()
	}
	if !handle.head {
		t.Error()
	}
}

func TestHasNext(t *testing.T) {
	handle := NewLinkedList(10)

	if handle.HasNext() {
		t.Error()
	}
}

func TestAppend(t *testing.T) {
	handle := NewLinkedList(10)
	handle2 := handle.Append(12)

	if handle.value != 10 {
		t.Error()
	}
	if handle2.value != 12 {
		t.Error()
	}
	if handle.next != handle2 {
		t.Error()
	}
	if handle2.HasNext() {
		t.Error()
	}
	if handle2.head {
		t.Error()
	}
}

func TestAppendMany(t *testing.T) {
	handle := NewLinkedList(10)
	handle.Append(12, 14, 16)

	if handle.value != 10 {
		t.Error()
	}
	if handle.next.value != 12 {
		t.Error()
	}
	if handle.next.next.value != 14 {
		t.Error()
	}
	if handle.next.next.next.value != 16 {
		t.Error()
	}
}

func TestNext(t *testing.T) {
	handle := NewLinkedList(10)
	handle.Append(12)

	if handle.Next().value != 12 {
		t.Error()
	}
}

func TestLast(t *testing.T) {
	handle := NewLinkedList(10)
	handle.Append(12, 14)

	if handle.Last().value != 14 {
		t.Error()
	}
}

func TestLength(t *testing.T) {
	handle := NewLinkedList(10)
	handle.Append(12, 14, 16)

	l, err := handle.Lenght()
	if err != nil {
		t.Error(err.Error())
	}
	if l != 4 {
		t.Error()
	}
}
