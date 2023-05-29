package main

import (
	forwardlist "algorithms/forwardlist"
	"fmt"
)

func printListElem[T any](val T) {
	fmt.Printf("%v\n", val)
}

func main() {
	list := forwardlist.NewForwardList[int]()
	list.AddMany(1, 2)
	list.Print(printListElem[int])
}
