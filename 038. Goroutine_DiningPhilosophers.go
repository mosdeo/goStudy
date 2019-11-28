package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DiningPhilosophers struct {
	wg *sync.WaitGroup
}

func (this *DiningPhilosophers) SetWaitGroup(wg *sync.WaitGroup) {
	this.wg = wg
}

func (this *DiningPhilosophers) WantToEat(philosopher int, pickLeftFork func(), pickRightFork func(), eat func(), putLeftFork func(), putRightFork func()) {
	fmt.Printf("Philosopher%d\n", philosopher)
	this.wg.Done()
}

func Eat() {
	rand.Seed(time.Now().Unix())
	<-time.After(time.Duration(rand.Int()%100) * time.Millisecond)
}

func PickLeftFork() {

}

func PickRightFork() {

}

func PutLeftFork() {

}

func PutRightFork() {

}

func main() {
	diningPhilosophers := DiningPhilosophers{}
	diningPhilosophers.SetWaitGroup(&sync.WaitGroup{})
	diningPhilosophers.wg.Add(5)
	go diningPhilosophers.WantToEat(0, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	go diningPhilosophers.WantToEat(1, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	go diningPhilosophers.WantToEat(2, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	go diningPhilosophers.WantToEat(3, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	go diningPhilosophers.WantToEat(4, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	diningPhilosophers.wg.Wait()
}
