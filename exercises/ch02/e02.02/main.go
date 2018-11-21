// cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/daved/gopl.io/ch02/e02.02/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		display(n)
	}
}

func display(n float64) {
	fns := []func(float64){temp, dist, wght}

	for _, fn := range fns {
		fn(n)
	}
}

func printAll(v0, v1, v2, v3 interface{}) {
	fmt.Printf("%s = %s, %s = %s\n", v0, v1, v2, v3)
}

func temp(n float64) {
	f := unitconv.Fahrenheit(n)
	c := unitconv.Celsius(n)
	printAll(f, unitconv.FToC(f), c, unitconv.CToF(c))
}

func dist(n float64) {
	f := unitconv.Feet(n)
	m := unitconv.Meter(n)
	printAll(f, unitconv.FToM(f), m, unitconv.MToF(m))
}

func wght(n float64) {
	p := unitconv.Pound(n)
	k := unitconv.Kilogram(n)
	printAll(p, unitconv.PToK(p), k, unitconv.KToP(k))
}
