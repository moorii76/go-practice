package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine() {
	fmt.Println("--------------------ゴルーチン--------------------")
	// for-select
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 1):
				fmt.Println("Waiting...")
			case <-ch:
				fmt.Println("Done!")
				return
			}
		}
	}()

	time.Sleep(time.Second * 5)
	ch <- struct{}{}
	time.Sleep(time.Second * 1)

	// WaitGroup
	fmt.Println("----------------------------------------")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Second * 1)
			fmt.Printf("waitGroup:%v\n", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("waitGroup:Done")
}
