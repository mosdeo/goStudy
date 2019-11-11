package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(-5 & math.MaxInt32)
	fmt.Println(int32(1<<31 - 1))
	fmt.Println(-5 & 0xFFFF)
}
