// 以 struct「貓」與「人」分別實現「改變體重」這個 inteface
// 來練習與示範 interface 的用法
// Lin Kao Yuan, 2019/11/13(Wed)

package main

import "fmt"

// 這是改變體重的 Interface
type BodyWeightChange interface {
	Eat() int
	Excretion() int
}

// 這是「貓」結構
type Cat struct {
	name   string
	weight int
}

// 這是「人」結構
type People struct {
	name   string
	weight int
}

//「貓」實現改變體重的 Interface
func (c *Cat) Eat() int {
	fmt.Printf("%s Eat() 吃飼料\n", c.name)
	c.weight++
	return c.weight
}

func (c *Cat) Excretion() int {
	fmt.Printf("%s Excretion() 拉在貓沙上\n", c.name)
	c.weight--
	return c.weight
}

//「人」實現改變體重的 Interface
func (p *People) Eat() int {
	fmt.Printf("%s Eat() 吃便當\n", p.name)
	p.weight++
	return p.weight
}

func (p *People) Excretion() int {
	fmt.Printf("%s Excretion() 去上廁所\n", p.name)
	p.weight--
	return p.weight
}

func main() {
	var firstCat Cat = Cat{name: "firstCat", weight: 5}
	var firstPeople People = People{name: "firstPeople", weight: 20}
	fmt.Printf("I'm %s, my body weight is %d\n", firstCat.name, firstCat.weight)
	fmt.Printf("I'm %s, my body weight is %d\n", firstPeople.name, firstPeople.weight)

	firstCat.Eat()
	firstPeople.Eat()
	fmt.Printf("I'm %s, my body weight is %d\n", firstCat.name, firstCat.weight)
	fmt.Printf("I'm %s, my body weight is %d\n", firstPeople.name, firstPeople.weight)

	firstCat.Excretion()
	firstCat.Excretion()
	firstPeople.Excretion()
	firstPeople.Excretion()
	fmt.Printf("I'm %s, my body weight is %d\n", firstCat.name, firstCat.weight)
	fmt.Printf("I'm %s, my body weight is %d\n", firstPeople.name, firstPeople.weight)
}
