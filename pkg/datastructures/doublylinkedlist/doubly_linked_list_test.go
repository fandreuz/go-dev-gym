package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	handle := New(10)

	if handle.value != 10 {
		t.FailNow()
	}
	if handle.next != nil {
		t.FailNow()
	}
	if handle.prev != nil {
		t.FailNow()
	}
}

func TestHasNext(t *testing.T) {
	handle := New(10)

	if handle.HasNext() {
		t.FailNow()
	}
}

func TestAppend(t *testing.T) {
	handle := New(10)
	handle2 := handle.Append(12)

	if handle.value != 10 {
		t.FailNow()
	}
	if handle2.value != 12 {
		t.FailNow()
	}
	if handle.next != handle2 {
		t.FailNow()
	}
	if handle2.prev != handle {
		t.FailNow()
	}
	if handle2.HasNext() {
		t.FailNow()
	}
}

func TestAppendMany(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	if handle.value != 10 {
		t.FailNow()
	}
	if handle.next.value != 12 {
		t.FailNow()
	}
	if handle.next.next.value != 14 {
		t.FailNow()
	}
	if handle.next.next.next.value != 16 {
		t.FailNow()
	}
	if handle.next.next.next.prev.value != 14 {
		t.FailNow()
	}
	if handle.next.next.next.prev.prev.value != 12 {
		t.FailNow()
	}
	if handle.next.next.next.prev.prev.prev.value != 10 {
		t.FailNow()
	}
}

func TestNext(t *testing.T) {
	handle := New(10)
	handle.Append(12)

	if handle.Next().value != 12 {
		t.FailNow()
	}
}

func TestPrevious(t *testing.T) {
	handle := New(10)
	handle.Append(12)

	if handle.Next().Previous().value != 10 {
		t.FailNow()
	}
}

func TestLast(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14)

	if handle.Last().value != 14 {
		t.FailNow()
	}
}

func TestLength(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	l := handle.Length()
	if l != 4 {
		t.FailNow()
	}
}

func TestAt(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	handle2, err := handle.Shift(0)
	if err != nil {
		t.Error(err.Error())
	}
	if handle2.value != 10 {
		t.FailNow()
	}

	handle2, err = handle.Shift(3)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if handle2.value != 16 {
		t.FailNow()
	}
}

func TestShiftError(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	_, err := handle.Shift(4)
	if err == nil {
		t.FailNow()
	}
}

func TestShift(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	handle2, _ := handle.Shift(1)
	handle3, err := handle2.Shift(2)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if handle3.value != 16 {
		t.FailNow()
	}
}

func TestAsArray(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)

	expected := []int{10, 12, 14, 16}
	if !reflect.DeepEqual(*(handle.AsArray()), expected) {
		t.Errorf("Expected: %v, actual %v", expected, *(handle.AsArray()))
		t.FailNow()
	}
	handle2, _ := handle.Shift(2)
	if !reflect.DeepEqual(*(handle2.AsArray()), expected[2:]) {
		t.Errorf("Expected: %v, actual %v", expected[2:], *(handle2.AsArray()))
		t.FailNow()
	}
}

func TestInsertBefore(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)
	handle2, _ := handle.Shift(2)

	handle2.InsertBefore(1, 2, 3, 4)

	expected := []int{10, 12, 1, 2, 3, 4, 14, 16}
	if !reflect.DeepEqual(*(handle.AsArray()), expected) {
		t.Errorf("Expected: %v, actual %v", expected, *(handle.AsArray()))
		t.FailNow()
	}
}

func TestInsertAfter(t *testing.T) {
	handle := New(10)
	handle.Append(12, 14, 16)
	handle2, _ := handle.Shift(2)

	handle2.InsertAfter(1, 2, 3, 4)

	expected := []int{10, 12, 14, 1, 2, 3, 4, 16}
	if !reflect.DeepEqual(*(handle.AsArray()), expected) {
		t.Errorf("Expected: %v, actual %v", expected, *(handle.AsArray()))
		t.FailNow()
	}
}
