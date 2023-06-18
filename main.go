package main

import (
	"fmt"
)

type interf interface {
	Method1() interf
	Show()
}

type struct1 struct {
	str string
	p1  *struct2
}

type struct2 struct {
	str string
	p2  *struct1
}

func (s *struct1) Show() {
	fmt.Println(s.p1.str)
}

func (s *struct2) Show() {
	fmt.Println(s.p2.str)
}

func (s *struct1) Method1() interf {
	return s.p1
}

func (s *struct2) Method1() interf {
	return s.p2
}

func f(i interf) {
	i.Show()
	i = i.Method1()
	i.Show()
	i = i.Method1()
	i.Show()
}

func main() {
	var d1 struct1 = struct1{str: "struct1"}
	var d2 struct2 = struct2{str: "struct2"}

	d1.p1 = &d2
	d2.p2 = &d1

	// var i interf = &d1
	// d11 := struct1{str: "struct11"}
	// d11 = i

	f(&d1)
}
