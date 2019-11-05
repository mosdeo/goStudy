package main

import (
	"fmt"
)

// Go 官方文件有一段話，是我看到目前為止最清楚描述 array 與 slice 關係的：
// A slice, once initialized, is always associated with an underlying array that holds its elements.
// A slice therefore shares storage with its array and with other slices of the same array;
// by contrast, distinct arrays always represent distinct storage.
// 所以這裡就來驗證一下。

func main() {
	var onlyOneArray = [12]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	var s0to3 = onlyOneArray[:4]
	var s4to8 = onlyOneArray[4:9]
	var s9to12 = onlyOneArray[9:]

	// 先讓三個 Slice 都 mapping 到同一個 array 上面
	fmt.Println("onlyOneArray=", onlyOneArray)
	fmt.Println("s0to3=", s0to3)
	fmt.Println("s4to8=", s4to8)
	fmt.Println("s9to12=", s9to12)
	fmt.Println()

	// 然後修改 array，觀察每一個 slice 的變化
	// onlyOneArray = sort.Reverse(onlyOneArray) //不能用 array 沒有實作需要的 interface
	var tempArray [len(onlyOneArray)]int = onlyOneArray
	for i := 0; i < len(onlyOneArray); i++ {
		onlyOneArray[i] = tempArray[len(onlyOneArray)-1-i]
	}
	fmt.Println("Srouce array was reversed.")
	fmt.Println()

	// 看結果
	fmt.Println("onlyOneArray=", onlyOneArray)
	fmt.Println("s0to3=", s0to3)
	fmt.Println("s4to8=", s4to8)
	fmt.Println("s9to12=", s9to12)

	// 結論：Slice 的確是 Array 的片段映射
}
