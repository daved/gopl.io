package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Shift ...
func Shift(x uint64) int {
	var acc int
	for i := uint64(0); i < 64; i++ {
		if x&(1<<i) > 0 {
			acc++
		}
	}
	return acc
}

// Loop ...
func Loop(x uint64) int {
	var acc byte
	for i := uint64(0); i < 8; i++ {
		acc += pc[byte(x>>(i*8))]
	}
	return int(acc)
}

// Long ...
func Long(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
