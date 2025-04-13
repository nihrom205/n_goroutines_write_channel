package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 7
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- fmt.Sprintf("горутина № %d", n)
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
