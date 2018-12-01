package hashdiff

func differences(a, b [32]byte) int {
	var c int

	for i, v := range a {
		c += setBits(v ^ b[i])
	}

	return c
}

func setBits(b byte) int {
	var c int

	for b > 0 {
		b = b & (b - 1)
		c++
	}

	return c
}
