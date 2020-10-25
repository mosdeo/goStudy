package main

import (
	"fmt"
)

type DemoBuiltInInterfaceType struct{}

// func (dbit DemoBuiltInInterfaceType) DivideByZero() {
// 	var zero = 0
// 	var impossible = 1 / zero
// 	fmt.Println(impossible)
// }

func (dbit DemoBuiltInInterfaceType) Error() string {
	return "The function implement built-in interface Error()"
}

func main() {
	var dbitAsError error = &DemoBuiltInInterfaceType{}
	fmt.Println(dbitAsError.Error())
}
