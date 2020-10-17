package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DiningPhilosophers struct {
	wg                     *sync.WaitGroup
	streamForks            [5]chan interface{}
	missingDoubleForkTimes int
}

func (this *DiningPhilosophers) WantToEat(philosopher int, pickLeftFork func(int), pickRightFork func(int), eat func(int), putLeftFork func(int), putRightFork func(int)) {
	defer this.wg.Done()

	var leftNum = (philosopher + 4) % 5  //取得該哲學家左邊的號碼
	var rightNum = (philosopher + 6) % 5 //取得該哲學家右邊的號碼

	for {
		select {
		case this.streamForks[leftNum] <- philosopher: //嘗試拿起左邊叉子
			PickLeftFork(philosopher) //成功拿起左邊叉子
			select {
			case this.streamForks[rightNum] <- philosopher: //嘗試拿起右邊叉子
				PickRightFork(philosopher)  //成功拿起又邊叉子
				Eat(philosopher)            //左右邊都拿到了，開始吃
				<-this.streamForks[leftNum] //吃完了，放下左邊叉子
				PutLeftFork(philosopher)
				<-this.streamForks[rightNum] //吃完了，放下右邊叉子
				PutRightFork(philosopher)
				return //吃飽離開
			default: //無法拿起右邊叉子
				fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, rightNum)
				<-this.streamForks[leftNum] //把已經拿起來的左邊叉子釋放出去
				PutLeftFork(philosopher)
			}
		default: //無法拿起左邊叉子
			fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, leftNum)
		}
		this.missingDoubleForkTimes++
		Think()
	}
}

func Eat(philosopher int) {
	fmt.Printf("===== Philosopher %d have eaten. =====\n", philosopher)
}

func Think() {
	Random := func(max int) int {
		rand.Seed(time.Now().Unix())
		return rand.Int() % (max + 1)
	}
	<-time.After(time.Millisecond * time.Duration(Random(50)))
}

func PickLeftFork(philosopher int) {
	var leftNum = (philosopher + 4) % 5
	fmt.Printf("Philosopher %d picked fork %d.\n", philosopher, leftNum)
}

func PickRightFork(philosopher int) {
	var rightNum = (philosopher + 6) % 5
	fmt.Printf("Philosopher %d picked fork %d.\n", philosopher, rightNum)
}

func PutLeftFork(philosopher int) {
	var leftNum = (philosopher + 4) % 5
	fmt.Printf("Philosopher %d putted fork %d.\n", philosopher, leftNum)

}

func PutRightFork(philosopher int) {
	var rightNum = (philosopher + 6) % 5
	fmt.Printf("Philosopher %d putted fork %d.\n", philosopher, rightNum)

}

func main() {
	diningPhilosophers := DiningPhilosophers{
		wg: &sync.WaitGroup{},
	}

	// Channel 初始化
	for i := range diningPhilosophers.streamForks {
		diningPhilosophers.streamForks[i] = make(chan interface{}, 1)
	}

	// 叫所有哲學家開始動作
	start := time.Now()
	for i := range diningPhilosophers.streamForks {
		diningPhilosophers.wg.Add(1)
		go diningPhilosophers.WantToEat(i, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	}

	//等待每一位哲學家都吃過
	diningPhilosophers.wg.Wait()
	fmt.Println("Spent time:", time.Now().Sub(start))
	fmt.Printf("Missing double forks %d times.", diningPhilosophers.missingDoubleForkTimes)
}
