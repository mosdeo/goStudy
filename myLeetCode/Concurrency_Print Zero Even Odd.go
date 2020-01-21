// LeetCode in Coucurrency：Print Zero Even Odd
// Solution by unbuffered chan, without time delay.

package main

import (
	"fmt"
	"runtime"
)

type ZeroEvenOdd struct {
	n                int
	streamEvenToZero chan interface{}
	streamOddToZero  chan interface{}
	streamZeroToEven chan interface{}
	streamZeroToOdd  chan interface{}
	streamZeroToEnd  chan interface{}
}

func (this *ZeroEvenOdd) Zero(printNumber func(int)) {
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
			//<-time.After(time.Microsecond)
			i--
		}
	}

	if 0 == this.n%2 {
		<-this.streamEvenToZero //等待 Even() 結束，自己再結束
	} else {
		<-this.streamOddToZero //等待 Odd() 結束，自己再結束
	}

	this.streamZeroToEnd <- nil
}

func (this *ZeroEvenOdd) Even(printNumber func(int)) {
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
	oddUpper := ((this.n + 1) - (this.n+1)%2) - 1
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
	var PrintZeroEvenOdd = func(testNum int) {
		var zeo = &ZeroEvenOdd{
			n:                testNum,
			streamEvenToZero: make(chan interface{}),
			streamOddToZero:  make(chan interface{}),
			streamZeroToEven: make(chan interface{}),
			streamZeroToOdd:  make(chan interface{}),
			streamZeroToEnd:  make(chan interface{}),
		}

		go func() { zeo.streamEvenToZero <- nil }() //給起頭的火種
		go zeo.Zero(PrintNumber)
		go zeo.Even(PrintNumber)
		go zeo.Odd(PrintNumber)
		<-zeo.streamZeroToEnd //等待 Zero() 送出結束訊號
		fmt.Println()
	}

	for testNum := range [14]int{} {
		fmt.Printf("Case %2d: ", testNum)
		PrintZeroEvenOdd(testNum)
	}
}
