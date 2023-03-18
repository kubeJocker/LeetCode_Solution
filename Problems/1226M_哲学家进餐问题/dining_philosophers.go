package _226M_哲学家进餐问题

import (
	"fmt"
	"sync"
)

type DiningPhilosophers struct {
	getkey    []chan int
	endsSgnal *sync.WaitGroup
}

func pickLeftFork(i int) {
	fmt.Printf("philosopher %d pickLeftFork\n", i)
}
func pickRightFork(i int) {
	fmt.Printf("philosopher %d pickRightFork\n", i)
}
func eat(i int) {
	fmt.Printf("philosopher %d eat\n", i)
}
func putLeftFork(i int) {
	fmt.Printf("philosopher %d putLeftFork\n", i)
}
func putRightFork(i int) {
	fmt.Printf("philosopher %d putRightFork\n", i)
}

func (s *DiningPhilosophers) wantsToEat(philosopher int, pickLeftFork func(int), pickRightFork func(int), eat func(int), putLeftFork func(int), putRightFork func(int)) {
	if philosopher%2 == 0 {
		<-s.getkey[philosopher]
		pickLeftFork(philosopher)
		<-s.getkey[(philosopher+1)%5]
		pickRightFork(philosopher)

		eat(philosopher)

		putRightFork(philosopher)
		s.getkey[(philosopher+1)%5] <- 1
		putLeftFork(philosopher)
		s.getkey[philosopher] <- 1
	} else {
		<-s.getkey[(philosopher+1)%5]
		pickRightFork(philosopher)
		<-s.getkey[philosopher]
		pickLeftFork(philosopher)

		eat(philosopher)

		putLeftFork(philosopher)
		s.getkey[philosopher] <- 1
		putRightFork(philosopher)
		s.getkey[(philosopher+1)%5] <- 1
	}
	s.endsSgnal.Done()
}

func main() {
	s := &DiningPhilosophers{
		getkey:    make([]chan int, 5),
		endsSgnal: &sync.WaitGroup{},
	}
	for i := 0; i < len(s.getkey); i++ {
		s.getkey[i] = make(chan int, 1)
		s.getkey[i] <- 1
	}
	n := 100
	for i := 0; i < n; i++ {
		s.endsSgnal.Add(5)
		go s.wantsToEat(0, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)
		go s.wantsToEat(1, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)
		go s.wantsToEat(2, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)
		go s.wantsToEat(3, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)
		go s.wantsToEat(4, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)
	}
	s.endsSgnal.Wait()
}
