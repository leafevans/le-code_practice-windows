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

package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("继续执行")
	// for i := range 2 {
	// 	for j := range 10 {
	// 		if j == 3 {
	// 			break
	// 		}
	// 		fmt.Printf("i = %v, j = %v\n", i, j)
	// 	}
	// }
	// 一次跳出多层循环
	// label1:
	// 	for i := range 2 {
	// 		for j := range 10 {
	// 			if j == 3 {
	// 				break label1
	// 			}
	// 			fmt.Printf("i = %v, j = %v\n", i, j)
	// 		}
	// 	}
	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
