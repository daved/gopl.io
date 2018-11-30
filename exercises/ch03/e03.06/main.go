// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		xinc                   = float64(xmax-xmin) / (width * 2)
		yinc                   = float64(ymax-ymin) / (height * 2)
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			ca := mandelbrot(complex(x-xinc, y-yinc))
			cb := mandelbrot(complex(x+xinc, y-yinc))
			cc := mandelbrot(complex(x-xinc, y+yinc))
			cd := mandelbrot(complex(x+xinc, y+yinc))
			c := average(ca, cb, cc, cd)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, c)
		}
	}

	_ = png.Encode(os.Stdout, img) //nolint
}

func average(cs ...color.Color) color.Color {
	var ar, ag, ab, aa uint64
	for _, c := range cs {
		cr, cg, cb, ca := c.RGBA()
		ar += uint64(cr * cr)
		ag += uint64(cg * cg)
		ab += uint64(cb * cb)
		aa += uint64(ca * ca)
	}

	l := float64(len(cs))

	r := uint8(math.Sqrt(float64(ar)/l) / 256)
	g := uint8(math.Sqrt(float64(ag)/l) / 256)
	b := uint8(math.Sqrt(float64(ab)/l) / 256)
	a := uint8(math.Sqrt(float64(aa)/l) / 256)

	return color.RGBA{r, g, b, a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return makeColorByGrayscale(contrast * n)
		}
	}
	return color.Black
}

func makeColorByGrayscale(n uint8) color.RGBA {
	return color.RGBA{
		mapByColorOffset(redOffset, n),
		mapByColorOffset(greenOffset, n),
		mapByColorOffset(blueOffset, n),
		255,
	}
}

type colorOffset uint8

const (
	scaleMax = 255
	center   = 127
	lowerBnd = 84
	upperBnd = 170
	extLower = 42
	extUpper = 212

	redOffset   colorOffset = extLower
	greenOffset             = center
	blueOffset              = extUpper
)

func mapByColorOffset(co colorOffset, n uint8) uint8 {
	diff := center - uint8(co)
	n += diff

	if !inRange(extLower, extUpper, n) {
		return 0
	}
	if inRange(extLower, lowerBnd, n) {
		return mapToFullScale(extLower, lowerBnd, n)
	}
	if inRange(upperBnd, extUpper, n) {
		return mapToFullScale(extUpper, upperBnd, n)
	}

	return scaleMax
}

func inRange(x, y, n uint8) bool {
	lower, upper := x, y
	if lower > upper {
		lower, upper = upper, lower
	}

	return n >= lower && n <= upper
}

func mapToFullScale(min, max, n uint8) uint8 {
	if !inRange(min, max, n) {
		return 0
	}

	mn, mx := min, max
	if mn > mx {
		n = mn - (n - mx)
		mn, mx = mx, mn
	}

	f := float32(n-mn) / float32(mx-mn)

	return uint8(f * float32(scaleMax))
}
