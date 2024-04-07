package skiplist

import (
	"reflect"
	"testing"
)

func TestInsert2(t *testing.T) {
	list, err := New[int32](10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list.Insert(20)
	list.Insert(10)
	expected := []int32{10, 20}
	if !reflect.DeepEqual(*(list.AsArray()), expected) {
		t.Errorf("Expected: %v, actual %v", expected, *(list.AsArray()))
		t.FailNow()
	}
}

func TestInsert3(t *testing.T) {
	list, err := New[int32](10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list.Insert(20)
	list.Insert(10)
	list.Insert(19)
	expected := []int32{10, 19, 20}
	if !reflect.DeepEqual(*(list.AsArray()), expected) {
		t.Errorf("Expected: %v, actual %v", expected, *(list.AsArray()))
		t.FailNow()
	}
}

func TestInsert5(t *testing.T) {
	list, err := New[int32](10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list.Insert(20)
	list.Insert(10)
	list.Insert(19)
	list.Insert(0)
	list.Insert(21)
	expected := []int32{0, 10, 19, 20, 21}
	if !reflect.DeepEqual(*(list.AsArray()), expected) {
		t.Errorf("Expected: %v, actual %v", expected, *(list.AsArray()))
		t.FailNow()
	}
}

func TestInsert10(t *testing.T) {
	list, err := New[int32](10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list.Insert(20)
	list.Insert(10)
	list.Insert(19)
	list.Insert(0)
	list.Insert(21)
	list.Insert(1)
	list.Insert(5)
	list.Insert(22)
	list.Insert(18)
	list.Insert(3)
	expected := []int32{0, 1, 3, 5, 10, 18, 19, 20, 21, 22}
	if !reflect.DeepEqual(*(list.AsArray()), expected) {
		t.Errorf("Expected: %v, actual %v", expected, *(list.AsArray()))
		t.FailNow()
	}
}

func TestContains(t *testing.T) {
	list, err := New[int32](10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	list.Insert(20)
	list.Insert(10)
	list.Insert(19)
	list.Insert(0)
	list.Insert(21)
	list.Insert(1)
	list.Insert(5)
	list.Insert(22)
	list.Insert(18)
	list.Insert(3)

	if !list.Contains(18) {
		t.Error("should contain 18")
		t.FailNow()
	}
	if !list.Contains(22) {
		t.Error("should contain 22")
		t.FailNow()
	}
	if list.Contains(2) {
		t.Error("should not contain 2")
		t.FailNow()
	}
	if list.Contains(23) {
		t.Error("should not contain 23")
		t.FailNow()
	}
}
