package main

import (
	"fmt"
	"strings"
)

//https://golang.org/ref/spec#Method_declarations

func add(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func mul(a float64, b float64) float64 {
	return a * b
}

func div(a float64, b float64) float64 {
	return a / b
}

func mulitReturn() (string, string) {
	return "return1", "return2"
}

func main() {
	a, b := 7.0, 8.0
	fmt.Printf("a+b=%f\n", add(a, b))
	fmt.Printf("a-b=%f\n", sub(a, b))
	fmt.Printf("a*b=%f\n", mul(a, b))
	fmt.Printf("a/b=%f\n", div(a, b))

	addCopy := add
	fmt.Printf("addCopy(a,b)=%f\n", addCopy(a, b))

	//multi return
	var arrStr [2]string
	arrStr[0], arrStr[1] = mulitReturn()
	arrStr0, arrStr1 := mulitReturn()
	s1 := arrStr[0:1]
	fmt.Println(strings.Join(arrStr, ","))

}
