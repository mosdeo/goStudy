package main

import (
	"fmt"
	"math/rand"
	"time"
)

const mostDelayMillisecond int = 99

func Print0(repeat int) {
	for i := 0; i < repeat; i++ {
		time.Sleep(time.Duration(TrueRandom(mostDelayMillisecond)) * time.Millisecond)
		fmt.Printf("0")
	}
	fmt.Println("End of Print0()")
}

func Print1(repeat int) {
	for i := 0; i < repeat; i++ {
		time.Sleep(time.Duration(TrueRandom(mostDelayMillisecond)) * time.Millisecond)
		fmt.Printf("1")
	}
	fmt.Println("End of Print1()")
}

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func main() {
	go Print0(99)
	go Print1(99)

	time.Sleep(10 * time.Second)
	// main() 一結束，所有的 goroutine 就會結束。
	// 所以這個延遲時間
	// 設定太短會來不及讓 goroutine 做完，
	// 設定太長會浪費等待時間。
	// 那要怎樣才能讓 main() 知道 goroutine 結束的準確時間？
}
