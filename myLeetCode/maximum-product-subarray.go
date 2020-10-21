package main

import (
	"fmt"
)

type TestCase struct {
	Qus []int
	Ans int
}

func main() {
	var testCases = []TestCase{
		TestCase{
			Qus: []int{2, 3, -2, 4},
			Ans: 6,
		},
		TestCase{
			Qus: []int{-2, 0, -1},
			Ans: 0,
		},
	}

	for _, testCase := range testCases {
		fmt.Println("")
		result := maxProduct(testCase.Qus)
		fmt.Println("computedTimes=", computedTimes)
		fmt.Println("Qus=", testCase.Qus)
		fmt.Println("True Ans=", testCase.Ans)

		//答案錯誤就暫停
		if testCase.Ans != result {
			fmt.Println("Mistake answer =", result)
			break
		}
	}
}

// 以上程式碼不進 leetcode

var tableComputed map[int]map[int]int
var computedTimes int

func maxProduct(nums []int) int {
	max := -2147483647

	neg_continue_only_pdt := 1
	pos_continue_only_pdt := 1

	for _, n := range nums {

		if 0 == n {
			neg_continue_only_pdt = 1
			pos_continue_only_pdt = 1
		} else if 0 > n {
			pos_continue_only_pdt = 1
			neg_continue_only_pdt *= n
		} else if 0 < n {
			pos_continue_only_pdt *= n
			neg_continue_only_pdt = 1
		}

		max = MaxInt(max, pos_continue_only_pdt)
		max = MaxInt(max, neg_continue_only_pdt)
	}

	fmt.Println(pos_continue_only_pdt)
	fmt.Println(neg_continue_only_pdt)

	return max
}

func MaxInt(lv int, rv int) int {
	if lv > rv {
		return lv
	} else {
		return rv
	}
}
