package main

import (
	"fmt"

	. "github.com/ahmetb/go-linq"
)

func main() {
	var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var words []string = []string{"Cat", "Dog", "People", "Cup", "Apple", "Boss", "Zoo", "X-ray", "Github"}

	fmt.Println(From(numbers).First())
	fmt.Println(From(words).First())
}
