package skiplist

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Probability of jumping to the next skip-list level, see decideLevel()
const p = .5
const intervalueSeparatorLength = 3
const intervalueSeparator = '-'
const crossChar = '|'

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

func (sliceByLevel SliceByLevel[T]) superNext() *node[T] {
	target := sliceByLevel.atLevel(1)
	if len(target.nexts) > 0 {
		return target.next(1)
	}
	return &(*placeholderNodes[T](1))[0]
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
			for !currentNode.placeholder && currentNode.value < value {
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
	return closestSmallers.superNext().hasValue(value)
}

func (list SkipList[T]) decideLevel() uint {
	level := uint(1)
	for v := rand.Float32(); v < p && level < list.maxLevel; v = rand.Float32() {
		level++
	}
	return level
}

// true when added
func (list *SkipList[T]) Insert(value T) bool {
	closestSmallers := list.closestSmallers(value)
	if len(closestSmallers.nodes) > 0 && closestSmallers.superNext().hasValue(value) {
		return false
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

	return true
}

func (list SkipList[T]) AsArray() *[]T {
	array := make([]T, 0)
	for currentNode := list.heads[0]; !currentNode.placeholder; currentNode = currentNode.nexts[0] {
		array = append(array, currentNode.value)
	}
	return &array
}

// true when removed
func (list SkipList[T]) Remove(value T) bool {
	closestSmallers := list.closestSmallers(value)
	if len(closestSmallers.nodes) > 0 && !closestSmallers.superNext().hasValue(value) {
		return false
	}

	target := *closestSmallers.superNext()
	l := uint(len(target.nexts))
	for level := uint(1); level <= l; level++ {
		oldNext := *target.next(level)
		smallerAtLevel := closestSmallers.atLevel(level)
		if !smallerAtLevel.placeholder && smallerAtLevel.next(level).hasValue(value) {
			smallerAtLevel.setNext(level, oldNext)
		}

		if list.head(level).hasValue(value) {
			list.setHead(level, oldNext)
		}
	}

	return true
}

func (list SkipList[T]) distancesPerLevel() [][]uint64 {
	output := make([][]uint64, list.topLevel)
	for level := uint(1); level <= list.topLevel; level++ {
		output[level-1] = make([]uint64, 1)
	}

	for currentNode := list.head(1); !currentNode.placeholder; currentNode = currentNode.next(1) {
		level := uint(1)
		for ; level <= uint(len(currentNode.nexts)); level++ {
			if !currentNode.next(level).placeholder {
				output[level-1] = append(output[level-1], 0)
			}
		}
		for ; level <= uint(list.topLevel); level++ {
			l := len(output[level-1])
			output[level-1][l-1]++
		}
	}

	return output
}

func printInternodeSeparator() {
	for i := 0; i < intervalueSeparatorLength; i++ {
		fmt.Printf("%c", intervalueSeparator)
	}
}

func printNodeEmptyPlaceholder(nodeMaxWidth uint8) {
	for i := uint8(0); i < nodeMaxWidth; i++ {
		fmt.Printf("%c", intervalueSeparator)
	}
}

func printNodeFullPlaceholder(nodeMaxWidth uint8) {
	for i := uint8(0); i < nodeMaxWidth/2; i++ {
		fmt.Printf("%c", intervalueSeparator)
	}
	fmt.Printf("%c", crossChar)
	for i := uint8(0); i < nodeMaxWidth/2; i++ {
		fmt.Printf("%c", intervalueSeparator)
	}
}

func centerString(str string, width int) string {
	spaces := int(float64(width-len(str)) / 2)
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", width-(spaces+len(str)))
}

func printNode[T any](value T, nodeMaxWidth uint8) {
	fmt.Print(centerString(fmt.Sprintf("%v", value), int(nodeMaxWidth)))
}

func (list SkipList[T]) widestNodeWidth() uint8 {
	maxWidth := uint8(0)
	for currentNode := list.head(1); !currentNode.placeholder; currentNode = currentNode.next(1) {
		maxWidth = max(maxWidth, uint8(len(fmt.Sprintf("%v", currentNode.value))))
	}
	return maxWidth
}

func (list SkipList[T]) PrettyPrint() {
	distances := list.distancesPerLevel()
	nodeMaxWidth := list.widestNodeWidth() + 2
	if nodeMaxWidth%2 == 0 {
		nodeMaxWidth++
	}

	for level := uint(len(distances)); level > 1; level-- {
		distancesAtLevel := distances[level-1]
		for i := uint(0); i < uint(len(distancesAtLevel)); i++ {
			for j := uint64(0); j < uint64(distancesAtLevel[i]); j++ {
				printNodeEmptyPlaceholder(nodeMaxWidth)
				printInternodeSeparator()
			}

			printNodeFullPlaceholder(nodeMaxWidth)
			if i != uint(len(distancesAtLevel)-1) {
				printInternodeSeparator()
			}
		}
		fmt.Println()
	}

	for currentNode := list.head(1); !currentNode.placeholder; currentNode = currentNode.next(1) {
		if currentNode != list.head(1) {
			printInternodeSeparator()
		}
		printNode(currentNode.value, nodeMaxWidth)
	}
	fmt.Println()
}
