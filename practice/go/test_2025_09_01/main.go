// package main

// import "fmt"

// import "fmt"

// func main() {
// 	num := 10.2 + 5.2i
// 	fmt.Printf("%f\n", num)
// }

/*
变量声明格式:
var 变量名 类型
*/

// var g = "全局变量"

// func getUserInfo() (string, int) {
// 	return "张三", 10
// }

// func main() {
// 	var username, _ = getUserInfo()
// 	fmt.Println(username)
// 	// fmt.Println(g)
// 	// var num int = 10
// 	// // 变量声明没有初始化，值为空
// 	// var username string
// 	// var a = "zhangsan"
// 	// username = "张三"
// 	// fmt.Println(num, username)
// 	// fmt.Println(a)
// 	// // 定义变量见名思义
// 	// var m_ = "李四"
// 	// fmt.Println(m_)
// 	// 第一种初始化变量的方式
// 	// var username string
// 	// var num int = 10
// 	// username = "哈基米"
// 	// fmt.Println(username, num)
// 	// 第二张初始化变量的方式
// 	// var username string = "曼波"
// 	// fmt.Println(username)
// 	// 自动类型推导
// 	// var username = "张三"
// 	// fmt.Println(username)

// 	// var username = "张三"
// 	// fmt.Println(username)
// 	// var age = 20
// 	// var sex = "男"
// 	// fmt.Println(username, age, sex)
// 	// username = "哈基米"
// 	// age = 10
// 	// sex = "牠"
// 	// fmt.Println(username, age, sex)

// 	// 一次生成多个变量
// 	// var (
// 	// 	username string = "哈基米"
// 	// 	age      int    = 20
// 	// 	sex      string = "牠"
// 	// )
// 	// fmt.Println(username, age, sex)
// 	// 短变量声明法
// 	// username := "张三"
// 	// fmt.Println(username)
// 	// fmt.Printf("Type: %T\n", username)
// 	// num := 10 + 10i
// 	// fmt.Printf("Type: %T\n", num)\
// 	// 短变量声明法只能在方法内使用
// 	// a, b, c := 12, 13, "c"
// 	// fmt.Println(a, b, c)
// }
// 匿名变量不占用命名空间，不会分配内存，所以匿名变量不存在重复声明

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var username = "张三"
// 	var age int = 20
// 	fmt.Println(username, age)
// }

// package main

// import "fmt"

// func main() {
// 	// var (
// 	// 	a string = "张三"
// 	// 	b int    = 10
// 	// 	c bool   = true
// 	// )
// 	// fmt.Println(a, b, c)
// 	var (
// 		a string
// 		b int
// 		c bool
// 	)
// 	a = "哈基米"
// 	b = 10
// 	c = true
// 	fmt.Println(a, b, c)
// }

// package main

// import "fmt"

// func getInfo() (string, int) {
// 	return "胖宝宝", 100
// }

// func main() {
// 	username, _ := getInfo()
// 	fmt.Println(username)
// }

// package main

// import "fmt"

// func main() {
// 	// var username = "哈基米"
// 	// username = "胖宝宝"
// 	// const PI = 3.1415926
// 	// fmt.Println(PI)
// 	// const (
// 	// 	A = "A"
// 	// 	B = 10
// 	// )
// 	// fmt.Println(A, B)
// 	// const (
// 	// 	Num = 10
// 	// 	Count
// 	// 	C
// 	// )
// 	// fmt.Println(Num, Count, C)
// 	// const A = iota
// 	// const (
// 	// 	B = iota
// 	// 	C
// 	// 	D
// 	// )
// 	// fmt.Println(A, B, C, D)
// 	// const (
// 	// 	A = iota
// 	// 	B = 100
// 	// 	C = iota
// 	// 	D
// 	// )
// 	// fmt.Println(A, B, C, D)
// 	const (
// 		// 1 2
// 		A, B = iota + 1, iota + 2
// 		// 2 3
// 		C, D
// 		// 3 4
// 		E, F
// 	)
// 	fmt.Println(A, B)
// 	fmt.Println(C, D)
// 	fmt.Println(E, F)
// }

// package main

// import "fmt"

// const (
// 	A = 100
// 	B
// 	C
// 	D = 200
// 	E
// )

// func main() {
// 	fmt.Println(A, B, C, D, E)
// }

// package main

// import (
// 	"fmt"
// )

// const (
// 	A = iota
// 	B
// 	_
// 	C
// 	D
// )

// func main() {
// 	fmt.Println(A, B, C, D)
// }

// package main

// import "fmt"

// const (
// 	A = iota
// 	B = 100
// 	C
// 	D = iota
// 	E
// )
// const F = iota

// func main() {
// 	fmt.Println(A, B, C, D, E, F)
// }

// package main

// import "fmt"

// const (
// 	B  = 1 << (10 * iota)
// 	KB // 2^10
// 	MB // 2^20
// 	GB // 2^30
// )

// func main() {
// 	fmt.Println(B, KB, MB, GB)
// }

// package main

// import "fmt"

// const (
// 	Read = 1 << iota
// 	Write
// 	Exec
// 	Admin
// )

//	func main() {
//		fmt.Println(Read | Write)
//	}
// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	// num := 10

// 	// fmt.Printf("%v %T\n", num, num)
// 	// -128~127
// 	// var num int8 = 127
// 	// fmt.Println(num)

