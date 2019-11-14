package main

import "fmt"

type IFuncAB interface {
	FuncA()
	FuncB()
}
type IFuncA interface {
	FuncA()
}

type Subject struct {
}

func (s Subject) FuncA() {
	fmt.Println("FuncA()")
}

func (s Subject) FuncB() {
	fmt.Println("FuncB()")
}

func main() {

	//實作某個介面不需要明確宣告
	//介面的轉換只看內涵，不看名稱
	//該介面具有的方法都具備就可以

	//由多入少
	var objAB IFuncAB = &Subject{}
	objAB.FuncA()
	var objA IFuncA = objAB
	objA.FuncA()

	//由少入多
	// var objA IFuncA = &Subject{}
	// objA.FuncA()
	// var objAB IFuncAB = objA
	// objAB.FuncA()
}
