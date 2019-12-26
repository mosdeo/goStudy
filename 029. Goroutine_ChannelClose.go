package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var streamSync = make(chan int, 3)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			<-time.After(time.Duration(50) * time.Millisecond)
			n, ok := <-streamSync
			fmt.Println(n, ok)

			if !ok {
				// channel 被關閉，數字消費完畢，就結束
				break
			}
		}
	}()

	go func() {
		for i := 0; i <= 5; i++ {
			<-time.After(time.Duration(50) * time.Millisecond)
			streamSync <- i
		}
		streamSync <- -1
		close(streamSync) // 產生數字發送完畢，關閉 channel
	}()

	wg.Wait()
}
