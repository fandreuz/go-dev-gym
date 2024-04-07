package skiplist

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"math/rand"
)

// Probability of jumping to the next skip-list level, see decideLevel()
const p = .5

type node[T cmp.Ordered] struct {
	nexts       []node[T]
	value       T
	placeholder bool
}

func placeholderNodes[T cmp.Ordered](count uint) *[]node[T] {
	nodes := make([]node[T], count)
	for i := uint(0); i < count; i++ {
		nodes[i] = node[T]{placeholder: true}
	}
	return &nodes
}

func newNode[T cmp.Ordered](level uint, value T) node[T] {
	return node[T]{*placeholderNodes[T](level), value, false}
}

type SkipList[T cmp.Ordered] struct {
	maxLevel uint
	topLevel uint
	heads    []node[T]
}

type SliceByLevel[T cmp.Ordered] struct {
	nodes []node[T]
}

func New[T cmp.Ordered](n uint64) (SkipList[T], error) {
	if n <= 0 {
		return SkipList[T]{}, errors.New("n must be positive")
	}
	maxLevel := uint(math.Ceil(math.Log(float64(n))))
	return SkipList[T]{maxLevel: maxLevel, topLevel: 0, heads: *placeholderNodes[T](maxLevel)}, nil
}

func (list SkipList[T]) head(level uint) *node[T] {
	return &list.heads[level-1]
}

func (list SkipList[T]) setHead(level uint, newHead node[T]) {
	list.heads[level-1] = newHead
}

func (node *node[T]) next(level uint) *node[T] {
	return &node.nexts[level-1]
}

func (node *node[T]) setNext(level uint, next node[T]) {
	node.nexts[level-1] = next
}

func (node *node[T]) hasValue(value T) bool {
	return !node.placeholder && node.value == value
}

func (byLevel SliceByLevel[T]) atLevel(level uint) *node[T] {
	return &byLevel.nodes[level-1]
}

func (list SkipList[T]) closestSmallers(value T) SliceByLevel[T] {
	ancestors := *placeholderNodes[T](list.topLevel)

	if list.topLevel > 0 {
		currentNode := list.head(list.topLevel)
		var prev node[T] = ancestors[0]

		for level := list.topLevel; level > 0; level-- {
			for !currentNode.placeholder && currentNode.value <= value {
				prev = *currentNode
				currentNode = currentNode.next(level)
			}

			ancestors[level-1] = prev

			if level > 1 { // Update for next loop
				if prev.placeholder {
					currentNode = list.head(level - 1)
				} else {
					if prev.next(level-1).value <= value {
						currentNode = prev.next(level - 1)
					}
				}
			}
		}
	}

	return SliceByLevel[T]{ancestors}
}

func (list SkipList[T]) Contains(value T) bool {
	closestSmallers := list.closestSmallers(value)
	return closestSmallers.atLevel(1).hasValue(value)
}

func (list SkipList[T]) decideLevel() uint {
	level := uint(1)
	for v := rand.Float32(); v < p && level < list.maxLevel; v = rand.Float32() {
		level++
	}
	return level
}

func (list *SkipList[T]) Insert(value T) error {
	closestSmallers := list.closestSmallers(value)
	if len(closestSmallers.nodes) > 0 && closestSmallers.atLevel(list.topLevel).hasValue(value) {
		return fmt.Errorf("value exists already: %v", value)
	}

	selectedLevel := list.decideLevel()
	newNode := newNode(selectedLevel, value)

	level := uint(1)
	for ; level <= min(selectedLevel, list.topLevel); level++ {
		nodeAtLevel := closestSmallers.atLevel(level)
		if !nodeAtLevel.placeholder {
			oldNext := *nodeAtLevel.next(level)
			newNode.setNext(level, oldNext)
			nodeAtLevel.setNext(level, newNode)
		} else {
			oldHead := *list.head(level)
			list.setHead(level, newNode)
			newNode.setNext(level, oldHead)
		}
	}

	if list.topLevel < selectedLevel {
		for ; level <= selectedLevel; level++ {
			list.setHead(level, newNode)
		}
		list.topLevel = selectedLevel
	}

	return nil
}

func (list SkipList[T]) AsArray() *[]T {
	array := make([]T, 0)
	for currentNode := list.heads[0]; !currentNode.placeholder; currentNode = currentNode.nexts[0] {
		array = append(array, currentNode.value)
	}
	return &array
}
