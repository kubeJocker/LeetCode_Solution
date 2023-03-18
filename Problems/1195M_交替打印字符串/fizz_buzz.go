package _195M_交替打印字符串

import (
	"fmt"
	"sync"
)

// 以number为核心，每次存在数量到条件时就signal放开对应的来线程函数来跑
type FizzBuzz struct {
	n            int
	fizzCond     *sync.Cond
	buzzCond     *sync.Cond
	fizzbuzzCond *sync.Cond
	numberCond   *sync.Cond
}

var waitGroup = sync.WaitGroup{}

func (s *FizzBuzz) fizz(f func()) {
	for {
		s.fizzCond.L.Lock()
		s.fizzCond.Wait()
		f()
		s.numberCond.Signal()
	}
}

func (s *FizzBuzz) buzz(f func()) {
	for {
		s.buzzCond.L.Lock()
		s.buzzCond.Wait()
		s.buzzCond.L.Unlock()
		f()
		s.numberCond.Signal()
	}
}

func (s *FizzBuzz) fizzbuzz(f func()) {
	for {
		s.fizzbuzzCond.L.Lock()
		s.fizzbuzzCond.Wait()
		s.fizzbuzzCond.L.Unlock()
		f()
		s.numberCond.Signal()
	}
}

func (s *FizzBuzz) number(f func(int)) {
	for i := 1; i <= s.n; i++ {
		s.numberCond.L.Lock()
		if i%3 == 0 && i%5 == 0 {
			s.fizzbuzzCond.Signal()

			s.numberCond.Wait()
		} else if i%3 == 0 {
			s.fizzCond.Signal()
			s.numberCond.Wait()
		} else if i%5 == 0 {
			s.buzzCond.Signal()
			s.numberCond.Wait()
		} else {
			f(i)
		}
		s.numberCond.L.Unlock()
		if i < s.n {
			fmt.Print(", ")
		}
	}
	waitGroup.Done()
}

func main() {
	fizzBuzz := &FizzBuzz{
		n:            15,
		fizzCond:     sync.NewCond(new(sync.Mutex)),
		buzzCond:     sync.NewCond(new(sync.Mutex)),
		fizzbuzzCond: sync.NewCond(new(sync.Mutex)),
		numberCond:   sync.NewCond(new(sync.Mutex)),
	}
	waitGroup.Add(1)
	go fizzBuzz.fizz(printFizz)
	go fizzBuzz.buzz(printBuzz)
	go fizzBuzz.fizzbuzz(printFizzBuzz)
	go fizzBuzz.number(printNumber)
	waitGroup.Wait()
}

func printFizz() {
	fmt.Print("fizz")
}

func printBuzz() {
	fmt.Print("buzz")
}

func printFizzBuzz() {
	fmt.Print("fizzbuzz")
}

func printNumber(i int) {
	fmt.Print(i)
}
