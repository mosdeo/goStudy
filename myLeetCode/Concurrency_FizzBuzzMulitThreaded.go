package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type FizzBuzz struct {
	n           int
	wg          *sync.WaitGroup
	streamBaton chan int
}

func (this *FizzBuzz) PrintLoop(passCondition func(int) bool, printString func(int)) {
	defer this.wg.Done()

	for i := 0; i <= this.n; i++ {
		if passCondition(i) {
			nextNum := <-this.streamBaton //接棒
			if i == nextNum {
				printString(i)
				this.streamBaton <- i + 1 //交棒
			} else {
				this.streamBaton <- nextNum //把數字還回去
				i--
			}
			runtime.Gosched()
		}
	}
}

func (this *FizzBuzz) PrintFizz() {
	PassCondition := func(i int) bool { return (0 == i%3) && (0 != i%5) }
	PrintString := func(i int) { fmt.Printf("Fizz(%d), ", i) }

	this.PrintLoop(PassCondition, PrintString)
}

func (this *FizzBuzz) PrintBuzz() {
	PassCondition := func(i int) bool { return (0 != i%3) && (0 == i%5) }
	PrintString := func(i int) { fmt.Printf("Buzz(%d), ", i) }

	this.PrintLoop(PassCondition, PrintString)
}

func (this *FizzBuzz) PrintFizzBuzz() {
	PassCondition := func(i int) bool { return 0 == i%(3*5) }
	PrintString := func(i int) { fmt.Printf("FizzBuzz(%d), ", i) }

	this.PrintLoop(PassCondition, PrintString)
}

func (this *FizzBuzz) PrintNumber() {
	PassCondition := func(i int) bool { return (0 != i%3) && (0 != i%5) }
	PrintString := func(i int) { fmt.Printf("%d, ", i) }

	this.PrintLoop(PassCondition, PrintString)
}

func main() {
	start := time.Now()

	for testCase := 0; testCase <= 20; testCase++ {

		fizzbuzz := &FizzBuzz{
			n:           testCase,
			wg:          &sync.WaitGroup{},
			streamBaton: make(chan int, 1),
		}

		fizzbuzz.wg.Add(4)
		go fizzbuzz.PrintFizz()
		go fizzbuzz.PrintBuzz()
		go fizzbuzz.PrintFizzBuzz()
		go fizzbuzz.PrintNumber()

		fizzbuzz.streamBaton <- 0 //啟動交棒
		fizzbuzz.wg.Wait()
		close(fizzbuzz.streamBaton)
		fmt.Println() //這個 Test Case 結束了，換行。
	}

	spentTime := time.Now().Sub(start)
	fmt.Println("Spent time:", spentTime)
}
