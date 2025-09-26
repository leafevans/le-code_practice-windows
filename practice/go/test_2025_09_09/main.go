// package main

// import "fmt"

// func main() {
// 	// var a = 34
// 	// var b = 10
// 	// a += b
// 	// b = a - b
// 	// a -= b
// 	// fmt.Printf("a = %v, b = %v\n", a, b)
// 	// days := 100
// 	// day := days / 7
// 	// week := days % 7
// 	// fmt.Println(day, week)
// 	// huashi := 100.0
// 	// sheshi := (huashi - 32) / 1.8
// 	// fmt.Printf("%.2f\n", sheshi)
// 	// var a = 0b101
// 	// var b = 0b010

// 	// fmt.Println("a & b =", a&b)
// 	// fmt.Println("a | b =", a|b)
// 	// fmt.Println("a ^ b =", a^b)
// 	// fmt.Println("a << b =", a<<b)
// 	// fmt.Println("a >> b =", a>>b)

// 	// a := 10
// 	// b := 15
// 	// if a < b {
// 	// 	fmt.Println("哈基米")
// 	// } else if a == b {
// 	// 	fmt.Println("曼波")
// 	// } else {
// 	// 	fmt.Println("叮咚鸡")
// 	// }

// 	// if age := 34; age > 20 {
// 	// 	fmt.Println("成年人")
// 	// }

// 	// score := 65

// 	// if score > 90 {
// 	// 	fmt.Println("A")
// 	// } else if score > 75 {
// 	// 	fmt.Println("B")
// 	// } else {
// 	// 	fmt.Println("C")
// 	// }

// 	// if score := 75; score >= 90 {
// 	// 	fmt.Println("A")
// 	// } else if score > 75 {
// 	// 	fmt.Println("B")
// 	// } else {
// 	// 	fmt.Println("C")
// 	// }

// 	// age := 30
// 	// if age > 20 {
// 	// 	fmt.Println("成年人")
// 	// } else {
// 	// 	fmt.Println("未成年")
// 	// }

// 	// a := 34
// 	// b := 24
// 	// fmt.Println(math.Max(float64(a), float64(b)))
// 	// i := 1
// 	// for ; i <= 10; i++ {
// 	// 	fmt.Printf("%v ", i)
// 	// }
// 	// fmt.Println()

// 	// i := 1
// 	// for i <= 10 {
// 	// 	fmt.Printf("%v ", i)
// 	// 	i++
// 	// }
// 	// fmt.Println()

// 	i := 1
// 	for {
// 		fmt.Printf("%v ", i)
// 		if i == 10 {
// 			break
// 		}
// 		i++
// 	}
// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	score := 65
// 	if score >= 90 {
// 		fmt.Println("A")
// 	} else if score > 75 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("C")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	if score := 56; score >= 90 {
// 		fmt.Println("a")
// 	} else if score > 75 {
// 		fmt.Println("b")
// 	} else {
// 		fmt.Println("c")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Printf("%v ", i)
// 		i++
// 	}
// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	k := 1
// 	for {
// 		if k <= 10 {
// 			fmt.Println("ok~~~", k)
// 		} else {
// 			break
// 		}
// 		k++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 50; i++ {
// 		if i%2 == 0 {
// 			fmt.Printf("%v ", i)
// 		}
// 	}
// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// i := 5
// 	// for {
// 	// 	fmt.Println(i)
// 	// 	if i == 10 {
// 	// 		break
// 	// 	}
// 	// }
// 	// k := 1
// 	// for {
// 	// 	if k <= 10 {
// 	// 		fmt.Println("ok~~~", k)
// 	// 	} else {
// 	// 		break
// 	// 	}
// 	// 	k++
// 	// }
// 	// for i := 0; i <= 50; i++ {
// 	// 	if i%2 == 0 {
// 	// 		fmt.Println(i)
// 	// 	}
// 	// }
// 	sum := 0
// 	count := 0
// 	for i := 1; i < 100; i++ {
// 		if i%9 == 0 {
// 			sum += i
// 			count++
// 		}
// 	}
// 	fmt.Println("count =", count)
// 	fmt.Println("sum =", sum)
// }

