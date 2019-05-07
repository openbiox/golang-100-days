## net/http

```go
// Get URL response
resp, err := http.Get(url)

// don't leak resources
resp.Body.Close()

// set timeout limitation
timeout := time.Duration(5 * time.Second)
client := http.Client{Timeout: timeout}
client.Get(url)
```

## time

```go
start := time.Now()

elapse := time.Since(start).Seconds()
```

## fmt.Fprintf and fmt.Sprintf

```go
fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
fmt.Sprintf("while reading %s: %v", url, err)
```

## chan and go

```go
package main

import "fmt"

func gotest (ch chan<- string) {
  // not equal the R '<-'
  ch <- fmt.Sprintf("%.2fs %7d %s", 3.3, 22 , "ss")
}

func main() {
    ch := make(chan string)
    go gotest(ch)
    // <- : required for chanel data flow 
    fmt.Println(<-ch)
}
```