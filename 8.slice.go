package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := make([]int, 5)
	s2 := s1
	fmt.Println(reflect.TypeOf(s1)) //[]int
	fmt.Println(reflect.TypeOf(s2)) //[]int

	s2[2] = 9
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Printf("&s1=%p\n", &s1)
	fmt.Printf("&s2=%p\n", &s2)
	// 由此得知，slice 是 reference

	// assign value at specific index
	s3 := []int{3: 3, 6: 6, 9: 9}
	fmt.Println("s3=", s3)
}
