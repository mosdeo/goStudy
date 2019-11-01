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
	var(//型別寫在變數之後
		a int = 10
		b string = "BBB" 
		c bool = true
	)

	fmt.Println("基本變數型別")
	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(x, reflect.TypeOf(x))
	fmt.Println(y, reflect.TypeOf(y))
	fmt.Println(z, reflect.TypeOf(z))
	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	
	// 進階變數宣告
	var(
		d = 9.9
		e = "EEE"
		f = false
	)

	fmt.Println("自動推斷型別")
	fmt.Println(d, reflect.TypeOf(d))
	fmt.Println(e, reflect.TypeOf(e))
	fmt.Println(f, reflect.TypeOf(f))


}

