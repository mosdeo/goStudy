package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
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
	for i := 0; i < this.n; i++ {
		select {
		case <-this.streamOddToZero:
			printNumber(0)
			this.streamZeroToEven <- nil
		case <-this.streamEvenToZero:
			printNumber(0)
			this.streamZeroToOdd <- nil
		}
	}

	// fmt.Println("Colsed streamZeroToOdd, streamZeroToOdd")
	// fmt.Println("Zero() Done")
	this.wg.Done()
}

func (this *ZeroEvenOdd) Even(printNumber func(int)) {

	for i := 2; i <= this.n; i += 2 {
		select {
		case <-this.streamZeroToEven:
			printNumber(i)
			this.streamEvenToZero <- nil //
		default:
			// fmt.Println("default")
			i -= 2
		}
	}

	// fmt.Println("Even() Done")
	this.wg.Done()
}

func (this *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= this.n; i += 2 {
		select {
		case <-this.streamZeroToOdd:
			printNumber(i)
			this.streamOddToZero <- nil
		default:
			// fmt.Println("default")
			i -= 2
		}
	}

	// fmt.Println("Odd() Done")
	this.wg.Done()
}

func PrintNumber(x int) {
	fmt.Printf("%d", x)
}

func main() {
	testNum, _ := strconv.Atoi(os.Args[1])
	var zeo = &ZeroEvenOdd{
		n:                testNum,
		streamEvenToZero: make(chan interface{}, 1),
		streamOddToZero:  make(chan interface{}, 1),
		streamZeroToEven: make(chan interface{}, 1),
		streamZeroToOdd:  make(chan interface{}, 1),
	}

	//設定同步
	wg := &sync.WaitGroup{}
	zeo.SetWaitGroup(wg)

	wg.Add(3)
	go func() { zeo.streamEvenToZero <- nil }() //給起頭的火種
	go zeo.Zero(PrintNumber)
	go zeo.Even(PrintNumber)
	go zeo.Odd(PrintNumber)
	wg.Wait()
	fmt.Println()
}
