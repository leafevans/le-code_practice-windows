package calc

var a = 100  // 私有变量
var Age = 20 // 公有变量

func Add[T int | float32 | float64](x, y T) T {
	return x + y
}

func Sub[T int | float32 | float64](x, y T) T {
	return x - y
}
