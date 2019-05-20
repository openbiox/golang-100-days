## encoding/json

```go
// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import "encoding/json"
import "fmt"
import "os"

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
type response1 struct {
    Page   int
    Fruits []string
}
type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    // First we'll look at encoding basic data types to
    // JSON strings. Here are some examples for atomic
    // values.
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // And here are some for slices and maps, which encode
    // to JSON arrays and objects as you'd expect.
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // The JSON package can automatically encode your
    // custom data types. It will only include exported
    // fields in the encoded output and will by default
    // use those names as the JSON keys.
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // You can use tags on struct field declarations
    // to customize the encoded JSON key names. Check the
    // definition of `response2` above to see an example
    // of such tags.
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // Now let's look at decoding JSON data into Go
    // values. Here's an example for a generic data
    // structure.
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // We need to provide a variable where the JSON
    // package can put the decoded data. This
    // `map[string]interface{}` will hold a map of strings
    // to arbitrary data types.
    var dat map[string]interface{}

    // Here's the actual decoding, and a check for
    // associated errors.
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // In order to use the values in the decoded map,
    // we'll need to cast them to their appropriate type.
    // For example here we cast the value in `num` to
    // the expected `float64` type.
    num := dat["num"].(float64)
    fmt.Println(num)

    // Accessing nested data requires a series of
    // casts.
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // We can also decode JSON into custom data types.
    // This has the advantages of adding additional
    // type-safety to our programs and eliminating the
    // need for type assertions when accessing the decoded
    // data.
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // In the examples above we always used bytes and
    // strings as intermediates between the data and
    // JSON representation on standard out. We can also
    // stream JSON encodings directly to `os.Writer`s like
    // `os.Stdout` or even HTTP response bodies.
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
```

## github.com/go-yaml/yaml

YAML support for the Go language.

```golang
package main

import (
        "fmt"
        "log"

        "gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
        A string
        B struct {
                RenamedC int   `yaml:"c"`
                D        []int `yaml:",flow"`
        }
}

func main() {
        t := T{}
    
        err := yaml.Unmarshal([]byte(data), &t)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- t:\n%v\n\n", t)
    
        d, err := yaml.Marshal(&t)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- t dump:\n%s\n\n", string(d))
    
        m := make(map[interface{}]interface{})
    
        err = yaml.Unmarshal([]byte(data), &m)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- m:\n%v\n\n", m)
    
        d, err = yaml.Marshal(&m)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
```
## github.com/go-ini/ini

Package ini provides INI file read and write functionality in Go.

```ini
# possible values : production, development
app_mode = development

[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = /home/git/grafana

[server]
# Protocol (http or https)
protocol = http

# The http port  to use
http_port = 9999

# Redirect to correct domain if host header does not match domain
# Prevents DNS rebinding attacks
enforce_domain = true
```

```go
package main

import (
    "fmt"
    "os"

    "gopkg.in/ini.v1"
)

func main() {
    cfg, err := ini.Load("my.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

    // 典型读取操作，默认分区可以使用空字符串表示
    fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
    fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

    // 我们可以做一些候选值限制的操作
    fmt.Println("Server Protocol:",
        cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
    // 如果读取的值不在候选列表内，则会回退使用提供的默认值
    fmt.Println("Email Protocol:",
        cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

    // 试一试自动类型转换
    fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
    fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))
    
    // 差不多了，修改某个值然后进行保存
    cfg.Section("").Key("app_mode").SetValue("production")
    cfg.SaveTo("my.ini.local")
}
```

## github.com/BurntSushi/toml

TOML parser for Golang with reflection.

```toml
[[song]]
name = "Thunder Road"
duration = "4m49s"

[[song]]
name = "Stairway to Heaven"
duration = "8m03s"
```

```go
type song struct {
  Name     string
  Duration duration
}
type songs struct {
  Song []song
}
var favorites songs
if _, err := toml.Decode(blob, &favorites); err != nil {
  log.Fatal(err)
}

for _, s := range favorites.Song {
  fmt.Printf("%s (%s)\n", s.Name, s.Duration)
}
```