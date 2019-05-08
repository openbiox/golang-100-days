## net/http

```golang
http.HandleFunc("/", handler)

// log.Fatal is equivalent to Print() followed by a call to os.Exit(1).
log.Fatal(http.ListenAndServe("localhost:8989", nil))


// Simple api handler
// Two tpyes: request and response
// gif.EncodeAll can use w as io.Writer
func handler(w http.ResponseWriter, r *http.Request) {
	// r contains various useful value and function
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
// several r *http.Request values:
// r.URL.Path
// r.Header
// r.Host
// r.Method: GET/POST
// r.ParseForm(): required for getting r.Form key-value pair
```

## sync

```golang
// type Mutex struct {
//   state    int32
//   sema    uint32
// }
// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
// A Mutex must not be copied after first use.
//
// To avoid different goroutine concorrent modified the count value
var mu sync.Mutex

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```

## switch

```golang
// just like in c/c++ switch
// another choice for if else-if and else
for k, v := range r.Form {
	switch k {
	case "num":
		{
			val, _ := strconv.Atoi(v[0])
			if val%2 == 0 {
				fmt.Fprintf(w, "%d is not odd\n", val)
			} else {
				fmt.Fprintf(w, "%d is odd\n", val)
			}
		}
	default:
		fmt.Println("landed on edge!")
	}
}
```
