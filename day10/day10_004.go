package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func main() {
	type song struct {
		Name     string
		Duration duration
	}
	type songs struct {
		Song []song
	}
	var favorites songs
	if _, err := toml.DecodeFile("config.toml", &favorites); err != nil {
		log.Fatal(err)
	}

	for _, s := range favorites.Song {
		fmt.Printf("%s (%s)\n", s.Name, s.Duration)
	}
}
