package datastructures

import "testing"

func TestNew(t *testing.T) {
	handle := NewLinkedList(10)

	if handle.value != 10 {
		t.Fail()
	}
	if handle.next != nil {
		t.Fail()
	}
}

func TestHasNext(t *testing.T) {
	handle := NewLinkedList(10)

	if handle.HasNext() {
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	handle := NewLinkedList(10)
	handle2, err := handle.Append(12)

	if err != nil {
		t.Fatalf(err.Error())
	}
	if handle.value != 10 {
		t.Fail()
	}
	if handle2.value != 12 {
		t.Fail()
	}
	if handle2.HasNext() {
		t.Fail()
	}
}

func TestNext(t *testing.T) {
	handle := NewLinkedList(10)
	handle.Append(12)

	if handle.Next().value != 12 {
		t.Fail()
	}
}

func TestLast(t *testing.T) {
	handle := NewLinkedList(10)
	handle2, _ := handle.Append(12)
	handle2.Append(14)

	if handle.Last().value != 14 {
		t.Fail()
	}
}
