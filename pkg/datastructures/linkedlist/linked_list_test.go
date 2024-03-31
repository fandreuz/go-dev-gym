package linkedlist

import "testing"

func TestNew(t *testing.T) {
	handle := New(10)

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
	handle := New(10)

	if handle.HasNext() {
		t.Error()
	}
}

func TestAppend(t *testing.T) {
	handle := New(10)
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
	handle := New(10)
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
	handle := New(10)
	handle.Append(12)

	if handle.Next().value != 12 {
		t.Error()
	}
}

func TestLast(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14)

	if handle.Last().value != 14 {
		t.Error()
	}
}

func TestLength(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	l, err := handle.Length()
	if err != nil {
		t.Error(err.Error())
	}
	if l != 4 {
		t.Error()
	}
}

func TestAt(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	handle2, err := handle.At(0)
	if err != nil {
		t.Error(err.Error())
	}
	if handle2.value != 10 {
		t.Error()
	}

	handle2, err = handle.At(3)
	if err != nil {
		t.Error(err.Error())
	}
	if handle2.value != 16 {
		t.Error()
	}
}

func TestAtError(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	_, err := handle.At(4)
	if err == nil {
		t.Error()
	}
}
