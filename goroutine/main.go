package goroutine

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func Goroutine() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error creating trace file:", err)
	}
	defer func() {
		if err := f.Close; err != nil {
			log.Fatalln("Error closing trace file:", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error starting trace:", err)
	}
	defer trace.Stop()
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	//task(ctx, "Task1")
	//task(ctx, "Task2")
	//task(ctx, "Task3")

	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	wg.Wait()

	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("main function end")

}
func task(ctx context.Context, name string) {
	// deferにチェーンして関数をつけた場合、最後の関数のみ遅延して実行され、それ以外の関数は即時実行される。
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second)
	fmt.Println("Hello from", name)
}
func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Hello from", name)
}
