package hashmap

const defaultArraySize = uint64(10)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type Hasher[K comparable] func(key K) uint32

type HashMap[K comparable, V any] struct {
	// TODO Replace with linked list
	entries [][]*entry[K, V]
	hasher  Hasher[K]
}

func New[K comparable, V any](hasher Hasher[K]) HashMap[K, V] {
	return NewWithSize[K, V](hasher, defaultArraySize)
}

func NewWithSize[K comparable, V any](hasher Hasher[K], array_size uint64) HashMap[K, V] {
	array := make([][]*entry[K, V], array_size)
	return HashMap[K, V]{entries: array, hasher: hasher}
}

func (hm HashMap[K, V]) arrayIndex(key K) uint32 {
	hash := hm.hasher(key)
	return hash % uint32(len(hm.entries))
}

func (hm HashMap[K, V]) Put(key K, value V) {
	index := hm.arrayIndex(key)
	currentEntries := hm.entries[index]
	for i := uint32(0); i < uint32(len(currentEntries)); i++ {
		// TODO
		if currentEntries[i].key == key {
			currentEntries[i].key = key
			currentEntries[i].value = value
			return
		}
	}
	hm.entries[index] = append(currentEntries, &entry[K, V]{key, value})
}

func (hm HashMap[K, V]) Get(key K) *V {
	index := hm.arrayIndex(key)
	currentEntries := hm.entries[index]
	for i := uint32(0); i < uint32(len(currentEntries)); i++ {
		if currentEntries[i].key == key {
			return &currentEntries[i].value
		}
	}
	return nil
}

func (hm HashMap[K, V]) Remove(key K) *V {
	index := hm.arrayIndex(key)
	currentEntries := hm.entries[index]
	for i := uint32(0); i < uint32(len(currentEntries)); i++ {
		if currentEntries[i].key == key {
			value := currentEntries[i].value
			hm.entries[index] = append(hm.entries[index][:i], hm.entries[index][i+1:]...)
			return &value
		}
	}
	return nil
}

func (hm HashMap[K, V]) Size() uint64 {
	counter := uint64(0)
	for i := uint64(0); i < uint64(len(hm.entries)); i++ {
		counter += uint64(len(hm.entries[i]))
	}
	return counter
}
