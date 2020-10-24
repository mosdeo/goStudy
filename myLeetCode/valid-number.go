package main

import (
	"fmt"
	"strings"
	"time"
)

type TestCase struct {
	Qus string
	Ans bool
}

var testCases = []TestCase{
	TestCase{Qus: ".e1", Ans: false},
	TestCase{Qus: "e9", Ans: false},
	TestCase{Qus: "1e", Ans: false},
	TestCase{Qus: "9-", Ans: false},
	TestCase{Qus: "0", Ans: true},
	TestCase{Qus: "11", Ans: true},
	TestCase{Qus: "1 ", Ans: true},
	TestCase{Qus: ".1 ", Ans: true},
	TestCase{Qus: "3.", Ans: true},
	TestCase{Qus: ". 1", Ans: false},
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
		result := isNumber(testCase.Qus)
		endTime := time.Now()
		fmt.Println("Spent time:", endTime.Sub(startTime))

		//答案錯誤就暫停
		if testCase.Ans != result {
			fmt.Println("Mistake answer =", result)
		}
	}
}

func isNumber(s string) bool {
	str0123456789 := func() string {
		str := ""
		for i := '0'; i <= '9'; i++ {
			str += string(i)
		}
		return str
	}()

	//如果只有一個字，只能是數字
	if 1 == len(s) {
		return strings.ContainsAny(s, str0123456789)
	}

	//找不到數字就 false
	if !strings.ContainsAny(s, str0123456789) {
		return false
	}

	// 這些符號最多一個
	most_one := ".e_+"
	for _, c := range most_one {
		if 1 < strings.Count(s, string(c)) {
			return false
		}
	}

	//不可出現在尾
	catnotlast := "e_+-"
	for _, c := range catnotlast {
		len := len(s)
		if s[len-1:] == string(c) {
			return false
		}
	}

	//不可出現在頭
	catnotfirst := "e"
	for _, c := range catnotfirst {
		if s[0:1] == string(c) {
			return false
		}
	}

	//有空格在中間
	if strings.Contains(strings.Trim(s, " "), " ") {
		return false
	}

	//有非法符號就 false
	illegalAlpha := func() string {
		str := ""
		for i := 0; i < 26; i++ {
			str += string('A' + i)
			str += string('a' + i)
		}
		str = strings.ReplaceAll(str, "e", "")
		return str
	}()
	if strings.ContainsAny(s, illegalAlpha) {
		return false
	}

	//e前後一定要有數字
	e_index := strings.Index(s, "e")
	if -1 != e_index {
		if 3 <= len(s) {
			if !strings.ContainsAny(s[e_index-1:e_index], str0123456789) {
				return false
			}
			if !strings.ContainsAny(s[e_index+1:e_index+2], str0123456789) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
