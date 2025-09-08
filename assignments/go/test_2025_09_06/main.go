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

// import "fmt"

// func main() {
// 	// var str string = "Hello World!"
// 	// var str2 = "Hello Golang!"
// 	// var str3 = "Hello Mambo!"
// 	// fmt.Printf("%v\n%v\n%v\n", str, str2, str3)
// 	// str1 := "What can I\n say?"
// 	// fmt.Println(str1)
// 	// str2 := "C:\\Go\\bin"
// 	// fmt.Println(str2)
// 	// str3 := "\"ABC\""
// 	// fmt.Println(str3)
// 	// 使用转义符的字符串示例
// 	str := "Hello\nWorld\t!\n\"Go\" is \\awesome\\.\nHex: \x41\nUnicode: \u4F60\u597D\nLong Unicode: \U0001F600"

// 	// 打印字符串
// 	fmt.Println(str)

// 	// 原始字符串（不处理转义符）
// 	rawStr := `Hello\nWorld\t!
// "Go" is \awesome\.
// Hex: \x41
// Unicode: \u4F60\u597D`

// 	fmt.Println("原始字符串：")
// 	fmt.Println(rawStr)
// 	fmt.Println("Name\tAge\vHeight") // \t 对齐列，\v 可能换行
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// fmt.Println("Na\tAge\nH")
// 	// 	str := `Hello
// 	// World!`
// 	// 	fmt.Println(str, len(str))
// 	// str := "你好！"
// 	// fmt.Println(str, len(str))
// 	// str1 := "Hello "
// 	// str2 := "World!"
// 	// str3 := fmt.Sprintf("%v%v", str1, str2)
// 	// fmt.Println(str1 + str2)
// 	// fmt.Println(str3)
// 	// str := ("Hello " +
// 	// 	"World!")
// 	// fmt.Println(str)
// 	// str := "123-456-789"
// 	// arr := strings.Split(str, "-")
// 	// fmt.Printf("%v %T\n", arr, arr)
// 	// str2 := strings.Join(arr, "*")
// 	// fmt.Printf("%v %T\n", str2, str2)
// 	// arr := []string{"php", "java", "golang"}
// 	// fmt.Println(arr)
// 	// str := strings.Join(arr, " ")
// 	// fmt.Printf("%v %T\n", str, str)
// 	// str1 := "this is str"
// 	// str2 := "this"
// 	// flag := strings.Contains(str1, str2)
// 	// fmt.Println(flag)
// 	str1 := "this is str"
// 	fmt.Println(strings.Index(str1, "9s"))
// 	fmt.Println(strings.LastIndex(str1, "is"))
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("C:\\Code\\demo\\main.go")
// }

// package main

// import "fmt"

// func main() {
// 	str := `第一行
// 第二行
// 第三行`
// 	fmt.Println(str)
// }

// package main

// import "fmt"

// func main() {
// 	str := "this is str"
// 	fmt.Println(len(str))
// }

// package main

// import "fmt"

// func main() {
// 	str1 := "Hello "
// 	str2 := "Golang!"

// 	fmt.Println(str1 + str2)

// 	var str3 string = fmt.Sprintf("%v%v", str1, str2)
// 	fmt.Println(str3)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "123-456-789"
// 	arr := strings.Split(str, "-")
// 	fmt.Println(arr)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var str = "this is golang"
// 	var flag = strings.Contains(str, "golang")
// 	fmt.Println(flag)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "this is golang"
// 	var flag bool = strings.HasPrefix(str, "this")
// 	fmt.Println(flag)

// 	flag = strings.HasSuffix(str, "go")
// 	fmt.Println(flag)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "this is golang"
// 	var index = strings.Index(str, "is")
// 	fmt.Println(index)

// 	index = strings.LastIndex(str, "is")
// 	fmt.Println(index)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var str1 = "123-456-789"
// 	arr := strings.Split(str1, "-")
// 	str2 := strings.Join(arr, "*")
// 	fmt.Println(str2)
// }

