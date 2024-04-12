package hashmap

import (
	"hash/fnv"
	"testing"
)

func hashFunc(value string) uint32 {
	hasher := fnv.New32a()
	hasher.Write([]byte(value))
	return hasher.Sum32()
}

func TestSize(t *testing.T) {
	hm := New[string, int32](hashFunc)
	hm.Put("ciao", 20)
	hm.Put("hello", 30)
	hm.Put("ciao2", 20)

	if hm.Size() != 3 {
		t.Errorf("Expected size %d, got %d", 3, hm.Size())
		t.FailNow()
	}
}

func TestGet(t *testing.T) {
	hm := New[string, int32](hashFunc)
	hm.Put("ciao", 20)
	hm.Put("hello", 30)
	hm.Put("ciao2", 20)

	result := hm.Get("ciao2")

	if result == nil {
		t.Errorf("Expected %d, got nil", 20)
		t.FailNow()
	}

	if *result != 20 {
		t.Errorf("Expected %d, got %d", 20, *result)
	}
}

func TestRemove(t *testing.T) {
	hm := New[string, int32](hashFunc)
	hm.Put("ciao", 20)
	hm.Put("hello", 30)
	hm.Put("ciao2", 20)

	result := hm.Remove("hello")

	if result == nil {
		t.Errorf("Expected %d, got nil", 20)
		t.FailNow()
	}

	if *result != 30 {
		t.Errorf("Expected %d, got %d", 20, *result)
		t.FailNow()
	}

	if hm.Size() != 2 {
		t.Errorf("Expected %d, got %d", 2, hm.Size())
		t.FailNow()
	}

	if hm.Get("hello") != nil {
		t.Errorf("Expected nil, got %d", *hm.Get("hello"))
		t.FailNow()
	}
}
