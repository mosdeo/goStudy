// LeetCode in Coucurrency：Print Zero Even Odd
// Solution by unbuffered chan, without time delay.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type ZeroEvenOdd struct {
	n                int
	wg               *sync.WaitGroup
	streamEvenToZero chan interface{}
	streamOddToZero  chan interface{}
	streamZeroToEven chan interface{}
	streamZeroToOdd  chan interface{}
}

func (this *ZeroEvenOdd) SetWaitGroup(wg *sync.WaitGroup) {
	this.wg = wg
}

func (this *ZeroEvenOdd) Zero(printNumber func(int)) {
	defer this.wg.Done()
	//defer fmt.Println("Zero() Done")

	for i := 0; i < this.n; i++ {
		select {
		case <-this.streamOddToZero:
			printNumber(0)
			this.streamZeroToEven <- nil
		case <-this.streamEvenToZero:
			printNumber(0)
			this.streamZeroToOdd <- nil
		default:
			runtime.Gosched()
			i--
		}
	}

	if 0 == this.n%2 {
		<-this.streamEvenToZero //等待 Even() 結束，自己再結束
	} else {
		<-this.streamOddToZero //等待 Odd() 結束，自己再結束
	}
}

func (this *ZeroEvenOdd) Even(printNumber func(int)) {
	defer this.wg.Done()

	evenUpper := this.n - this.n%2
	// fmt.Println("evenUpper:", evenUpper)
	for i := 2; i <= evenUpper; {
		<-this.streamZeroToEven
		printNumber(i)
		i += 2
		this.streamEvenToZero <- nil
	}
}

func (this *ZeroEvenOdd) Odd(printNumber func(int)) {
	defer this.wg.Done()

	oddUpper := ((this.n + 1) - (this.n+1)%2) - 1
	// fmt.Println("oddUpper:", oddUpper)
	for i := 1; i <= oddUpper; i += 2 {
		<-this.streamZeroToOdd
		printNumber(i)
		this.streamOddToZero <- nil
	}
}

func PrintNumber(x int) {
	fmt.Printf("%d", x)
}

func main() {
	testCases := []int{0, 1, 2, 3, 7, 10, 11, 13, 14}

	var PrintZeroEvenOdd = func(testNum int) {
		var zeo = &ZeroEvenOdd{
			n:                testNum,
			streamEvenToZero: make(chan interface{}),
			streamOddToZero:  make(chan interface{}),
			streamZeroToEven: make(chan interface{}),
			streamZeroToOdd:  make(chan interface{}),
		}

		//設定同步
		wg := &sync.WaitGroup{}
		zeo.SetWaitGroup(wg)

		wg.Add(3)
		//go func() { zeo.streamEvenToZero <- nil }() //給起頭的火種
		go zeo.Zero(PrintNumber)
		go zeo.Even(PrintNumber)
		go zeo.Odd(PrintNumber)
		<-time.After(time.Microsecond)
		zeo.streamEvenToZero <- nil
		wg.Wait()
		fmt.Println()
	}

	for _, testNum := range testCases {
		fmt.Printf("Case %2d: ", testNum)
		PrintZeroEvenOdd(testNum)
	}
}