// package main

// import "fmt"

// func main() {
// 	// a := 97
// 	// // 97 int32
// 	// fmt.Printf("%c %T\n", a, a)
// 	// var str = "this"
// 	// fmt.Printf("值：%v，类型：%T\n", str[2], str[2])
// 	// str := "this"
// 	// fmt.Println(len(str))
// 	// str2 := "你好"
// 	// fmt.Println(len(str2))
// 	// ch := '国'
// 	// fmt.Printf("%c %T\n", ch, ch)
// 	// str := "你好 World!"
// 	// // for i := 0; i < len(str); i++ {
// 	// // 	fmt.Printf("%v(%c) ", str[i], str[i])
// 	// // }
// 	// for _, s := range str {
// 	// 	fmt.Printf("%v(%c) ", s, s)
// 	// }
// 	s1 := "big"
// 	byteS1 := []byte(s1)
// 	byteS1[0] = 'p'
// 	fmt.Println(string(byteS1))
// 	s2 := "白萝卜"
// 	runeS1 := []rune(s2)
// 	runeS1[0] = '红'
// 	fmt.Println(string(runeS1))
// }

// package main

// import "fmt"

// func main() {
// 	a := 'a'
// 	b := 'b'
// 	// 直接输出 byte 的时候输出的是这个字符对应的码值
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	// 如果要输出该字符，需要格式化输出
// 	// %c 相应 Unicode 码点所表示的字符
// 	fmt.Printf("%c %c\n", a, b)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := "m"
// 	fmt.Println(len(a))
// 	b := "张"
// 	fmt.Println(len(b))
// }

// package main

// import "fmt"

// func main() {
// 	str := "哈基米Mambo"
// 	for i := 0; i < len(str); i++ {
// 		fmt.Printf("%v(%c) ", str[i], str[i])
// 	}
// 	fmt.Println()

// 	for _, ch := range str {
// 		fmt.Printf("%v(%c) ", ch, ch)
// 	}
// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	ch1 := "营"
// 	ch2 := '营'
// 	fmt.Printf("ch1 的类型为 %T，ch2 的类型为 %T\n", ch1, ch2)
// }

// package main

// import "fmt"

// func main() {
// 	str1 := "big"
// 	byteStr1 := []byte(str1)
// 	byteStr1[0] = 'p'
// 	fmt.Println(string(byteStr1))

// 	str2 := "红萝卜"
// 	runeStr2 := []rune(str2)
// 	runeStr2[0] = '白'
// 	fmt.Println(string(runeStr2))
// }

// package main

// import "fmt"

// func main() {
// 	str1 := "big"
// 	byteStr1 := []byte(str1)
// 	byteStr1[0] = 'p'
// 	fmt.Println(string(byteStr1))

// 	str2 := "白萝卜"
// 	runeStr2 := []rune(str2)
// 	runeStr2[0] = '红'
// 	fmt.Println(string(runeStr2))
// }

// package main

// import "fmt"

// func main() {
// 	str := "哈基米Mambo"
// 	for _, ch := range str {
// 		fmt.Printf("%v(%c) ", ch, ch)
// 	}
// }

// package main

