package main

import (
	"fmt"
	"time"
)

func main() {

	//預約好拋出錯誤訊息
	var PutRecoverMsg = func() {
		fmt.Printf("recover(): %s\n", recover())
	}
	defer PutRecoverMsg()

	// 1. defer 延遲並逆序執行
	defer fmt.Println("This statment has been defer at 1st")
	defer fmt.Println("This statment has been defer at 2nd")
	for i := 0; i < 7; i++ {
		defer fmt.Printf("This statment has been defer in loop, i=%d\n", i)
	}
	fmt.Println("This statment after defer")

	// 2. 故意製造當機（兩種可能隨機擇一）
	if 0 == time.Now().Second()&1 {
		panic("Randomly selceted run panic()")
	} else {
		var zero = 1 - 1
		fmt.Println(zero)
		var impossible = 1 / zero
		fmt.Println(impossible)
	}

	// 這裡會發生「panic: runtime error: integer divide by zero」，
	// 並將所有已經被 defer 的敘述執行，展示 golang 對應例外處理的方式，
	// 最常用到的場合，是在讀檔錯誤時還能關閉檔案善後。
}
