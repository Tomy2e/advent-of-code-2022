package common

import "strconv"

func MustAtoi(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return r
}
