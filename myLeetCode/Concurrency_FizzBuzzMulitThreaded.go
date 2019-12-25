package main

import (
	"fmt"
	"sync"
	"time"
)

type FizzBuzz struct {
	n                int
	wg               *sync.WaitGroup
	streamToFizz     chan struct{}
	streamToBuzz     chan struct{}
	streamToFizzBuzz chan struct{}
	streamToNumber   chan struct{}
	streamExit       chan struct{}
}

// printFizz() outputs "fizz".
func (this *FizzBuzz) PrintFizz() {
	defer this.wg.Done()

	for {
		select {
		case <-this.streamToFizz:
			fmt.Print("Fizz, ")
			this.streamToNumber <- struct{}{}
		case <-this.streamExit:
			return
		default:
			<-time.After(time.Duration(20) * time.Millisecond)
		}
	}
}

// printBuzz() outputs "buzz".
func (this *FizzBuzz) PrintBuzz() {
	defer this.wg.Done()

	for {
		select {
		case <-this.streamToBuzz:
			fmt.Print("Buzz, ")
			this.streamToNumber <- struct{}{}
		case <-this.streamExit:
			return
		default:
			<-time.After(time.Duration(20) * time.Millisecond)
		}
	}
}

// printFizzBuzz() outputs "fizzbuzz".
func (this *FizzBuzz) PrintFizzBuzz() {
	defer this.wg.Done()

	for {
		select {
		case <-this.streamToFizzBuzz:
			fmt.Print("FizzBuzz, ")
			this.streamToNumber <- struct{}{}
		case <-this.streamExit:
			return
		case <-time.After(time.Duration(20) * time.Millisecond):
		}
	}
}

// You may call global function `void printNumber(int x)`
// to output "x", where x is an integer.
func (this *FizzBuzz) PrintNumber() {
	defer this.wg.Done()

	for i := 1; i <= this.n; {
		select {
		case <-this.streamToNumber:
			if 0 == i%3 && 0 == i%5 {
				this.streamToFizzBuzz <- struct{}{}
			} else if 0 == i%3 {
				this.streamToFizz <- struct{}{}
			} else if 0 == i%5 {
				this.streamToBuzz <- struct{}{}
			} else {
				fmt.Printf("%d, ", i)
				this.streamToNumber <- struct{}{}
			}
			i++
		case <-time.After(time.Duration(20) * time.Millisecond):
		}
	}

	// 有三個 thread 要接收訊號，訊號會被消費，所以要發三次
	this.streamExit <- struct{}{}
	this.streamExit <- struct{}{}
	this.streamExit <- struct{}{}
}

func main() {
	for i := 0; i <= 25; i++ {
		fizzbuzz := &FizzBuzz{
			n:                i,
			wg:               &sync.WaitGroup{},
			streamToFizz:     make(chan struct{}),
			streamToBuzz:     make(chan struct{}),
			streamToFizzBuzz: make(chan struct{}),
			streamToNumber:   make(chan struct{}, 1),
			streamExit:       make(chan struct{}),
		}
		fizzbuzz.wg.Add(4)
		go fizzbuzz.PrintFizz()
		go fizzbuzz.PrintBuzz()
		go fizzbuzz.PrintFizzBuzz()
		go fizzbuzz.PrintNumber()

		fizzbuzz.streamToNumber <- struct{}{} //啟動
		fizzbuzz.wg.Wait()
		fmt.Println()
	}
}
