// package main

// import "fmt"

// func main() {
// 	defer fmt.Println("Hello World!")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	fmt.Println(3)
// 	fmt.Println("Hello Golang!")
// }

// package main

// import "fmt"

// func fn() {
// 	fmt.Println("开始！")

// 	defer func() {
// 		fmt.Println("Hachimi")
// 		fmt.Println("Mambo")
// 	}()

// 	fmt.Println("结束！")
// }

// func fn2() int {
// 	var a int

// 	defer func() {
// 		a++
// 	}()

// 	fmt.Println("结束！")

// 	return a
// }

// func fn3() (a int) {
// 	defer func(a int) {
// 		fmt.Println(a)
// 		a++
// 	}(a)

// 	fmt.Printf("a = %d\n", a)

// 	return 9
// }

// func main() {
// 	fmt.Println(fn2())
// 	fmt.Println(fn3())
// }

// package main

// import "fmt"

// func calc(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

// func main() {
// 	x := 1
// 	y := 2
// 	defer calc("AA", x, calc("A", x, y))
// 	x = 10
// 	defer calc("BB", x, calc("B", x, y))
// 	y = 20
// }

// package main

// import "fmt"

// func fnA() {
// 	fmt.Println("fnA")
// }

// func fnB(a, b int) {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("err:", err)
// 		}
// 	}()
// 	fmt.Println(a / b)
// }

// func main() {
// 	fnB(10, 0)
// 	fmt.Println("Mambo")
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func readFile(filename string) error {
// 	if filename == "main.go" {
// 		return nil
// 	} else {
// 		return errors.New("读取文件失败")
// 	}
// }

// func fn() {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("给管理员发送邮件")
// 		}
// 	}()
// 	err := readFile("xxx.go")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {
// 	fn()
// 	fmt.Println("继续执行")
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Start!")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)
// 	fmt.Println("结束！")
// }

// package main

// import "fmt"

// func f1() int {
// 	// 局部变量 x = 5。
// 	x := 5

// 	// defer 执行时 x 变为 6，但不影响返回值。
// 	defer func() {
// 		x++
// 	}()

// 	// 返回 x 的值为 5。
// 	return x
// }

// // x 是命名返回值，初始为 0。
// func f2() (x int) {
// 	// defer 执行时，x 变为 6。
// 	defer func() {
// 		x++
// 	}()

// 	// 设置 x = 5，然后 defer 修改为 6。
// 	return 5
// }

// // y 是命名返回值，初始为 0。
// func f3() (y int) {
// 	// 局部变量 x = 5。
// 	x := 5

// 	// defer 修改局部变量 x 为 6，不影响 y。
// 	defer func() {
// 		x++
// 	}()

// 	// 设置 y = x，此时 x 为 5。
// 	return x
// }

// // x 是命名返回值，初始为 0。
// func f4() (x int) {
// 	// defer 接受 x 的副本（0）。
// 	defer func(x int) {
// 		// 修改副本为 1，不影响原 x。
// 		x++
// 	}(x)

// 	// 设置 x = 5。
// 	return 5
// }

// func main() {
// 	fmt.Println(f1())
// 	fmt.Println(f2())
// 	fmt.Println(f3())
// 	fmt.Println(f4())
// }

// package main

// import "fmt"

// func calc(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

// func main() {
// 	x := 1
// 	y := 1
// 	defer calc("AA", x, calc("A", x, y))
// 	x = 10
// 	defer calc("BB", x, calc("B", x, y))
// 	y = 20
// }

// package main

// import "fmt"

// func fnA() {
// 	fmt.Println("fn A")
// }

// func fnB() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("recover in B")
// 		}
// 	}()
// 	panic("panic in B")
// }

// func fnC() {
// 	fmt.Println("fn C")
// }

// func main() {
// 	fnA()
// 	fnB()
// 	fnC()
// }

package main

import "fmt"

func fn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("抛出异常给管理员发送邮件。")
			fmt.Println(err)
		}
	}()

	a := 10
	b := 0
	res := a / b
	fmt.Println("res =", res)
}

func main() {
	fn()
}
