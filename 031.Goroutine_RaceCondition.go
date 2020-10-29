package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 999; i++ {
		go RaceCondition()
	}
}

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func RaceCondition() {
	//	模擬應用中的計算延遲
	var DoSomethingCalc = func() {
		<-time.After(time.Microsecond * time.Duration(TrueRandom(5)))
	}

	var data int
	go func() {
		data++
	}()

	if 0 == data {
		DoSomethingCalc()
		fmt.Printf("%d", data)
	}
}
