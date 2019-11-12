package main

import (
	"fmt"
)

// Go 語言沒有 class，要自定義變數的集合只有 struct，而且方法不能寫在原本的 scope 內，要另外寫。
type Cat struct {
	age    int
	weight int
	color  string
}

func (c *Cat) Bark() string {
	return "Meow"
}

func (c *Cat) GetAge() int {
	return c.age
}

func (c Cat) GetWeight() int {
	return c.weight
}

// 這個不能改變 weight
func (c Cat) SetWeightByValueMethod(weight int) {
	c.weight = weight
}

// 這個才可以改變 weight
func (c *Cat) SetWeightByPointerMethod(weight int) {
	c.weight = weight
}

// 參考
// https://golang.org/doc/faq#methods_on_values_or_pointers

func (c Cat) GetColor() string {
	return c.color
}

func main() {
	var firstCat Cat = Cat{age: 2, weight: 5, color: "Black"}
	fmt.Println(firstCat.Bark())
	fmt.Printf("I'm %d years old.\n", firstCat.GetAge())
	fmt.Printf("My skin color is %s .\n", firstCat.GetColor())

	// Demo different set method between ByValue and ByPointer.
	fmt.Printf("My body weight is %d kg .\n", firstCat.GetWeight())
	firstCat.SetWeightByValueMethod(7)
	fmt.Printf("My body weight is %d kg .\n", firstCat.GetWeight())
	firstCat.SetWeightByPointerMethod(7)
	fmt.Printf("My body weight is %d kg .\n", firstCat.GetWeight())
}