// package main

// import "fmt"

// func main() {
// 	line := 5
// 	for i := 1; i <= line; i++ {
// 		for j := 0; j < line-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		for j := 0; j < 2*i-1; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// n := 15
// 	// for i := 1; i <= n; i++ {
// 	// 	for j := 0; j < n-i; j++ {
// 	// 		fmt.Print(" ")
// 	// 	}
// 	// 	for j := 0; j < 2*i-1; j++ {
// 	// 		fmt.Print("*")
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	// for i := 1; i <= 9; i++ {
// 	// 	for j := 1; j <= i; j++ {
// 	// 		fmt.Printf("%v x %v = %v\t", i, j, i*j)
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	// n := 5
// 	// for i := 1; i <= n; i++ {
// 	// 	for j := 0; j < n-i; j++ {
// 	// 		fmt.Print(" ")
// 	// 	}
// 	// 	for j := 0; j < 2*i-1; j++ {
// 	// 		fmt.Print("*")
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	// str := "哈基米"
// 	// for _, ch := range str {
// 	// 	fmt.Printf("%c ", ch)
// 	// }
// 	// fmt.Println()
// 	// arr := []string{"php", "java", "node", "golang"}
// 	// for _, str := range arr {
// 	// 	fmt.Println(str)
// 	// }

// 	// switch ext := ".html"; ext {
// 	// case ".html":
// 	// 	fmt.Println("text/html")
// 	// case ".css":
// 	// 	fmt.Println("text/css")
// 	// case ".js":
// 	// 	fmt.Println("text/javascript")
// 	// default:
// 	// 	fmt.Println("找不到此后缀")
// 	// }

// 	// var n = 5
// 	// switch n {
// 	// case 2, 4, 6, 8, 10:
// 	// 	fmt.Println("偶数")
// 	// case 1, 3, 5, 7, 9:
// 	// 	fmt.Println("奇数")
// 	// }
// 	// switch score := "A"; score {
// 	// case "A", "B", "C":
// 	// 	fmt.Println("及格")
// 	// case "D":
// 	// 	fmt.Println("不及格")
// 	// }
// 	switch age := 18; {
// 	case age < 24:
// 		fmt.Println("好好学习")
// 		fallthrough
// 	case age < 60:
// 		fmt.Println("好好赚钱")
// 		fallthrough
// 	case age >= 60:
// 		fmt.Println("注意身体")
// 		fallthrough
// 	default:
// 		fmt.Println("输入错误")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	str := "Mambo 哈基米"
// 	for index, val := range str {
// 		fmt.Printf("index = %d, val = %c\n", index, val)
// 	}
// 	for _, val := range str {
// 		fmt.Printf("val = %c\n", val)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	extname := ".html"
// 	switch extname {
// 	case ".html":
// 		fmt.Println("text/html")
// 	case ".css":
// 		fmt.Println("text/css")
// 	case ".js":
// 		fmt.Println("text/javascript")
// 	default:
// 		fmt.Println("格式错误")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	n := 2
// 	switch n {
// 	case 1, 3, 5, 7, 9:
// 		fmt.Println("奇数")
// 	case 2, 4, 6, 8:
// 		fmt.Println("偶数")
// 	default:
// 		fmt.Println(n)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	switch n := 7; n {
// 	case 1, 3, 5, 7, 9:
// 		fmt.Println("奇数")
// 	case 2, 4, 6, 8:
// 		fmt.Println("偶数")
// 	default:
// 		fmt.Println(n)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	age := 56
// 	switch {
// 	case age < 25:
// 		fmt.Println("好好学习！")
// 	case age <= 60:
// 		fmt.Println("好好工作！")
// 	case age > 60:
// 		fmt.Println("好好享受！")
// 	default:
// 		fmt.Println("好好活着！")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// for i := 1; i <= 10; i++ {
// 	// 	if i == 2 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(i)
// 	// }
// 	// fmt.Println("继续执行")
// 	// for i := range 2 {
// 	// 	for j := range 10 {
// 	// 		if j == 3 {
// 	// 			break
// 	// 		}
// 	// 		fmt.Printf("i = %v, j = %v\n", i, j)
// 	// 	}
// 	// }
// 	// 一次跳出多层循环
// 	// label1:
// 	// 	for i := range 2 {
// 	// 		for j := range 10 {
// 	// 			if j == 3 {
// 	// 				break label1
// 	// 			}
// 	// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 	// 		}
// 	// 	}
// 	// for i := range 10 {
// 	// 	if i%2 == 0 {
// 	// 		continue
// 	// 	}
// 	// 	fmt.Println(i)
// 	// }
// 	// 	for i := range 2 {
// 	// 		for j := range 10 {
// 	// 			if j == 3 {
// 	// 				continue
// 	// 			}
// 	// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 	// 		}
// 	// 	}
// 	// 	fmt.Println()

