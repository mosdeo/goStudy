package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type TestCase struct {
	Qus string
	Ans int
}

var testCases = []TestCase{
	TestCase{
		Qus: "-91283472332",
		Ans: -2147483648},
	TestCase{
		Qus: "-42",
		Ans: -42},
	TestCase{
		Qus: "42",
		Ans: 42},
	TestCase{
		Qus: "4193 with words",
		Ans: 4193},
	TestCase{
		Qus: "",
		Ans: 0},
	TestCase{
		Qus: "+1",
		Ans: 1},
	TestCase{
		Qus: "+-12",
		Ans: 0},
	TestCase{
		Qus: "9223372036854775808",
		Ans: 2147483647},
	TestCase{
		Qus: " ++1",
		Ans: 0},
	TestCase{
		Qus: "21474836++",
		Ans: 21474836},
	TestCase{
		Qus: "+",
		Ans: 0},
}

func main() {

	for _, testCase := range testCases {
		fmt.Println("")
		if len(testCase.Qus) > 99 {
			fmt.Println("len(Qus)=", len(testCase.Qus))
		} else {
			fmt.Println("Qus=", testCase.Qus)
		}
		fmt.Println("True Ans=", testCase.Ans)

		startTime := time.Now()
		result := myAtoi(testCase.Qus)
		endTime := time.Now()
		fmt.Println("Spent time:", endTime.Sub(startTime))

		//答案錯誤就暫停
		if testCase.Ans != result {
			fmt.Println("Mistake answer =", result)
		}
	}
}

func myAtoi(s string) int {
	pn0123456789 := "+-0123456789"

	s = strings.Trim(s, " ") //去除左右空白

	// 無法歸納的 edge case 集中處理
	illegalPatterns := []string{"++", "+-", "-+", "--"}
	if len(s) > 1 {
		for _, ip := range illegalPatterns {
			if strings.Contains(s[0:2], ip) {
				return 0
			}
		}
	}

	s = strings.TrimLeft(s, "+") //去除開頭正號

	if "" == s {
		return 0
	}

	//開頭是否為必需字元？
	if !strings.ContainsAny(s[0:1], pn0123456789) {
		return 0
	}

	// 負號處理
	dir := 1
	if "-" == s[0:1] {
		dir = -1
		s = s[1:]
	}

	num := 0
	for _, c := range s {
		if '0' > c || c > '9' {
			break
		}
		num = num*10 + int(c-'0')
		if num < math.MinInt32 || math.MaxInt32 < num {
			break
		}
	}

	//下限處理
	if math.MinInt32 > dir*num {
		return math.MinInt32
	}
	//上限處理
	if math.MaxInt32 < dir*num {
		return math.MaxInt32
	}

	return dir * num
}
