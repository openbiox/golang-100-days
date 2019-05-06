package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
)

func main() {
	const (
		cycles     = 5     // number of complete x oscillator revolutions
		res        = 0.001 // angular resolution
		size       = 100   // image canvas covers [-size..+size]
		nframes    = 64    // number of animation frames
		delay      = 8     // delay between frames in 10ms units
		whiteIndex = 0     // first color in palette
		blackIndex = 1     // next color in palette
	)

	// go run day3_002.go see debug for day3_001.go
	fmt.Println(rand.Float64())

	var palette = []color.Color{color.White, color.Black}
	fmt.Println(palette)
	// test := rand.Seed(time.Now().UTC().UnixNano())
	// fmt.Println(test)
	// will cause error: "rand.Seed(time.Now().UTC().UnixNano()) used as value"

	anim := gif.GIF{LoopCount: nframes}
	fmt.Println(anim)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	phase := 0.0                 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
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
	fmt.Println(anim)

	// get RGB color color.RGBA{0xRR, 0xGG, 0xBB, 0xff}
	fmt.Println(color.RGBA{0x00, 0x80, 0x00, 0xff})

	// test append
	sliceGo := []int{1, 2, 3, 4}
	sliceGo = append(sliceGo, 4)
	fmt.Println(sliceGo)
}
