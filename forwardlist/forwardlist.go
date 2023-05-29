package forwardList

import (
	"algorithms/containers"
)

type ForwardListOperations[T any] interface {
	AddMany(val ...T)
	Add(val ...T)

	containers.Container
}

type forwardList[T any] struct {
	size int
	head *nodeForwardList[T]
	tail *nodeForwardList[T]
}

type nodeForwardList[T any] struct {
	data T
	next *nodeForwardList[T]
}

func NewForwardList[T any]() *forwardList[T] {
	return &(forwardList[T]{})
}

func (list *forwardList[T]) Print(show func(val T)) {
	for currentNode := list.tail; currentNode != nil; currentNode = currentNode.next {
		show(currentNode.data)
	}
}

func (list *forwardList[T]) AddMany(val ...T) {
	for _, v := range val {
		list.size++
		newNode := &(nodeForwardList[T]{data: v})

		if list.head == nil {
			list.head = newNode
			list.tail = newNode
			continue
		}

		newNode.next = list.tail
		list.tail = newNode
	}
}

func (list *forwardList[T]) Delete(ind int) {

}

func (list *forwardList[T]) Contains(val any) {

}
