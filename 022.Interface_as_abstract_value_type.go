// https://golang.org/ref/spec#Type_assertions

package main

import (
	"fmt"
)

func DemoTypeAssertion(unknownTypeV interface{}) {
	//a, b := unknownTypeV.(type) //use of .(type) outside type switch
	switch unknownTypeV.(type) {
	case int:
		fmt.Printf("Type is int, value = %d\n", unknownTypeV.(int))
	case string:
		fmt.Printf("Type is string, value = %s\n", unknownTypeV.(string))
	default:
		fmt.Println("Type not found, value = ", unknownTypeV)
	}
}

func main() {
	var a interface{}
	var b interface{}
	var c interface{}
	a = 77
	b = 2.33
	c = "Oh!"
	fmt.Println(a, b, c)

	DemoTypeAssertion(5566)
	DemoTypeAssertion("得第一")
	DemoTypeAssertion([]int{0: 87})
}
