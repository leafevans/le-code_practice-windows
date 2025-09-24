package tools

func Mul[T int | float32 | float64](x, y T) T {
	return x * y
}
