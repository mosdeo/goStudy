package main

import (
	"fmt"
	"reflect"
)

// 匿名函式的型別要先有
type ConditionOfInt = func(int) bool

func Filter(Decider ConditionOfInt, numbers []int) (result []int) {
	for _, number := range numbers {
		if Decider(number) {
			result = append(result, number)
		}
	}

	return result
}

func main() {
	var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//使用的同時，才宣告過濾規則函式
	fmt.Println(Filter(func(elem int) bool { return elem > 5 }, array))
	fmt.Println(Filter(func(elem int) bool { return elem < 5 }, array))

	fmt.Println("TypeOf(Foo_I_have_an_AnonymousFunction)=", reflect.TypeOf(Foo_I_have_an_AnonymousFunction))
	Foo_I_have_an_AnonymousFunction()
}

// func foo() {}

func Foo_I_have_an_AnonymousFunction() {
	var func_ref func() = func() {
		fmt.Println("This is func_ref()")
	}

	func_ref()
}
