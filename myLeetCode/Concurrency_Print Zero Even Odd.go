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
	// streamBreakFlag  chan interface{}
	isZeroReturned bool
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
	// fmt.Println("Zero() Done")
	this.isZeroReturned = true
	this.wg.Done()
}

func (this *ZeroEvenOdd) Even(printNumber func(int)) {

	for i := 2; i <= this.n; i += 2 {
		<-this.streamZeroToEven
		printNumber(i)
		if !this.isZeroReturned {
			// 確保發送之前消費者還存活著
			this.streamEvenToZero <- nil
		}
	}

	// fmt.Println("Even() Done")
	this.wg.Done()
}

func (this *ZeroEvenOdd) Odd(printNumber func(int)) {
	// defer func() { this.breakFlag = true }() // 通知其他兩個 goroutine 結束

	for i := 1; i <= this.n; i += 2 {
		<-this.streamZeroToOdd
		printNumber(i)
		if !this.isZeroReturned {
			// 確保發送之前消費者還存活著
			// 這裡還是有點危險，有可能判斷的時候存活、要寫入 channel 的時候卻已經結束
			// 需要加上延遲確保流程上萬無一失
			// time.Sleep(time.Duration(100 * time.Millisecond))
			// 加上好像也沒差？
			this.streamOddToZero <- nil
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
		zeo.streamEvenToZero <- nil
	}()
	go zeo.Zero(PrintNumber)
	go zeo.Even(PrintNumber)
	go zeo.Odd(PrintNumber)
	wg.Wait()
	fmt.Println()
}