// 	// label2:
// 	// 	for i := range 2 {
// 	// 		for j := range 10 {
// 	// 			if j == 3 {
// 	// 				continue label2
// 	// 			}
// 	// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 	// 		}
// 	// 	}
// 	var n = 30
// 	if n > 24 {
// 		fmt.Println("成年人")
// 		goto label
// 	}
// 	fmt.Println("aaa")
// 	fmt.Println("bbb")
// label:
// 	fmt.Println("ccc")
// 	fmt.Println("ddd")
// }

// package main

// import "fmt"

// func main() {
// 	for i := range 2 {
// 		for j := range 10 {
// 			if j == 3 {
// 				break
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	k := 1
// 	for {
// 		if k <= 10 {
// 			fmt.Println("ok~~~", k)
// 		} else {
// 			break
// 		}
// 		k++
// 	}
// }

// package main

// import "fmt"

// func main() {
// label:
// 	for i := range 2 {
// 		for j := range 10 {
// 			if j == 3 {
// 				break label
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// label:
// 	for i := range 2 {
// 		for j := range 4 {
// 			if j == 2 {
// 				continue label
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	n := 30
// 	fmt.Println("ok1")
// 	if n > 20 {
// 		goto label
// 	}
// 	fmt.Println("ok2")
// 	fmt.Println("ok2")
// 	fmt.Println("ok4")
// label:
// 	fmt.Println("ok5")
// 	fmt.Println("ok6")
// 	fmt.Println("ok7")
// }

// package main

// import "fmt"

// func main() {
// 	for i := range 10 {
// 		for j := range 10 {
// 			if j == 2 {
// 				goto breakTag
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// 	return
// breakTag:
// 	fmt.Println("结束循环")
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unsafe"
// )

// func main() {
// 	str := `哈基米喔南北绿豆
// 波压库奶龙`
// 	fmt.Println(str)
// 	fmt.Println(len(str))
// 	str += "\n曼波欧马吉利曼波"
// 	fmt.Println(str)
// 	str1 := "123-456-789"
// 	arr := strings.Split(str1, "-")
// 	fmt.Printf("%v %T\n", arr, arr)
// 	fmt.Println(strings.Contains(str, "哈基米"))
// 	fmt.Println(strings.HasPrefix(str, "哈基米"))
// 	fmt.Println(strings.HasSuffix(str, "曼波"))
// 	fmt.Println(strings.Index(str, "曼波"))
// 	fmt.Println(strings.LastIndex(str, "曼波"))
// 	str2 := strings.Join(arr, "*")
// 	fmt.Printf("%v %T\n", str2, str2)
// 	fmt.Println()
// 	for _, ch := range str {
// 		fmt.Printf("%v(%c) ", ch, ch)
// 	}
// 	fmt.Println()
// 	ch := '张'
// 	fmt.Println(unsafe.Sizeof(ch))
// 	runeStr := []rune(str)
// 	runeStr[0] = '叮'
// 	runeStr[1] = '咚'
// 	runeStr[2] = '鸡'
// 	fmt.Println(string(runeStr))
// }

