package forwardlist

import (
	"algorithms/containers"
)

type ForwardListOperations[T any] interface {
	Add(...T)
	Pop()
	AccessToEachElem(func(T))
	Delete(int)
	Front() *T
	Back() *T
	Begin() iterator[T]
	End() iterator[T]
	// Sort()
	Insert(int, ...T)

	containers.Container[T]
	containers.Iterator[T]
}

type forwardList[T any] struct {
	size containers.ContainerSize
	head *nodeForwardList[T]
	tail *nodeForwardList[T]
}

type nodeForwardList[T any] struct {
	data T
	next *nodeForwardList[T]
}

type iterator[T any] struct {
	ptr *nodeForwardList[T]
}

func (list *forwardList[T]) Begin() iterator[T] {
	return iterator[T]{ptr: list.head}
}

func (list *forwardList[T]) End() iterator[T] {
	return iterator[T]{ptr: list.tail}
}

func (iter *iterator[T]) GetVal() *T {
	return &iter.ptr.data
}

func (iter *iterator[T]) Next() containers.Iterator[T] {
	if iter.ptr == nil {
		return nil
	}
	iter.ptr = iter.ptr.next
	return iter
}

func (list *forwardList[T]) Front() *T {
	if list.head != nil {
		return &list.head.data
	}
	return nil
}

func (list *forwardList[T]) Back() *T {
	if list.tail != nil {
		return &list.tail.data
	}
	return nil
}

func (list *forwardList[T]) Size() containers.ContainerSize {
	return list.size
}

func NewForwardList[T any]() *forwardList[T] {
	return &(forwardList[T]{})
}

func (list *forwardList[T]) AccessToEachElem(handler func(*T)) {
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		handler(&currentNode.data)
	}
}

func (list *forwardList[T]) Add(val ...T) {
	for i := 0; i < len(val); i++ {
		list.size++
		newNode := &(nodeForwardList[T]{data: val[i]})

		if list.head == nil {
			list.head = newNode
			list.tail = newNode
			continue
		}

		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *forwardList[T]) Delete(ind containers.ContainerSize) {
	if ind >= list.size {
		panic("delete index out of range")
	}

	if ind == 0 {
		list.head = list.head.next
		return
	}

	var cnt containers.ContainerSize = 0
	prevNode := list.head
	for ; cnt < ind; prevNode = prevNode.next {
		cnt++
	}

	if ind == list.size-1 {
		list.tail = prevNode
		prevNode.next = nil
		return
	}

	currNode := prevNode.next
	prevNode.next = currNode.next
}

func (list *forwardList[T]) Contains(val T, comp func(*T, *T) bool) bool {
	for currNode := list.head; currNode != nil; currNode = currNode.next {
		if comp(&val, &currNode.data) {
			return true
		}
	}
	return false
}

func (list *forwardList[T]) Insert(ind containers.ContainerSize, val ...T) {
	if ind >= list.size {
		panic("insert index out of range")
	}

	var tempHead, tempTail *nodeForwardList[T]

	for i := 0; i < len(val); i++ {
		tempNode := &nodeForwardList[T]{data: val[i]}

		if tempHead == nil {
			tempHead = tempNode
			tempTail = tempNode
			continue
		}

		tempTail.next = tempNode
		tempTail = tempNode
	}

	if ind == 0 {
		tempTail.next = list.head
		list.head = tempHead
		return
	}

	currNode := list.head
	for cnt := containers.ContainerSize(0); cnt < ind-containers.ContainerSize(1); cnt++ {
		currNode = currNode.next
	}

	nextNode := currNode.next
	currNode.next = tempHead
	tempTail.next = nextNode
}

func (list *forwardList[T]) Pop() {
	if list.head == nil {
		return
	}

	list.head = list.head.next
}
