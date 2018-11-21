package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.HandleFunc("/", handle)

	if err := http.ListenAndServe(":8383", mux); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	if _, err := writeSVG(buf); err != nil {
		http.Error(w, "cannot write svg", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	if _, err := w.Write(buf.Bytes()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func writeSVG(w io.Writer) (int, error) {
	n, err := fmt.Fprintf(w,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>",
		width, height,
	)
	if err != nil {
		return n, err
	}

	hFn := swellcone
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			color := "blue"

			ax, ay, aOver := corner(hFn, i+1, j)
			bx, by, bOver := corner(hFn, i, j)
			cx, cy, cOver := corner(hFn, i, j+1)
			dx, dy, dOver := corner(hFn, i+1, j+1)

			if !isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			if aOver || bOver || cOver || dOver {
				color = "red"
			}

			jn, jerr := fmt.Fprintf(w,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"style='fill:%s;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color,
			)
			n += jn
			if jerr != nil {
				return n, err
			}
		}
	}

	xn, err := fmt.Fprintln(w, "</svg>")
	n += xn
	return n, err
}

type heightFunc func(x, y float64) float64

func corner(fn heightFunc, i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := fn(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z > 0
}

func swellcone(x, y float64) float64 {
	r := math.Atan2(x, y) // distance from (0,0)
	return (math.Sin(r) / r) - .25
}

func topfin(x, y float64) float64 {
	r := math.Atan2(-x, y) // distance from (0,0)
	return ((math.Cbrt(r) / r) / 8) - .1
}

func standard(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func isValid(fs ...float64) bool {
	for _, f := range fs {
		if math.IsInf(f, 0) || math.IsNaN(f) {
			return false
		}
	}

	return true
}
