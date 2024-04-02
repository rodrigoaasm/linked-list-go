package structs

import (
	"fmt"
)

type Pred[T comparable] func(content T) bool

type LinkedList[T comparable] struct {
	head *Node[T]
	size uint32
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	list := new(LinkedList[T])
	list.head = nil
	list.size = 0

	return list
}

func (list *LinkedList[T]) findNode(value T) (*Node[T], *Node[T], bool) {
	var prevNode *Node[T]
	currentNode := list.head
	for currentNode != nil {
		if currentNode.content == value {
			return prevNode, currentNode, true
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
	return nil, nil, false
}

func (list *LinkedList[T]) Push(value T) error {
	if list.head == nil {
		list.head = NewNode(value)
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = NewNode(value)
	}
	list.size += 1

	return nil
}

func (list *LinkedList[T]) Remove(value T) bool {
	prev, current, ok := list.findNode(value)
	if !ok {
		return ok
	}

	if prev == nil {
		list.head = current.next
	} else {
		prev.next = current.next
	}
	return true
}

func (list *LinkedList[T]) FindEqual(value T) (*T, bool) {
	if _, node, ok := list.findNode(value); ok {
		return &node.content, true
	} else {
		return nil, false
	}
}

func (list *LinkedList[T]) Find(fn Pred[T]) (*T, bool) {
	currentNode := list.head
	for currentNode != nil {
		if fn(currentNode.content) {
			return &currentNode.content, true
		}
		currentNode = currentNode.next
	}
	return nil, false
}

func (list *LinkedList[T]) Filter(fn Pred[T]) *LinkedList[T] {
	newList := NewLinkedList[T]()
	currentNode := list.head
	for currentNode != nil {
		if fn(currentNode.content) {
			newList.Push(currentNode.content)
		}
		currentNode = currentNode.next
	}
	return newList
}

func (list *LinkedList[T]) Map(fn func(content T) interface{}) *LinkedList[interface{}] {
	newList := NewLinkedList[interface{}]()
	currentNode := list.head
	for currentNode != nil {
		newList.Push(fn(currentNode.content))
		currentNode = currentNode.next
	}
	return newList
}

func (list *LinkedList[T]) Sort(less func(contentA T, contentB T) bool, parallel bool) *LinkedList[T] {
	if list.size < 200 {
		list.head = insertionSort(list.head, less)
	} else if parallel {
		ch := makeMergeSortWorker(list.head, list.size, less, true)
		list.head = <-ch
	} else {
		list.head = mergeSort(list.head, list.size, less, true)
	}
	return list
}

func (list *LinkedList[T]) Size() uint32 {
	return list.size
}

func (list *LinkedList[T]) Print() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Print(currentNode.content)
		fmt.Print(" -> ")
		currentNode = currentNode.next
	}
	fmt.Println("NIL")
}
