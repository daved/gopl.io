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
	_ = png.Encode(os.Stdout, img) //nolint
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
