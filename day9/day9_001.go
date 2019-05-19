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
}