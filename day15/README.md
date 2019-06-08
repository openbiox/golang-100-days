## sync

See <https://studygolang.com/articles/11038?fr=sidebar>

```golang
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var a = 0

    // 100 thread
    // var lock sync.Mutex
    for i := 0; i < 100; i++ {
        go func(idx int) {
            // lock.Lock()
            // defer lock.Unlock()
            a += 1
            fmt.Printf("goroutine %d, a=%d\n", idx, a)
        }(i)
    }
    time.Sleep(time.Second)
}
```

```golang
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
        fmt.Println("groutine2: wait lock")
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
```

```golang
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
```

```golang
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    seconds := []int{1, 2, 3, 4}

    for i, s := range seconds {
        // counter plus 1
        wg.Add(1)
        go func(i, s int) {
            // counter minus 1
            defer wg.Done()
            fmt.Printf("goroutine%d ending\n", i)
        }(i, s)
    }
    
    // waiting
    wg.Wait()
    fmt.Println("all goroutine finished")
}
```

```golang
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
```

```golang
package main

import (
    "bytes"
    "io"
    "os"
    "sync"
    "time"
)

var bufPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func timeNow() time.Time {
    return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
    // get temp obj, if not then create
    b := bufPool.Get().(*bytes.Buffer)
    b.Reset()
    b.WriteString(timeNow().UTC().Format(time.RFC3339))
    b.WriteByte(' ')
    b.WriteString(key)
    b.WriteByte('=')
    b.WriteString(val)
    w.Write(b.Bytes())
    // put temp obj to pool
    bufPool.Put(b)
}

func main() {
    Log(os.Stdout, "path", "/search?q=flowers")
}
```

```golang
package main

import (
    "fmt"
    "sync"
)

func main() {
    var once sync.Once
    onceBody := func() {
        fmt.Println("Only once")
    }
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            once.Do(onceBody)
            done <- true
        }()
    }
    for i := 0; i < 10; i++ {
        <-done
    }
}
```
