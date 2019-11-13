// Lin Kao Yuan, 2019/11/13(Wed)
// 以 struct「貓」與「人」分別實現「改變體重」這個 inteface
// 來練習與示範 interface 的用法

package main

import (
	"fmt"
)

// 這是改變體重的 Interface
type BodyWeightChange interface {
	Eat(anyFood string) error
	Excretion() error
	Test() //沒有實現也不會發生錯誤
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
func (c *Cat) Eat(catFood string) error {
	if 10 <= c.weight {
		return fmt.Errorf("已經 %d kg，太胖了不可以再吃", c.weight)
	}
	fmt.Printf("%s Eat() 吃了飼料, 裡面有:%s\n", c.name, catFood)
	c.weight++
	return nil
}

func (c *Cat) Excretion() error {
	if 1 >= c.weight {
		return fmt.Errorf("已經 %d kg，太瘦了不可以再拉", c.weight)
	}
	fmt.Printf("%s Excretion() 拉完在貓沙上\n", c.name)
	c.weight--
	return nil
}

//「人」實現改變體重的 Interface
func (p *People) Eat() error {
	if 99 <= p.weight {
		return fmt.Errorf("已經 %d kg，太胖了不可以再吃", p.weight)
		//這裡若用 fmt.Sprintf() 會被編輯器提醒應該改用 fmt.Errorf()，但硬要用也可以跑。
		//等等來看看有什麼差別
		//懂了，原來是 errors.New(fmt.Sprintf()) == fmt.Errorf()
	}
	fmt.Printf("%s Eat() 吃了便當\n", p.name)
	p.weight++
	return nil
}

func (p *People) Excretion() error {
	if 1 >= p.weight {
		return fmt.Errorf("已經 %d kg，太瘦了不可以再上廁所", p.weight)
	}
	fmt.Printf("%s Excretion() 上完廁所\n", p.name)
	p.weight--
	return nil
}

func main() {
	fmt.Println("\n========= structs 初始化與狀態展示 =========")
	var firstCat Cat = Cat{name: "firstCat", weight: 5}
	var firstPeople People = People{name: "firstPeople", weight: 97}
	fmt.Printf("I'm %s, my body weight is %d\n", firstCat.name, firstCat.weight)
	fmt.Printf("I'm %s, my body weight is %d\n", firstPeople.name, firstPeople.weight)

	fmt.Println("\n========= 執行體重增加介面 =========")
	firstCat.Eat("肉")
	firstPeople.Eat()
	fmt.Printf("I'm %s, my body weight is %d\n", firstCat.name, firstCat.weight)
	fmt.Printf("I'm %s, my body weight is %d\n", firstPeople.name, firstPeople.weight)

	fmt.Println("\n========= 執行體重減少介面 =========")
	firstCat.Excretion()
	firstCat.Excretion()
	firstPeople.Excretion()
	firstPeople.Excretion()
	fmt.Printf("I'm %s, my body weight is %d\n", firstCat.name, firstCat.weight)
	fmt.Printf("I'm %s, my body weight is %d\n", firstPeople.name, firstPeople.weight)

	fmt.Println("\n========= 在超過範圍的邊緣瘋狂試探 =========")
	fmt.Println(firstCat.Excretion())
	fmt.Println(firstCat.Excretion())
	fmt.Println(firstCat.Excretion())
	fmt.Println(firstCat.Excretion())
	fmt.Println(firstCat.Excretion())
	fmt.Println(firstCat.Excretion())
	fmt.Println(firstCat.Excretion())
	fmt.Println(firstPeople.Eat())
	fmt.Println(firstPeople.Eat())
	fmt.Println(firstPeople.Eat())
	fmt.Println(firstPeople.Eat())
	fmt.Println(firstPeople.Eat())
}
