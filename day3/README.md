## image/color

Image and color process

```go
color.White
color.Black
color.RGBA{0x00, 0x80, 0x00, 0xff}
```

## math/rand

```go
// the pseudo-random number generator using the current time
rand.Seed(time.Now().UTC().UnixNano())
rand.Float64()
```

## io.Writer

```go
// os.Stdout is one of io.Writer
```

## const

Define the constant value in the Golang package or function.

## image/gif

```go
nframes := 64
gif.GIF{LoopCount: nframes}
```

## image

```go
// draw an rectangle
// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. 
// The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.
size := 100
var palette = []color.Color{
	color.White, color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x8B, 0x8B, 0xff},
	color.RGBA{0xFF, 0x45, 0x00, 0xff}}
image.Rect(0, 0, 2*size+1, 2*size+1)

// image.NewPaletted returns a new Paletted image with the given width, height and color palette.
img := image.NewPaletted(rect, palette)
// 2 in palette: color.RGBA{0x00, 0x8B, 0x8B, 0xff}
colorIndex = 2
img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					colorIndex)
```

## append

Process slice object in golang that is similar to Python (e.g. a=[1,2,3]; a.append(4))

```golang
sliceGo := []int{1, 2, 3, 4}
sliceGo = append(sliceGo, 4)
fmt.Println(sliceGo)
```