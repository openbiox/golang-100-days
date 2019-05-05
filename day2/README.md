## bufio pkg

```go

// From https://studygolang.com/articles/11905?fr=sidebar
// io.Stdin => strings.NewReader
// input := bufio.NewScanner(os.Stdin)
input := "foo  bar   baz"
scanner := bufio.NewScanner(strings.NewReader(input))
scanner.Split(bufio.ScanWords)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

```

## ioutil.ReadFile()

```go
data, err := ioutil.ReadFile(filename)
for _, line := range strings.Split(string(data), "\n") {
	counts[line]++
}
```

## os.Open()

```go
f = os.Open(filename)
input := bufio.NewScanner(f)
```

## input.Scan() and input.Text()

```go
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
  fmt.Println(input.Text())
}
```