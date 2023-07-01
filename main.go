package main

func Addint(x, y int) int {
	return x + y
}

func Subint(x, y int) int {
	return x - y
}

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
