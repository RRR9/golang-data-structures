package forwardList

import (
	"algorithms/containers"
)

type ForwardListOperations[T any] interface {
	Add(...T)
	AccessToEachElem(func(T))
	Delete(int)
	Front() *T
	Back() *T
	GetIter() iterator[T]
	// Insert(T, int)

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
	ptrCurrent *nodeForwardList[T]
	ptrBegin   *nodeForwardList[T]
	ptrEnd     *nodeForwardList[T]
}

func (list iterator[T]) Begin() containers.Iterator[T] {
	return nil
}

func (list iterator[T]) End() containers.Iterator[T] {
	return nil
}

func (list *iterator[T]) GetVal() *T {
	return nil
}

func (list iterator[T]) Next() containers.Iterator[T] {

	return nil
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
	if ind > list.size-containers.ContainerSize(1) {
		return
	}

	if ind == 0 {
		list.head = list.head.next
		return
	}

	var cnt containers.ContainerSize = 0
	prevNode := list.head
	for ; cnt < ind-containers.ContainerSize(1); prevNode = prevNode.next {
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
