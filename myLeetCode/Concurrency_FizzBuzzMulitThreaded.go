// 改為「不採用 channel，讓各個 goroutine 自行負擔檢查整除條件的責任」的版本。
// 不過這個版本好像太簡單了，下一個版本要用各自邊跑邊交棒的。

package main

import (
	"fmt"
	"sync"
	"time"
)

type FizzBuzz struct {
	n  int
	wg *sync.WaitGroup
}

func (this *FizzBuzz) PrintFizz() {
	defer this.wg.Done()

	if 0 == this.n%3 && 0 != this.n%5 {
		fmt.Print("Fizz, ")
	}
}

func (this *FizzBuzz) PrintBuzz() {
	defer this.wg.Done()

	if 0 != this.n%3 && 0 == this.n%5 {
		fmt.Print("Buzz, ")
	}
}

func (this *FizzBuzz) PrintFizzBuzz() {
	defer this.wg.Done()

	if 0 == this.n%(3*5) {
		fmt.Print("FizzBuzz, ")
	}
}

func (this *FizzBuzz) PrintNumber() {
	defer this.wg.Done()

	if 0 != this.n%3 && 0 != this.n%5 {
		fmt.Printf("%d, ", this.n)
	}
}

func main() {

	start := time.Now()

	fizzbuzz := &FizzBuzz{
		wg: &sync.WaitGroup{},
	}

	for testCase := 0; testCase <= 19; testCase++ {
		for i := 0; i <= testCase; i++ {
			fizzbuzz.n = i
			fizzbuzz.wg.Add(4)
			go fizzbuzz.PrintFizz()
			go fizzbuzz.PrintBuzz()
			go fizzbuzz.PrintFizzBuzz()
			go fizzbuzz.PrintNumber()
			fizzbuzz.wg.Wait()
		}
		fmt.Println() //這個 Test Case 結束了，換行。
	}

	spentTime := time.Now().Sub(start)
	fmt.Println("Spent time:", spentTime)
}
