package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 2)

	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: lock about 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: I unlocked, get it")
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println("goroutine2: wait lock")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: haha，I locaked")
		ch <- struct{}{}
	}()

	// wait goroutine finish
	for i := 0; i < 2; i++ {
		<-ch
	}
}
