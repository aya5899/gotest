package compare

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if y < x {
		return y
	}
	return x
}
