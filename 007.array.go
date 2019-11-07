package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Array
	var intArray [9]int
	intArray[0] = 55
	intArray[1] = 66
	fmt.Println(intArray[0])
	fmt.Println(intArray[1])
	fmt.Println(intArray[2])
	fmt.Println("len(intArray)=", len(intArray))

	// Slice
	arr1 := []int{0, 1, 2}
	fmt.Println(arr1)      // [0,1,2]
	fmt.Println(arr1[0:2]) // [0,1]
	fmt.Println(arr1[2:])  // [1,2]
	//fmt.Println(arr[0:5]) // exit status 2

	// 不定長度的另一種宣告法
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("len(arr2)=", len(arr2))

	// 只取陣列的一個元素，是什麼型別？有兩種答案
	fmt.Println(arr2[3:4])
	fmt.Println(reflect.TypeOf(arr2[3:4]))
	fmt.Println(arr2[3])
	fmt.Println(reflect.TypeOf(arr2[3]))

	// 多維陣列
	fmt.Println("============ 多維陣列 ============")

	// intitial style 1
	var twoDimArray [][]int = [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}
	fmt.Println("initial style 1:", twoDimArray)

	// intitial style 2
	twoDimArray_ := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}
	fmt.Println("initial style 2:", twoDimArray_)

	// 走訪陣列
	fmt.Println("Basic style:")
	for i := 0; i < len(twoDimArray); i++ {
		fmt.Println(twoDimArray[i])
	}

	fmt.Println("Range style without index:")
	for _, element := range twoDimArray {
		fmt.Println(element)
	}

	fmt.Println("Range style within index:")
	for index, element := range twoDimArray {
		fmt.Println(index, element)
	}

}
