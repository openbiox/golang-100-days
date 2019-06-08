package main

import (
    "fmt"
    "math/rand"
    "sync"
)

var count int
var rw sync.RWMutex

func main() {
    ch := make(chan struct{}, 10)
    for i := 0; i < 5; i++ {
        go read(i, ch)
    }
    for i := 0; i < 5; i++ {
        go write(i, ch)
    }

    for i := 0; i < 10; i++ {
        <-ch
    }
}

func read(n int, ch chan struct{}) {
    rw.RLock()
    fmt.Printf("goroutine %d enter reading mode...\n", n)
    v := count
    fmt.Printf("goroutine %d Reading finished, value：%d\n", n, v)
    rw.RUnlock()
    ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
    rw.Lock()
    fmt.Printf("goroutine %d enter writing mode...\n", n)
    v := rand.Intn(1000)
    count = v
    fmt.Printf("goroutine %d write ending，new value：%d\n", n, v)
    rw.Unlock()
    ch <- struct{}{}
}
