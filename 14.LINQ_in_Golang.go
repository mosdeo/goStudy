package main

import (
	"fmt"
	"strings"

	"github.com/ahmetb/go-linq"
	. "github.com/ahmetb/go-linq"
)

func LinqQueryToSliceInt(linqQuery linq.Query) (outSlice []int) {
	linqQuery.ToSlice(&outSlice)
	return outSlice
}

func main() {
	var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var words []string = []string{"Cat", "Dog", "People", "Cup", "Apple", "Boss", "Zoo", "X-ray", "Github"}

	// LINQ 基本用法
	fmt.Println(LinqQueryToSliceInt(From(numbers).Reverse())) //反轉
	fmt.Println(LinqQueryToSliceInt(Range(99, 3)))            //產生連續數列
	fmt.Println(LinqQueryToSliceInt(Range(100, -3)))          //產生連續數列
	fmt.Println(From(numbers).SumInts())                      //加總
	fmt.Println(From(numbers).Average())                      //平均
	fmt.Println(From(numbers).Min(), From(numbers).Max())     //最大最小
	fmt.Println(From(words).Last())                           //取最後一個

	//所有字母改大寫
	var onlyUpperWords = make([]string, 0)
	From(words).Select(func(s interface{}) interface{} {
		return strings.ToUpper(s.(string))
	}).ToSlice(&onlyUpperWords)
	fmt.Println(onlyUpperWords)

	//ForEach 印出每個單字加上間隔
	From(words).ForEach(func(s interface{}) {
		fmt.Print("-", s.(string))
	})
	fmt.Println()

	//挑選出含有“o”的單字
	var wordsConatinsAlphao = make([]string, 0)
	From(words).Where(func(s interface{}) bool {
		return strings.Contains(s.(string), "o")
	}).ToSlice(&wordsConatinsAlphao)
	fmt.Println(wordsConatinsAlphao)
}
