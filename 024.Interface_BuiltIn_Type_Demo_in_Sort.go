// Lin Kao Yuan, 2019/11/13(Wed)
// 藉由實現 sort.Interface 排序中文數字
// 來練習對現成 library 的擴充

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type ChineseNumber struct {
	value string
}

type ChineseNumberSlice []ChineseNumber

func (p ChineseNumberSlice) Len() int { return len(p) }
func (p ChineseNumberSlice) Less(i, j int) bool {
	var chineseNumMappingTable = map[string]int{
		"零": 0, "一": 1, "二": 2, "三": 3, "四": 4,
		"五": 5, "六": 6, "七": 7, "八": 8, "九": 9}

	return chineseNumMappingTable[p[i].value] > chineseNumMappingTable[p[j].value]
}
func (p ChineseNumberSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	var sliceChineseNumber = []ChineseNumber{
		0: ChineseNumber{value: "零"},
		1: ChineseNumber{value: "一"},
		2: ChineseNumber{value: "二"},
		3: ChineseNumber{value: "三"},
		4: ChineseNumber{value: "四"},
		5: ChineseNumber{value: "五"},
		6: ChineseNumber{value: "六"},
		7: ChineseNumber{value: "七"},
		8: ChineseNumber{value: "八"},
		9: ChineseNumber{value: "九"},
	}
	fmt.Println("初始  :", sliceChineseNumber)

	rand.Seed(int64(time.Now().Second()))
	rand.Shuffle(len(sliceChineseNumber), func(i, j int) {
		{
			sliceChineseNumber[i], sliceChineseNumber[j] =
				sliceChineseNumber[j], sliceChineseNumber[i]
		}
	})
	fmt.Println("打亂後:", sliceChineseNumber)

	//這裡是一個地雷，要轉型成已經明確定義的型別，不可以傳入[]T，
	//否則會得到「does not implement sort.Interface (missing Len method)」的編譯錯誤
	sort.Sort(ChineseNumberSlice(sliceChineseNumber))
	fmt.Println("排序後:", sliceChineseNumber)
}
