// package main

// import "fmt"

//	func main() {
//		var name string = "Hachimi"
//		var age = 25
//		fmt.Println(name, age)
//	}
// package main

// import "fmt"

//	func main() {
//		x := 10
//		y, x := 30, 20
//		fmt.Println(x, y)
//	}
// package main

// import "fmt"

//	func main() {
//		a, b, c := 10, 20, 30
//		fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
//		fmt.Printf("%#o\n", a)
//		fmt.Printf("%5d\n", 42)
//		fmt.Printf("%0*d\n", 6, 42)
//		fmt.Printf("%04.1F\n", 3.1415926)
//		fmt.Printf("%4.5s\n", "Hello World!")
//		fmt.Printf("%#.8x\n", 10001)
//		fmt.Printf("%.5e\n", 1414213421.1)
//		fmt.Printf("%.5G\n", 12.1)
//	}
// package main

// import "fmt"

// type Human struct {
// 	Name string
// }

//	func main() {
//		people := Human{Name: "张三"}
//		fmt.Printf("%v %+v %#v %T %%\n", people, people, people, people)
//		fmt.Printf("%t\n", true)
//		fmt.Printf("%b %c %d %o %x %X %q %U\n", 5, 0x4e2d, 0x12, 10, 13, 13, '闹', '😂')
//		fmt.Printf("%b %e %E %f %g %G\n", 3.14, 3.14159, 3.14159, 3.14159, 3.14159, 3.14159)
//		fmt.Printf("%s %q %x %X\n", []byte("Go 语言"), "Go 语言", "golang", "golang")
//		fmt.Printf("%p", &people)
//	}
// package main

// import "fmt"

//	func main() {
//		s := "中"
//		fmt.Println(len(s))
//		for _, r := range s {
//			fmt.Printf("%c\n", r)
//		}
//	}
// package main

// import (
// 	"fmt"

// 	"github.com/rivo/uniseg"
// )

//	func main() {
//		s := "👨‍👩‍👧‍👦"
//		gr := uniseg.NewGraphemes(s)
//		i := 0
//		for gr.Next() {
//			fmt.Printf("Char %d: %s\n", i, gr.Str())
//			i++
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"

// 	"github.com/rivo/uniseg"
// )

// func main() {
// 	s := "啧啧啧👨‍👩‍👧‍👦"

//		fmt.Println(utf8.RuneCountInString(s))
//		fmt.Println(uniseg.GraphemeClusterCount(s))
//	}
// package main

// import "fmt"

//	func main() {
//		var username = "Hachimi"
//		var age int = 20
//		fmt.Println(username, age)
//		var a, b, c, d = 1, 2, 3, false
//		fmt.Println(a, b, c, d)
//	}
// package main

// import "fmt"

//	func main() {
//		var (
//			a string
//			b int
//			c bool
//		)
//		a = "哈基米"
//		b = 10
//		c = true
//		fmt.Println(a, b, c)
//	}
// package main

// import "fmt"

// var m = 100

//	func main() {
//		n := 10
//		m := 200
//		fmt.Println(m, n)
//	}
// package main

// import "fmt"

// func Info() (string, int) {
// 	return "胖宝宝", 100
// }

//	func main() {
//		username, _ := Info()
//		fmt.Println(username)
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
// 	C
// 	D
// )

//	func main() {
//		fmt.Println(A, B, C, D)
//	}
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

//	func main() {
//		fmt.Println(A, B, C, D, E, F)
//	}
// package main

// import "fmt"

// const (
// 	A, B = iota + 1, iota + 2
// 	C, D
// 	E, F
// 	G = 100
// )

//	func main() {
//		fmt.Println(A, B, C, D, E, F, G)
//	}
// package main

// import "fmt"

//	func main() {
//		var num int64
//		num = 123
//		fmt.Printf("值：%v，类型：%T", num, num)
//	}
// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

//	func main() {
//		var a int8 = 120
//		var s string = "哈基米"
//		fmt.Println(unsafe.Sizeof(a))
//		fmt.Println(unsafe.Sizeof(s), len(s))
//	}
// package main

// import "fmt"

//	func main() {
//		var num1 int8 = 127
//		num2 := int32(num1)
//		fmt.Printf("值：%v, 类型：%T\n", num2, num2)
//	}
// package main

// import "fmt"

//	func main() {
//		var a int = 10
//		fmt.Printf("%d\n", a)
//		fmt.Printf("%#b\n", a)
//		var b int = 0o77
//		fmt.Printf("%#o\n", b)
//		var c int = 0xff
//		fmt.Printf("%#x\n", c)
//		fmt.Printf("%#X\n", c)
//		fmt.Printf("%d\n", c)
//	}
// package main

// import (
// 	"fmt"
// 	"math"
// )

//	func main() {
//		fmt.Printf("%f\n", math.Pi)
//		fmt.Printf("%.2f\n", math.Pi)
//		num := 1.1
//		fmt.Printf("值：%v，类型：%T\n", num, num)
//	}
// package main

// import "fmt"

//	func main() {
//		num := 5.1234e2
//		num2 := 5.1234e2
//		num3 := 5.1234e-2
//		fmt.Printf("num = %v, num2 = %v, num3 = %v\n", num, num2, num3)
//	}
// package main

// import "fmt"

//	func main() {
//		d := 1129.6
//		fmt.Println(d * 100)
//		m1, m2 := 8.2, 3.8
//		fmt.Println(m1 - m2)
//	}
// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

//	func main() {
//		var b = true
//		fmt.Printf("%v 占用字节：%v", b, unsafe.Sizeof(b))
//	}
// package main

// import "fmt"

//	func main() {
//		fmt.Println("C:\\Code\\demo\\main.go")
//		str := `第一行
//
// 第二行
// 第三行`
//
//		fmt.Println(str)
//	}
// package main

// import "fmt"

//	func main() {
//		str := "this is str"
//		fmt.Println(len(str))
//	}
// package main

// import "fmt"

//	func main() {
//		str1 := "Hello "
//		str2 := "Golang!"
//		fmt.Println(str1 + str2)
//		var str3 string = fmt.Sprintf("%v%v", str1, str2)
//		fmt.Println(str3)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		str := "123-456-789"
//		arr := strings.Split(str, "-")
//		fmt.Println(arr)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		str := "this is golang"
//		flag := strings.Contains(str, "golang")
//		fmt.Println(flag)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "this is golang"
// 	flag := strings.HasPrefix(str, "this")
// 	fmt.Println(flag)

//		flag = strings.HasSuffix(str, "go")
//		fmt.Println(flag)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "this is golang"
// 	index := strings.Index(str, "is")
// 	fmt.Println(index)

//		index = strings.LastIndex(str, "is")
//		fmt.Println(index)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		str := "123-456-789"
//		arr := strings.Split(str, "-")
//		str2 := strings.Join(arr, "-*-")
//		fmt.Println(str2)
//	}
// package main

// import "fmt"

//	func main() {
//		a := 'a'
//		b := 'b'
//		c := 99
//		fmt.Println(a)
//		fmt.Println(b)
//		fmt.Printf("%c %c %c\n", a, b, c)
//	}
// package main

// import "fmt"

//	func main() {
//		a := "m"
//		fmt.Println(len(a))
//		b := "张"
//		fmt.Println(len(b))
//	}
// package main

// import "fmt"

//	func main() {
//		str := "哈基米Mambo"
//		for i := 0; i < len(str); i++ {
//			fmt.Printf("%v(%c) ", str[i], str[i])
//		}
//		fmt.Println()
//		for _, ch := range str {
//			fmt.Printf("%v(%c) ", ch, ch)
//		}
//		fmt.Println()
//	}
// package main

// import "fmt"

//	func main() {
//		str1 := "big"
//		byteStr1 := []byte(str1)
//		byteStr1[0] = 'p'
//		fmt.Println(string(byteStr1))
//		str2 := "红萝卜"
//		runeStr2 := []rune(str2)
//		runeStr2[0] = '白'
//		fmt.Println(string(runeStr2))
//	}
// package main

// import "fmt"

//	func main() {
//		str1 := "big"
//		byteStr1 := []byte(str1)
//		byteStr1[0] = 'p'
//		fmt.Println(string(byteStr1))
//		str2 := "红萝卜"
//		runeStr2 := []rune(str2)
//		runeStr2[0] = '白'
//		fmt.Println(string(runeStr2))
//	}
// package main

// import "fmt"

//	func main() {
//		str1 := "big"
//		byteStr1 := []byte(str1)
//		byteStr1[0] = 'p'
//		fmt.Println(string(byteStr1))
//		str2 := "红萝卜"
//		byteStr2 := []rune(str2)
//		byteStr2[0] = '白'
//		fmt.Println(string(byteStr2))
//	}
// package main

// import "fmt"

//	func main() {
//		var a int8 = 20
//		var b int16 = 40
//		var c = int16(a) + b
//		fmt.Printf("值：%v，类型：%T\n", c, c)
//	}
// package main

// import "fmt"

//	func main() {
//		var a float32 = 3.2
//		var b int16 = 6
//		c := a + float32(b)
//		fmt.Printf("值：%v，类型：%T\n", c, c)
//	}
// package main

// import "fmt"

//	func main() {
//		var a int16 = 128
//		var b = int8(a)
//		fmt.Printf("b = %v\n", b)
//	}
// package main

// import (
// 	"fmt"
// 	"math"
// )

//	func main() {
//		a, b := 3, 4
//		c := math.Sqrt(float64(a*a + b*b))
//		fmt.Println(c)
//	}
// package main

// import "fmt"

//	func main() {
//		var i int = 20
//		var f float64 = 3.1415926
//		var t bool = true
//		var b byte = 'a'
//		str := fmt.Sprintf("%d", i)
//		fmt.Printf("值：%v，类型：%T\n", str, str)
//		str = fmt.Sprintf("%f", f)
//		fmt.Printf("值：%v，类型：%T\n", str, str)
//		str = fmt.Sprintf("%t", t)
//		fmt.Printf("值：%v，类型：%T\n", str, str)
//		str = fmt.Sprintf("%c", b)
//		fmt.Printf("值：%v，类型：%T\n", str, str)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

//	func main() {
//		var num int = 10
//		s := strconv.Itoa(num)
//		fmt.Printf("值：%v，类型：%T\n", s, s)
//		s2 := strings.ToUpper(strconv.FormatInt(int64(num), 16))
//		fmt.Printf("值：%v，类型：%T\n", s2, s2)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		var num float32 = 1.14514
//		s := strconv.FormatFloat(float64(num), 'g', 2, 64)
//		fmt.Println(s)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		str := strconv.FormatBool(true)
//		fmt.Printf("值：%v，类型：%T\n", str, str)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		str := strconv.FormatBool(true)
//		fmt.Printf("值为 %v，类型为 %T\n", str, str)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		var num float32 = 1.14514
//		s := strconv.FormatFloat(float64(num), 'f', 2, 64)
//		fmt.Println(s)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		str := "1234"
//		i64, _ := strconv.ParseInt(str, 16, 64)
//		fmt.Printf("值：%v，类型：%T\n", i64, i64)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		str := "1234"
//		i64, _ := strconv.ParseInt(str, 10, 64)
//		fmt.Printf("值：%v，类型：%T\n", i64, i64)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		str := "3.1415926"
//		f32, _ := strconv.ParseFloat(str, 32)
//		f64, _ := strconv.ParseFloat(str, 64)
//		fmt.Printf("值为 %v，类型为 %T\n", f32, f32)
//		fmt.Printf("值为 %v，类型为 %T\n", f64, f64)
//	}
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		b, _ := strconv.ParseBool("true")
//		fmt.Printf("值为 %v，类型为 %T\n", b, b)
//	}
// package main

// import "fmt"

//	func main() {
//		str := "Hello 张三"
//		for _, ch := range str {
//			fmt.Printf("%v(%c)", ch, ch)
//		}
//		fmt.Println()
//	}
// package main

// import "fmt"

//	func main() {
//		fmt.Println("10 + 3 =", 10+3)
//		fmt.Println("10 - 3 =", 10-3)
//		fmt.Println("10 * 3 =", 10*3)
//		fmt.Println("10 / 3 =", 10/3)
//		fmt.Println("10.0 / 3.0 =", 10.0/3.0)
//		fmt.Println("10 % 3 =", 10%3)
//		fmt.Println("-10 % 3 =", -10%3)
//		fmt.Println("10 % -3 =", 10%-3)
//		fmt.Println("-10 % -3 =", -10%-3)
//	}
// package main

// import "fmt"

//	func main() {
//		var age int = 40
//		if age > 30 && age < 50 {
//			fmt.Println("哈基米")
//		}
//		if age > 30 && age < 40 {
//			fmt.Println("曼波")
//		}
//		if age > 30 || age < 50 {
//			fmt.Println("叮咚鸡")
//		}
//		if age > 30 || age < 40 {
//			fmt.Println("大狗叫")
//		}
//		if age > 30 {
//			fmt.Println("胖宝宝")
//		}
//		if !(age > 30) {
//			fmt.Println("小白手套")
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		a := 9
//		b := 2
//		t := a
//		a = b
//		b = t
//		fmt.Printf("交换后的情况为 a = %v，b = %v\n", a, b)
//	}
// package main

// import "fmt"

//	func main() {
//		days := 100
//		week := days / 7
//		day := days % 7
//		fmt.Printf("%d 个星期零 %d 天\n", week, day)
//	}
// package main

// import "fmt"

//	func main() {
//		days := 100
//		week := days / 7
//		day := days % 7
//		fmt.Printf("%d 个星期零 %d 天\n", week, day)
//	}
// package main

// import "fmt"

//	func main() {
//		var huashi float32 = 134.2
//		var sheshi float32 = 5.0 / 9 * (huashi - 100)
//		fmt.Println(huashi, sheshi)
//	}
// package main

// import "fmt"

//	func main() {
//		a := 5                        // 101
//		b := 2                        // 010
//		fmt.Println("a & b =", a&b)   // 000
//		fmt.Println("a | b =", a|b)   // 111
//		fmt.Println("a ^ b =", a^b)   // 111
//		fmt.Println("a >> b =", a>>b) // 1
//		fmt.Println("a << b =", a<<b) // 10100
//		fmt.Println("a << 1 =", a<<1) // 1010
//		fmt.Println("a >> 1 =", a>>1) // 10
//		fmt.Println("7 >> 2 =", 7>>2) // 1
//	}
// package main

// import "fmt"

