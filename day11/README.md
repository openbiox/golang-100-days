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
	extraOption := params.Reffa.GenerateDictExtra
	wd, _ := os.Getwd()
	logPath := fmt.Sprintf("%s/log/%s.GenerateDict.log", wd, reffa.Id)
	var cmd *exec.Cmd

	CreateFileParDir(logPath)
	if err := CreateFileParDir(outDict); err != nil {
		log.Fatalf("Can not create output directory of file %s", outDict)
		return false
	}

	args := []string{"-jar", cfg.Picard, "CreateSequenceDictionary", "R=" + reffa.Path, "O=" + outDict}
	if len(extraOption) >= 1 && extraOption[0] != "" {
		args = MergeSlice(args, extraOption)
	}
	cmd = exec.Command(cfg.Java, args...)

	status := !RunExecCmd(logPath, cmd)
	hasFile, _ := PathExists(outDict)
	if status && !hasFile {
		return false
	}
	return true
}
```
