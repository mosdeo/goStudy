// The same instance of FooBar will be passed to two different threads.
// Thread A will call foo() while thread B will call bar().
// Modify the given program to output "foobar" n times.
// Ex1:
//  Input: n = 1
//  Output: "foobar"
// Ex1:
//  Input: n = 2
//  Output: "foobarfoobar"

package main

import "fmt"

type FooBar struct {
	n              int
	streamFooToBar chan struct{}
	streamBarToFoo chan struct{}
	streamEnd      chan struct{}
}

func (this *FooBar) Foo(printFoo func()) {

	for i := 0; i < this.n; {
		// printFoo() outputs "foo". Do not change or remove this line.
		<-this.streamBarToFoo
		printFoo()
		i++
		this.streamFooToBar <- struct{}{}
	}

	<-this.streamBarToFoo //等待 Bar() 印完最後一次
}

func (this *FooBar) Bar(printBar func()) {

	for i := 0; i < this.n; {
		// printBar() outputs "bar". Do not change or remove this line.
		<-this.streamFooToBar
		printBar()
		i++
		this.streamBarToFoo <- struct{}{}
	}

	this.streamEnd <- struct{}{}
}

func main() {

	var PrintFooBar = func(times int) {

		fooBar := &FooBar{
			n:              times,
			streamFooToBar: make(chan struct{}),
			streamBarToFoo: make(chan struct{}),
			streamEnd:      make(chan struct{}),
		}

		go fooBar.Foo(func() { fmt.Printf("Foo") })
		go fooBar.Bar(func() { fmt.Printf("Bar ") })
		fooBar.streamBarToFoo <- struct{}{} //啟動
		<-fooBar.streamEnd                  //as wg.Wait()
		fmt.Println()
	}

	testCase := []int{0, 1, 2, 3, 4, 5, 7}

	for _, repeat := range testCase {
		fmt.Printf("Repeat %d: ", repeat)
		PrintFooBar(repeat)
	}
}
