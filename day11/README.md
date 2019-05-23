## CreateFile

```go
func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
```

## ConnectFile

```go
func ConnectFile(name string) (*os.File, error) {
	var fn *os.File
	var err error
	if hasFile, _ := PathExists(name); !hasFile {
		CreateFile(name)
	}
	fn, err = os.OpenFile(name, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0664)
	return fn, err
}
```

## Advance of log and os/exec

```go
package log

import (
	"fmt"
	"log"
	"path/filepath"
	_ "reflect"
	"runtime"
	"strings"
)

func DEBUG(formating string, args ...interface{}) {
	LOG("DEBUG", formating, args...)
}

func INFO(formating string, args ...interface{}) {
	LOG("INFO", formating, args...)
}

func LOG(level string, formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
		funcname = filepath.Ext(funcname)           // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
	}
	log.SetPrefix("["+ level + "] ")
	log.Printf("%s:%d:%s %s \n", filename, line, funcname, fmt.Sprintf(formating, args...))
}
```

```go
// See my package reffa https://github.com/JhuangLab/libtai/blob/master/reffa/reffa.go
func GenerateDict(reffa *Reffa, outDict string, config *TaiFiles, params *TaiParams) bool {
	INFO("")
	cfg := config.Default
	//extraOption := params.Reffa.GenerateDictExtra
	wd, _ := os.Getwd()
	logPath := fmt.Sprintf("%s/log/%s.generate_dict.log", wd, reffa.Id)
	if isExists, _ := PathExists(path.Dir(logPath)); !isExists {
		os.MkdirAll(path.Dir(logPath), os.ModePerm)
	}
	if isExists, _ := PathExists(path.Dir(outDict)); !isExists {
		os.MkdirAll(path.Dir(outDict), os.ModePerm)
	}
	log.SetPrefix("[BASH] ")
	cmd := exec.Command(cfg.Java, "-jar", cfg.Picard, "CreateSequenceDictionary", "R="+reffa.Path, "O="+outDict)
	var stdout, sterr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &sterr
	logFn, err := ConnectFile(logPath)
	if err != nil {
		return false
	}

	cmdStr := strings.Join(cmd.Args, " ") + " &> " + logPath
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return false
	}
	cmd.Wait()
	log.SetOutput(io.MultiWriter(os.Stderr, logFn))
	log.Println(cmdStr)
	log.Println(sterr.String())
	return true
}
```