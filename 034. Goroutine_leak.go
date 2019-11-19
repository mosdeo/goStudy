package main

import (
	"fmt"
)

func main() {

	// 使 strings 得到空字串
	var DoWork = func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("DoWork exited.")
			defer close(completed)

			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	DoWork(nil)
	fmt.Println("Done")
}
