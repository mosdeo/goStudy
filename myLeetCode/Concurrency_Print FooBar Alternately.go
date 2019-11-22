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
	for i := 0; i < this.n; i++ {

		// printFoo() outputs "foo". Do not change or remove this line.
		<-this.streamBarToFoo
		printFoo()
		this.streamFooToBar <- struct{}{}
	}
}

func (this *FooBar) Bar(printBar func()) {
	for i := 0; i < this.n; i++ {

		// printBar() outputs "bar". Do not change or remove this line.
		<-this.streamFooToBar
		printBar()
		this.streamBarToFoo <- struct{}{}
	}

	this.streamEnd <- struct{}{}
}

func main() {
	fmt.Println()

	fooBar := &FooBar{
		n:              4,
		streamFooToBar: make(chan struct{}),
		streamBarToFoo: make(chan struct{}),
		streamEnd:      make(chan struct{}),
	}

	go fooBar.Foo(func() { fmt.Printf("Foo") })
	go fooBar.Bar(func() { fmt.Printf("Bar ") })
	fooBar.streamBarToFoo <- struct{}{} //啟動
	<-fooBar.streamEnd                  //as wg.Wait()

	//關閉各個 worker

}
