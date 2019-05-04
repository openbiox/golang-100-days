# Day 1

## Loop

```go
for initialization; condition; post {
  // zero or more statements
}

// a traditional "while" loop
for condition {
// ...
}

// a traditional infinite loop
for {
// ...
}
```

## Benchmark test

*_test.go

```go
func BenchmarkFuncName(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// expression
	}
}
```
