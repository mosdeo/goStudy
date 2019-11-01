package main

import(
	"fmt"
	"os"
)

func main(){
	var s, sep string //型別“”string"寫在後面

	// for 是 golang 唯一的迴圈語法
	// for 三個元件都沒有括號
	// for 的 scope 大括號是必要的，而且開頭要與最後一個元件在同一行
	for i := 1; i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
