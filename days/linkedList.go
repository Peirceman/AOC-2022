package days

import (
	"errors"
	"fmt"
	"strings"
)

type doublyLinkedNode[T any] struct {
	previous, next *doublyLinkedNode[T]
	value *T
}

type doublyLinkedList[T any] struct {
	start, end *doublyLinkedNode[T]
	length int
}

func newDoublyLinkedNode[T any](value *T) *doublyLinkedNode[T] {
	return &doublyLinkedNode[T] {
		nil, nil, value,
	}
}

func (this *doublyLinkedNode[T]) copy() *doublyLinkedNode[T] {
	return newDoublyLinkedNode(this.value)
}

func NewDoublyLinkedListConstants[T any](intialValues ...T) *doublyLinkedList[T] {
	result := &doublyLinkedList[T] {
		nil, nil, 0,
	}

	if (len(intialValues) == 0) {
		return result
	}

	for _, value := range intialValues {
		a := value
		result.addLast(&a)
	}

	return result
}

func NewDoublyLinkedList[T any](intialValues ...*T) *doublyLinkedList[T] {
	result := &doublyLinkedList[T] {
		nil, nil, 0,
	}

	if len(intialValues) == 0 {
		return result
	}

	for _, value := range intialValues {
		result.addLast(value)
	}

	return result
}

// if index < 0 or index >= length, return default value for node
func (this *doublyLinkedList[T]) Get(index int) (*T, error) {
	node, getNodeError := this.getNode(index)
	if getNodeError == nil {
		return node.value, nil
	}

	var defaultValue T
	return &defaultValue, getNodeError
}

// if index < 0 or index >= length, return nil value for node
func (this *doublyLinkedList[T]) getNode(index int) (*doublyLinkedNode[T], error) {
	if index < 0 || index >= this.length {
		return nil, errors.New(fmt.Sprintf("index %d out of bounds for length %d\n", index, this.length))
	}

	if (index == 0) {
		return this.start, nil
	}

	if (index == this.length - 1) {
		return this.end, nil
	}

	if (index <= this.length / 2) {
		currentNode := this.start
		for i := 0; i < index; i++ {
			currentNode = currentNode.next
		}
		return currentNode, nil
	}


	// 0 1 5 2 3 4
	//           ^
	// 6
	// 5
	index++
	currentNode := this.end
	for i := this.length; i > index; i-- {
		currentNode = currentNode.previous
	}
	return currentNode, nil
}

func (this *doublyLinkedList[T]) removeFirst() {
	if this.length == 1 {
		this.end = nil
		this.start = nil
		this.length = 0
		return
	}

	newStart := this.start.next
	newStart.previous = nil
	this.start = newStart
	this.length--
}

func (this *doublyLinkedList[T]) removeLast() {
	if this.length == 1 {
		this.removeFirst()
		return
	}

	newEnd := this.end.previous
	newEnd.next = nil
	this.end = newEnd
	this.length--
}

func (this *doublyLinkedList[T]) Remove(index int) {
	if this.length < 1 {
		return
	}

	if index < 1 {
		this.removeFirst()
		return
	}

	if index >= this.length - 1 {
		this.removeLast()
		return
	}

	nodeToRemove,_ := this.getNode(index)
	nodeToRemove.previous.next = nodeToRemove.next
	nodeToRemove.next.previous = nodeToRemove.previous
	this.length--
}

// deletes from the end so the list.length <= length
func (this *doublyLinkedList[T]) MaxLength(length int) {
	if (length <= 0) {
		this.start = nil
		this.end = nil
		this.length = 0
		return
	}

	if (length >= this.length) {
		return
	}

	newEnd,_ := this.getNode(length - 1)
	newEnd.next = nil
	this.end = newEnd
	this.length = length
}

func (this *doublyLinkedList[T]) Reverse() *doublyLinkedList[T] {
	if this.length == 0 {
		return NewDoublyLinkedList[T]()
	}

	result := NewDoublyLinkedList[T]()
	result.start = this.end.copy()
	thisCurrentNode := this.end
	resultCurrentNode := result.start
	for thisCurrentNode.previous != nil {
		resultCurrentNode.next = thisCurrentNode.previous.copy()
		resultCurrentNode.next.previous = resultCurrentNode

		thisCurrentNode = thisCurrentNode.previous
		resultCurrentNode = resultCurrentNode.next
	}
	result.end = resultCurrentNode
	result.length = this.length
	return result
}

