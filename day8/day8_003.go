// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	// NewRGBA func(r Rectangle) *RGBA
	// NewRGBA returns a new RGBA image with the given bounds.
	// Rect func(x0, y0, x1, y1 int) Rectangle
	// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	// Encode writes the Image m to w in PNG format. Any Image may be encoded, but images that are not image.NRGBA might be encoded lossily.
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
