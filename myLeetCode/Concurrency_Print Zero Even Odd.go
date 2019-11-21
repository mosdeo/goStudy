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
	streamBreakFlag  chan interface{}
	breakFlag        bool
}

func (this *ZeroEvenOdd) SetWaitGroup(wg *sync.WaitGroup) {
	this.wg = wg
}

func (this *ZeroEvenOdd) Zero(printNumber func(int)) {
	defer this.wg.Done()

	for i := 0; i < this.n; i++ {
		select {
		case <-this.streamOddToZero:
			printNumber(0)
			this.streamZeroToEven <- nil
		case <-this.streamEvenToZero:
			printNumber(0)
			this.streamZeroToOdd <- nil
		case <-this.streamBreakFlag:
			fmt.Println("Zero() Done")
			goto RETURN
		default:
			i--
			// 	if this.breakFlag {
			// 		break
			// 	}
		}
	}
RETURN:
}

func (this *ZeroEvenOdd) Even(printNumber func(int)) {
	// defer func() { this.breakFlag = true }() // 通知其他兩個 goroutine 結束

	for i := 1; i <= this.n; i += 2 {
		<-this.streamZeroToEven
		if this.breakFlag {
			break
		}
		printNumber(i)
		this.streamEvenToZero <- nil
	}

	this.streamBreakFlag <- nil
	this.streamBreakFlag <- nil
	fmt.Println("Even() Done")
	this.wg.Done()
}

func (this *ZeroEvenOdd) Odd(printNumber func(int)) {
	// defer func() { this.breakFlag = true }() // 通知其他兩個 goroutine 結束

	for i := 2; i <= this.n; i += 2 {
		<-this.streamZeroToOdd
		if this.breakFlag {
			break
		}
		printNumber(i)
		this.streamOddToZero <- nil
	}

	this.streamBreakFlag <- nil
	this.streamBreakFlag <- nil
	fmt.Println("Odd() Done")
	this.wg.Done()
}

func PrintNumber(x int) {
	fmt.Printf("%d", x)
}

func main() {
	testNum, _ := strconv.Atoi(os.Args[1])
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
	go func() {
		zeo.streamOddToZero <- nil
	}()
	go zeo.Zero(PrintNumber)
	go zeo.Even(PrintNumber)
	go zeo.Odd(PrintNumber)
	wg.Wait()
	fmt.Println()
}
