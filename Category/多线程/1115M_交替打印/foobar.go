package _115M_交替打印

import (
	"context"
	"fmt"
)

type FooBar struct {
	n       int
	barchan chan int
	foochan chan int
}

func (s *FooBar) foo(ctx context.Context) {
	for i := 0; i < s.n; i++ {
		select {
		case <-ctx.Done():
			return
		case <-s.foochan:
			fmt.Print("Foo")
			s.barchan <- 1
		}
	}
}

var end = make(chan int)

func (s *FooBar) bar(ctx context.Context) {
	for i := 0; i < s.n; i++ {
		select {
		case <-ctx.Done():
			return
		case <-s.barchan:
			fmt.Print("Bar")
			s.foochan <- 1
		}
	}
	end <- 1
}

func main() {
	n := 3
	foobar := &FooBar{n: n, foochan: make(chan int), barchan: make(chan int)}
	ctx, cancel := context.WithCancel(context.Background())
	foobar.foochan <- 1
	defer cancel()
	go foobar.foo(ctx)
	go foobar.bar(ctx)
	<-end
}
