// ref:https://openhome.cc/Gossip/Go/VariableConstantDeclaration.html
package main

import(
	"fmt"
	"reflect"
)

func main(){

	// 基本變數宣告
	var i int = 10
	var x, y, z int = 1, 2, 3;
	var(
		a = 10
		b = "BBB"
		c = true
	)

	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(x, reflect.TypeOf(x))
	fmt.Println(y, reflect.TypeOf(y))
	fmt.Println(z, reflect.TypeOf(z))
	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	
	// 進階變數宣告
}