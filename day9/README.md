## github.com/codeskyblue/go-sh

```go
package main

import "github.com/codeskyblue/go-sh"

func main() {
    sh.Command("echo", "hello\tworld").Command("cut", "-f2").Run()


    // sh: echo hello
	sh.Command("echo", "hello").Run()

	// sh: export BUILD_ID=123
	// s := sh.NewSession().SetEnv("BUILD_ID", "123")

	// sh: alias ll='ls -l'
	// s = sh.NewSession().Alias('ll', 'ls', '-l')

	// sh: (cd /; pwd)
	sh.Command("pwd", sh.Dir("/")).Run()

	// sh: test -d /tmp/data || mkdir data
	if ! sh.Test("dir", "/tmp/data") { sh.Command("mkdir", "/tmp/data").Run() }

	// sh: cat first second | awk '{print $1}'
	sh.Command("cat", "first", "second").Command("awk", "{print $1}").Run()

	// sh: count=$(echo "one two three" | wc -w)
	// count, err := sh.Echo("one two three").Command("wc", "-w").Output()

	// sh(in ubuntu): timeout 1s sleep 3
	// c := sh.Command("sleep", "3"); c.Start(); c.WaitTimeout(time.Second) # default SIGKILL
	// out, err := sh.Command("sleep", "3").SetTimeout(time.Second).Output() # set session timeout and get output)

	// sh: echo hello | cat
	// out, err := sh.Command("cat").SetInput("hello").Output()

	// sh: cat # read from stdin
	// out, err := sh.Command("cat").SetStdin(os.Stdin).Output()

	// sh: ls -l > /tmp/listing.txt # write stdout to file
	// err := sh.Command("ls", "-l").WriteStdout("/tmp/listing.txt")

    // to keep env and dir
    session := sh.NewSession()
    session.SetEnv("BUILD_ID", "123")
    session.SetDir("/")
    # then call cmd
    session.Command("echo", "hello").Run()
    # set ShowCMD to true for easily debug
    session.ShowCMD = true

    // By default, pipeline returns error only if the last command exit with a non-zero status. However, you can also enable pipefail option like bash. In that case, pipeline returns error if any of the commands fail and for multiple failed commands, it returns the error of rightmost failed command.
    session := sh.NewSession()
    session.PipeFail = true
    session.Command("cat", "unknown-file").Command("echo").Run()
}
```

By default, pipelines's std-error is set to last command's std-error. However, you can also combine std-errors of all commands into pipeline's std-error using session.PipeStdErrors = true.

## os/exec

[Ref1.](https://colobu.com/2017/06/19/advanced-command-execution-in-Go-with-os-exec/)

```go
func main() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
```

Process Stdout and Stderr.

```go
func main() {
	cmd := exec.Command("ls", "-lah")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}
```

Async

```go
cmd := exec.Command("cat", "/dev/random")
randomBytes := &bytes.Buffer{}
cmd.Stdout = randomBytes

// Start command asynchronously
err := cmd.Start()
if err != nil {
  os.Stderr.WriteString(err.Error())
}

ticker := time.NewTicker(time.Second)
go func(ticker *time.Ticker) {
  now := time.Now()
  for _ = range ticker.C {
    fmt.Printf("Ticker: %s\n", []byte(fmt.Sprintf("%s", time.Since(now))))
  }
}(ticker)

// Kill the process after 4 seconds
timer := time.NewTimer(time.Second * 4)
go func(timer *time.Timer, ticker *time.Ticker, cmd *exec.Cmd) {
  for _ = range timer.C {
    err := cmd.Process.Signal(os.Kill)
    if err != nil {
      os.Stderr.WriteString(err.Error())
    }
    ticker.Stop()
  }
}(timer, ticker, cmd)

// Wait for the command to finish
cmd.Wait()
fmt.Printf("Result: %d\n", []byte(fmt.Sprintf("%d bytes generated", len(randomBytes.Bytes()))))
```