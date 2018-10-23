// Package unitconv performs miscellaneous unit conversions.
package unitconv

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

type Kelvin float64

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

type Feet float64

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

type Meter float64

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

type Pound float64

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

type Kilogram float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func FToM(f Feet) Meter {
	return Meter(f * .3048)
}

func MToF(m Meter) Feet {
	return Feet(m / .3048)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p * .45359237)
}

func KToP(k Kilogram) Pound {
	return Pound(k / .45359237)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)
