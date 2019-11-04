package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [10]int{}
	// func make([]T, len, cap) []T
	slice := make([]int, 10)

	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(arr))   //[10]int 比較沒彈性
	fmt.Println(reflect.TypeOf(slice)) //[]int 比較有彈性

	// Slice copy
	// 這邊還不能理解，很玄
	s1 := []int{0, 1, 2, 3, 4}
	s2 := reflect.ValueOf(s1)

	fmt.Println(s1)
	fmt.Println(s2)                 //與s1相同
	fmt.Println(reflect.TypeOf(s1)) //[]int
	fmt.Println(reflect.TypeOf(s2)) //relect.Value

	// 還沒找到 runtime 才決定型別的 slice 產生方法
	// 不能以 TypeOf 的 return 當作型別用
	s3 := make(reflect.TypeOf(s1), len(s1)) //會錯
	// 目前找到的答案看起來很複雜
	// https://stackoverflow.com/questions/39363798/how-to-create-a-slice-of-variable-type-in-go
}