// func main() {
// 	// var a int8 = 20
// 	// var b int16 = 40
// 	// fmt.Println(int16(a) + b)
// 	// var a float32 = 19.9
// 	// var b float64 = 40.1
// 	// fmt.Println(float64(a) + b)
// 	// var a float32 = 20.23
// 	// var b int = 40
// 	// fmt.Println(a + float32(b))
// 	// var a int8 = 20
// 	// var b int16 = 140
// 	// fmt.Println(int16(a) + b)
// 	// var i int = 20
// 	// var f float64 = 12.456
// 	// var t bool = true
// 	// var b byte = 'a'
// 	// strI := fmt.Sprintf("%d", i)
// 	// strF := fmt.Sprintf("%f", f)
// 	// strT := fmt.Sprintf("%t", t)
// 	// strB := fmt.Sprintf("%c", b)
// 	// fmt.Println(strI)
// 	// fmt.Println(strF)
// 	// fmt.Println(strT)
// 	// fmt.Println(strB)
// 	// i := 20
// 	// strI := strconv.FormatInt(int64(i), 10)
// 	// fmt.Printf("值：%v，类型：%T\n", strI, strI)
// 	// f := 3.1415926
// 	// strF := strconv.FormatFloat(float64(f), 'f', 2, 32)
// 	// fmt.Printf("值：%v，类型：%T\n", strF, strF)
// 	// t := true
// 	// strT := strconv.FormatBool(t)
// 	// fmt.Printf("值：%v，类型：%T\n", strT, strT)
// 	// a := 'a'
// 	// str4 := strconv.FormatUint(uint64(a), 10)
// 	// fmt.Printf("值：%v，类型：%T\n", str4, str4)
// 	// str := "123456"
// 	// fmt.Printf("%v %T\n", str, str)
// 	// num, _ := strconv.ParseFloat(str, 64)
// 	// fmt.Printf("%v %T\n", num, num)
// 	// b, _ := strconv.ParseBool("true")
// 	// fmt.Printf("%v %T\n", b, b)

// }

// package main

// import "fmt"

// func main() {
// 	var a int8 = 20
// 	var b int16 = 40
// 	var c = int16(a) + b
// 	fmt.Printf("值：%v，类型：%T\n", c, c)
// }

// package main

// import "fmt"

// func main() {
// 	var a float32 = 3.2
// 	var b int16 = 6
// 	c := a + float32(b)
// 	fmt.Printf("值：%v，类型：%T\n", c, c)
// }

// package main

// import "fmt"

// func main() {
// 	var a int16 = 129
// 	var b = int8(a)
// 	fmt.Printf("b = %v\n", b)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	a, b := 3, 4
// 	// math.Sqrt 接收的参数是 float64 类型，需要强制转换
// 	c := int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var i int = 20
// 	var f float64 = 3.14159
// 	var t bool = true
// 	var b byte = 'a'

// 	str := fmt.Sprintf("%d", i)
// 	fmt.Printf("值：%v，类型：%T\n", str, str)

// 	str = fmt.Sprintf("%f", f)
// 	fmt.Printf("值：%v，类型：%T\n", str, str)

// 	str = fmt.Sprintf("%t", t)
// 	fmt.Printf("值：%v，类型：%T\n", str, str)

// 	str = fmt.Sprintf("%c", b)
// 	fmt.Printf("值：%v，类型：%T\n", str, str)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// int 转换为 string
// 	var num1 int = 20
// 	s1 := strconv.Itoa(num1)
// 	fmt.Printf("值：%v，类型：%T\n", s1, s1)

// 	// 或者使用 FormatInt，Itoa 是对其的便携封装
// 	s1a := strconv.FormatInt(int64(num1), 10)
// 	fmt.Printf("值：%v，类型：%T\n", s1a, s1a)

// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var num float32 = 1.14514
// 	s := strconv.FormatFloat(float64(num), 'g', 100, 64)
// 	fmt.Printf("值：%v，类型：%T\n", s, s)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var num int64 = 20
// 	s := strconv.FormatInt(num, 10)
// 	fmt.Printf("值为 %v，类型 %T\n", s, s)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	str := "1234"
// 	// 第二个参数返回的是错误信息
// 	i64, _ := strconv.ParseInt(str, 10, 64)
// 	fmt.Printf("值为 %v，类型为 %T\n", i64, i64)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	str := "3.1415926"
// 	f32, _ := strconv.ParseFloat(str, 32)
// 	f64, _ := strconv.ParseFloat(str, 64)

// 	fmt.Printf("值为 %v，类型为 %T\n", f32, f32)
// 	fmt.Printf("值为 %v，类型为 %T\n", f64, f64)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	b, _ := strconv.ParseBool("true")
// 	fmt.Printf("值为 %v，类型为 %T\n", b, b)
// }