//	func main() {
//		score := 65
//		if score >= 90 {
//			fmt.Println("A")
//		} else if score > 75 {
//			fmt.Println("B")
//		} else {
//			fmt.Println("C")
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		if score := 56; score >= 90 {
//			fmt.Println("a")
//		} else if score > 75 {
//			fmt.Println("b")
//		} else {
//			fmt.Println("c")
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		for i := 1; i <= 10; i++ {
//			fmt.Printf("%v ", i)
//		}
//		fmt.Println()
//	}
// package main

// import "fmt"

//	func main() {
//		i := 1
//		for ; i <= 10; i++ {
//			fmt.Printf("%v ", i)
//		}
//		fmt.Println()
//	}
// package main

// import "fmt"

//	func main() {
//		i := 1
//		for i <= 10 {
//			fmt.Printf("%v ", i)
//			i++
//		}
//		fmt.Println()
//	}
// package main

// import "fmt"

//	func main() {
//		k := 1
//		for {
//			if k <= 10 {
//				fmt.Println("ok~~~", k)
//			} else {
//				break
//			}
//			k++
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		for i := 0; i < 50; i++ {
//			if i%2 == 0 {
//				fmt.Printf("%v ", i)
//			}
//		}
//		fmt.Println()
//	}
// package main

// import "fmt"

//	func main() {
//		sum := 0
//		count := 0
//		for i := 0; i <= 100; i++ {
//			if i%9 == 0 {
//				sum += i
//				count++
//			}
//		}
//		fmt.Println(sum, count)
//	}
// package main

// import "fmt"

//	func main() {
//		sum := 0
//		for i := 1; i <= 100; i++ {
//			sum += i
//		}
//		fmt.Println(sum)
//	}
// package main

// import "fmt"

//	func main() {
//		n := 1
//		for i := 1; i <= 5; i++ {
//			n *= i
//		}
//		fmt.Println(n)
//	}
// package main

// import "fmt"

//	func main() {
//		for i := 1; i <= 16; i++ {
//			fmt.Print("* ")
//			if i%4 == 0 {
//				fmt.Println()
//			}
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		for range 4 {
//			for range 4 {
//				fmt.Print("* ")
//			}
//			fmt.Println()
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		for i := 1; i <= 5; i++ {
//			for j := 1; j <= i; j++ {
//				fmt.Print("* ")
//			}
//			fmt.Println()
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		for i := 1; i <= 9; i++ {
//			for j := 1; j <= i; j++ {
//				fmt.Printf("%v x %v = %v \t", i, j, i*j)
//			}
//			fmt.Println()
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		for i := 1; i <= 9; i++ {
//			for j := 1; j <= i; j++ {
//				fmt.Printf("%v x %v = %v\t", i, j, i*j)
//			}
//			fmt.Println()
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		str := "Mambo 哈基米"
//		for idx, val := range str {
//			fmt.Printf("idx = %v, val = %c\n", idx, val)
//		}
//		for _, val := range str {
//			fmt.Printf("val = %c\n", val)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		extname := ".html"
//		switch extname {
//		case ".html":
//			fmt.Println("text/html")
//		case ".css":
//			fmt.Println("text/css")
//		case ".js":
//			fmt.Println("text/javascript")
//		default:
//			fmt.Println("格式错误")
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		n := 2
//		switch n {
//		case 1, 3, 5, 7, 9:
//			fmt.Println("奇数")
//		case 0, 2, 4, 6, 8:
//			fmt.Println("偶数")
//		default:
//			fmt.Println(n)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		switch n := 7; n {
//		case 1, 3, 5, 7, 9:
//			fmt.Println("奇数")
//		case 0, 2, 4, 6, 8:
//			fmt.Println("偶数")
//		default:
//			fmt.Println(n)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		age := 56
//		switch {
//		case age < 25:
//			fmt.Println("好好学习")
//		case age <= 60:
//			fmt.Println("好好工作")
//		case age > 60:
//			fmt.Println("好好享受")
//		default:
//			fmt.Println("好好活着")
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		str := "a"
//		switch str {
//		case "a":
//			fmt.Println("a")
//			fallthrough
//		case "b":
//			fmt.Println("b")
//			fallthrough
//		case "c":
//			fmt.Println("c")
//			fallthrough
//		default:
//			fmt.Println("...")
//		}
//	}
// package main

// import "fmt"

// func main() {
// label:
//
//		for i := range 2 {
//			for j := range 10 {
//				if j == 3 {
//					break label
//				}
//				fmt.Printf("i = %v, j = %v\n", i, j)
//			}
//		}
//	}
// package main

// import "fmt"

// func main() {
// label:
//
//		for i := range 2 {
//			for j := range 4 {
//				if j == 2 {
//					continue label
//				}
//				fmt.Printf("i = %v, j = %v\n", i, j)
//			}
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		n := 30
//		fmt.Println("ok1")
//		if n > 20 {
//			goto label
//		}
//		fmt.Println("ok2")
//		fmt.Println("ok3")
//		fmt.Println("ok4")
//
// label:
//
//		fmt.Println("ok5")
//		fmt.Println("ok6")
//		fmt.Println("ok7")
//	}
// package main

// import "fmt"

// func main() {
// 	for i := range 10 {
// 		for j := range 10 {
// 			if j == 10 {
// 				goto breakTag
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}

//	}
//	return
//
// breakTag:
//
//		fmt.Println("结束循环")
//	}
// package main

// import "fmt"

//	func main() {
//		var arr [3]int
//		arr[0] = 80
//		arr[1] = 100
//		arr[2] = 96
//		fmt.Println(arr)
//	}
// package main

// import "fmt"

//	func main() {
//		var testArray [3]int
//		var numArray = [3]int{1, 2}
//		var cityArray = [3]string{"北京", "上海"}
//		fmt.Println(testArray)
//		fmt.Println(numArray)
//		fmt.Println(cityArray)
//	}
// package main

// import "fmt"

//	func main() {
//		var testArray [3]int
//		var numArray = [...]int{1, 2}
//		cityArray := [...]string{"北京", "上海", "深圳"}
//		fmt.Println(testArray)
//		fmt.Println(numArray)
//		fmt.Printf("numArray 类型为 %T\n", numArray)
//		fmt.Println(cityArray)
//		fmt.Printf("cityArray 类型为 %T\n", cityArray)
//	}
// package main

// import "fmt"

//	func main() {
//		arr := [...]int{0: 1, 3: 5, 6: 8, 4: 2}
//		fmt.Println(arr)
//		fmt.Printf("Type of arr: %T\n", arr)
//	}
// package main

// import "fmt"

//	func main() {
//		arr := [...]string{"北京", "上海", "深圳"}
//		for i := 0; i < len(arr); i++ {
//			fmt.Println(arr[i])
//		}
//		for idx, val := range arr {
//			fmt.Println(idx, val)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		arr := [3]int{10, 20, 30}
//		arr2 := arr
//		arr2[0] = 40
//		fmt.Println(arr)
//		fmt.Println(arr2)
//		fmt.Println()
//		arr3 := [3][2]int{
//			{1, 1},
//			{1, 1},
//			{1, 1},
//		}
//		arr4 := arr3
//		arr4[2][0] = 100
//		fmt.Println(arr3)
//		fmt.Println(arr4)
//	}
// package main

// import "fmt"

//	func main() {
//		a, b, c := 10, 20, 30
//		ptrArr := [3]*int{&a, &b, &c}
//		fmt.Println("指针数组本身（值类型，数组地址）：", &ptrArr)
//		fmt.Println("指针数组的元素（存储的是指针地址）：", ptrArr)
//		for i, ptr := range ptrArr {
//			fmt.Printf("ptrArr[%d] 指向的值：%d\n", i, *ptr)
//		}
//		ptrArr2 := ptrArr
//		ptrArr2[0] = &c
//		fmt.Println("ptrArr[0] 指向的值：", *ptrArr[0])
//		fmt.Println("ptrArr2[0] 指向的值：", *ptrArr2[0])
//	}
// package main

// import "fmt"

// func main() {
// 	arr := [3][2]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"成都", "重庆"},
// 	}
// 	fmt.Println(arr)
// 	fmt.Println(arr[2][1])

//		for _, v1 := range arr {
//			for _, v2 := range v1 {
//				fmt.Printf("%s\t", v2)
//			}
//			fmt.Println()
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		arr := [...]int{1, -4, 232, 22, 54, 98, -34}
//		sum := 0
//		for _, val := range arr {
//			sum += val
//		}
//		mean := float64(sum) / float64(len(arr))
//		fmt.Printf("sum = %v, mean = %v\n", sum, mean)
//	}
// package main

// import "fmt"

//	func main() {
//		arr := []int{1, -1, 12, 65, 11}
//		max := arr[0]
//		maxIdx := 0
//		for idx, val := range arr {
//			if max < val {
//				max = val
//				maxIdx = idx
//			}
//		}
//		fmt.Printf("max = %v, maxIdx = %v\n", max, maxIdx)
//	}
// package main

// import "fmt"

//	func main() {
//		arr := []int{1, 3, 5, 7, 8}
//		for i := range arr {
//			for j := i + 1; j < len(arr); j++ {
//				if arr[i]+arr[j] == 8 {
//					fmt.Printf("(%v, %v)\n", i, j)
//				}
//			}
//		}
//	}
// package main

// func arraySum(x [4]int) int {
// 	sum := 0
// 	for _, v := range x {
// 		sum += v
// 	}
// 	return sum
// }

//	func main() {
//		a := [4]int{1, 2, 3, 4}
//		println(arraySum(a))
//	}
// package main

// import "fmt"

//	func main() {
//		var a []string
//		var b = []int{}
//		var c = []bool{false, true}
//		fmt.Println(a)
//		fmt.Println(b)
//		fmt.Println(c)
//		fmt.Println(a == nil)
//		fmt.Println(b == nil)
//		fmt.Println(c == nil)
//	}
// package main

// import "fmt"

//	func main() {
//		arr := []string{"北京", "上海", "深圳"}
//		for _, val := range arr {
//			fmt.Println(val)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		arr := [5]int{55, 56, 57, 58, 59}
//		arr2 := arr[1:4]
//		fmt.Println(arr2)
//		fmt.Printf("Type of arr2: %T\n", arr2)
//		arr2[0] = 1
//		fmt.Println(arr)
//	}
// package main

// import "fmt"

//	func main() {
//		arr := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
//		fmt.Printf("arr: %v\ttype: %T\tlen: %v\tcap: %v\n", arr, arr, len(arr), cap(arr))
//		arr2 := arr[1:3]
//		fmt.Printf("arr2: %v\ttype: %T\tlen: %v\tcap: %v\n", arr2, arr2, len(arr2), cap(arr2))
//		arr3 := arr2[1:5]
//		fmt.Printf("arr3: %v\ttype: %T\tlen: %v\tcap: %v\n", arr3, arr3, len(arr3), cap(arr3))
//	}
// package main

// import "fmt"

//	func main() {
//		arr := []int{2, 3, 5, 7, 11, 13}
//		fmt.Println(arr)
//		fmt.Printf("len = %v, cap = %v\n", len(arr), cap(arr))
//		arr2 := arr[:2]
//		fmt.Println(arr2)
//		fmt.Printf("len = %v, cap = %v\n", len(arr2), cap(arr2))
//		arr3 := arr[1:3]
//		fmt.Println(arr3)
//		fmt.Printf("len = %v, cap = %v\n", len(arr3), cap(arr3))
//	}
// package main

// import "fmt"

//	func main() {
//		slice := make([]int, 2, 10)
//		fmt.Printf("val = %v, len = %v, cap = %v\n", slice, len(slice), cap(slice))
//	}
// package main

// import "fmt"

//	func main() {
//		slice := make([]int, 2, 10)
//		fmt.Printf("val = %v, len = %d, cap = %d\n", slice, len(slice), cap(slice))
//	}
// package main

// import "fmt"

//	func main() {
//		var slice []int
//		slice2 := []int{}
//		slice3 := make([]int, 0)
//		fmt.Printf("bool = %v, len = %v, cap = %v\n", slice == nil, len(slice), cap(slice))
//		fmt.Printf("bool = %v, len = %v, cap = %v\n", slice2 == nil, len(slice2), cap(slice2))
//		fmt.Printf("bool = %v, len = %v, cap = %v\n", slice3 == nil, len(slice3), cap(slice3))
//	}
// package main

// import "fmt"

//	func main() {
//		slice := make([]int, 3)
//		slice2 := slice
//		slice2[0] = 100
//		fmt.Println(slice)
//		fmt.Println(slice2)
//	}
// package main

// import "fmt"

//	func main() {
//		var numSlice []int
//		for i := range 10 {
//			numSlice = append(numSlice, i)
//			fmt.Printf("%v\tlen: %v\tcap: %v\tptr: %p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		var citySlice []string
//		citySlice = append(citySlice, "北京")
//		citySlice = append(citySlice, "上海", "广州", "深圳")
//		cityArr := [2]string{"成都", "杭州"}
//		citySlice = append(citySlice, cityArr[:]...)
//		fmt.Println(citySlice)
//	}
// package main

// import "fmt"

//	func main() {
//		slice := []int{100, 200, 300}
//		slice2 := []int{400, 500, 600}
//		slice3 := append(slice, slice2...)
//		fmt.Println(slice3)
//	}
// package main

// import "fmt"

//	func main() {
//		numSlice := []int{30, 31, 32, 33, 34, 35, 36, 37}
//		numSlice = append(numSlice[:2], numSlice[3:]...)
//		fmt.Println(numSlice)
//	}
// package main

// import "fmt"

//	func main() {
//		slice := make([]string, 5, 10)
//		for i := range 12 {
//			slice = append(slice, fmt.Sprintf("%v", i))
//		}
//		fmt.Println(slice)
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		slice := []int{3, 7, 8, 9, 1}
//		sort.Ints(slice)
//		fmt.Println(slice)
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		sort.Ints(intSlice)
//		sort.Float64s(floatSlice)
//		sort.Strings(stringSlice)
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
//		sort.Sort(sort.Reverse(sort.Float64Slice(floatSlice)))
//		sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//		sort.Sort(sort.IntSlice(intSlice))
//		sort.Sort(sort.Float64Slice(floatSlice))
//		sort.Sort(sort.StringSlice(stringSlice))
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		sort.Sort(sort.IntSlice(intSlice))
//		sort.Sort(sort.Float64Slice(floatSlice))
//		sort.Sort(sort.StringSlice(stringSlice))
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//		sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
//		sort.Sort(sort.Reverse(sort.Float64Slice(floatSlice)))
//		sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"cmp"
// 	"fmt"
// 	"slices"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		slices.Sort(intSlice)
//		slices.Sort(floatSlice)
//		slices.Sort(stringSlice)
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//		slices.SortFunc(intSlice, func(a, b int) int { return cmp.Compare(b, a) })
//		slices.SortFunc(floatSlice, func(a, b float64) int { return cmp.Compare(b, a) })
//		slices.SortFunc(stringSlice, func(a, b string) int { return cmp.Compare(b, a) })
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import "fmt"

//	func main() {
//		scoreMap := make(map[string]int, 8)
//		scoreMap["哈基米"] = 90
//		scoreMap["曼波"] = 100
//		fmt.Println(scoreMap)
//		fmt.Println(scoreMap["曼波"])
//		fmt.Printf("Type of scoreMap: %T\n", scoreMap)
//	}
// package main

// import "fmt"

//	func main() {
//		userinfo := map[string]string{
//			"username": "叮咚鸡",
//			"password": "大狗叫",
//		}
//		fmt.Println(userinfo)
//	}
// package main

// import "fmt"

//	func main() {
//		scoreMap := make(map[string]int, 8)
//		scoreMap["哈基米"] = 90
//		scoreMap["曼波"] = 100
//		val, ok := scoreMap["哈基米"]
//		if ok {
//			fmt.Println(val)
//		} else {
//			fmt.Println("查无此人")
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		scoreMap := make(map[string]int)
//		scoreMap["孙连城"] = 90
//		scoreMap["李达康"] = 100
//		scoreMap["沙瑞金"] = 60
//		delete(scoreMap, "李达康")
//		for key := range scoreMap {
//			fmt.Println(key)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"slices"
// )

//	func main() {
//		scoreMap := make(map[string]int, 200)
//		for i := range 100 {
//			key := fmt.Sprintf("stu%02d", i)
//			val := rand.Intn(100)
//			scoreMap[key] = val
//		}
//		keys := make([]string, 0, 200)
//		for key := range scoreMap {
//			keys = append(keys, key)
//		}
//		slices.Sort(keys)
//		for _, key := range keys {
//			fmt.Println(key, scoreMap[key])
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"slices"
// )

//	func main() {
//		scoreMap := make(map[string]int, 200)
//		for i := range 100 {
//			key := fmt.Sprintf("stu%02d", i)
//			val := rand.Intn(100)
//			scoreMap[key] = val
//		}
//		keys := make([]string, 0, 200)
//		for key := range scoreMap {
//			keys = append(keys, key)
//		}
//		slices.Sort(keys)
//		for _, key := range keys {
//			fmt.Println(key, scoreMap[key])
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		str := "Hello 哈基米"
//		for idx, ch := range str {
//			fmt.Printf("%v(%c)\n", idx, ch)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		mapSlice := make([]map[string]string, 3)
//		for i, v := range mapSlice {
//			fmt.Printf("i = %v, v = %v\n", i, v)
//		}
//		fmt.Println("After init...")
//		mapSlice[0] = make(map[string]string, 10)
//		mapSlice[0] = make(map[string]string, 10)
//		mapSlice[0]["name"] = "哈基米"
//		mapSlice[0]["password"] = "123456"
//		mapSlice[0]["address"] = "双流区"
//		for i, v := range mapSlice {
//			fmt.Printf("i = %v, v = %v\n", i, v)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		sliceMap := make(map[string][]string, 3)
//		fmt.Println(sliceMap)
//		fmt.Println("After init...")
//		key := "中国"
//		val, ok := sliceMap[key]
//		if !ok {
//			val = make([]string, 0, 2)
//		}
//		val = append(val, "北京", "上海")
//		sliceMap[key] = val
//		fmt.Println(sliceMap)
//	}
// package main

// import "fmt"

//	func main() {
//		sliceMap := make(map[string][]string, 3)
//		fmt.Println(sliceMap)
//		fmt.Println("After init...")
//		key := "中国"
//		val, ok := sliceMap[key]
//		if !ok {
//			val = make([]string, 0, 2)
//		}
//		val = append(val, "北京", "上海")
//		sliceMap[key] = val
//		fmt.Println(sliceMap)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		wordMap := make(map[string]int)
//		str := "how do you do"
//		for _, word := range strings.Split(str, " ") {
//			if word == "" {
//				continue
//			}
//			wordMap[word]++
//		}
//		fmt.Println(wordMap)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		wordMap := make(map[string]int)
//		str := "how   do you  do"
//		for word := range strings.FieldsSeq(str) {
//			wordMap[word]++
//		}
//		fmt.Println(wordMap)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		wordMap := make(map[string]int)
//		str := "how  do you do"
//		for _, word := range strings.Fields(str) {
//			wordMap[word]++
//		}
//		fmt.Println(wordMap)
//	}
// package main

// import "fmt"

// func sayHello() {
// 	fmt.Println("Hello World!")
// }

// func add(x, y int) int {
// 	return x + y
// }

//	func main() {
//		sayHello()
//		fmt.Println(add(10, 20))
//	}
// package main

// import "fmt"

// func add(nums ...int) (sum int) {
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return
// }

//	func main() {
//		fmt.Println(add())
//		fmt.Println(add(10))
//		fmt.Println(add(10, 20))
//		fmt.Println(add(10, 20, 30))
//		fmt.Println(add(10, 20, 30, 40))
//	}
// package main

// import "fmt"

// func add(x int, y ...int) (sum int) {
// 	sum = x
// 	for _, num := range y {
// 		sum += num
// 	}
// 	return
// }

//	func main() {
//		fmt.Println(add(10))
//		fmt.Println(add(10, 20))
//		fmt.Println(add(10, 20, 30))
//		fmt.Pri ntln(add(10, 20, 30, 40))
//	}
// package main

//	func calc(x, y int) (sum, diff int) {
//		sum = x + y
//		diff = x - y
//		return
//	}
// package main

// import "fmt"

// var num = 10

// func testGlobal() {
// 	fmt.Printf("num = %v\n", num)
// }

//	func main() {
//		testGlobal()
//	}
// package main

// import "fmt"

//	func testLocal() {
//		var x int64 = 100
//		fmt.Printf("x = %d\n", x)
//	}
// package main

// import "fmt"

// var num = 10

// func testNum() {
// 	num := 100
// 	fmt.Printf("num = %d\n", num)
// }

//	func main() {
//		testNum()
//		fmt.Printf("num = %d\n", num)
//	}
// package main

// import "fmt"

// func testLocal(x, y int) {
// 	fmt.Println(x, y)
// 	if x > 0 {
// 		z := 100
// 		fmt.Println(z)
// 	}
// 	for i := range 10 {
// 		fmt.Println(i)
// 	}
// }

//	func main() {
//		testLocal(1, 2)
//	}
// package main

// import "fmt"

// type calc func(int, int) int

// func add(x, y int) int {
// 	return x + y
// }

// func sub(x, y int) int {
// 	return x - y
// }

//	func main() {
//		var c calc = add
//		fmt.Printf("Type of c: %T\n", c)
//		fmt.Println(c(1, 2))
//		f := add
//		fmt.Printf("Type of f: %T\n", f)
//		fmt.Println(f(10, 20))
//	}
// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func calc(x, y int, fn func(x, y int) int) int {
// 	return fn(x, y)
// }

//	func main() {
//		fmt.Println(calc(10, 20, add))
//	}
// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func calc(x, y int, fn func(x, y int) int) int {
// 	return fn(x, y)
// }

//	func main() {
//		fmt.Println(calc(20, 10, add))
//	}
// package main

// import "fmt"

// func add(x, y int) (sum int) {
// 	sum = x + y
// 	return
// }

// func sub(x, y int) (diff int) {
// 	diff = x - y
// 	return
// }

// func calc(option byte) func(int, int) int {
// 	switch option {
// 	case '+':
// 		return add
// 	case '-':
// 		return sub
// 	case '*':
// 		return func(x, y int) int {
// 			return x * y
// 		}
// 	default:
// 		return nil
// 	}
// }

//	func main() {
//		fmt.Println(calc('+')(10, 20))
//		fmt.Println(calc('-')(10, 20))
//		fmt.Println(calc('*')(10, 20))
//	}
// package main

// import "fmt"

//	func main() {
//		add := func(x, y int) {
//			fmt.Println(x + y)
//		}
//		add(10, 20) // 通过变量调用匿名函数
//		// 自执行函数指匿名函数定义完加 () 直接执行
//		func(x, y int) {
//			fmt.Println(x + y)
//		}(10, 20)
//	}
// package main

// import "fmt"

//	func main() {
//		add := func(x, y int) {
//			fmt.Println(x + y)
//		}
//		add(100, 10)
//		func(x, y int) {
//			fmt.Println(x + y)
//		}(20, 30)
//	}
// package main

// import "fmt"

// func adder() func(int) int {
// 	var x int
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

//	func main() {
//		var fn = adder()
//		fmt.Println(fn(10))
//		fmt.Println(fn(20))
//		fmt.Println(fn(30))
//		fn2 := adder()
//		fmt.Println(fn2(40))
//		fmt.Println(fn2(50))
//		fmt.Printf("%T\n", fn)
//	}
// package main

// import "fmt"

// func adder(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

//	func main() {
//		fn := adder(10)
//		fmt.Println(fn(10))
//		fmt.Println(fn(20))
//		fmt.Println(fn(30))
//		fn2 := adder(20)
//		fmt.Println(fn2(40))
//		fmt.Println(fn2(50))
//	}
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

//	func main() {
//		jpgFn := makeSuffix(".jpg")
//		txtFn := makeSuffix(".txt")
//		fmt.Println(jpgFn("test"))
//		fmt.Println(txtFn("test"))
//	}
// package main

// import "strings"

//	func makeSuffix(suffix string) func(string) string {
//		return func(name string) string {
//			if !strings.HasSuffix(name, suffix) {
//				return name + suffix
//			}
//			return name
//		}
//	}
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

//	func main() {
//		jpgFn := makeSuffix(".jpg")
//		txtFn := makeSuffix(".txt")
//		fmt.Println(jpgFn("test"))
//		fmt.Println(txtFn("test"))
//	}
// package main

// import "fmt"

// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}
// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return add, sub
// }

//	func main() {
//		fn, fn2 := calc(10)
//		fmt.Println(fn(1), fn2(2))
//		fmt.Println(fn(3), fn2(4))
//		fmt.Println(fn(5), fn2(6))
//	}
// package main

// import "fmt"

// func calc(base int) (add func(int) int, sub func(int) int) {
// 	add = func(i int) int {
// 		base += i
// 		return base
// 	}
// 	sub = func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return
// }

//	func main() {
//		fn, fn2 := calc(10)
//		fmt.Println(fn(1), fn2(2))
//		fmt.Println(fn(3), fn2(4))
//		fmt.Println(fn(5), fn2(6))
//	}
// package main

// import "fmt"

//	func main() {
//		fmt.Println("Start!")
//		defer fmt.Println(1)
//		defer fmt.Println(2)
//		defer fmt.Println(3)
//		fmt.Println("结束！")
//	}
// package main

// import "fmt"

// func f1() int {
// 	// 局部变量 x = 5
// 	x := 5
// 	// defer 执行时 x 变为 6，但不影响返回值
// 	defer func() {
// 		x++
// 	}()
// 	// 返回 x 的值为 5
// 	return x
// }

// // x 是命名返回值，初始值为 0
// func f2() (x int) {
// 	// defer 执行时，x 变为 6
// 	defer func() {
// 		x++
// 	}()
// 	// 设置 x = 5，然后 defer 修改为 6
// 	return 5
// }

// // y 是命名返回值，初始为 0
// func f3() (y int) {
// 	// 局部变量 x = 5
// 	x := 5
// 	// defer 修改局部变量 x 为 6，不影响 y
// 	defer func() {
// 		x++
// 	}()
// 	// 设置 y = x，此时 x 为 5
// 	return x
// }

// // x 是命名返回值，初始为 0
// func f4() (x int) {
// 	// defer 接受 x 的副本（0）
// 	defer func(x int) {
// 		// 修改副本为 1，不影响原 x
// 		x++
// 	}(x)
// 	// 设置 x = 5
// 	return 5
// }

//	func main() {
//		fmt.Println(f1())
//		fmt.Println(f2())
//		fmt.Println(f3())
//		fmt.Println(f4())
//	}
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

//	func main() {
//		fmt.Println(f1())
//		fmt.Println(f2())
//		fmt.Println(f3())
//		fmt.Println(f4())
//	}
// package main

// import "fmt"

// func calc(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

//	func main() {
//		x, y := 1, 1
//		defer calc("AA", x, calc("A", x, y))
//		x = 10
//		defer calc("BB", x, calc("B", x, y))
//		y = 20
//	}
// package main

// import "fmt"

// func calc(idx string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(idx, a, b, ret)
// 	return ret
// }

//	func main() {
//		x, y := 1, 1
//		defer calc("AA", x, calc("A", x, y))
//		x = 10
//		defer calc("BB", x, calc("B", x, y))
//		y = 20
//	}
// package main

// import "fmt"

// func fnA() {
// 	fmt.Println("fn A")
// }

// func fnB() {
// 	panic("panic in B")
// }

// func fnC() {
// 	fmt.Println("fn C")
// }

//	func main() {
//		fnA()
//		fnB()
//		fnC()
//	}
// package main

// import "fmt"

// func fnA() {
// 	fmt.Println("fn A")
// }

// func fnB() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("recover in B", err)
// 		}
// 	}()
// 	panic("panic in B")
// }

// func fnC() {
// 	fmt.Println("fn C")
// }

//	func main() {
//		fnA()
//		fnB()
//		fnC()
//	}
// package main

// import "fmt"

// func fn() {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("抛出异常给管理员发生邮件")
// 			fmt.Println(err)
// 		}
// 	}()
// 	a := 10
// 	b := 0
// 	res := a / b
// 	fmt.Println("res =", res)
// }

//	func main() {
//		fn()
//	}
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func readFile(filename string) error {
// 	if filename == "main.go" {
// 		return nil
// 	}
// 	return errors.New("读取文件错误")
// }

// func fn() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("抛出异常给管理员发送邮件")
// 		}
// 	}()
// 	err := readFile("xxx.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("继续执行")
// }

//	func main() {
//		fn()
//	}
// package main

// import (
// 	"fmt"
// )

// func readFile(filename string) {
// 	if filename == "main.go" {
// 		return
// 	}
// 	panic("读取文件错误")
// }

// func fn() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("捕获到 panic：", err)
// 			fmt.Println("抛出异常给管理员发送邮件")
// 		}
// 	}()

// 	readFile("xxx.go")
// 	fmt.Println("继续执行")
// }

//	func main() {
//		fn()
//		fmt.Println("程序正常退出")
//	}
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func readFile(filename string) error {
// 	if filename == "main.go" {
// 		return nil
// 	}
// 	return errors.New("读取文件错误")
// }

// func fn() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("抛出异常给管理员发生邮件")
// 		}
// 	}()
// 	err := readFile("xxx.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("继续执行")
// }

//	func main() {
//		fn()
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		// 获取当前时间
//		now := time.Now()
//		fmt.Printf("Current time: %v\n", now)
//		year := now.Year()     // 年
//		month := now.Month()   // 月
//		day := now.Day()       // 天
//		hour := now.Hour()     // 时
//		minute := now.Minute() // 分
//		second := now.Second() // 秒
//		fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
//			year, month, day, hour, minute, second)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		now := time.Now()
//		fmt.Println(now.Format("2006-01-02 15:04:05"))
//		fmt.Println(now.Format("2006-01-02 03:04:05"))
//		fmt.Println(now.Format("2006/01/02 15:04"))
//		fmt.Println(now.Format("15:04 2006/01/02"))
//		fmt.Println(now.Format("2006/01/02"))
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		// 获取当前时间
//		now := time.Now()
//		unixTime := now.Unix()       // 秒级时间戳（10 位）
//		unixMilli := now.UnixMilli() // 毫秒级时间戳（13 位）
//		unixMicro := now.UnixMicro() // 微秒级时间戳（16 位）
//		unixNano := now.UnixNano()   // 纳秒级时间戳（19 位）
//		// 打印所有精度的时间戳
//		fmt.Printf("Current Unix Time (秒): %v\n", unixTime)
//		fmt.Printf("Current Unix Milli (毫秒): %v\n", unixMilli)
//		fmt.Printf("Current Unix Micro (微秒): %v\n", unixMicro)
//		fmt.Printf("Current Unix Nano (纳秒): %v\n", unixNano)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		// 获取当前时间
//		now := time.Now()
//		unixTime := now.Unix()       // 秒级时间戳（10 位）
//		unixMilli := now.UnixMilli() // 毫秒级时间戳（13 位）
//		unixMicro := now.UnixMicro() // 微秒级时间戳（16 位）
//		unixNano := now.UnixNano()   // 纳秒级时间戳（19 位）
//		// 打印所有精度的时间戳
//		fmt.Printf("Current Unix Time (秒): %v\n", unixTime)
//		fmt.Printf("Current Unix Milli (毫秒): %v\n", unixMilli)
//		fmt.Printf("Current Unix Micro (微秒): %v\n", unixMicro)
//		fmt.Printf("Current Unix Nano (纳秒): %v\n", unixNano)
//	}
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

//	func main() {
//		unixToTime(1758348998)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func unixToTime(timestamp int64) {
// 	timeObj := time.Unix(timestamp, 123456789)
// 	timeStr := timeObj.Format("2006-01-02 15:04:05")
// 	fmt.Println(timeStr)
// }

//	func main() {
//		unixToTime(1958948998)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		timeStr := "2026-02-17 22:12:22"
//		timeTemple := "2006-01-02 15:04:05"
//		timeObj, _ := time.ParseInLocation(timeTemple, timeStr, time.Local)
//		fmt.Println(timeObj.Unix())
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		timeStr := "2011-02-18 22:12:22"
//		timeTemple := "2006-01-02 15:04:05"
//		timeObj, _ := time.ParseInLocation(timeTemple, timeStr, time.Local)
//		fmt.Println(timeObj.Unix())
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		now := time.Now()
//		later := now.Add(time.Hour*4 + time.Minute*20)
//		fmt.Println(later.Hour(), later.Minute())
//		fmt.Println(later.Sub(now))
//		fmt.Println(now.Equal(later))
//		fmt.Println(now.Before(later))
//		fmt.Println(now.After(later))
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		ticker := time.NewTicker(time.Second)
//		n := 0
//		for v := range ticker.C {
//			fmt.Println(v.Second())
//			n++
//			if n > 5 {
//				ticker.Stop()
//				return
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		ticker := time.NewTicker(time.Second * 2)
//		n := 0
//		for v := range ticker.C {
//			fmt.Println(v.Second())
//			n++
//			if n > 5 {
//				ticker.Stop()
//				break
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		n := 0
//		for {
//			if n >= 5 {
//				break
//			}
//			time.Sleep(time.Second * 2)
//			fmt.Println(time.Now().Second())
//			n++
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		n := 0
//		ticker := time.NewTicker(time.Second)
//		for v := range ticker.C {
//			fmt.Println(v.Second())
//			n++
//			if n > 5 {
//				ticker.Stop()
//				break
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		n := 0
//		ticker := time.NewTicker(time.Second)
//		for v := range ticker.C {
//			fmt.Println(v.Format("2006-01-02 15:04:05"))
//			n++
//			if n > 5 {
//				ticker.Stop()
//				break
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		n := 0
//		for {
//			time.Sleep(time.Second)
//			n++
//			fmt.Println(n)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		a := 10
//		b := &a
//		fmt.Printf("a: %v, ptr(a): %p\n", a, &a)
//		fmt.Printf("b: %v, type: %T, ptr(b): %p\n", b, b, &b)
//	}
// package main

// import "fmt"

//	func main() {
//		a := 10
//		b := &a
//		fmt.Printf("Type of b: %T\n", b)
//		c := *b
//		fmt.Printf("Type of c: %T\n", c)
//		fmt.Printf("Value of c: %v\n", c)
//	}
// package main

// import "fmt"

// func modify1(x int) {
// 	x = 100
// }

// func modify2(x *int) {
// 	*x = 100
// }

//	func main() {
//		a := 10
//		modify1(a)
//		fmt.Println(a)
//		modify2(&a)
//		fmt.Println(a)
//	}
// package main

// import "fmt"

//	func main() {
//		var userinfo map[string]string
//		userinfo["username"] = "张三"
//		fmt.Println(userinfo)
//	}
// package main

// import "fmt"

//	func main() {
//		var a *int
//		*a = 100
//		fmt.Println(*a)
//	}
// package main

// import "fmt"

// func adder() func(int) int {
// 	var x int
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

//	func main() {
//		var fn = adder()
//		fmt.Println(fn(10))
//		fmt.Println(fn(20))
//		fmt.Println(fn(30))
//		fn2 := adder()
//		fmt.Println(fn2(40))
//		fmt.Println(fn2(50))
//	}
// package main

// import "fmt"

// func adder(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

//	func main() {
//		fn := adder(10)
//		fmt.Println(fn(10))
//		fmt.Println(fn(20))
//		fmt.Println(fn(30))
//		fn2 := adder(20)
//		fmt.Println(fn2(40))
//		fmt.Println(fn2(50))
//	}
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

//	func main() {
//		jpgFn := makeSuffix(".jpg")
//		txtFn := makeSuffix(".txt")
//		fmt.Println(jpgFn("test"))
//		fmt.Println(txtFn("test"))
//	}
// package main

// import "fmt"

// func calc(base int) (add func(int) int, sub func(int) int) {
// 	add = func(i int) int {
// 		base += i
// 		return base
// 	}
// 	sub = func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return
// }

//	func main() {
//		a := 10
//		fn, fn2 := calc(a)
//		fmt.Println(fn(1), fn2(2))
//		fmt.Println(fn(3), fn2(4))
//		fmt.Println(fn(5), fn2(6))
//	}
// package main

// import "fmt"

//	func main() {
//		var funcs []func()
//		for i := 0; i < 5; i++ {
//			funcs = append(funcs, func() {
//				fmt.Println(i)
//			})
//		}
//		for _, f := range funcs {
//			f()
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	var funcs []func()

// 	for i := range 5 {
// 		funcs = append(funcs, func(num int) func() {
// 			return func() {
// 				fmt.Println(i)
// 			}
// 		}(i))
// 	}

//		for _, f := range funcs {
//			f()
//		}
//	}
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

//	func main() {
//		fmt.Println(f1())
//		fmt.Println(f2())
//		fmt.Println(f3())
//		fmt.Println(f4())
//	}
// package main

// import "fmt"

// func fnA() {
// 	fmt.Println("fn A")
// }

// func fnB() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("recover in B", err)
// 		}
// 	}()
// 	panic("panic in B")
// }

// func fnC() {
// 	fmt.Println("fn C")
// }

//	func main() {
//		fnA()
//		fnB()
//		fnC()
//	}
// package main

// import "fmt"

// func fn() {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("抛出异常给管理员发送邮件")
// 			fmt.Println(err)
// 		}
// 	}()

// 	a := 10
// 	b := 0
// 	res := a / b
// 	fmt.Println("res =", res)
// }

//	func main() {
//		fn()
//	}
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func readFile(filename string) error {
// 	if filename == "main.go" {
// 		return nil
// 	}
// 	return errors.New("读取文件错误")
// }

// func fn() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("抛出异常给管理员发送邮件")
// 		}
// 	}()
// 	err := readFile("xxx.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("继续执行")
// }

//	func main() {
//		fn()
//	}
// package main

// import "fmt"

//	func main() {
//		a := new(int)
//		b := new(bool)
//		fmt.Printf("%T\n", a)
//		fmt.Printf("%T\n", b)
//		fmt.Println(*a)
//		fmt.Println(*b)
//	}
// package main

// import "fmt"

//	func main() {
//		var a *int = new(int)
//		*a = 10
//		fmt.Println(*a)
//	}
// package main

// import "fmt"

//	func main() {
//		var userinfo map[string]string
//		userinfo = make(map[string]string, 1e2)
//		userinfo["username"] = "哈基米"
//		fmt.Println(userinfo)
//	}
// package main

// import "fmt"

// type newInt int  // 类型定义
// type myInt = int // 类型别名

//	func main() {
//		var a newInt
//		var b myInt
//		fmt.Printf("Val: %v, Type: %T\n", a, a)
//		fmt.Printf("Val: %v, Type: %T\n", b, b)
//	}
// package main

//	type person struct {
//		name string
//		city string
//		age  int8
//	}
// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age int
// }

//	func main() {
//		var p1 person
//		p1.name = "哈基米"
//		p1.city = "北京"
//		p1.age = 18
//		fmt.Printf("p1 = %v\n", p1)
//		fmt.Printf("p1 = %#v\n", p1)
//	}
// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int8
// }

//	func main() {
//		var p2 = new(person)
//		p2.name = "曼波"
//		p2.age = 18
//		p2.city = "上海"
//		fmt.Printf("%T\n", p2)
//		fmt.Printf("p2 = %#v\n", p2)
//	}
// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

//	func main() {
//		var p2 = new(person)
//		(*p2).name = "manbo"
//		p2.age = 18
//		p2.city = "shanghai"
//		fmt.Printf("%T\n", p2)
//		fmt.Printf("p2 = %#v\n", p2)
//	}
// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

//	func main() {
//		p3 := &person{}
//		fmt.Printf("%T\n", p3)
//		fmt.Printf("p3 = %#v\n", p3)
//		p3.name = "hachimi"
//		p3.age = 30
//		p3.city = "杭州"
//		(*p3).age = 50
//		fmt.Printf("p3 = %#v\n", p3)
//	}
// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

//	func main() {
//		p4 := &person{
//			name: "manba",
//			city: "yili",
//			age:  18,
//		}
//		fmt.Printf("p4 = %#v\n", p4)
//	}
// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

//	func main() {
//		p6 := &person{city: "北京"}
//		fmt.Printf("%#v\n", p6)
//	}
// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

//	func main() {
//		p := person{
//			"LBJ",
//			"New York",
//			28,
//		}
//		fmt.Printf("p = %#v\n", p)
//	}
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) PrintInfo() {
// 	fmt.Printf("Name: %v, Age: %v\n", p.name, p.age)
// }

//	func main() {
//		p1 := Person{
//			"曼波",
//			25,
//		}
//		p1.PrintInfo()
//	}
// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int8
// }

// func (p Person) printInfo() {
// 	fmt.Printf("Name: %v\tAge: %d\n", p.name, p.age)
// }

// func (p *Person) setInfo(name string, age int) {
// 	p.name = name
// 	p.age = int8(age)
// }

//	func main() {
//		p1 := &Person{
//			name: "哈牛魔",
//			age:  29,
//		}
//		p1.printInfo()
//		p1.setInfo("ccn", 111)
//		p1.printInfo()
//	}
// package main

// import "fmt"

// type myInt int

// func (m myInt) SayHello() {
// 	fmt.Println("Hello int!")
// }

//	func main() {
//		var m1 myInt
//		m1.SayHello()
//		m1 = 100
//		fmt.Printf("%#v\t%T\n", m1, m1)
//	}
// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	name   string
// 	Sno    string
// }

//	func main() {
//		s := &Student{
//			Id:     1,
//			Gender: "女",
//			name:   "平泽唯",
//			Sno:    "s001",
//		}
//		fmt.Printf("%#v\n", s)
//		jsonByte, _ := json.Marshal(s)
//		jsonStr := string(jsonByte)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	name   string
// 	Sno    string
// }

//	func main() {
//		s := &Student{
//			Id:     1,
//			Gender: "女",
//			name:   "平泽唯",
//			Sno:    "s0001",
//		}
//		fmt.Printf("%#v\n", s)
//		jsonByte, _ := json.Marshal(s)
//		jsonStr := string(jsonByte)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// 	Sno    string
// }

//	func main() {
//		jsonStr := `{"Id":1, "Gender":"女", "Name":"平泽唯", "Sno":"s0001"}`
//		s := new(Student)
//		err := json.Unmarshal([]byte(jsonStr), s)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Printf("%#v\n%v\n", s, s.Name)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

//	func main() {
//		s := &Student{Id: 1, Gender: "女", Name: "平泽唯", Sno: "s0001"}
//		fmt.Printf("%#v\n", s)
//		jsonByte, _ := json.Marshal(s)
//		jsonStr := string(jsonByte)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // 学生
// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// // 班级
// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	c := &Class{
// 		Title:    "0301",
// 		Students: make([]Student, 0, 200),
// 	}

// 	for i := range 10 {
// 		s := Student{
// 			Name:   fmt.Sprintf("stu_%02d", i),
// 			Gender: "女",
// 			Id:     i,
// 		}
// 		c.Students = append(c.Students, s)
// 	}

// 	data, err := json.Marshal(c)

// 	if err != nil {
// 		fmt.Println("JSON marshal failed.")
// 		return
// 	}

//		fmt.Printf("JSON: %s\n", data)
//	}
// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// // 学生
// type Student struct {
// 	Id     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// }

// // 班级
// type Class struct {
// 	Title    string    `json:"title"`
// 	Students []Student `json:"students"`
// }

// func main() {
// 	c := &Class{
// 		Title:    "0302",
// 		Students: make([]Student, 0, 200),
// 	}

// 	for i := range 10 {
// 		s := Student{
// 			Name:   fmt.Sprintf("stu_%02d", i),
// 			Gender: "女",
// 			Id:     i,
// 		}
// 		c.Students = append(c.Students, s)
// 	}

// 	data, err := json.Marshal(c)

// 	if err != nil {
// 		fmt.Println("JSON marshal failed.")
// 		return
// 	}

//		fmt.Printf("JSON: %s\n", data)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	str := `{
//     "Title": "0302",
//     "Students": [
//         {
//             "Id": 0,
//             "Gender": "女",
//             "Name": "stu_00"
//         },
//         {
//             "Id": 1,
//             "Gender": "女",
//             "Name": "stu_01"
//         },
//         {
//             "Id": 2,
//             "Gender": "女",
//             "Name": "stu_02"
//         },
//         {
//             "Id": 3,
//             "Gender": "女",
//             "Name": "stu_03"
//         },
//         {
//             "Id": 4,
//             "Gender": "女",
//             "Name": "stu_04"
//         },
//         {
//             "Id": 5,
//             "Gender": "女",
//             "Name": "stu_05"
//         },
//         {
//             "Id": 6,
//             "Gender": "女",
//             "Name": "stu_06"
//         },
//         {
//             "Id": 7,
//             "Gender": "女",
//             "Name": "stu_07"
//         },
//         {
//             "Id": 8,
//             "Gender": "女",
//             "Name": "stu_08"
//         },
//         {
//             "Id": 9,
//             "Gender": "女",
//             "Name": "stu_09"
//         }
//     ]
// }
// 	`
// 	c := new(Class)
// 	err := json.Unmarshal([]byte(str), c)

//		if err != nil {
//			fmt.Println("JSON unmarshal failed!")
//			return
//		}
//		fmt.Printf("%#v\n", c)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	name   string
// 	Sno    string
// }

// func main() {
// 	s := &Student{
// 		Id:     1,
// 		Gender: "女",
// 		name:   "平泽唯",
// 		Sno:    "s0001",
// 	}
// 	fmt.Printf("%#v\n", s)

//		jsonByte, _ := json.Marshal(s)
//		jsonStr := string(jsonByte)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// 	Sno    string
// }

// func main() {
// 	jsonStr := `{"Id":1, "Gender":"女", "Name":"平泽唯", "Sno":"s0001"}`
// 	s := Student{}
// 	err := json.Unmarshal([]byte(jsonStr), &s)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

//		fmt.Printf("%#v\n%v\n", s, s.Name)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

// func main() {
// 	s := &Student{
// 		Id:     1,
// 		Gender: "女",
// 		Name:   "平泽唯",
// 		Sno:    "s0001",
// 	}
// 	fmt.Printf("%#v\n", s)

//		jsonByte, _ := json.Marshal(s)
//		jsonStr := string(jsonByte)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	c := &Class{
// 		Title:    "0302",
// 		Students: make([]Student, 0, 200),
// 	}

// 	for i := range 10 {
// 		s := Student{
// 			Name:   fmt.Sprintf("stu_%02d", i),
// 			Gender: "女",
// 			Id:     i,
// 		}
// 		c.Students = append(c.Students, s)
// 	}

// 	data, err := json.Marshal(c)

// 	if err != nil {
// 		fmt.Println("JSON marshal failed.")
// 		return
// 	}

//		fmt.Printf("JSON: %s\n", data)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	c := &Class{
// 		Title:    "0302",
// 		Students: make([]Student, 0, 200),
// 	}

// 	for i := range 10 {
// 		s := Student{
// 			Name:   fmt.Sprintf("stu_%02d", i),
// 			Gender: "女",
// 			Id:     i,
// 		}
// 		c.Students = append(c.Students, s)
// 	}

// 	data, err := json.Marshal(c)

// 	if err != nil {
// 		fmt.Println("JSON marshal failed.")
// 		return
// 	}

//		fmt.Printf("JSON: %s\n", data)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	str := `
// {
//     "Title": "0302",
//     "Students": [
//         {
//             "Id": 0,
//             "Gender": "女",
//             "Name": "stu_00"
//         },
//         {
//             "Id": 1,
//             "Gender": "女",
//             "Name": "stu_01"
//         },
//         {
//             "Id": 2,
//             "Gender": "女",
//             "Name": "stu_02"
//         },
//         {
//             "Id": 3,
//             "Gender": "女",
//             "Name": "stu_03"
//         },
//         {
//             "Id": 4,
//             "Gender": "女",
//             "Name": "stu_04"
//         },
//         {
//             "Id": 5,
//             "Gender": "女",
//             "Name": "stu_05"
//         },
//         {
//             "Id": 6,
//             "Gender": "女",
//             "Name": "stu_06"
//         },
//         {
//             "Id": 7,
//             "Gender": "女",
//             "Name": "stu_07"
//         },
//         {
//             "Id": 8,
//             "Gender": "女",
//             "Name": "stu_08"
//         },
//         {
//             "Id": 9,
//             "Gender": "女",
//             "Name": "stu_09"
//         }
//     ]
// }
// `

// 	c := new(Class)

// 	err := json.Unmarshal([]byte(str), c)

// 	if err != nil {
// 		fmt.Println("JSON unmarshal failed!")
// 		return
// 	}

//		fmt.Printf("%#v\n", c)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // 学生
// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// // 班级
// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	str := `
// {
//     "Title": "0302",
//     "Students": [
//         {
//             "Id": 0,
//             "Gender": "女",
//             "Name": "stu_00"
//         },
//         {
//             "Id": 1,
//             "Gender": "女",
//             "Name": "stu_01"
//         },
//         {
//             "Id": 2,
//             "Gender": "女",
//             "Name": "stu_02"
//         },
//         {
//             "Id": 3,
//             "Gender": "女",
//             "Name": "stu_03"
//         },
//         {
//             "Id": 4,
//             "Gender": "女",
//             "Name": "stu_04"
//         },
//         {
//             "Id": 5,
//             "Gender": "女",
//             "Name": "stu_05"
//         },
//         {
//             "Id": 6,
//             "Gender": "女",
//             "Name": "stu_06"
//         },
//         {
//             "Id": 7,
//             "Gender": "女",
//             "Name": "stu_07"
//         },
//         {
//             "Id": 8,
//             "Gender": "女",
//             "Name": "stu_08"
//         },
//         {
//             "Id": 9,
//             "Gender": "女",
//             "Name": "stu_09"
//         }
//     ]
// }
// `

// 	c := &Class{}

// 	err := json.Unmarshal([]byte(str), c)

// 	if err != nil {
// 		fmt.Println("JSON unmarshal failed!")
// 		return
// 	}

//		fmt.Printf("%#v\n", c)
//	}
// package main

// import (
// 	"fmt"
// )

// var x int8 = 20

// const pi = 3.14

// func init() {
// 	fmt.Println(x)
// }

//	func main() {
//		fmt.Println("Hello World!")
//	}
// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) Start() {
// 	fmt.Println(p.Name, "开机")
// }

// func (p Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// type Camera struct{}

// func (c Camera) Start() {
// 	fmt.Println("相机 开机")
// }

// func (c Camera) Stop() {
// 	fmt.Println("相机 关机")
// }

// func main() {
// 	var phone Usber = Phone{
// 		Name: "一加手机",
// 	}
// 	phone.Start()
// 	phone.Stop()

//		camera := Camera{}
//		var c Usber = camera
//		c.Start()
//		c.Stop()
//	}
// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) Start() {
// 	fmt.Println(p.Name, "开机")
// }

// func (p Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// type Camera struct{}

// func (c Camera) Start() {
// 	fmt.Println("相机 开机")
// }

// func (c Camera) Stop() {
// 	fmt.Println("相机 关机")
// }

// func main() {
// 	var phone Usber = Phone{
// 		Name: "一加手机",
// 	}
// 	phone.Start()
// 	phone.Stop()

//		camera := Camera{}
//		var c Usber = camera
//		c.Start()
//		c.Stop()
//	}
// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) Start() {
// 	fmt.Println(p.Name, "开机")
// }

// func (p Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// type Computer struct {
// 	Name string
// }

// func (c Computer) Work(usb Usber) {
// 	usb.Start()
// 	usb.Stop()
// }

//	func main() {
//		phone := Phone{
//			Name: "小米手机",
//		}
//		computer := Computer{}
//		computer.Work(phone)
//	}
// package main

// import "fmt"

// func main() {
// 	var x any
// 	s := "Hello World!"
// 	x = s
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

// 	i := 100
// 	x = i
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

//		b := true
//		x = b
//		fmt.Printf("Type: %T, Value: %v\n", x, x)
//	}
// package main

// import "fmt"

// func show(a any) {
// 	fmt.Printf("Type: %T, Value: %v\n", a, a)
// }

// func main() {
// 	var x any

// 	s := "Hello World!"
// 	x = s
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

// 	i := 100
// 	x = i
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

//		b := true
//		x = b
//		fmt.Printf("Type: %T, Value: %v\n", x, x)
//	}
// package main

// import "fmt"

//	func main() {
//		var studentInfo = make(map[string]any)
//		studentInfo["name"] = "平泽唯"
//		studentInfo["age"] = 18
//		studentInfo["married"] = false
//		fmt.Println(studentInfo)
//	}
// package main

// import "fmt"

//	func main() {
//		var studentInfo = make(map[string]any)
//		studentInfo["name"] = "平泽唯"
//		studentInfo["age"] = 18
//		studentInfo["married"] = false
//		fmt.Println(studentInfo)
//	}
// package main

// import "fmt"

//	func main() {
//		var slice = []any{"平泽唯", 20, true, 32.2}
//		fmt.Println(slice)
//	}
// package main

// import "fmt"

// func main() {
// 	var x any = "Hello Golang!"
// 	v, ok := x.(string)

//		if ok {
//			fmt.Printf("Type: %T, Value: %v\n", v, v)
//		} else {
//			fmt.Println("类型断言失败")
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	var x any = "Hello Golang!"
// 	v, ok := x.(string)

//		if ok {
//			fmt.Println(v)
//		} else {
//			fmt.Println("类型断言失败")
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	var x any = "Hello Golang!"
// 	v, ok := x.(string)

//		if ok {
//			fmt.Println(v)
//		} else {
//			fmt.Println("类型断言失败")
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	var x any = "Hello Golang!"
// 	v, ok := x.(string)

//		if ok {
//			fmt.Println(v)
//		} else {
//			fmt.Println("类型断言失败")
//		}
//	}
// package main

// import "fmt"

// func justifyType(x any) {
// 	switch x.(type) {
// 	case string:
// 		fmt.Println("string 类型")
// 	case int:
// 		fmt.Println("int 类型")
// 	case bool:
// 		fmt.Println("bool 类型")
// 	case float32:
// 		fmt.Println("float32 类型")
// 	case float64:
// 		fmt.Println("float64 类型")
// 	default:
// 		fmt.Println("Unsupport type!")
// 	}
// }

//	func main() {
//		justifyType(10)
//		justifyType("Hello World!")
//	}
// package main

// import "fmt"

// func justifyType(x any) {
// 	switch x.(type) {
// 	case string:
// 		fmt.Println("string 类型")
// 	case int:
// 		fmt.Println("int 类型")
// 	case bool:
// 		fmt.Println("bool 类型")
// 	case float32:
// 		fmt.Println("float32 类型")
// 	case float64:
// 		fmt.Println("float64 类型")
// 	default:
// 		fmt.Println("Unsupport type!")
// 	}
// }

//	func main() {
//		justifyType(10)
//		justifyType("Hello World")
//	}
// package main

// import "fmt"

// func justifyType(x any) {
// 	switch x.(type) {
// 	case string:
// 		fmt.Println("string 类型")
// 	case int:
// 		fmt.Println("int 类型")
// 	case bool:
// 		fmt.Println("bool 类型")
// 	case float32:
// 		fmt.Println("float32 类型")
// 	case float64:
// 		fmt.Println("float64 类型")
// 	default:
// 		fmt.Println("Unsopport type!")
// 	}
// }

//	func main() {
//		justifyType(10)
//		justifyType("Hello World")
//	}
// package main

// import "fmt"

// type Address struct {
// 	Name  string
// 	Phone int
// }

// func main() {
// 	userinfo := make(map[string]any)
// 	userinfo["username"] = "平泽唯"
// 	userinfo["age"] = 18
// 	userinfo["hobby"] = []string{"偷闲", "吃甜品"}

// 	fmt.Println(userinfo["hobby"].([]string)[1])

// 	address := Address{
// 		Name:  "田中井律",
// 		Phone: 1234567890,
// 	}
// 	userinfo["address"] = address

//		a, ok := userinfo["address"].(Address)
//		if ok {
//			fmt.Println(a.Name)
//		}
//	}
// package main

// import "fmt"

// type Address struct {
// 	Name  string
// 	Phone int
// }

// func main() {
// 	userinfo := make(map[string]any)
// 	userinfo["username"] = "平泽唯"
// 	userinfo["age"] = 18
// 	userinfo["hobby"] = []string{"偷闲", "吃甜品"}

// 	fmt.Println(userinfo["hobby"].([]string)[1])

// 	address := Address{
// 		Name:  "田中井律",
// 		Phone: 1234567890,
// 	}
// 	userinfo["address"] = address

//		a, ok := userinfo["address"].(Address)
//		if ok {
//			fmt.Println(a.Name)
//		}
//	}
// package main

// import "fmt"

// type Address struct {
// 	Name  string
// 	Phone int
// }

// func main() {
// 	userinfo := make(map[string]any)
// 	userinfo["username"] = "平泽唯"
// 	userinfo["age"] = 18
// 	userinfo["hobby"] = []string{"偷闲", "吃甜品"}

// 	fmt.Println(userinfo["hobby"].([]string)[1])

// 	address := Address{
// 		Name:  "田中井律",
// 		Phone: 1234567890,
// 	}

//		userinfo["address"] = address
//		a, ok := userinfo["address"].(Address)
//		if ok {
//			fmt.Println(a.Name)
//		}
//	}
// package main

// import "fmt"

// func justifyType(x any) {
// 	switch x.(type) {
// 	case string:
// 		fmt.Println("string 类型")
// 	case int:
// 		fmt.Println("int 类型")
// 	case bool:
// 		fmt.Println("bool 类型")
// 	case float32:
// 		fmt.Println("float32 类型")
// 	case float64:
// 		fmt.Println("float64 类型")
// 	default:
// 		fmt.Println("Unsupport type!")
// 	}
// }

//	func main() {
//		justifyType(10)
//		justifyType(3.14)
//		justifyType("Hello World")
//	}
// package main

// import "fmt"

// type Address struct {
// 	Name  string
// 	Phone int
// }

// func main() {
// 	userinfo := make(map[string]any)
// 	userinfo["username"] = "平泽唯"
// 	userinfo["age"] = 18
// 	userinfo["hobby"] = []string{"偷闲", "吃甜品"}

// 	fmt.Println(userinfo["hobby"].([]string)[1])

// 	address := Address{
// 		Name:  "田中井律",
// 		Phone: 1234567890,
// 	}
// 	userinfo["address"] = address

//		a, ok := userinfo["address"].(Address)
//		if ok {
//			fmt.Println(a.Name)
//		}
//	}
// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) Start() {
// 	fmt.Println(p.Name, "开机")
// }

// func (p Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// func main() {
// 	phone1 := Phone{
// 		Name: "小米手机",
// 	}
// 	var p1 Usber = phone1
// 	p1.Start()

//		var p2 Usber = &Phone{
//			Name: "苹果手机",
//		}
//		p2.Start()
//	}
// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p *Phone) Start() { fmt.Println(p.Name, "开机") }

// func (p *Phone) Stop() { fmt.Println(p.Name, "关机") }

//	func main() {
//		var p2 Usber = &Phone{
//			Name: "苹果手机",
//		}
//		p2.Start()
//	}
// package main

// import "fmt"

// type Animaler interface {
// 	SetName(string)
// }

// type Peter interface {
// 	GetName() string
// }

// type Dog struct {
// 	Name string
// }

// func (d *Dog) SetName(name string) {
// 	d.Name = name
// }

// func (d Dog) GetName() string {
// 	return d.Name
// }

// func main() {
// 	dog := &Dog{Name: "布鲁斯"}

// 	var d1 Animaler = dog
// 	var d2 Peter = dog

//		fmt.Println(d2.GetName())
//		d1.SetName("布莱恩")
//		fmt.Println(d2.GetName())
//	}
// package main

// import "fmt"

// type Animaler interface {
// 	SetName(string)
// }

// type Peter interface {
// 	GetName() string
// }

// type Dog struct {
// 	Name string
// }

// func (d *Dog) SetName(name string) {
// 	d.Name = name
// }

// func (d Dog) GetName() string {
// 	return d.Name
// }

// func main() {
// 	dog := &Dog{Name: "布鲁斯"}

// 	var d1 Animaler = dog
// 	var d2 Peter = dog

//		fmt.Println(d2.GetName())
//		d1.SetName("布莱恩")
//		fmt.Println(d2.GetName())
//	}
// package main

// import "fmt"

// type Runner interface {
// 	Run()
// }

// type Swimer interface {
// 	Swim()
// }

// type Animal interface {
// 	Runner
// 	Swimer
// }

// type Cat struct {
// 	Name string
// }

// func (c Cat) Run() {
// 	fmt.Println("基米跑步")
// }

// func (c Cat) Swim() {
// 	fmt.Println("基米游泳")
// }

//	func main() {
//		var cat Animal = &Cat{Name: "哈基米"}
//		cat.Run()
//		cat.Swim()
//	}
// package main

// import "fmt"

// type Runner interface {
// 	Run()
// }

// type Swimer interface {
// 	Swim()
// }

// type Animal interface {
// 	Runner
// 	Swimer
// }

// type Cat struct {
// 	Name string
// }

// func (c Cat) Run() {
// 	fmt.Println("基米跑步")
// }

// func (c Cat) Swim() {
// 	fmt.Println("基米游泳")
// }

//	func main() {
//		var cat Animal = &Cat{Name: "哈基米"}
//		cat.Run()
//		cat.Swim()
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func fn() {
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		go fn()
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(50 * time.Millisecond)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		timeStr := "2026-02-17 22:12:22"
//		timeTemple := "2006-01-02 15:04:05"
//		timeObj, _ := time.ParseInLocation(timeTemple, timeStr, time.Local)
//		fmt.Println(timeObj.Unix())
//	}
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

//			if n > 5 {
//				ticker.Stop()
//				return
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		n := 0
//		for {
//			n++
//			time.Sleep(time.Second)
//			fmt.Println(time.Now())
//			if n > 5 {
//				return
//			}
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		scoreMap := make(map[string]int, 8)
//		scoreMap["哈基米"] = 90
//		scoreMap["曼波"] = 100
//		fmt.Println(scoreMap)
//		fmt.Println(scoreMap["曼波"])
//		fmt.Printf("Type of scoreMap: %T\n", scoreMap)
//	}
// package main

// import "fmt"

//	func main() {
//		userinfo := map[string]string{
//			"username": "叮咚鸡",
//			"password": "大狗叫",
//		}
//		fmt.Println(userinfo)
//	}
// package main

// import "fmt"

//	func main() {
//		scoreMap := map[string]int{
//			"孙连城": 90,
//			"李达康": 100,
//			"沙瑞金": 60,
//		}
//		fmt.Println(scoreMap)
//		delete(scoreMap, "沙瑞金")
//		fmt.Println(scoreMap)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func fn() {
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		go fn()
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(50 * time.Millisecond)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func fn() {
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		go fn()
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(50 * time.Millisecond)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(1)
//		go fn(&wg)
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(10 * time.Millisecond)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(1)
//		go fn(&wg)
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(50 * time.Millisecond)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ticker := time.NewTicker(10 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 5 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(1)
//		go fn(&wg)
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(50 * time.Millisecond)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 5 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(1)
//		go fn(&wg)
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(50 * time.Millisecond)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go fn(&wg)

//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(50 * time.Millisecond)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func fn(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello Goroutine!", num)
// }

//	func main() {
//		var wg sync.WaitGroup
//		for i := range 10 {
//			wg.Add(1)
//			go fn(i, &wg)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func fn(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello Goroutine!", num)
// }

//	func main() {
//		var wg sync.WaitGroup
//		for i := range 10 {
//			wg.Add(1)
//			go fn(i, &wg)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func fn(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello Goroutine!", num)
// }

//	func main() {
//		var wg sync.WaitGroup
//		for i := range 10 {
//			wg.Add(1)
//			go fn(i, &wg)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func fn(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello Goroutine!", num)
// }

// func main() {
// 	var wg sync.WaitGroup

//		for i := range 10 {
//			wg.Add(1)
//			go fn(i, &wg)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	cpuNum := runtime.NumCPU()
// 	fmt.Println("cpuNum =", cpuNum)

//		runtime.GOMAXPROCS(cpuNum - 1)
//		fmt.Println("OK!")
//	}
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	cpuNum := runtime.NumCPU()
// 	fmt.Println("cpuNum =", cpuNum)

//		runtime.GOMAXPROCS(cpuNum - 1)
//		fmt.Println("OK!")
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	primeCount := 0

// 	for num := 2; num < 1.2e5; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 			}
// 		}
// 		if isPrime {
// 			primeCount++
// 		}
// 	}
// 	elapsed := time.Since(start)

//		fmt.Println(primeCount, elapsed)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	primeCount := 0

// 	for num := 2; num < 1.2e5; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 			}
// 		}
// 		if isPrime {
// 			primeCount++
// 		}
// 	}
// 	elapsed := time.Since(start)

//		fmt.Println(primeCount, elapsed)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func countPrime(start, end int, wg *sync.WaitGroup, res *int) {
// 	defer wg.Done()
// 	count := 0
// 	for num := start; num < end; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 			}
// 		}
// 		if isPrime {
// 			count++
// 		}
// 	}
// 	*res = count
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	parts := 10
// 	interval := 120000 / parts
// 	res := make([]int, parts)

// 	for i := range parts {
// 		wg.Add(1)
// 		s := i * interval
// 		e := (i + 1) * interval
// 		if s < 2 {
// 			s = 2
// 		}
// 		go countPrime(s, e, &wg, &res[i])
// 	}

// 	wg.Wait()
// 	total := 0
// 	for _, v := range res {
// 		total += v
// 	}

//		fmt.Println(total, time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func countPrime(start, end int, wg *sync.WaitGroup, res *int) {
// 	defer wg.Done()
// 	count := 0
// 	for num := start; num < end; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			count++
// 		}
// 	}
// 	*res = count
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	parts := 10
// 	intervals := 120000 / parts
// 	res := make([]int, parts)

// 	for i := range parts {
// 		s := max(2, i*intervals)
// 		e := (i + 1) * intervals
// 		wg.Add(1)
// 		go countPrime(s, e, &wg, &res[i])
// 	}
// 	wg.Wait()

// 	count := 0
// 	for _, v := range res {
// 		count += v
// 	}

//		fmt.Println(count, time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func countPrime(start, end int, wg *sync.WaitGroup, res *int) {
// 	defer wg.Done()

// 	for num := start; num < end; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			(*res)++
// 		}
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	parts := 10
// 	intervals := 1200000 / parts
// 	res := make([]int, parts)

// 	for i := range parts {
// 		s := max(2, i*intervals)
// 		e := (i + 1) * intervals
// 		wg.Add(1)
// 		go countPrime(s, e, &wg, &res[i])
// 	}
// 	wg.Wait()

// 	count := 0
// 	for _, v := range res {
// 		count += v
// 	}

//		fmt.Println(count, time.Since(start))
//	}
// package main

// import "fmt"

//	func main() {
//		ch := make(chan int, 1)
//		ch <- 10
//		<-ch
//		ch <- 12
//		<-ch
//		fmt.Println("Success")
//	}
// package main

// import "fmt"

//	func main() {
//		ch := make(chan int, 1)
//		ch <- 10
//		<-ch
//		ch <- 12
//		<-ch
//		ch <- 17
//		fmt.Println("Success")
//	}
// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 5)

// 	for i := range 5 {
// 		ch <- i + 1
// 	}
// 	close(ch)

//		for v := range ch {
//			fmt.Print(v, " ")
//		}
//		fmt.Println()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func write(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := range 10 {
// 		fmt.Printf("写入数据：%v\n", i+1)
// 		ch <- i + 1
// 	}

// 	close(ch)
// }

// func read(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for v := range ch {
// 		fmt.Printf("读取数据：%v\n", v)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(2)
// 	go write(ch, &wg)
// 	go read(ch, &wg)
// 	wg.Wait()

//		fmt.Println("读取完毕")
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func write(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := range 10 {
// 		fmt.Printf("写入数据：%v\n", i+1)
// 		ch <- i + 1
// 	}

// 	close(ch)
// }

// func read(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for v := range ch {
// 		fmt.Printf("读取数据：%v\n", v)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(2)
// 	go write(ch, &wg)
// 	go read(ch, &wg)
// 	wg.Wait()

//		fmt.Println("读取完毕")
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func write(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := range 10 {
// 		fmt.Printf("写入数据：%v\n", i+1)
// 		ch <- i + 1
// 	}

// 	close(ch)
// }

// func read(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for v := range ch {
// 		fmt.Printf("读取数据：%v\n", v)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(2)
// 	go write(ch, &wg)
// 	go read(ch, &wg)
// 	wg.Wait()

//		fmt.Println("读取完毕")
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func write(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := range 10 {
// 		fmt.Printf("写入数据：%v\n", i+1)
// 		ch <- i + 1
// 	}

// 	close(ch)
// }

// func read(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for v := range ch {
// 		fmt.Printf("读取数据：%v\n", v)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(2)
// 	go write(ch, &wg)
// 	go read(ch, &wg)
// 	wg.Wait()

//		fmt.Println("读取完毕")
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func genNums(nums chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := 2; num < 120; num++ {
// 		nums <- num
// 	}
// 	close(nums)
// }

// func filterPrime(nums, primes chan int, exit chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range nums {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			primes <- num
// 		}
// 	}
// 	exit <- true
// }

// func printPrime(primes chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for prime := range primes {
// 		fmt.Print(prime, " ")
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workers := 5
// 	nums := make(chan int, 10)
// 	primes := make(chan int, 10)
// 	exit := make(chan bool, workers)

// 	wg.Add(1)
// 	go genNums(nums, &wg)

// 	for range workers {
// 		wg.Add(1)
// 		go filterPrime(nums, primes, exit, &wg)
// 	}

// 	wg.Add(1)
// 	go printPrime(primes, &wg)

// 	wg.Go(func() {
// 		for range workers {
// 			<-exit
// 		}
// 		close(primes)
// 	})

// 	wg.Wait()

//		fmt.Println()
//		fmt.Println(time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func genNums(nums chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := 2; num < 120; num++ {
// 		nums <- num
// 	}
// 	close(nums)
// }

// func filterPrime(nums, primes chan int, exit chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range nums {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			primes <- num
// 		}
// 	}
// 	exit <- true
// }

// func printPrime(primes chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for prime := range primes {
// 		fmt.Print(prime, " ")
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workers := 5
// 	nums := make(chan int, 10)
// 	primes := make(chan int, 10)
// 	exit := make(chan bool, workers)

// 	wg.Add(1)
// 	go genNums(nums, &wg)

// 	for range workers {
// 		wg.Add(1)
// 		go filterPrime(nums, primes, exit, &wg)
// 	}

// 	wg.Add(1)
// 	go printPrime(primes, &wg)

// 	wg.Go(func() {
// 		for range workers {
// 			<-exit
// 		}
// 		close(primes)
// 	})

// 	wg.Wait()

//		fmt.Println()
//		fmt.Println(time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func genNums(nums chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := 2; num < 120; num++ {
// 		nums <- num
// 	}
// 	close(nums)
// }

// func filterPrime(nums, primes chan int, exit chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range nums {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			primes <- num
// 		}
// 	}
// 	exit <- true
// }

// func printPrime(primes chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for prime := range primes {
// 		fmt.Print(prime, " ")
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workers := 5
// 	nums := make(chan int, 10)
// 	primes := make(chan int, 10)
// 	exit := make(chan bool, workers)

// 	wg.Add(1)
// 	go genNums(nums, &wg)

// 	for range workers {
// 		wg.Add(1)
// 		go filterPrime(nums, primes, exit, &wg)
// 	}

// 	wg.Add(1)
// 	go printPrime(primes, &wg)

// 	wg.Go(func() {
// 		for range workers {
// 			<-exit
// 		}
// 		close(primes)
// 	})

// 	wg.Wait()

//		fmt.Println()
//		fmt.Println(time.Since(start))
//	}
// package main

// import "fmt"

// func main() {
// 	ch1 := make(chan int, 2)
// 	ch1 <- 10
// 	ch1 <- 12
// 	m1 := <-ch1
// 	m2 := <-ch1
// 	fmt.Println(m1, m2)

//		ch2 := make(chan<- int, 2)
//		ch2 <- 10
//		ch2 <- 12
//		ch3 := make(<-chan int, 2)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func genNums(nums chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := 2; num < 120; num++ {
// 		nums <- num
// 	}
// 	close(nums)
// }

// func filterPrime(nums <-chan int, primes chan<- int, exit chan<- bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range nums {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			primes <- num
// 		}
// 	}
// 	exit <- true
// }

// func printPrime(primes <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for prime := range primes {
// 		fmt.Println(prime)
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workers := 5
// 	nums := make(chan int, 10)
// 	primes := make(chan int, 10)
// 	exit := make(chan bool, workers)

// 	wg.Add(1)
// 	go genNums(nums, &wg)

// 	for range workers {
// 		wg.Add(1)
// 		go filterPrime(nums, primes, exit, &wg)
// 	}

// 	wg.Add(1)
// 	go printPrime(primes, &wg)

// 	wg.Go(func() {
// 		for range workers {
// 			<-exit
// 		}
// 		close(primes)
// 	})

// 	wg.Wait()

//		fmt.Println()
//		fmt.Println(time.Since(start))
//	}
// package main

// import "fmt"

//	func main() {
//		switch n := 7; n {
//		case 1, 3, 5, 7, 9:
//			fmt.Println("奇数")
//		case 2, 4, 6, 8:
//			fmt.Println("偶数")
//		default:
//			fmt.Println(n)
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		age := 56
//		switch {
//		case age < 25:
//			fmt.Println("好好学习")
//		case age <= 60:
//			fmt.Println("好好工作")
//		case age > 60:
//			fmt.Println("好好享受")
//		default:
//			fmt.Println("好好活着")
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		str := "a"
//		switch str {
//		case "a":
//			fmt.Println("a")
//			fallthrough
//		case "b":
//			fmt.Println("b")
//			fallthrough
//		case "c":
//			fmt.Println("c")
//			fallthrough
//		default:
//			fmt.Println("...")
//		}
//	}
// package main

// import "fmt"

// func main() {
// label:
//
//		for i := range 2 {
//			for j := range 3 {
//				if j == 3 {
//					break label
//				}
//				fmt.Printf("i = %v, j = %v\n", i, j)
//			}
//		}
//	}
// package main

// import "fmt"

// func main() {
// label:
//
//		for i := range 2 {
//			for j := range 10 {
//				if j == 3 {
//					break label
//				}
//				fmt.Printf("i = %v, j = %v\n", i, j)
//			}
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		for i := range 2 {
//			for j := range 4 {
//				if j == 2 {
//					continue
//				}
//				fmt.Printf("i = %v, j = %v\n", i, j)
//			}
//		}
//	}
// package main

// import "fmt"

// func main() {
// label:
//
//		for i := range 2 {
//			for j := range 4 {
//				if j == 2 {
//					continue label
//				}
//				fmt.Printf("i = %v, j = %v\n", i, j)
//			}
//		}
//	}
// package main

// import "fmt"

// func main() {
// label:
//
//		for i := range 2 {
//			for j := range 4 {
//				if j == 2 {
//					continue label
//				}
//				fmt.Printf("i = %v, j = %v\n", i, j)
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		sort.Ints(intSlice)
//		sort.Float64s(floatSlice)
//		sort.Strings(stringSlice)
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
//		sort.Sort(sort.Reverse(sort.Float64Slice(floatSlice)))
//		sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
//		sort.Sort(sort.Reverse(sort.Float64Slice(floatSlice)))
//		sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main() {
// 	intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
// 	floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
// 	stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

// 	slices.Sort(intSlice)
// 	slices.Sort(floatSlice)
// 	slices.Sort(stringSlice)

//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main() {
// 	intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
// 	floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
// 	stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

// 	slices.SortFunc(intSlice, func(a, b int) int {
// 		if a > b {
// 			return 1
// 		} else if a < b {
// 			return -1
// 		} else {
// 			return 0
// 		}
// 	})
// 	slices.SortFunc(floatSlice, func(a, b float64) int {
// 		if a > b {
// 			return -1
// 		} else if a < b {
// 			return 1
// 		} else {
// 			return 0
// 		}
// 	})
// 	slices.SortFunc(stringSlice, func(a, b string) int {
// 		if a > b {
// 			return 1
// 		} else if a < b {
// 			return -1
// 		} else {
// 			return 0
// 		}
// 	})

//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import "fmt"

// func main() {
// 	intChan := make(chan int, 10)
// 	for i := range 10 {
// 		intChan <- i
// 	}

// 	stringChan := make(chan string, 5)
// 	for i := range 5 {
// 		stringChan <- "Hello, " + fmt.Sprintf("%d", i)
// 	}

//		for {
//			select {
//			case v := <-intChan:
//				fmt.Printf("从 intChan 读取到数据：%d\n", v)
//			case v := <-stringChan:
//				fmt.Printf("从 stringChan 读取到数据：%s\n", v)
//			default:
//				fmt.Printf("取完所有数据")
//				return
//			}
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	intChan := make(chan int, 10)
// 	for i := range 10 {
// 		intChan <- i
// 	}

// 	stringChan := make(chan string, 5)
// 	for i := range 5 {
// 		stringChan <- "Hello, " + fmt.Sprintf("%d", i)
// 	}

//		for {
//			select {
//			case v := <-intChan:
//				fmt.Printf("从 intChan 读取的数据：%d\n", v)
//			case v := <-stringChan:
//				fmt.Printf("从 stringChan 读取的数据：%s\n", v)
//			default:
//				fmt.Printf("取完所有数据")
//				return
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHello() {
// 	for range 10 {
// 		time.Sleep(time.Millisecond * 100)
// 		fmt.Println("Hello World!")
// 	}
// }

// func test() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Printf("发生错误：%v\n", err)
// 		}
// 	}()
// 	var myMap map[int]string
// 	myMap[0] = "Golang"
// }

// func main() {
// 	go sayHello()
// 	go test()

//		for i := range 10 {
//			fmt.Println("main() ok =", i)
//			time.Sleep(time.Millisecond * 100)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func sayHello(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for range 10 {
// 		time.Sleep(time.Millisecond * 100)
// 		fmt.Println("Hello World!")
// 	}
// }

// func test(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Printf("发生错误：%v\n", err)
// 		}
// 	}()
// 	var myMap map[int]string
// 	myMap[0] = "Golang"
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go sayHello(&wg)
// 	go test(&wg)

//		for i := range 10 {
//			fmt.Println("main() ok =", i)
//			time.Sleep(time.Millisecond * 100)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap map[int]int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	numMap[num] = res
// }

// func main() {
// 	numMap := make(map[int]int)
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 60; i++ {
// 		wg.Add(1)
// 		go factorial(i, numMap, &wg)
// 	}

// 	wg.Wait()

//		for i, v := range numMap {
//			fmt.Printf("map[%d] = %d\n", i, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap map[int]int, wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	mutex.Lock()
// 	numMap[num] = res
// 	mutex.Unlock()
// }

// func main() {
// 	numMap := make(map[int]int)
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go factorial(i, numMap, &wg, &mutex)
// 	}

// 	wg.Wait()

//		for i, v := range numMap {
//			fmt.Printf("map[%d] = %d\n", i, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap map[int]int, wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	mutex.Lock()
// 	numMap[num] = res
// 	mutex.Unlock()
// }

// func main() {
// 	numMap := make(map[int]int)
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go factorial(i, numMap, &wg, &mutex)
// 	}

// 	wg.Wait()

//		for i, v := range numMap {
//			fmt.Printf("map[%d] = %v\n", i, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func write(wg *sync.WaitGroup, mutex *sync.RWMutex) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	fmt.Println("---执行写操作")
// 	time.Sleep(time.Millisecond * 100)
// 	mutex.Unlock()
// }

// func read(wg *sync.WaitGroup, mutex *sync.RWMutex) {
// 	defer wg.Done()
// 	mutex.RLock()
// 	fmt.Println("***执行读操作")
// 	time.Sleep(time.Microsecond * 100)
// 	mutex.RUnlock()
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	var mutex sync.RWMutex

// 	for range 10 {
// 		wg.Add(1)
// 		go write(&wg, &mutex)
// 	}

// 	for range 10 {
// 		wg.Add(1)
// 		go read(&wg, &mutex)
// 	}

// 	wg.Wait()

//		fmt.Println(time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func write(wg *sync.WaitGroup, mutex *sync.RWMutex) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	fmt.Println("---执行写操作")
// 	time.Sleep(100 * time.Millisecond)
// 	mutex.Unlock()
// }

// func read(wg *sync.WaitGroup, mutex *sync.RWMutex) {
// 	defer wg.Done()
// 	mutex.RLock()
// 	fmt.Println("***执行读操作")
// 	time.Sleep(100 * time.Millisecond)
// 	mutex.RUnlock()
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	var mutex sync.RWMutex

// 	for range 10 {
// 		wg.Add(1)
// 		go write(&wg, &mutex)
// 	}

// 	for range 10 {
// 		wg.Add(1)
// 		go read(&wg, &mutex)
// 	}

// 	wg.Wait()

//		fmt.Println(time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap *sync.Map, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	numMap.Store(num, res)
// }

// func main() {
// 	var numMap sync.Map
// 	var wg sync.WaitGroup

// 	wg.Add(10)
// 	for i := 1; i <= 10; i++ {
// 		go factorial(i, &numMap, &wg)
// 	}
// 	wg.Wait()

//		for k, v := range numMap.Range {
//			fmt.Printf("map[%d] = %d\n", k, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap *sync.Map, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	numMap.Store(num, res)
// }

// func main() {
// 	var numMap sync.Map
// 	var wg sync.WaitGroup

// 	wg.Add(10)
// 	for i := 1; i <= 10; i++ {
// 		go factorial(i, &numMap, &wg)
// 	}
// 	wg.Wait()

//		for k, v := range numMap.Range {
//			fmt.Printf("map[%d] = %d\n", k, v)
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	str1 := "Hello "
// 	str2 := "Golang!"

// 	fmt.Println(str1 + str2)

//		var str3 string = fmt.Sprintf("%v%v", str1, str2)
//		fmt.Println(str3)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		str := "123-456-789"
//		arr := strings.Split(str, "-")
//		fmt.Println(arr)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		var str = "this is golang"
//		var flag = strings.Contains(str, "golang")
//		fmt.Println(flag)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		str := "this is golang"
//		var flag bool = strings.HasPrefix(str, "this")
//		fmt.Println(flag)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "this is golang"
// 	var flag bool = strings.HasPrefix(str, "this")
// 	fmt.Println(flag)

//		flag = strings.HasSuffix(str, "go")
//		fmt.Println(flag)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "this is golang"
// 	var index = strings.Index(str, "is")
// 	fmt.Println(index)

//		index = strings.LastIndex(str, "is")
//		fmt.Println(index)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		var str = "123-456-789"
//		var arr = strings.Split(str, "-")
//		var str2 = strings.Join(arr, "*")
//		fmt.Println(str2)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	ID     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

//	func main() {
//		s1 := Student{
//			ID:     1,
//			Gender: "男",
//			Name:   "李四",
//			Sno:    "s0001",
//		}
//		data, _ := json.Marshal(s1)
//		jsonStr := string(data)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	ID     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

//	func main() {
//		s1 := Student{
//			ID:     1,
//			Gender: "男",
//			Name:   "李四",
//			Sno:    "s0001",
//		}
//		data, _ := json.Marshal(s1)
//		jsonStr := string(data)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectType(x any) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("Type: %v\n", t)
// }

//	func main() {
//		var f float32 = 12.5
//		reflectType(f)
//		var i int64 = 100
//		reflectType(i)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectType(x any) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("Type: %v\n", t)
// }

//	func main() {
//		var f float32 = 12.5
//		reflectType(f)
//		var i int64 = 100
//		reflectType(i)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectType(x any) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("Type: %v\n", t)
// }

//	func main() {
//		var f float32 = 12.5
//		reflectType(f)
//		var i int64 = 100
//		reflectType(i)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type MyInt int

// type MyStruct struct{ Name string }

// var a MyInt
// var b *MyStruct

//	func main() {
//		fmt.Println(reflect.TypeOf(a).Name())
//		fmt.Println(reflect.TypeOf(b).Name())
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectType(x any) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("Type: %v\tName: %v\tKind: %v\n", t, t.Name(), t.Kind())
// }

// type MyInt int64

// type Person struct {
// 	Name string
// 	Age  int
// }

// type Animal struct {
// 	Name string
// }

// func main() {
// 	var ptr *float32
// 	var mi MyInt
// 	var c rune
// 	var p = Person{
// 		Name: "平泽唯",
// 		Age:  17,
// 	}
// 	var a = Animal{"后藤一里"}
// 	var slice = []int{1, 2, 3, 4, 5}

//		reflectType(ptr)
//		reflectType(mi)
//		reflectType(c)
//		reflectType(p)
//		reflectType(a)
//		reflectType(slice)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectValue(x any) {
// 	v := reflect.ValueOf(x)
// 	c := v.Int() + 6
// 	fmt.Println(c)
// }

//	func main() {
//		var num int32 = 100
//		reflectValue(num)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectValue(x any) {
// 	v := reflect.ValueOf(x)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Int64:
// 		fmt.Printf("Type: int64\tValue: %d\n", v.Int())
// 	case reflect.Float32:
// 		fmt.Printf("Type: float32\tValue: %f\n", v.Float())
// 	case reflect.Float64:
// 		fmt.Printf("Type: float64\tValue: %f\n", v.Float())
// 	default:
// 		fmt.Println("Type: unknown")
// 	}
// }

// func main() {
// 	var a float32 = 3.14
// 	var b int64 = 100
// 	reflectValue(a)
// 	reflectValue(b)

//		c := reflect.ValueOf(10)
//		fmt.Printf("Type(c): %T\n", c)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectValue(x any) {
// 	v := reflect.ValueOf(x)
// 	if v.Elem().Kind() == reflect.Int64 {
// 		v.Elem().SetInt(200)
// 	}
// }

//	func main() {
//		var a int64 = 100
//		reflectValue(&a)
//		fmt.Println(a)
//	}
// package main

// import "reflect"

//	func reflectValue(x any) {
//		v := reflect.ValueOf(x)
//		if v.Elem().Kind() == reflect.Int64 {
//			v.Elem().SetInt(200)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectSetValue(x any) {
// 	v := reflect.ValueOf(x)
// 	if v.Elem().Kind() == reflect.Int64 {
// 		v.Elem().SetInt(200)
// 	}
// }

//	func main() {
//		var a int64 = 100
//		reflectSetValue(&a)
//		fmt.Println(a)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// func (s Student) GetInfo() string {
// 	return fmt.Sprintf("Name: %v\tAge: %v\tScore: %v\n", s.Name, s.Age, s.Score)
// }

// func (s *Student) SetInfo(name string, age, score int) {
// 	s.Name = name
// 	s.Age = age
// 	s.Score = score
// }

// func (s *Student) Print() {
// 	fmt.Println("打印方法")
// }

// func PrintStructField(s any) {
// 	fmt.Println("----结构体字段————")
// 	t := reflect.TypeOf(s)

// 	if (t.Kind() != reflect.Struct) && (t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体")
// 		return
// 	}

// 	field0 := t.Field(0)
// 	fmt.Println(field0.Name)
// 	fmt.Println(field0.Type)
// 	fmt.Println(field0.Tag.Get("json"))
// 	fmt.Println("----")

// 	field1, ok := t.FieldByName("Age")
// 	if ok {
// 		fmt.Println(field1.Name)
// 		fmt.Println(field1.Type)
// 		fmt.Println(field1.Tag.Get("json"))
// 		fmt.Println("----")
// 	}

// 	fmt.Println("字段数量：", t.NumField())
// }

// func PrintStructFunc(s any) {
// 	fmt.Println("----结构体方法----")
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if (t.Kind() != reflect.Struct) &&
// 		(t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体")
// 	}

// 	method0 := t.Method(0)
// 	fmt.Println(method0.Name)
// 	fmt.Println(method0.Type)
// 	fmt.Println("----")

// 	fmt.Println(t.NumMethod())
// 	fmt.Println("----")

// 	v.MethodByName("Print").Call(nil)
// 	fmt.Println("----")

// 	v.MethodByName("SetInfo").Call([]reflect.Value{
// 		reflect.ValueOf("后藤独"),
// 		reflect.ValueOf(16),
// 		reflect.ValueOf(60),
// 	})
// 	fmt.Println(v.MethodByName("GetInfo").Call(nil))
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		PrintStructField(stu)
//		PrintStructFunc(&stu)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// func (s Student) GetInfo() string {
// 	return fmt.Sprintf("Name: %v\tAge: %v\tScore: %v\n", s.Name, s.Age, s.Score)
// }

// func (s *Student) SetInfo(name string, age, score int) {
// 	s.Name = name
// 	s.Age = age
// 	s.Score = score
// }

// func (s *Student) Print() {
// 	fmt.Println("打印方法")
// }

// func PrintStructField(s any) {
// 	fmt.Println("----结构体字段----")
// 	t := reflect.TypeOf(s)

// 	if (t.Kind() != reflect.Struct) &&
// 		(t.Elem().Kind() != reflect.Struct) {
// 	}

// 	field0 := t.Field(0)
// 	fmt.Println(field0.Name)
// 	fmt.Println(field0.Type.Name())
// 	fmt.Println(field0.Type.Kind())
// 	fmt.Println(field0.Tag.Get("json"))
// 	fmt.Println("----")

// 	field1, ok := t.FieldByName("Age")
// 	if ok {
// 		fmt.Println(field1.Name)
// 		fmt.Println(field1.Type.Name())
// 		fmt.Println(field1.Type.Kind())
// 		fmt.Println(field1.Tag.Get("json"))
// 		fmt.Println("----")
// 	}

// 	fmt.Println("字段数量：", t.NumField())
// }

// func PrintStructFunc(s any) {
// 	fmt.Println("----结构体方法----")
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if (t.Kind() != reflect.Struct) && (t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体")
// 		return
// 	}

// 	method0 := t.Method(0)
// 	fmt.Println(method0.Name)
// 	fmt.Println(method0.Type.Name())
// 	fmt.Println(method0.Type.Kind())
// 	fmt.Println("----")

// 	fmt.Println(t.NumMethod())
// 	fmt.Println("----")

// 	v.MethodByName("Print").Call(nil)
// 	fmt.Println("----")

// 	v.MethodByName("SetInfo").Call([]reflect.Value{
// 		reflect.ValueOf("后藤独"),
// 		reflect.ValueOf(16),
// 		reflect.ValueOf(60),
// 	})
// 	fmt.Println(v.MethodByName("GetInfo").Call(nil))
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		PrintStructField(stu)
//		PrintStructFunc(&stu)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// func (s Student) GetInfo() string {
// 	return fmt.Sprintf(
// 		"Name: %v\tAge: %v\tScore: %v", s.Name, s.Age, s.Score,
// 	)
// }

// func reflectChangeStruct(s any) {
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if t.Elem().Kind() != reflect.Struct {
// 		fmt.Println("传入的不是结构体指针类型")
// 		return
// 	}

// 	name := v.Elem().FieldByName("Name")
// 	name.SetString("后藤独")

// 	age := v.Elem().FieldByName("Age")
// 	age.SetInt(17)
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		fmt.Println(stu.GetInfo())
//		reflectChangeStruct(&stu)
//		fmt.Println(stu.GetInfo())
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// func (s Student) GetInfo() string {
// 	return fmt.Sprintf("Name: %v\tAge: %v\tScore: %v", s.Name, s.Age, s.Score)
// }

// func reflectChangeStruct(s any) {
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if t.Elem().Kind() != reflect.Struct {
// 		fmt.Println("传入的不是结构体指针类型")
// 		return
// 	}

// 	name := v.Elem().FieldByName("Name")
// 	name.SetString("后藤独")

// 	age := v.Elem().FieldByName("Age")
// 	age.SetInt(18)
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		fmt.Println(stu.GetInfo())
//		reflectChangeStruct(&stu)
//		fmt.Println(stu.GetInfo())
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// func (s Student) GetInfo() string {
// 	return fmt.Sprintf("Name: %v\tAge: %v\tScore: %v", s.Name, s.Age, s.Score)
// }

// func reflectChangeStruct(s any) {
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if t.Elem().Kind() != reflect.Struct {
// 		fmt.Println("传入的不是结构体指针类型")
// 		return
// 	}

// 	name := v.Elem().FieldByName("Name")
// 	name.SetString("后藤一里")

// 	age := v.Elem().FieldByName("Age")
// 	age.SetInt(17)
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		fmt.Println(stu.GetInfo())
//		reflectChangeStruct(&stu)
//		fmt.Println(stu.GetInfo())
//	}
// package main

// import "fmt"

// func toSlice[T any](args ...T) []T {
// 	return args
// }

//	func main() {
//		strings := toSlice("hello", "world")
//		nums := toSlice(1, 2, 3)
//		fmt.Println(strings)
//		fmt.Println(nums)
//	}
// package main

// import "fmt"

// func toSlice[T any](args ...T) []T {
// 	return args
// }

//	func main() {
//		strs := toSlice("hello", "world")
//		nums := toSlice(1, 2, 3)
//		fmt.Println(strs)
//		fmt.Println(nums)
//	}
// package main

// import "fmt"

// type Stack[T any] struct {
// 	data []T
// }

// func (s *Stack[T]) Push(x T) {
// 	s.data = append(s.data, x)
// }

// func (s *Stack[T]) Pop() T {
// 	n := len(s.data)
// 	x := s.data[n-1]
// 	s.data = s.data[:n-1]
// 	return x
// }

// func main() {
// 	var s Stack[int]
// 	s.Push(10)
// 	s.Push(20)
// 	s.Push(30)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	s.Push(40)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())

//		var s2 Stack[string]
//		s2.Push("hello")
//		s2.Push("world")
//		fmt.Println(s2.Pop())
//		fmt.Println(s2.Pop())
//	}
// package main

// import "fmt"

// type Stack[T any] struct {
// 	data []T
// }

// func (s *Stack[T]) Push(x T) {
// 	s.data = append(s.data, x)
// }

// func (s *Stack[T]) Pop() T {
// 	n := len(s.data)
// 	x := s.data[n-1]
// 	s.data = s.data[:n-1]
// 	return x
// }

// func main() {
// 	var s Stack[int]
// 	s.Push(10)
// 	s.Push(20)
// 	s.Push(30)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	s.Push(40)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())

//		var s2 Stack[string]
//		s2.Push("hello")
//		s2.Push("world")
//		fmt.Println(s2.Pop())
//		fmt.Println(s2.Pop())
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// type Container[T any] interface {
// 	Len() int
// 	Add(T)
// 	Remove() T
// }

// func Swap[T any](x, y *T) {
// 	*x, *y = *y, *x
// }

//	func Print[T fmt.Stringer, U io.Reader](x T, y U) {
//		fmt.Println(x.String())
//		io.Copy(os.Stdout, y)
//	}
// package main

// import "fmt"

// func toSlice[T any](args ...T) []T {
// 	return args
// }

//	func main() {
//		strings := toSlice("hello", "world")
//		nums := toSlice(1, 2, 3)
//		fmt.Println(strings)
//		fmt.Println(nums)
//	}
// package main

// import "fmt"

// func toSlice[T any](args ...T) []T {
// 	return args
// }

//	func main() {
//		strings := toSlice("hello", "world")
//		nums := toSlice(1, 2, 3)
//		fmt.Println(strings)
//		fmt.Println(nums)
//	}
// package main

// import "fmt"

// type Stack[T any] struct {
// 	data []T
// }

// func (s *Stack[T]) Push(x T) {
// 	s.data = append(s.data, x)
// }

// func (s *Stack[T]) Pop() T {
// 	n := len(s.data)
// 	x := s.data[n-1]
// 	s.data = s.data[:n-1]
// 	return x
// }

// func main() {
// 	var s Stack[int]
// 	s.Push(10)
// 	s.Push(20)
// 	s.Push(30)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	s.Push(40)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())

//		var s2 Stack[string]
//		s2.Push("hello")
//		s2.Push("world")
//		fmt.Println(s2.Pop())
//		fmt.Println(s2.Pop())
//	}
// package main

//	type Container[T any] interface {
//		Len() int
//		Add(T)
//		Remove() T
//	}
// package main

// import (
// 	"cmp"
// 	"fmt"
// 	"slices"
// )

// func Sort[T cmp.Ordered](s []T) {
// 	slices.Sort(s)
// }

//	func main() {
//		nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
//		Sort(nums)
//		fmt.Println(nums)
//	}
// package main

// import (
// 	"fmt"
// 	"slices"
// )

//	func main() {
//		nums := []int{1, 2, 3, 3, 3, 4, 5, 6}
//		x := 3
//		idx, ok := slices.BinarySearch(nums, x)
//		if ok {
//			fmt.Println(idx)
//		}
//	}
// package main

// import (
// 	"cmp"
// 	"fmt"
// 	"slices"
// )

//	func main() {
//		nums := []int{6, 5, 4, 4, 3, 3, 3, 2, 1}
//		x := 3
//		idx, ok := slices.BinarySearchFunc(nums, x, func(a, b int) int { return cmp.Compare(b, a) })
//		if ok {
//			fmt.Println(idx)
//		}
//	}
// package main

// import "fmt"

// type Number interface {
// 	int | float64
// }

// func Add[T Number](a, b T) T {
// 	return a + b
// }

//	func main() {
//		fmt.Println(Add(10, 20))
//		fmt.Println(Add(3.14, 2.86))
//	}
// package main

// import "fmt"

// type AnyInt interface {
// 	~int
// }

// func PrintDouble[T AnyInt](val T) {
// 	fmt.Println(val * 2)
// }

// type MyInt int

// func main() {
// 	var a int = 10
// 	var b MyInt = 20

//		PrintDouble(a)
//		PrintDouble(b)
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		fmt.Println(file)
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		fmt.Println(file)
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		byteSlice := make([]byte, 128)
//		n, err := file.Read(byteSlice)
//		if err == io.EOF {
//			fmt.Println("文件读取完毕")
//			return
//		}
//		fmt.Printf("读取了 %d 字节数据\n", n)
//		fmt.Println(string(byteSlice[:n]))
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		byteSlice := make([]byte, 128)
//		n, err := file.Read(byteSlice)
//		if err == io.EOF {
//			fmt.Println("文件读取完毕")
//			return
//		}
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Printf("读取了 %d 字节数据\n", n)
//		fmt.Println(string(byteSlice[:n]))
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		var content []byte
//		var byteSlice = make([]byte, 128)
//		for {
//			n, err := file.Read(byteSlice)
//			if err == io.EOF {
//				fmt.Println("文件读取完毕")
//				break
//			}
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			content = append(content, byteSlice[:n]...)
//		}
//		fmt.Println(string(content))
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		var content []byte
//		var byteSlice = make([]byte, 128)
//		for {
//			n, err := file.Read(byteSlice)
//			if err == io.EOF {
//				fmt.Println("文件读取完毕")
//				break
//			}
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			content = append(content, byteSlice[:n]...)
//		}
//		fmt.Println(string(content))
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		var content []byte
//		var byteSlice = make([]byte, 128)
//		for {
//			n, err := file.Read(byteSlice)
//			content = append(content, byteSlice[:n]...)
//			if err == io.EOF {
//				fmt.Println("文件读取完毕")
//				break
//			}
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//		fmt.Println(string(content))
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		reader := bufio.NewReader(file)
//		fileStr := ""
//		for {
//			line, err := reader.ReadString('\n')
//			fileStr += line
//			if err == io.EOF {
//				fmt.Println("文件读取完毕")
//				break
//			}
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//		fmt.Println(fileStr)
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		reader := bufio.NewReader(file)
//		fileStr := ""
//		for {
//			line, err := reader.ReadString('\n')
//			fileStr += line
//			if err == io.EOF {
//				fmt.Println("文件读取完成")
//				break
//			}
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//		fmt.Println(fileStr)
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

//	func main() {
//		file, err := os.Open("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		reader := bufio.NewReader(file)
//		fileStr := ""
//		for {
//			line, err := reader.ReadString('\n')
//			fileStr += line
//			if err == io.EOF {
//				fmt.Println("文件读取完毕")
//				break
//			}
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//		fmt.Println(fileStr)
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		content, err := os.ReadFile("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(string(content))
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		content, err := os.ReadFile("./data/test.txt")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(string(content))
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		file, err := os.OpenFile("./data/test1.txt",
//			os.O_CREATE|os.O_RDWR, 0o644)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		str := "小 Go，你好！"
//		file.Write([]byte(str))
//		file.WriteString("写入数据")
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		file, err := os.OpenFile("./data/test1.txt", os.O_CREATE|os.O_RDWR, 0o644)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		str := "小 Go，你好！"
//		file.Write([]byte(str))
//		file.WriteString("写入数据")
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		file, err := os.OpenFile("./data/test1.txt", os.O_CREATE|os.O_RDWR, 0o644)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		str := "小 Go，你好！"
//		file.Write([]byte(str))
//		file.WriteString("写入数据")
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

//	func main() {
//		file, err := os.OpenFile("./data/test2.txt",
//			os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		writer := bufio.NewWriter(file)
//		for range 10 {
//			writer.WriteString("小 Go，你好！\n")
//		}
//		writer.Flush()
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

//	func main() {
//		file, err := os.OpenFile("./data/test.txt",
//			os.O_CREATE|os.O_RDWR, 0o644)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		writer := bufio.NewWriter(file)
//		for range 10 {
//			writer.WriteString("小 Go，你好！\n")
//		}
//		writer.Flush()
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		str := "Hello World!"
//		err := os.WriteFile("./data/test1.txt", []byte(str), 0o644)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		str := "Hello World!"
//		err := os.WriteFile("./testdata/test.txt", []byte(str), 0o644)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func CopyFile(dst, src string) error {
// 	content, err := os.ReadFile(src)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.WriteFile(dst, content, 0o644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//	func main() {
//		err := CopyFile("./data/test.txt", "main.go")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("拷贝成功")
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func CopyFile(dst, src string) error {
// 	content, err := os.ReadFile(src)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.WriteFile(dst, content, 0o644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//	func main() {
//		err := CopyFile("./data/test.txt", "./main.go")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("拷贝成功")
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func CopyFile(src, dst string) error {
// 	srcFile, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer srcFile.Close()
// 	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
// 	if err != nil {
// 		return err
// 	}
// 	defer dstFile.Close()

// 	buf := make([]byte, 32*1024)
// 	for {
// 		n, err := srcFile.Read(buf)
// 		if n > 0 {
// 			if _, werr := dstFile.Write(buf[:n]); werr != nil {
// 				return werr
// 			}
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//	func main() {
//		if err := CopyFile("./main.go", "./data/test.txt"); err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("拷贝成功")
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// func (s Student) GetInfo() string {
// 	return fmt.Sprintf("Name: %v\tAge: %v\tScore: %v\n", s.Name, s.Age, s.Score)
// }

// func (s *Student) SetInfo(name string, age, score int) {
// 	s.Name = name
// 	s.Age = age
// 	s.Score = score
// }

// func (s *Student) Print() {
// 	fmt.Println("打印方法")
// }

// func PrintStructField(s any) {
// 	fmt.Println("---- 结构体字段 ----")
// 	t := reflect.TypeOf(s)

// 	if (t.Kind() != reflect.Struct) && (t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体")
// 		return
// 	}

// 	field0 := t.Field(0)
// 	fmt.Println(field0.Name)
// 	fmt.Println(field0.Type.Name())
// 	fmt.Println(field0.Type.Kind())
// 	fmt.Println(field0.Tag.Get("json"))
// 	fmt.Println("----")

// 	field1, ok := t.FieldByName("Age")
// 	if ok {
// 		fmt.Println(field1.Name)
// 		fmt.Println(field1.Type)
// 		fmt.Println(field1.Tag.Get("json"))
// 		fmt.Println("----")
// 	}

// 	fmt.Printf("字段数量：%v\n", t.NumField())
// }

// func PrintStructFunc(s any) {
// 	fmt.Println("---- 结构体方法 ----")
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if (t.Kind() != reflect.Struct) &&
// 		(t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体")
// 		return
// 	}

// 	method0 := t.Method(0)
// 	fmt.Println(method0.Name)
// 	fmt.Println(method0.Type.Name())
// 	fmt.Println(method0.Type.Kind())
// 	fmt.Println("----")

// 	fmt.Println(t.NumMethod())
// 	fmt.Println("----")

// 	v.MethodByName("Print").Call(nil)
// 	fmt.Println("----")

// 	v.MethodByName("SetInfo").Call([]reflect.Value{
// 		reflect.ValueOf("后藤独"),
// 		reflect.ValueOf(16),
// 		reflect.ValueOf(60),
// 	})
// 	fmt.Println(v.MethodByName("GetInfo").Call(nil))
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		PrintStructField(stu)
//		PrintStructFunc(&stu)
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func CopyFile(dst, src string) error {
// 	content, err := os.ReadFile(src)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.WriteFile(dst, content, 0o644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//	func main() {
//		err := CopyFile("./data/test5.txt", "./main.go")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("拷贝成功")
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func CopyFile(src, dst string) error {
// 	srcFile, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer srcFile.Close()
// 	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
// 	if err != nil {
// 		return err
// 	}
// 	defer dstFile.Close()

// 	buf := make([]byte, 32*1024)
// 	for {
// 		n, err := srcFile.Read(buf)
// 		if n > 0 {
// 			if _, werr := dstFile.Write(buf[:n]); werr != nil {
// 				return werr
// 			}
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//	func main() {
//		if err := CopyFile("./main.go", "./data/test.txt"); err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("拷贝完成")
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func CopyFile(src, dst string) error {
// 	srcFile, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer srcFile.Close()
// 	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
// 	if err != nil {
// 		return err
// 	}
// 	defer dstFile.Close()

// 	buf := make([]byte, 32*1024)
// 	for {
// 		n, err := srcFile.Read(buf)
// 		if n > 0 {
// 			if _, werr := dstFile.Write(buf[:n]); werr != nil {
// 				return werr
// 			}
// 		}
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//	func main() {
//		if err := CopyFile("./main.go", "./data/test6.txt"); err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("拷贝完成")
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	err := os.MkdirAll("./data", 0o755)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	src, rerr := os.Open("./main.go")
// 	if rerr != nil {
// 		fmt.Println(rerr)
// 		return
// 	}
// 	defer src.Close()

// 	dst, werr := os.Create("./data/test.txt")
// 	if werr != nil {
// 		fmt.Println(werr)
// 		return
// 	}
// 	defer dst.Close()

//		n, err := io.Copy(dst, src)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Printf("成功拷贝 %v 字节\n", n)
//	}
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	err := os.MkdirAll("./data", 0o755)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	src, rerr := os.Open("./main.go")
// 	if rerr != nil {
// 		fmt.Println(rerr)
// 		return
// 	}
// 	defer src.Close()

// 	dst, werr := os.Create("./data/test.txt")
// 	if werr != nil {
// 		fmt.Println(werr)
// 		return
// 	}
// 	defer dst.Close()

// 	n, err := io.Copy(dst, src)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("成功拷贝了 %d 个字节\n", n)
// }
