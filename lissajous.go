// lissajous generates gif animations

package main

import (
	"image/color"
	"image/gif"
	"io"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first colour in the palette
	blackIndex = 1 // next colour in the palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out, io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 8
		delay   = 8
	)
	freq := rand.Float64() * 0.3
	anim := gif.GIF{LoopCount: nframes}

}
