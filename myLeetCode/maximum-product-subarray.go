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
		TestCase{
			Qus: []int{3, -1, 4},
			Ans: 4,
		},
		TestCase{
			Qus: []int{-3, -1, -1},
			Ans: 3,
		},
		TestCase{
			Qus: []int{2, -5, -2, -4, 3},
			Ans: 24,
		},
		TestCase{
			Qus: []int{1, 0, -1, 2, 3, -5, -2},
			Ans: 60,
		},
		TestCase{
			Qus: []int{-1, -2, -3, 0},
			Ans: 6,
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
			//break
		}
	}
}

// 以上程式碼不進 leetcode

// var tableComputed map[int]map[int]int
var computedTimes int

func maxProduct(nums []int) int {
	max, pdt := nums[0], nums[0]

	first_neg_pdt := func() int {
		if 0 > nums[0] {
			return nums[0]
		} else {
			return 1
		}
	}()
	pdt_is_n := false

	for _, n := range nums[1:] {

		// 自己是0就更新，不是就連乘
		if 0 == pdt {
			pdt = n
			pdt_is_n = true
		} else {
			pdt *= n
			pdt_is_n = false
		}

		//在不為零的連續段中，紀錄第一個負數積
		//遇到0重置
		if 0 > pdt && 1 == first_neg_pdt {
			first_neg_pdt = pdt
		} else if 0 == n {
			first_neg_pdt = 1
		}

		if !pdt_is_n {
			max = MaxInt(max, pdt/first_neg_pdt)
		}
		max = MaxInt(max, pdt)
		max = MaxInt(max, n)

		fmt.Printf("n=%d, max=%d, pdt=%d, first_neg_pdt=%d\n", n, max, pdt, first_neg_pdt)
	}

	return max
}

func MaxInt(lv int, rv int) int {
	if lv > rv {
		return lv
	} else {
		return rv
	}
}
