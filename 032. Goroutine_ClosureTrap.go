package main

import (
	"fmt"
	"sync"
)

var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	syncWaitGroupStyle()
	syncWaitGroupStylePassTrap()
}

func manulCounterStye() {
	var goroutineCount int

	for _, num := range numbers {

		// 並發的匿名函式
		goroutineCount++
		go func(num int) {
			fmt.Println(num)
			goroutineCount--
		}(num)
	}

	for 0 != goroutineCount {
	}
}

func syncWaitGroupStylePassTrap() {
	wg := sync.WaitGroup{}

	for _, num := range numbers {

		// 並發的匿名函式
		wg.Add(1)
		go func(num int) {
			fmt.Println(num)
			wg.Done()
		}(num)
	}

	wg.Wait()
}

func syncWaitGroupStyle() {
	wg := sync.WaitGroup{}

	for _, num := range numbers {

		// 並發的匿名函式
		wg.Add(1)
		go func() {
			fmt.Println(num)
			wg.Done()
		}()
	}

	wg.Wait()
}
