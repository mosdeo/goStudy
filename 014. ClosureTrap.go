package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var refToPrintInt []func() = nil
	for _, num := range numbers {

		// 寫一個匿名函式
		refToPrintInt = append(refToPrintInt, func() {
			localNum := num
			fmt.Println(localNum)
		})
		refToPrintInt[len(refToPrintInt)-1]()
	}

	fmt.Println(reflect.TypeOf(refToPrintInt))

	for _, f := range refToPrintInt {
		f()
	}

	// 多年前寫 C# 踩過的坑，我這邊再重現一遍
	// 這個例子說明閉包（Closure）會吃進外部變數的位址(address)，而非值(value)。
	// 這有點像是 LINQ 只是先建立「預約命令」，特定指令才會強制執行。
}
