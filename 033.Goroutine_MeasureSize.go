package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{} //只能被「取」的單向通道
	var wg sync.WaitGroup
	var noop = func() {
		wg.Done()
		<-c // unbuffered & 沒放東西進去，所以會永遠卡死在這裡。
	}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	after := memConsumed()
	fmt.Printf("%.3f kb\n", float64(after-before)/numGoroutines/1000)
}
