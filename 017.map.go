package main

import "fmt"

func main() {
	var table map[string]int = make(map[string]int)

	//加入
	table["Apple"] = 10
	table["Bob"] = 87
	table["Cat"] = 9
	table["Daddy"] = 69
	ListMapInfo(table) //刪除前

	//刪除
	delete(table, "Bob")
	ListMapInfo(table) //刪除後

	//查詢不存在的 key
	//會根據對應型別回傳一個空值
	fmt.Println(table["SDCFVGYBHUNJ"]) //0
	fmt.Println(table["sdfghjkl"])     //0
	//這會造成誤解，以為這兩個 key 對應到的 Value 就是 0

	//更好的空檢測方法
	value, ok := table["Apple"]
	fmt.Println(value, ok)
	value, ok = table["sdfghjkl"]
	fmt.Println(value, ok)

	//------------------------------------
	//大量初始化 map 的方式
	myDict := map[string]string{
		"H": "h",
		"T": "t",
		"X": "x",
		"Y": "y",
		"Z": "z",
	}

	fmt.Println(myDict)
}

func ListMapInfo(theMap map[string]int) {
	fmt.Println(theMap)
	fmt.Printf("len = %d\n", len(theMap))
}
