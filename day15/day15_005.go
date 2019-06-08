package main

import (
    "fmt"
    "sync"
    "time"
)

var count int = 4

func main() {
    ch := make(chan struct{}, 5)

    // new cond
    var l sync.Mutex
    cond := sync.NewCond(&l)

    for i := 0; i < 5; i++ {
        go func(i int) {
            // Grab Mutual Exclusive Locks
            cond.L.Lock()
            defer func() {
                cond.L.Unlock()
                ch <- struct{}{}
            }()

            // check condition
            for count > i {
                cond.Wait()
                fmt.Printf("get broadcast goroutine%d\n", i)
            }

            fmt.Printf("goroutine%d finished\n", i)
        }(i)
    }

    // ensure all goroutine started
    time.Sleep(time.Millisecond * 20)
    // lock and change count
    fmt.Println("broadcast...")
    cond.L.Lock()
    count -= 1
    cond.Broadcast()
    cond.L.Unlock()

    time.Sleep(time.Second)
    fmt.Println("signal...")
    cond.L.Lock()
    count -= 2
    cond.Signal()
    cond.L.Unlock()

    time.Sleep(time.Second)
    fmt.Println("broadcast...")
    cond.L.Lock()
    count -= 1
    cond.Broadcast()
    cond.L.Unlock()

    for i := 0; i < 5; i++ {
        <-ch
    }
}
