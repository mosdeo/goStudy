package main

import "fmt"

type Human struct {
	sex string
	age int
}

type Labor struct {
	Human
	Job       string
	Seniority int
}

// Override To String
func (labor Labor) String() string {
	return fmt.Sprintf("Sex:%s, Age:%d, Job:%s, Seniority:%d", labor.sex, labor.age, labor.Job, labor.Seniority)
}

func main() {
	labor := Labor{Human: Human{sex: "M", age: 22}, Job: "Coder", Seniority: 3}
	fmt.Println(labor)
	fmt.Println(labor.String())
}
