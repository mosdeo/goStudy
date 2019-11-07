package main

import (
	"fmt"
	"reflect"
	"strings"
)

//https://golang.org/ref/spec#Method_declarations

//官方建議寫法：回傳值要給名稱
func Add(a float64, b float64) (sum float64) {
	sum = a + b
	return sum
}

func Sub(a float64, b float64) float64 {
	return a - b
}

// MulitReturn...
func MulitReturn() (string, string) {
	return "return1", "return2"
}

func main() {
	a, b := 7.0, 8.0
	fmt.Printf("a+b=%f\n", Add(a, b))
	fmt.Printf("a-b=%f\n", Sub(a, b))

	//multi return
	var arrStr [2]string
	arrStr[0], arrStr[1] = MulitReturn()
	//arrStr0, arrStr1 := MulitReturn()
	slice_form_arrStr := arrStr[0:]
	fmt.Println("Multi return Joined:", strings.Join(slice_form_arrStr, ","))

	//可變參數用法
	fmt.Println("Sum 1~3=", Sum(1, 2, 3))
	fmt.Println("Sum 1~6=", Sum(1, 2, 3, 4, 5, 6))
	fmt.Println("Sum 1~9=", Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sum 1~9=", Sum(numbers...))

	//指向函式的參考
	AddCopy := Add
	fmt.Printf("addCopy(a,b)=%f\n", AddCopy(a, b))

	fmt.Println(reflect.TypeOf(Sum))
	var GetSum func(...int) int //空函式參考
	fmt.Println(GetSum)
	GetSum = Sum
	fmt.Println(GetSum(1, 2, 3, 4))

	//自定義新的函式型別、別名
	// type [定義型別] [底層型別]
	type TwoStrFunc func(string, string) string          //這是定義新的型別
	type AliasOfTwoStrFunc = func(string, string) string //這是給底層型別取別名
	var LoaclJoinTwoStr TwoStrFunc = CombineTwoStr
	var LoaclJoinTwoStrToo AliasOfTwoStrFunc = CombineTwoStr
	fmt.Println(LoaclJoinTwoStr("go", "lang"))
	fmt.Println("TypeOf(LoaclJoinTwoStr   )=", reflect.TypeOf(LoaclJoinTwoStr))
	fmt.Println("TypeOf(LoaclJoinTwoStrToo)=", reflect.TypeOf(LoaclJoinTwoStrToo))
}

// 可變參數函式寫法
func Sum(inputs ...int) (sum int) {
	for _, input := range inputs {
		sum += input
	}
	return sum
}

func CombineTwoStr(strA string, strB string) (strAB string) {
	strAB = strA + strB
	return strAB
}