// package main

// import "fmt"

// func main() {
// 	str := "Hello 张三"
// 	for _, ch := range str {
// 		fmt.Printf("%v(%c) ", ch, ch)
// 	}
// 	fmt.Println()
// }

// package main

// func main() {
// 	// a, b := 3, 4
// 	// c := int(math.Sqrt(float64(a*a + b*b)))
// 	// fmt.Println(c)
// 	// var i int = 20
// 	// var f float64 = 3.14159
// 	// var t bool = true
// 	// var b byte = 'a'

// 	// strI := fmt.Sprintf("%d", i)
// 	// strF := fmt.Sprintf("%f", f)
// 	// strT := fmt.Sprintf("%t", t)
// 	// strB := fmt.Sprintf("%c", b)

// 	// fmt.Println(strI, strF, strT, strB)
// 	// var num int = 20
// 	// s := strconv.Itoa(num)
// 	// fmt.Println(s)
// 	// s1 := strconv.FormatInt(int64(num), 10)
// 	// fmt.Println(s1)
// 	// var num = 3.1415926
// 	// str := strconv.FormatFloat(num, 'g', 6, 64)
// 	// fmt.Println(str)

// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	str := strconv.FormatBool(true)
// 	fmt.Printf("值为 %v，类型为 %T\n", str, str)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	str := "1234"
// 	i64, _ := strconv.ParseInt(str, 10, 64)
// 	fmt.Printf("值为 %v，类型为 %T\n", i64, i64)
// 	str = "3.14159 "
// 	f64, _ := strconv.ParseFloat(str, 64)
// 	fmt.Printf("值为 %v，类型为 %T\n", f64, f64)
// 	str = "true"
// 	b, _ := strconv.ParseBool(str)
// 	fmt.Printf("值为 %v，类型为 %T\n", b, b)
// 	str = "Hello 哈基米"
// 	for _, ch := range str {
// 		fmt.Printf("%v(%c) ", ch, ch)
// 	}
// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	a := 'a'
// 	b := 'b'
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Printf("%c %c\n", a, b)
// }

// package main

// import "fmt"

// func main() {
// 	str := "哈基米Mambo"
// 	// for i := 0; i < len(str); i++ {
// 	// 	fmt.Printf("%v(%c) ", str[i], str[i])
// 	// }
// 	// fmt.Println()

// 	// for _, ch := range str {
// 	// 	fmt.Printf("%v(%c) ", ch, ch)
// 	// }
// 	// fmt.Println()
// 	runeStr := []rune(str)
// 	runeStr[0] = '叮'
// 	fmt.Println(string(runeStr))
// }

// package main

// import "fmt"

// func test() bool {
// 	fmt.Println("test...")
// 	return true
// }

// func main() {
// 	// a := 6
// 	// b := 3
// 	// fmt.Println(a + b)
// 	// fmt.Println(a - b)
// 	// fmt.Println(a * b)
// 	// fmt.Println(a / b)
// 	// 余数 = 被除数 - (被除数 / 除数) * 除数
// 	// var a = -10
// 	// var b = -3
// 	// fmt.Println(a % b)
// 	// a++
// 	// a := 1
// 	// b := 2
// 	// fmt.Println(a == b)
// 	// fmt.Println(a != b)
// 	// fmt.Println(a > b)
// 	// fmt.Println(a >= b)
// 	// fmt.Println(a <= b)
// 	// fmt.Println(a < b)
// 	// if a < b {
// 	// 	fmt.Printf("Hello World!\n")
// 	// }
// 	// a := 23
// 	// b := 8
// 	// fmt.Println(a > 10 && b > 10)
// 	// fmt.Println(a > 10 || b > 10)
// 	// fmt.Println(a < 10 || b > 10)
// 	// flag := true
// 	// fmt.Println(!flag)
// 	// var a = 10
// 	// if a < 9 && test() {
// 	// 	fmt.Println("哈基米")
// 	// }
// 	var a = 10
// 	a += 2
// 	fmt.Println(a)
// 	a *= 3
// 	fmt.Println(a)
// 	a %= 3
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("10 + 3 =", 10+3)
// 	fmt.Println("10 - 3 =", 10-3)
// 	fmt.Println("10 * 3 =", 10*3)

