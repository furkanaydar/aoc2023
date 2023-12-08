package utils

func GCD(a int, b int) int {
	if a == 0 {
		return b
	}

	return GCD(b%a, a)
}

func LCM(a int, b int) int {
	return a * b / LCM(a, b)
}
