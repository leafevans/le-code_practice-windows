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

// package main

// import "fmt"

// func fn() {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("抛出异常给管理员发送邮件。")
// 			fmt.Println(err)
// 		}
// 	}()

// 	a := 10
// 	b := 0
// 	res := a / b
// 	fmt.Println("res =", res)
// }

// func main() {
// 	fn()
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func readFile(filename string) error {
// 	if filename == "main.go" {
// 		return nil
// 	}

// 	return errors.New("读取文件错误。")
// }

// func fn() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("抛出异常给管理员发送邮件。")
// 		}
// 	}()

// 	err := readFile("xxx.go")

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("继续执行！")
// }

// func main() {
// 	fn()
// 	fmt.Println("结束！")
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func readFile(filename string) error {
// 	if filename == "main.go" {
// 		return nil
// 	}

// 	return errors.New("读取文件错误。")
// }

// func fn() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("抛出异常给管理员发送邮件。")
// 		}
// 	}()

// 	err := readFile("xxx.go")

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("继续执行！")
// }

// func main() {
// 	fn()
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()

// 	year := now.Year()
// 	mouth := now.Month()
// 	day := now.Day()
// 	hour := now.Hour()
// 	minute := now.Minute()
// 	second := now.Second()

// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d",
// 		year, mouth, day,
// 		hour, minute, second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	// str := now.Format("2006/01/02 03:04:05")
// 	// fmt.Println(str)

// 	unixTime := now.Unix()
// 	fmt.Println(unixTime)

// 	unixMilli := now.UnixMilli()
// 	fmt.Println(unixMilli)

// 	unixNano := now.UnixNano()
// 	fmt.Println(unixNano)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	unixTime := 1758343396
// 	now := time.Unix(int64(unixTime), 0)
// 	str := now.Format("2006-01-02 15:04:05")
// 	fmt.Println(str)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	str := "2025-09-20 15:38:04"

// 	now, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)

// 	fmt.Println(now.Unix())
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// fmt.Println(time.Millisecond)
// 	// fmt.Println(time.Microsecond)
// 	// fmt.Println(time.Second)

// 	now := time.Now()

// 	fmt.Println(now.Hour())

// 	now = now.Add(time.Hour)

// 	fmt.Println(now.Hour())
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	n := 5

// 	for t := range ticker.C {
// 		n--
// 		if n == 0 {
// 			ticker.Stop()
// 			break
// 		}
// 		fmt.Println(t)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("aaa")
// 	time.Sleep(time.Second)
// 	fmt.Println("aaa2")
// 	time.Sleep(time.Second * 5)
// 	fmt.Println("aaa3")
// 	time.Sleep(time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// 获取当前时间。
// 	now := time.Now()
// 	fmt.Printf("Current time: %v\n", now)

// 	year := now.Year()     // 年
// 	month := now.Month()   // 月
// 	day := now.Day()       // 天
// 	hour := now.Hour()     // 时
// 	minute := now.Minute() // 分
// 	second := now.Second() // 秒

// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
// 		year, month, day,
// 		hour, minute, second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()

// 	// 24 小时制
// 	fmt.Println(now.Format("2006-01-02 15:04:05"))

// 	// 12 小时制
// 	fmt.Println(now.Format("2006-01-02 03:04:05"))

// 	fmt.Println(now.Format("2006/01/02 15:04"))
// 	fmt.Println(now.Format("15:04 2006/01/02"))
// 	fmt.Println(now.Format("2006/01/02"))
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// 获取当前时间。
// 	now := time.Now()
// 	unixTime := now.Unix()     // 时间戳
// 	unixNano := now.UnixNano() // 纳米时间戳

// 	fmt.Printf("Current Unix time: %v\n", unixTime)
// 	fmt.Printf("Current Unix nano: %v\n", unixNano)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func unixToTime(timestamp int64) {
// 	// 将时间戳转换为时间格式。
// 	timeObj := time.Unix(timestamp, 0)
// 	// 对时间进行格式化处理。
// 	str := timeObj.Format("2006-01-02 15:04:05")
// 	fmt.Println(str)
// }

// func main() {
// 	unixToTime(1758348998)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	str := "2025-09-20 14:24:00"
// 	timeTemple := "2006-01-02 15:04:05"
// 	timestamp, _ := time.ParseInLocation(timeTemple, str, time.Local)
// 	fmt.Println(timestamp.Unix())
// 	fmt.Printf("%T\n", time.Second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	later := now.Add(time.Hour)
// 	fmt.Println(now.Sub(later))
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	n := 0

// 	for v := range ticker.C {
// 		fmt.Println(v)
// 		n++
// 		if n > 5 {
// 			ticker.Stop()
// 			return
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	for {
// 		time.Sleep(time.Second)
// 		fmt.Println("我在定时执行任务！")
// 	}
// }

// package main

// import "fmt"

// func add(nums ...int) int {
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }

// func calc(x, y int) (int, int) {
// 	sum := x + y
// 	diff := x - y
// 	return sum, diff
// }

// func main() {
// 	fmt.Println(add())
// 	fmt.Println(add(10))
// 	fmt.Println(add(10, 20, 30))
// }

// package main

// func calc(x, y int) (sum, diff int) {
// 	sum = x + y
// 	diff = x - y
// 	return
// }

// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// type calc func(int, int) int

// func main() {
// 	var c calc = add
// 	fmt.Printf("Type of c: %T\n", c)
// 	fmt.Println(c(1, 2))

// 	f := add
// 	fmt.Printf("Type of f: %T\n", f)
// 	fmt.Println(f(10, 20))
// }

// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func calc(x, y int, fn func(x, y int) int) int {
// 	return fn(x, y)
// }

// func main() {
// 	fmt.Println(calc(10, 20, add))
// }

// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func sub(x, y int) int {
// 	return x - y
// }

// func calc(option string) func(int, int) int {
// 	switch option {
// 	case "+":
// 		return add
// 	case "-":
// 		return sub
// 	case "*":
// 		return func(x, y int) int {
// 			return x * y
// 		}
// 	default:
// 		return nil
// 	}
// }

// func main() {
// 	fmt.Println(calc("+")(10, 20))
// 	fmt.Println(calc("-")(10, 20))
// 	fmt.Println(calc("*")(10, 20))
// }

// package main

// import "fmt"

// func adder() func(int) int {
// 	var x int
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	fn := adder()
// 	fmt.Println(fn(10))
// 	fmt.Println(fn(20))
// 	fmt.Println(fn(30))

// 	fn2 := adder()
// 	fmt.Println(fn2(40))
// 	fmt.Println(fn2(50))
// }

// package main

// import "fmt"

// func adder(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	fn := adder(10)
// 	fmt.Println(fn(10))
// 	fmt.Println(fn(20))
// 	fmt.Println(fn(30))

// 	fn2 := adder(20)
// 	fmt.Println(fn2(40))
// 	fmt.Println(fn2(50))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func makeSuffix(suffix string) func(string) string {
// 	return func(name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// func main() {
// 	jpgFn := makeSuffix(".jpg")
// 	txtFn := makeSuffix(".txt")
// 	fmt.Println(jpgFn("test"))
// 	fmt.Println(txtFn("test"))
// }

// package main

// import "fmt"

// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(num int) int {
// 		base += num
// 		return base
// 	}

// 	sub := func(num int) int {
// 		base -= num
// 		return base
// 	}

// 	return add, sub
// }

// func main() {
// 	fn1, fn2 := calc(10)
// 	fmt.Println(fn1(1), fn2(2))
// 	fmt.Println(fn1(3), fn2(4))
// 	fmt.Println(fn1(5), fn2(6))
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Start!")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)
// 	fmt.Println("End!")
// }

// package main

// import "fmt"

// func f1() int {
// 	x := 5
// 	defer func() {
// 		x++
// 	}()
// 	return x
// }

// func f2() (x int) {
// 	defer func() {
// 		x++
// 	}()
// 	return 5
// }

// func f3() (y int) {
// 	x := 5
// 	defer func() {
// 		x++
// 	}()
// 	return x
// }

// func f4() (x int) {
// 	defer func(x int) {
// 		x++
// 	}(x)
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
// 	x, y := 1, 1
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
// 			fmt.Println("jjj")
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

// package main

// import "fmt"

// func fn() {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Print("jj")
// 			fmt.Println(err)
// 		}
// 	}()

// 	a := 10
// 	b := 0
// 	fmt.Println(a / b)
// }

// func main() {
// 	fn()
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func readFile(filename string) error {
// 	if filename == "main.go" {
// 		return nil
// 	}

// 	return errors.New("读取文件错误。")
// }

// func fn() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("抛出异常给管理员发生邮件！")
// 		}
// 	}()

// 	err := readFile("xxx.go")

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("继续执行！")
// }

// func main() {
// 	fn()
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	unixTime := now.Unix()
// 	unixNano := now.UnixNano()
// 	unixMilli := now.UnixMilli()
// 	fmt.Println(now, now.Hour())
// 	fmt.Println(unixTime)
// 	fmt.Println(unixNano)
// 	fmt.Println(unixMilli)
// 	str := now.Format("2006-01-02 03:04:05")
// 	fmt.Println(str)
// 	fmt.Println(now.Format("2006/01/02 15:04:05"))
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	unixTime := now.Unix()
// 	unixNano := now.UnixNano()

// 	fmt.Printf("Current Unix time: %v\n", unixTime)
// 	fmt.Printf("Current Unix nano: %v\n", unixNano)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	unixTime := now.Unix()
// 	unixNano := now.UnixNano()

// 	fmt.Printf("Current Unix time: %v\n", unixTime)
// 	fmt.Printf("Current Unix nano: %v\n", unixNano)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func unixToTime(timestamp int64) {
// 	timeObj := time.Unix(timestamp, 0)
// 	str := timeObj.Format("2006-01-02 15:04:05")
// 	fmt.Println(str)
// }

// func main() {
// 	unixToTime(1758348998)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	str := "2025-09-20 18:50:02"
// 	timestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
// 	fmt.Println(timestamp.Hour())
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	later := now.Add(time.Hour)
// 	fmt.Println(later)
// 	fmt.Println(now.Sub(later))
// 	fmt.Println(now.Equal(later))
// 	fmt.Println(now.Before(later))
// 	fmt.Println(later.After(now))
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ticker := time.NewTicker(time.Hour)
// 	n := 0

// 	for v := range ticker.C {
// 		fmt.Println(v)
// 		n++
// 		if n == 5 {
// 			ticker.Stop()
// 			break
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	n := 0
// 	for {
// 		time.Sleep(time.Second)
// 		fmt.Println("我在定时执行任务！")
// 		n++
// 		if n == 5 {
// 			break
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ticker := time.NewTicker(time.Hour)
// 	n := 0

// 	for v := range ticker.C {
// 		fmt.Println(v)
// 		n++
// 		if n == 5 {
// 			ticker.Stop()
// 			break
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 10
// 	p := &a
// 	fmt.Printf("a 的值为 %v，a 的类型为 %T，a 的地址为 %p\n",
// 		a, a, &a)
// 	fmt.Printf("p 的值为 %v，p 的类型为 %T\n",
// 		p, p)
// }