// 	fmt.Println("10 / 3 =", 10/3)
// 	fmt.Println("10.0 / 3.0 =", 10.0/3.0)

// 	fmt.Println("10 % 3 =", 10%3)
// 	fmt.Println("-10 % 3 =", -10%3)
// 	fmt.Println("10 % -3 =", 10%-3)
// 	fmt.Println("-10 % -3=", -10%-3)
// }

// package main

// import "fmt"

// func main() {
// 	a := 9
// 	b := 8
// 	fmt.Println(a == b) // false
// 	fmt.Println(a != b) // true
// 	fmt.Println(a > b)  // true
// 	fmt.Println(a >= b) // true
// 	fmt.Println(a < b)  // false
// 	fmt.Println(a <= b) // false
// 	flag := a > b
// 	fmt.Println("flag =", flag)
// }

// package main

// import "fmt"

// func main() {
// 	// 演示 &&
// 	var age int = 40
// 	if age > 30 && age < 50 {
// 		fmt.Println("哈基米")
// 	}
// 	if age > 30 && age < 40 {
// 		fmt.Println("曼波")
// 	}

// 	// 演示 ||
// 	if age > 30 || age < 50 {
// 		fmt.Println("叮咚鸡")
// 	}
// 	if age > 30 || age < 40 {
// 		fmt.Println("大狗叫")
// 	}

// 	// 演示 !
// 	if age > 30 {
// 		fmt.Println("胖宝宝")
// 	}
// 	if !(age > 30) {
// 		fmt.Println("小白手套")
// 	}
// }

// package main

// import "fmt"

// func test() bool {
// 	fmt.Println("test...")
// 	return true
// }

// func main() {
// 	var i int = 10
// 	if i < 9 && test() {
// 		fmt.Println("ok...")
// 	}

// 	if i > 9 || test() {
// 		fmt.Println("hello...")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	d := 8 + 2*8
// 	fmt.Println(d)

// 	x := 10
// 	x += 5
// 	fmt.Println("x + 5 =", x)

// 	x = 10
// 	x -= 5
// 	fmt.Println("x - 5 =", x)

// 	x = 10
// 	x *= 5
// 	fmt.Println("x * 5 =", x)

// 	x = 10
// 	x /= 5
// 	fmt.Println("x / 5 =", x)

// 	x = 10
// 	x %= 3
// 	fmt.Println("x % 3 =", x)
// }

// package main

// import "fmt"

// func main() {
// 	a := 9
// 	b := 2

// 	t := a
// 	a = b
// 	b = t
// 	fmt.Printf("交换后的情况是 a = %v，b = %v\n", a, b)
// }

// package main

// import "fmt"

// func main() {
// 	days := 100
// 	week := days / 7
// 	day := days % 7
// 	fmt.Printf("%d 个星期零 %d 天\n", week, day)
// }

// package main

// import "fmt"

// func main() {
// 	a := 5 // 101
// 	b := 2 // 010

// 	fmt.Println("a & b =", a&b)   // 000
// 	fmt.Println("a | b =", a|b)   // 111
// 	fmt.Println("a ^ b =", a^b)   // 111
// 	fmt.Println("5 >> 2 =", a>>b) // 1
// 	fmt.Println("5 << 2 =", a<<b) // 10100
// 	fmt.Println("5 << 1 =", 5<<1) // 1010
// 	fmt.Println("5 >> 1 =", 5>>1) // 10
// 	fmt.Println("7 >> 2 =", 7>>2) // 1
// }

// package main

// import "fmt"

// const (
// 	A = 100
// 	B = iota
// 	C
// 	D = iota
// )

// func main() {
// 	fmt.Println(A, B, C, D)
// }

package main

func main() {

}
