package common

// Abs returns the absolute value of v.
func Abs(v int) int {
	if v < 0 {
		return v * -1
	}

	return v
}
