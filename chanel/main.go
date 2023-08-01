package chanel

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Chanel() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 10
		time.Sleep(500 * time.Millisecond)
	}()
	fmt.Println(<-ch)
	wg.Wait()

	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	fmt.Println("num of working goroutines:", runtime.NumGoroutine())

	ch2 := make(chan int, 1)
	ch2 <- 2
	fmt.Println(<-ch2)

	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	close(ch1)
	v, ok := <-ch1
	fmt.Println(v, ok)
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	v, ok = <-ch2
	fmt.Println(v, ok)
	v, ok = <-ch2
	fmt.Println(v, ok)
	v, ok = <-ch2
	fmt.Println(v, ok)

	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
	}

	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("goroutine", i, "start")
			<-nCh
			fmt.Println("goroutine", i, "done")
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh)
	fmt.Println("unblock all goroutines")

	wg.Wait()
	fmt.Println("all goroutines done")
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
		}
	}()
	return ch
}
