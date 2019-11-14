package main

import "fmt"

type I父 interface {
	Func父()
}
type I子 interface {
	I父
	Func子()
}

type Subject struct {
}

func (s Subject) Func父() {
	fmt.Println("Func父()")
}

func (s Subject) Func子() {
	fmt.Println("Func子()")
}

func main() {

	//由多入少
	var obj子 I子 = &Subject{}
	obj子.Func父()

	var obj父 I父 = obj子
	obj父.Func父()
}
