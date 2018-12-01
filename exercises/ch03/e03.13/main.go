package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	k  = 1000
	KB = k
	MB = k * KB
	GB = k * MB
	TB = k * GB
	PB = k * TB
	EB = k * PB
	ZB = k * EB
	YB = k * ZB
)

func main() {
	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)
	fmt.Println(KB, MB, GB, TB, PB, EB)
}
