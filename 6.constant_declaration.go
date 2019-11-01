package main

import(
	"fmt"
)

func main(){
	const(
		nine = 9.9/1.1;
		const_str string = "read only string"
		falseConst bool = false;
	)

	// const 宣告了沒用到，也可以過編譯
	fmt.Println(nine)
	fmt.Println(const_str)

	//itoa列舉
	//應該是類似 enum 的語法吧
	const(
		c0 = iota
		c1
		c2
	)
	fmt.Println( c0, c1, c2) //0 1 2

	const(
		x0 = 9
		x1
		x2
	)
	fmt.Println( x0, x1, x2) // 9 9 9 
}