// if startInclusive < 0, startInclusive = 0
// if endExclusive > length, endExclusive = length
// if startInclusive > endExclusive, return nil
func (this *doublyLinkedList[T]) SubList(startInclusive, endExclusive int) *doublyLinkedList[T] {
	if startInclusive < 0 {
		startInclusive = 0
	}

	if endExclusive > this.length {
		endExclusive = this.length
	}

	if startInclusive > endExclusive {
		return nil
	}

	if startInclusive == 0 && endExclusive == this.length {
		return this.Copy()
	}

	result := NewDoublyLinkedList[T]()
	currentNode,_ := this.getNode(startInclusive)
	for i := startInclusive; i < endExclusive; i++ {
		result.addLastNode(currentNode.copy())
		currentNode = currentNode.next
	}

	return result
}

func (this *doublyLinkedList[T]) InsertConstant(index int, value T) {
	this.Insert(index, &value)
}

// if index <= 0, insert at start
// if index >= length, insert at end
func (this *doublyLinkedList[T]) Insert(index int, value *T) {
	nodeToInsert := newDoublyLinkedNode(value)

	if this.length == 0 {
		this.start = nodeToInsert
		this.end = nodeToInsert
		this.length = 1
		return
	}

	if index <= 0 {
		nodeToInsert.next = this.start
		this.start.previous = nodeToInsert
		this.start = nodeToInsert
		this.length++
		return
	}

	if index > this.length {
		nodeToInsert.previous = this.end
		this.end.next = nodeToInsert
		this.end = nodeToInsert
		this.length++
		return
	}

	nextNode,_ := this.getNode(index)
	previousNode := nextNode.previous
	previousNode.next = nodeToInsert
	nodeToInsert.previous = previousNode
	nodeToInsert.next = nextNode
	nextNode.previous = nodeToInsert
	this.length++
}



func (this *doublyLinkedList[T]) addLast(value *T) {
	this.addLastNode(newDoublyLinkedNode(value))
}

func (this *doublyLinkedList[T]) addLastNode(newNode *doublyLinkedNode[T]) {
	if (this.length == 0) {
		this.start = newNode
		this.end = newNode
		this.length++
		return
	}

	this.end.next = newNode;
	newNode.previous = this.end;
	this.end = newNode;
	this.length++
}

func (this *doublyLinkedList[T]) Copy() *doublyLinkedList[T] {
	result := NewDoublyLinkedList[T]()

	if (this.length == 0) {
		return result
	}

	result.start = this.start.copy()
	thisCurrentNode := this.start
	resultCurrentNode := result.start

	for thisCurrentNode.next != nil {
		resultCurrentNode.next = thisCurrentNode.next.copy()
		resultCurrentNode.next.previous = resultCurrentNode

		thisCurrentNode = thisCurrentNode.next
		resultCurrentNode = resultCurrentNode.next
	}
	result.end = resultCurrentNode
	result.length = this.length

	return result
}

func (this *doublyLinkedList[T]) Append(other *doublyLinkedList[T]) *doublyLinkedList[T] {
	result := this.Copy()

	if (other.length == 0) {
		return result
	}

	result.addLastNode(other.start.copy())
	otherCurrentNode := other.start;
	resultCurrentNode := result.end;

	for otherCurrentNode.next != nil {
		resultCurrentNode.next = otherCurrentNode.next.copy()
		resultCurrentNode.next.previous = resultCurrentNode

		otherCurrentNode = otherCurrentNode.next
		resultCurrentNode = resultCurrentNode.next
	}

	result.end = resultCurrentNode
	result.length += other.length - 1

	return result
}

func (this *doublyLinkedList[T]) String() string{
	if this.length == 0 {
		return "[]"
	}

	builder := strings.Builder{}
	builder.WriteString("[")
	for currentNode := this.start; currentNode != this.end; currentNode = currentNode.next {
		builder.WriteString(fmt.Sprintf("%v, ", *currentNode.value))
	}

	builder.WriteString(fmt.Sprintf("%v]", *this.end.value))
	return builder.String()
}
