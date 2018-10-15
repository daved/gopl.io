// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.White,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleLissajous)

	if err := http.ListenAndServe(":8383", mux); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func handleLissajous(w http.ResponseWriter, r *http.Request) {
	cs := 3

	if css, ok := r.URL.Query()["cycles"]; ok {
		s, err := strconv.Atoi(css[0])
		if err != nil {
			http.Error(w, "cycles param is malformed", http.StatusBadRequest)
			return
		}
		cs = s
	}

	lissajous(w, cs)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 400   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	if cycles <= 0 {
		cycles = 5
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			colorIndex := uint8(
				(int(math.Abs(t)) % (len(palette) - 2)) + 2,
			)
			img.SetColorIndex(
				size+int(x*size+0.5), size+int(y*size+0.5), colorIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	_ = gif.EncodeAll(out, &anim) //nolint
}
