// package main

// import "fmt"

// type Human struct {
// 	Name string
// 	Age  int
// }

//	func main() {
//		people := Human{Name: "张三", Age: 18}
//		fmt.Printf("%v %+v %#v %T %%\n", people, people, people, people)
//		fmt.Printf("%b", 5)
//		fmt.Printf("%c", 0x4E2D)
//		fmt.Printf("%d", 0b101)
//		fmt.Printf("%o", 0o122)
//		fmt.Printf("%s", "He said \"Hi\"")
//	}
// package main

// import "fmt"

// func main() {
// 	s := "Hello, 世界"

//		fmt.Printf("[%10.5s]\n", s)  // [     Hell]
//		fmt.Printf("[%-10.5s]\n", s) // [Hell      ]
//		fmt.Printf("[%010.5s]\n", s) // [00000Hell]  ← 注意：0 填充对字符串通常不推荐，但语法允许
//	}
// package main

// import "fmt"

// type Human struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	people := Human{"哈基米", 18}
// 	fmt.Printf("%v\n", people)
// 	fmt.Printf("%+v\n", people)
// 	fmt.Printf("%#v\n", people)
// 	fmt.Printf("%T\n", people)
// 	fmt.Printf("%%\n")
// 	fmt.Println()

// 	fmt.Printf("%t %t\n", true, false)
// 	fmt.Println()

// 	fmt.Printf("%b\n", 22)
// 	fmt.Printf("%c\n", 0x4E2D)
// 	fmt.Printf("%d\n", 0x12)
// 	fmt.Printf("%o\n", 10)
// 	fmt.Printf("%q\n", 0x4E2D)
// 	fmt.Printf("%x\n", 22)
// 	fmt.Printf("%X\n", 22)
// 	fmt.Printf("%U\n", '😂')
// 	fmt.Println()

// 	fmt.Printf("%b\n", 1.0)
// 	fmt.Printf("%e\n", 10.2)
// 	fmt.Printf("%.2f\n", 10.22321312)
// 	fmt.Printf("%.1E\n", 3.1415926)
// 	fmt.Printf("%g\n", 1.23123)
// 	fmt.Printf("%.6G\n", 123.231231)
// 	fmt.Println()

// 	fmt.Printf("%s\n", "Go 语言")
// 	fmt.Printf("%q\n", "Go 语言")
// 	fmt.Printf("%x\n", "你好")
// 	fmt.Printf("%X\n", "Golang")
// 	fmt.Println()

//		fmt.Printf("%p\n", &people)
//		fmt.Println([]byte("你好"))
//	}
// package main

// import "fmt"

// func main() {
// 	var username = "哈基米"
// 	var age int = 20
// 	fmt.Println(username, age)

//		var a, b, c, d = 1, 2, 3, 4
//		fmt.Println(a, b, c, d)
//	}
// package main

// import "fmt"

//	func main() {
//		var (
//			a string = "哈基米"
//			b int    = 120
//			c bool   = true
//		)
//		fmt.Println(a, b, c)
//		var name, age = "叮咚鸡", 20
//		fmt.Println(name, age)
//	}
// package main

// import "fmt"

// var m = 100

//	func main() {
//		n := 10
//		m := 200
//		fmt.Println(n, m)
//		a, b := "哈基米", 15
//		fmt.Println(a, b)
//	}
// package main

// import "fmt"

// const (
// 	Pi = 3.1415926
// 	E  = 2.1782
// )

// func Info() (int, string) {
// 	return 18, "魔丸"
// }

//	func main() {
//		x := 0
//		fmt.Println(x)
//		x, y := Info()
//		fmt.Println(x, y, Pi, E)
//	}
// package main

// import "fmt"

// const (
// 	A = 100
// 	B
// 	C
// 	D = 200
// 	E
// )

//	func main() {
//		fmt.Println(A, B, C, D, E)
//	}
// package main

// import (
// 	"fmt"
// 	"math"
// )

//	func main() {
//		const Pi = math.Pi
//		const E = math.E
//		fmt.Println(Pi)
//		fmt.Println(E)
//	}
// package main

// import "fmt"

// const A = iota

// const (
// 	B = iota
// 	_
// 	C
// 	D = 100
// 	E = iota
// )

// const F = iota

//	func main() {
//		fmt.Println(A, B, C, D, E, F)
//	}
package main

import "fmt"

const (
	A, B = iota + 1, iota + 2
	C, D
	E, F
	H = 10
)

func main() {
	fmt.Println(A, B, C, D, E, F)
}
