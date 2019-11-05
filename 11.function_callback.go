package main

import "fmt"

type ConditionOfInt = func(int) bool

func IsOdd(number int) bool {
	return 1 == number&1
}

func IsEven(number int) bool {
	return 0 == number&1
}

func Filter(Decider ConditionOfInt, numbers []int) (result []int) {
	for _, number := range numbers {
		if Decider(number) {
			result = append(result, number)
		}
	}

	return result
}

func main() {
	var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(Filter(IsEven, array))
	fmt.Println(Filter(IsOdd, array))
}
