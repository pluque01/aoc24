package math

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func EqualSign(x int, y int) bool {
	return (x >= 0 && y >= 0) || (x <= 0 && y <= 0)
}
