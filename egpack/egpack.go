package egpack

// Reverse reverses the string (as runes)
func Reverse(s string) string {
	r := []rune(s)
	// deliberately poor and unidiomatic swapper
	for i := 0; i < len(r)/2; i++ {
		j := (len(r) - 1) - i
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
