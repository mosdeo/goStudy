package main

import (
	"fmt"
	"strings"

	"github.com/ahmetb/go-linq"
)

func main() {
	var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var words []string = []string{"Cat", "Dog", "People", "Cup", "Apple", "Boss", "Zoo", "X-ray", "Github"}

	// LINQ 基本用法
	fmt.Println(linq.Range(99, 3).Results())                                      //產生連續數列
	fmt.Println(linq.From(numbers).SumInts())                                     //加總
	fmt.Println(linq.From(numbers).Average())                                     //平均
	fmt.Println(linq.From(numbers).Min(), linq.From(numbers).Max())               //最大最小
	fmt.Println(linq.From(words).Last())                                          //取最後一個
	fmt.Println(linq.From([]int{1, 2, 3, 4, 4, 4, 4, 5}).Distinct().Results()...) //去重
	fmt.Println(linq.From([]int{1, 2, 3, 4, 5, 6, 7, 8}).Reverse().Results()...)  //反轉

	//所有字母改大寫
	var onlyUpperWords = make([]string, 0)
	linq.From(words).Select(func(s interface{}) interface{} {
		return strings.ToUpper(s.(string))
	}).ToSlice(&onlyUpperWords)
	fmt.Println(onlyUpperWords)

	//ForEach 印出每個單字加上間隔
	linq.From(words).ForEach(func(s interface{}) {
		fmt.Print("-", s.(string))
	})
	fmt.Println()

	//找 Index
	fmt.Printf("Boss at %d\n", linq.From(words).IndexOf(func(item interface{}) bool { return item == "Boss" }))

	//挑選出含有“o”的單字
	var wordsConatinsAlphao = make([]string, 0)
	linq.From(words).Where(func(s interface{}) bool {
		return strings.Contains(s.(string), "o")
	}).ToSlice(&wordsConatinsAlphao)
	fmt.Println(wordsConatinsAlphao)
}
