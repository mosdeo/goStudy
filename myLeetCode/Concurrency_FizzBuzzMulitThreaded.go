// 交接棒不指定對象的版本

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

func (this *FizzBuzz) PrintFizz() {
	defer this.wg.Done()

	for i := 0; i <= this.n; i++ {
		if (0 == i%3) && (0 != i%5) {
			nextNum := <-this.streamBaton //接棒
			if i == nextNum {
				fmt.Printf("Fizz(%d), ", i)
				this.streamBaton <- i + 1 //交棒
			} else {
				this.streamBaton <- nextNum //把數字還回去
				i--
			}
			runtime.Gosched()
		}
	}
}

func (this *FizzBuzz) PrintBuzz() {
	defer this.wg.Done()

	for i := 0; i <= this.n; i++ {
		if (0 != i%3) && (0 == i%5) {
			nextNum := <-this.streamBaton //接棒
			if i == nextNum {
				fmt.Printf("Buzz(%d), ", i)
				this.streamBaton <- i + 1 //交棒
			} else {
				this.streamBaton <- nextNum //把數字還回去
				i--
			}
			runtime.Gosched()
		}
	}
}

func (this *FizzBuzz) PrintFizzBuzz() {
	defer this.wg.Done()

	for i := 0; i <= this.n; i++ {
		if 0 == i%(3*5) {
			nextNum := <-this.streamBaton //接棒
			if i == nextNum {
				fmt.Printf("FizzBuzz(%d), ", i)
				this.streamBaton <- i + 1 //交棒
			} else {
				this.streamBaton <- nextNum //把數字還回去
				i--
			}
			runtime.Gosched()
		}
	}
}

func (this *FizzBuzz) PrintNumber() {
	defer this.wg.Done()

	for i := 0; i <= this.n; i++ {
		if (0 != i%3) && (0 != i%5) {
			nextNum := <-this.streamBaton //接棒
			if i == nextNum {
				fmt.Printf("%d, ", i)
				this.streamBaton <- i + 1 //交棒
			} else {
				this.streamBaton <- nextNum //把數字還回去
				i--
			}
			runtime.Gosched()
		}
	}
}

func main() {

	start := time.Now()

	fizzbuzz := &FizzBuzz{
		wg:          &sync.WaitGroup{},
		streamBaton: make(chan int, 1),
	}

	for testCase := 0; testCase <= 20; testCase++ {
		fizzbuzz.n = testCase

		fizzbuzz.wg.Add(4)
		go fizzbuzz.PrintFizz()
		go fizzbuzz.PrintBuzz()
		go fizzbuzz.PrintFizzBuzz()
		go fizzbuzz.PrintNumber()

		fizzbuzz.streamBaton <- 0 //啟動交棒
		fizzbuzz.wg.Wait()
		<-fizzbuzz.streamBaton

		fmt.Println() //這個 Test Case 結束了，換行。
	}

	spentTime := time.Now().Sub(start)
	fmt.Println("Spent time:", spentTime)
}
