package main

import (
	"fmt"
	"sync"
)

type ZeroEvenOdd struct {
	n  int
	wg *sync.WaitGroup
}

// func(this *ZeroEvenOdd) SetN(n int){
// 	this.n = n
// }

func (this *ZeroEvenOdd) SetWaitGroup(wg *sync.WaitGroup) {
	this.wg = wg
}

func (this *ZeroEvenOdd) Zero(printNumber func(int)) {
	defer this.wg.Done()

}

func (this *ZeroEvenOdd) Even(printNumber func(int)) {
	defer this.wg.Done()

}

func (this *ZeroEvenOdd) Odd(printNumber func(int)) {
	defer this.wg.Done()

}

func PrintNumber(x int) {
	fmt.Printf("%d", x)
}

func main() {
	var zeo = &ZeroEvenOdd{n: 5}

	//設定同步
	wg := &sync.WaitGroup{}
	zeo.SetWaitGroup(wg)

	wg.Add(3)
	go zeo.Zero(PrintNumber)
	go zeo.Even(PrintNumber)
	go zeo.Odd(PrintNumber)
	wg.Wait()
}
