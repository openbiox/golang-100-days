package biogolang

import (
	"log"
	"os"
)

type People struct {
	Name   string
	Age    int
	Gender string
}

var Jianfneg string = "110"

var Cwd string

func init() {
	var err error
	Cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

/* Wrong example
var Cwd string

func init() {
	Cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cCwd)
}

*/
