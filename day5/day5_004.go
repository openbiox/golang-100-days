// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// if access /count see 0, then access / and then /count is 3
	// access /count will activiate two goroutine for handler and counter
	// so see /count is 0, the real value of count is 1 now
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if len(r.Form["cycles"]) > 0 {
		cycles, err := strconv.ParseFloat(r.Form["cycles"][0], 64)
		if err != nil {
		}
		lissajous(w, cycles)
	}

}

var palette = []color.Color{
	color.White, color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x8B, 0x8B, 0xff},
	color.RGBA{0xFF, 0x45, 0x00, 0xff}}

// define constant in this package
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// io.Writer for various output stream
func lissajous(out io.Writer, cycles float64) {
	// define constant in this function
	fmt.Println(cycles)
	if cycles == 0 {
		cycles = 5.0
	}
	fmt.Println(cycles)
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	// gif.GIF get struct object
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// set palette
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// write anim to stand output stream
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
