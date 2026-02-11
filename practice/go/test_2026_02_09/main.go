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

// func main() {
// 	n := 2
// 	switch n {
// 	case 1, 3, 5, 7, 9:
// 		fmt.Println("奇数")
// 	case 0, 2, 4, 6, 8:
// 		fmt.Println("偶数")
// 	default:
// 		fmt.Println(n)
// 	}
// }