// package main

// import "fmt"

// func main() {
// 	// var a int8 = 20
// 	// var b int16 = 40
// 	// var c = int16(a) + b
// 	// fmt.Println(c)
// 	// a, b := 3, 4
// 	// c := int(math.Sqrt(float64(a*a + b*b)))
// 	// fmt.Println(c)
// 	// i := 20
// 	// f := 3.1415926
// 	// t := true
// 	// b = 'a'

// 	// str := fmt.Sprintf("%d", i)
// 	// fmt.Printf("%v %T\n", str, str)

// 	// str = fmt.Sprintf("%f", f)
// 	// fmt.Printf("%v %T\n", str, str)

// 	// str = fmt.Sprintf("%t", t)
// 	// fmt.Printf("%v %T\n", str, str)

// 	// str = fmt.Sprintf("%c", b)
// 	// fmt.Printf("%v %T\n", str, str)
// 	// num := 20
// 	// str := strconv.Itoa(num)
// 	// fmt.Println(str)
// 	// num := 20
// 	// str := strconv.FormatInt(int64(num), 10)
// 	// fmt.Println(str)
// 	// num2 := 1.14514
// 	// str2 := strconv.FormatFloat(num2, 'g', 4, 64)
// 	// fmt.Println(str2)
// 	// str3 := strconv.FormatBool(true)
// 	// fmt.Println(str3)
// 	// str := "1234"
// 	// i64, _ := strconv.ParseInt(str, 10, 64)
// 	// fmt.Printf("%v, %T\n", i64, i64)
// 	// str := "1.23e4"
// 	// f, _ := strconv.ParseFloat(str, 64)
// 	// fmt.Println(f)
// 	// a := 0b101
// 	// b := 0b010

// 	// fmt.Println("a & b =", a&b)
// 	// fmt.Println("a | b =", a|b)
// 	// fmt.Println("a ^ b =", a^b)
// 	// fmt.Println("a << b =", a<<b)
// 	// fmt.Println("a >> b =", a>>b)
// 	if score := 56; score >= 90 {
// 		fmt.Println("A")
// 	} else if score > 75 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("C")
// 	}

// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("%v ", i)
// 	}
// 	fmt.Println()

// 	hachi := 1
// 	for ; hachi <= 10; hachi++ {
// 		fmt.Printf("%v ", hachi)
// 	}
// 	fmt.Println()

// 	mambo := 1
// 	for mambo <= 10 {
// 		fmt.Printf("%v ", mambo)
// 		mambo++
// 	}
// 	fmt.Println()

// 	ccb := 1
// 	for {
// 		if ccb <= 10 {
// 			fmt.Println("曼波~~~", ccb)
// 		} else {
// 			break
// 		}
// 		ccb++
// 	}
// 	fmt.Println()

