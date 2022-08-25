package acdb

func Acdb(a, b int) int {
	return a + b
}

func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
