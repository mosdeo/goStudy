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
	TestCase{Qus: ". 1", Ans: false},
	TestCase{Qus: "6+1", Ans: false},
	TestCase{Qus: "-1.", Ans: true},
	TestCase{Qus: "46.e3", Ans: true},
	TestCase{Qus: "6e6.5", Ans: false},
	TestCase{Qus: ".1 ", Ans: true},
	TestCase{Qus: "3.", Ans: true},
	TestCase{Qus: " 005047e+6", Ans: true},
	TestCase{Qus: "-1.e49046 ", Ans: true},
	TestCase{Qus: "8..e4", Ans: false},
	TestCase{Qus: ".2e81", Ans: true},
	TestCase{Qus: "-e58 ", Ans: false},
	TestCase{Qus: "+42e+76125", Ans: true},
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

	//降低處理情況複雜度
	s = strings.ReplaceAll(s, "e+", "e")
	s = strings.ReplaceAll(s, "e-", "e")
	s = strings.ReplaceAll(s, ".e", "e")
	s = strings.Trim(s, " ")

	// 這些符號最多一個
	most_one := ".e-+"
	for _, c := range most_one {
		if 1 < strings.Count(s, string(c)) {
			return false
		}
	}

	//不可出現在尾
	catnotlast := "e+-"
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

	//不可出現在中間
	catnotmiddle := "+- "
	for _, c := range catnotmiddle {
		if strings.Contains(strings.Trim(s, string(c)), string(c)) {
			return false
		}
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

	e_index := strings.Index(s, "e")

	// 有發現 e 在字串中
	if -1 != e_index {
		//e前面一個字元不能是正負號，要有係數
		if !strings.ContainsAny(s[e_index-1:e_index], str0123456789) {
			return false
		}

		//判斷e後面的字串是否只有數字？
		if "" != strings.Trim(s[e_index+1:], str0123456789) {
			return false //數字截完不是空字串，有錯
		}

		//e前面是可以是小數點或數字
		if "" != strings.Trim(s[:e_index], str0123456789+"+-.") {
			return false
		}
	}

	return true
}
