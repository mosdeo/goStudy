package main

import (
	"fmt"
	"math/rand"
	"time"
)

const mostDelayMillisecond int = 1

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func P(numsCHx chan int, nameOfP string) {
	defer close(numsCHx)

	for i := 0; i < 6; i++ {
		time.Sleep(time.Duration(TrueRandom(mostDelayMillisecond)) * time.Millisecond)
		numsCHx <- i
		fmt.Printf("%s生產了:%d\n", nameOfP, i)
	}
	fmt.Println("End of", nameOfP)
}

func MultiChannelConsumerWithoutSelect(numsCHs ...chan int) {
	receiverFalseCount := 0
	defer fmt.Printf("End of C(), receiverFalseCount=%d\n", receiverFalseCount)

	for {
		for _, numsCH := range numsCHs {
			n, ok := <-numsCH
			if ok {
				fmt.Println("消費了:", n)
			} else {
				receiverFalseCount++
				// fmt.Println("Channel is closed and empty.")
			}
		}
	}
}

func MultiChannelConsumerWithinSelect(numsCH0 chan int, numsCH1 chan int) {
	receiverToNum0, receiverToNum1, defaultCount := 0, 0, 0

	defer fmt.Printf("End of C(), receiverToNum0=%d, receiverToNum1=%d, defaultCount=%d\n",
		receiverToNum0, receiverToNum1, defaultCount)

	for {
		select {
		case num0 := <-numsCH0:
			receiverToNum0++
			fmt.Printf("numsCH0消費了:%d, 未消費數量%d\n", num0, len(numsCH0)+len(numsCH1))
		case num1 := <-numsCH1:
			receiverToNum1++
			fmt.Printf("numsCH1消費了:%d, 未消費數量%d\n", num1, len(numsCH0)+len(numsCH1))
		default:
			defaultCount++
			fmt.Println("default")
		}
	}
	fmt.Println("End of C()")
}

func main() {

	numsChannel0 := make(chan int, 3)
	numsChannel1 := make(chan int, 3)
	fmt.Println("len(numsChannel0)+len(numsChannel1)=", len(numsChannel0)+len(numsChannel1))

	go P(numsChannel0, "P0")
	go P(numsChannel1, "P1")
	time.Sleep(1 * time.Second)
	fmt.Println("len(numsChannel0)+len(numsChannel1)=", len(numsChannel0)+len(numsChannel1))
	time.Sleep(1 * time.Second)
	go MultiChannelConsumerWithoutSelect(numsChannel0, numsChannel1)
	// go MultiChannelConsumerWithinSelect(numsChannel0, numsChannel1)

	// time.Sleep(2 * time.Second)
	// Polling 檢查未被消費的元素數量
	for {
		time.Sleep(50 * time.Millisecond)
		if 0 == (len(numsChannel0) + len(numsChannel1)) {
			break
		}
	}
}