// 	// fmt.Println(unsafe.Sizeof(-1230))
// 	// var a int8 = 15
// 	// fmt.Printf("num = %v, type: %T\n", a, a)
// 	// 占用 1 个字节
// 	// fmt.Println(unsafe.Sizeof(a))
// 	var a int32 = 15
// 	fmt.Println(unsafe.Sizeof(a))
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	var a uint8 = 1
// 	var b uint64 = 1
// 	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
// }

// package main

// import "fmt"

// func main() {
// 	// var a1 int32 = 10
// 	// var a2 int64 = 21
// 	// // fmt.Println(int64(a1) + a2)
// 	// fmt.Println(a1 + int32(a2))
// 	// var a1 int16 = 130
// 	// fmt.Println(int8(a1))
// 	num := 0b010100101
// 	fmt.Printf("%X %T\n", num, num)
// }

// package main

// import "fmt"

// func main() {
// 	var num int64
// 	num = 123
// 	fmt.Printf("值：%v，类型：%T", num, num)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	var a int8 = 120
// 	fmt.Printf("%T\n", a)
// 	fmt.Println(unsafe.Sizeof(a))
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	var a int8 = 120
// 	fmt.Printf("%T\n", a)
// 	fmt.Println(unsafe.Sizeof(a))
// }

// package main

// import "fmt"

// func main() {
// 	var num1 int8 = 127
// 	num2 := int32(num1)
// 	fmt.Printf("值：%v，类型：%T", num2, num2)
// }

// package main

// import "fmt"

// const (
// 	A = 10
// 	B = iota
// 	C
// )

// func main() {
// 	// 十进制
// 	var a int = 10
// 	fmt.Printf("%d\n", a)
// 	// 占位符 %b 表示二进制
// 	fmt.Printf("%b\n", a)

// 	// 八进制以 0 开头（或 0o）
// 	var b int = 077
// 	fmt.Printf("%o\n", b)

// 	// 十六进制以 0x 开头
// 	var c int = 0xff
// 	fmt.Printf("%x\n", c)
// 	fmt.Printf("%X\n", c)
// 	fmt.Printf("%d\n", c)

// 	fmt.Println(A, B, C)
// }

// package main

// import "fmt"

// func main() {
// 	// num1 := 0b0_0_1_0_1_101
// 	// fmt.Println(num1)

// 	// var num2 = 0o377
// 	// fmt.Printf("%o\n", num2)

// 	// var num3 = 0x2e - 2
// 	// fmt.Printf("%v %T", num3, num3)
// 	num1 := 2.5e-3
// 	fmt.Printf("%v %T", num1, num1)
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("|%10d|\n", 42)
// 	fmt.Printf("|%-10d|\n", 42)
// 	fmt.Printf("|%010d|\n", 42)
// 	fmt.Printf("|%010s|\n", "Go")
// 	fmt.Printf("%.2f", 3.14159)
// 	fmt.Printf("%.3s\n", "golang")
// 	fmt.Printf("%.4d\n", 42)
// 	fmt.Printf("%.3g\n", 12345.678)
// 	fmt.Printf("|%8.2f|\n", 3.14)
// 	fmt.Printf("|%08.2f|\n", 3.14)
// 	fmt.Printf("|%6.7s|\n", "golang")
// 	fmt.Printf("|%6.4d|\n", 42)
// }

// package main

// import "fmt"

// func main() {
// 	// var num float64 = 3.1415926
// 	// fmt.Printf("%v - %f", num, num)
// 	// var f2 = 3.14e-2
// 	// fmt.Println(f2)
// 	// m1 := 8.2_0
// 	// m2 := 3.8_0
// 	// fmt.Println(m1 - m2)
// 	// var num float64 = 1129.6
// 	// fmt.Println(num * 1e2)
// 	// num := decimal.NewFromFloat(1129.6)
// 	// result := num.Mul(decimal.NewFromFloat(1e2))
// 	// fmt.Println(result)
// 	// num2 := decimal.RequireFromString("1129.6")
// 	// result2 := num2.Mul(decimal.NewFromInt(100))
// 	// fmt.Println(result2)
// 	// var num float64 = 8.2
// 	// var num2 float64 = 3.8
// 	// result := decimal.NewFromFloat(num).Sub(decimal.NewFromFloat(num2))
// 	// fmt.Printf("%v %T\n", result, result)
// 	a := 10
// 	b := float64(a)
// 	fmt.Printf("a 的类型是 %T，b 的类型是 %T\n", a, b)
// 	var c float32 = 23.45
// 	d := int(c)
// 	fmt.Println(d)
// }

// package main

// import "fmt"

// func main() {
// 	var flag bool = true
// 	fmt.Printf("%v %T\n", flag, flag)
// 	var str string
// 	fmt.Printf("%v\n", str)
// 	var count int
// 	fmt.Printf("%d\n", count)
// 	var num float32
// 	fmt.Printf("%v\n", num)
// 	if flag {
// 		fmt.Println("Hello World!")
// 	} else {
// 		fmt.Println("Hello Gemini!")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	num := 5.1234e2
// 	num2 := 5.1234e2
// 	num3 := 5.1234e-2
// 	fmt.Printf("num = %v, num2 = %v, num3 = %v", num, num2, num3)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	var b = true
// 	fmt.Printf("%v 占用字节：%v", b, unsafe.Sizeof(b))
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// num := 131
// 	// fmt.Printf("%+#X", num)
// 	fmt.Println(math.MaxFloat32)
// 	fmt.Println(math.MaxFloat64)
// }

package main

import "fmt"

func main() {
	flag := true
	if flag {
		fmt.Println("Hello World!")
	}
}
