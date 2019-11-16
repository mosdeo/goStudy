package main

import (
	"fmt"
)

func main() {
	DemoSelectIsRandomToEashCase()
}

func DemoSelectIsRandomToEashCase() {
	//優先順序
	// Closed Channel > default > Open Channel
	// c1, c2 close 之前，會先跑 default
	// c1, c2 close 之後，剩下就只跑 c1, c2 次數會隨機平分。
	// 書上有一句描述很傳神：「Go 語言 runtime 無法解析 select 的意圖」
	var CloseChannels = func(channels ...chan interface{}) {
		for _, channel := range channels {
			close(channel)
		}
	}

	c1 := make(chan interface{})
	c2 := make(chan interface{})
	// CloseChannels(c2)
	var c1Count, c2Count, defaultCount int
	//defer fmt.Printf("c1Count:%d, c2Count:%d, defaultCount:%d\n", c1Count, c2Count, defaultCount)

	fmt.Println("Into for loop")
	for i := 1; i < 1000; i++ {
		select {
		case <-c1:
			c1Count++
			// if 600 == i {
			// 	// CloseChannels(c1, c2)
			// }
		case <-c2:
			c2Count++
			// if 600 == i {
			// 	// CloseChannels(c1, c2)
			// }
		default:
			defaultCount++
			if 600 == i {
				CloseChannels(c1, c2)
				fmt.Printf("c1Count:%d, c2Count:%d, defaultCount:%d\n", c1Count, c2Count, defaultCount)
			}
		}
	}

	// CloseChannels(c1, c2)
	fmt.Printf("c1Count:%d, c2Count:%d, defaultCount:%d\n", c1Count, c2Count, defaultCount)
}