// 	for i := range 50 {
// 		if i%2 == 0 {
// 			fmt.Printf("%v ", i)
// 		}
// 	}
// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	// count := 0
// 	// sum := 0
// 	// for i := range 100 {
// 	// 	if i%9 == 0 {
// 	// 		sum += i
// 	// 		count++
// 	// 	}
// 	// }
// 	// fmt.Println(sum, count)
// 	// sum := 0
// 	// for i := 1; i <= 100; i++ {
// 	// 	sum += i
// 	// }
// 	// fmt.Println(sum)
// 	// for range 4 {
// 	// 	for range 4 {
// 	// 		fmt.Print("*")
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	// for range 4 {
// 	// 	for range 4 {
// 	// 		fmt.Print("*")
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	for i := 1; i <= 5; i++ {
// 		for j := 1; j <= 5; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	n := 5
// 	for i := 1; i <= n; i++ {
// 		for j := 0; j < n-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		for j := 0; j < 2*i-1; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	str := "Mambo 哈基米"
// 	for i, v := range str {
// 		fmt.Printf("i = %d, v = %c\n", i, v)
// 	}
// 	for _, v := range str {
// 		fmt.Printf("v = %c\n", v)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	str := "Mambo 哈基米"
// 	for _, v := range str {
// 		fmt.Printf("v = %c\n", v)
// 	}
// 	fmt.Println()
// }

// package main

// import "fmt"

// func main() {
// 	extname := ".html"
// 	switch extname {
// 	case ".html":
// 		fmt.Println("text/html")
// 	case ".css":
// 		fmt.Println("text/css")
// 	case ".js":
// 		fmt.Println("text/javascript")
// 	default:
// 		fmt.Println("格式错误")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	n := 2
// 	switch n {
// 	case 1, 3, 5, 7, 9:
// 		fmt.Println("奇数")
// 	case 2, 4, 6, 8:
// 		fmt.Println("偶数")
// 	default:
// 		fmt.Println(n)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// switch n := 7; n {
// 	// case 1, 3, 5, 7, 9:
// 	// 	fmt.Println("奇数")
// 	// case 2, 4, 6, 8:
// 	// 	fmt.Println("偶数")
// 	// default:
// 	// 	fmt.Println(n)
// 	// }
// 	age := 56
// 	switch {
// 	case age < 25:
// 		fmt.Println("好好学习！")
// 	case age <= 60:
// 		fmt.Println("好好工作！")
// 	case age > 60:
// 		fmt.Println("好好享受！")
// 	default:
// 		fmt.Println("好好活着！")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	str := "a"
// 	switch str {
// 	case "a":
// 		fmt.Println("a")
// 		fallthrough
// 	case "b":
// 		fmt.Println("b")
// 		fallthrough
// 	case "c":
// 		fmt.Println("c")
// 		fallthrough
// 	default:
// 		fmt.Println("...")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	for i := range 2 {
// 		for j := range 10 {
// 			if j == 3 {
// 				break
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	k := 1
// 	for {
// 		if k <= 10 {
// 			fmt.Println("ok~~~", k)
// 		} else {
// 			break
// 		}
// 		k++
// 	}
// }

// package main

// import "fmt"

// func main() {
// label:
// 	for i := range 2 {
// 		for j := range 10 {
// 			if j == 3 {
// 				break label
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// label:
// 	for i := range 2 {
// 		for j := range 10 {
// 			if j == 3 {
// 				break label
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// }

// package main

// func main() {
// 	// label:
// 	// 	for {
// 	// 		fmt.Println(100)
// 	// 		for i := range 10 {
// 	// 			fmt.Printf("%v ", i)
// 	// 			if i == 5 {
// 	// 				break label
// 	// 			}
// 	// 		}
// 	// 	}
// 	// 	fmt.Println()
// 	// 	for range 3 {
// 	// 		for j := range 5 {
// 	// 			if j == 3 {
// 	// 				goto label
// 	// 			}
// 	// 			fmt.Printf("%v ", j)
// 	// 		}
// 	// 		fmt.Println()
// 	// 	}
// 	// label:
// 	// 	fmt.Println("哈基米")
// 	// n := 5
// 	// switch {
// 	// case n >= 5, n <= 4:
// 	// 	fmt.Println("哈基米")
// 	// 	fallthrough
// 	// default:
// 	// 	fmt.Println("曼波")
// 	// }
// 	// str1 := "哈基米"
// 	// str2 := "曼波"
// 	// str := fmt.Sprintf("%v %v", str1, str2)
// 	// fmt.Println(str)
// }

package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
