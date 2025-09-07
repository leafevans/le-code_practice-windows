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

package main
