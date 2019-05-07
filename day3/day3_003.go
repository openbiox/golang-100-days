// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// []color.Color get color slice object
// color.White is the constant var in image/color
// get RGB color color.RGBA{0xRR, 0xGG, 0xBB, 0xff}
// Green : color.RGBA{0x00, 0x80, 0x00, 0xff}
var palette = []color.Color{
	color.White, color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x8B, 0x8B, 0xff},
	color.RGBA{0xFF, 0x45, 0x00, 0xff}}

// define constant in this package
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

// io.Writer for various output stream
func lissajous(out io.Writer) {
	// define constant in this function
	const (
		cycles  = 5     // number of complete x oscillator revolutions
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
			// fill different color
			if t < 1 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					2)
			} else if t < 2 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					3)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					2)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// write anim to stand output stream
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
