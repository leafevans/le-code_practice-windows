// package main

// import "fmt"

// func main() {
// 	scoreMap := make(map[string]int, 8)
// 	scoreMap["哈基米"] = 90
// 	scoreMap["曼波"] = 100
// 	fmt.Println(scoreMap)
// 	fmt.Println(scoreMap["曼波"])
// 	fmt.Printf("Type of scoreMap: %T\n", scoreMap)
// 	value, ok := scoreMap["哈基米"]
// 	if ok {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("...")
// 	}
// 	delete(scoreMap, "哈基米")
// 	value, ok = scoreMap["哈基米"]
// 	if ok {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("...")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	mapSlice := make([]map[string]string, 3)
// 	for k, v := range mapSlice {
// 		fmt.Printf("k = %v, v = %v\n", k, v)
// 	}
// 	fmt.Println("After: ")
// 	mapSlice[0] = make(map[string]string, 10)
// 	mapSlice[0]["name"] = "哈基米"
// 	mapSlice[0]["password"] = "123456"
// 	mapSlice[0]["address"] = "双流区"
// 	for i, v := range mapSlice {
// 		fmt.Printf("i = %v, v = %v\n", i, v)
// 	}
// 	fmt.Println(mapSlice[1] == nil)
// }

// package main

// import "fmt"

// func main() {
// 	sliceMap := make(map[string][]string)
// 	fmt.Println(sliceMap)
// 	fmt.Println("After:")
// 	key := "中国"
// 	value, ok := sliceMap[key]
// 	if !ok {
// 		value = make([]string, 0, 2)
// 	}
// 	value = append(value, "北京", "上海")
// 	sliceMap[key] = value
// 	fmt.Println(sliceMap)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	wordMap := make(map[string]int)
// 	str := "how   do you do"
// 	for word := range strings.FieldsSeq(str) {
// 		wordMap[word]++
// 	}
// 	fmt.Println(wordMap)
// }

// package main

// import "fmt"

// func heapify(arr []int, n, i int) {
// 	largest := i
// 	left := 2*i + 1
// 	right := 2*i + 2

// 	if left < n && arr[largest] < arr[left] {
// 		largest = left
// 	}

// 	if right < n && arr[largest] < arr[right] {
// 		largest = right
// 	}

// 	if largest != i {
// 		arr[i], arr[largest] = arr[largest], arr[i]
// 		heapify(arr, n, largest)
// 	}
// }

// func heapSort(arr []int) {
// 	n := len(arr)

// 	for i := n/2 - 1; i >= 0; i-- {
// 		heapify(arr, n, i)
// 	}

// 	for i := n - 1; i > 0; i-- {
// 		arr[0], arr[i] = arr[i], arr[0]
// 		heapify(arr, i, 0)
// 	}
// }

// func main() {
// 	arr := []int{23, 4, 11, 2323, 3, 444, 9, 57}
// 	heapSort(arr)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func bruteForceSearch(text, pattern string) int {
// 	m := len(text)
// 	n := len(pattern)

// 	for i := 0; i <= m-n; i++ {
// 		j := 0
// 		for j < n && text[i+j] == pattern[j] {
// 			j++
// 		}
// 		if j == n {
// 			return i
// 		}
// 	}

// 	return -1
// }

// func buildNext(pattern string) []int {
// 	next := make([]int, len(pattern))
// 	j := 0

// 	for i := 1; i < len(pattern); i++ {
// 		for j > 0 && pattern[i] != pattern[j] {
// 			j = next[j-1]
// 		}
// 		if pattern[i] == pattern[j] {
// 			j++
// 		}
// 		next[i] = j
// 	}

// 	return next
// }

// func kmpSearch(text, pattern string) int {
// 	if len(pattern) == 0 {
// 		return 0
// 	}

// 	next := buildNext(pattern)
// 	j := 0

// 	for i := range len(text) {
// 		for j > 0 && text[i] != pattern[j] {
// 			j = next[j-1]
// 		}
// 		if text[i] == pattern[j] {
// 			j++
// 		}
// 		if j == len(pattern) {
// 			return i - j + 1
// 		}
// 	}

// 	return -1
// }

// func main() {
// 	text := "ababcabcacbab"
// 	pattern := "abcac"
// 	index := bruteForceSearch(text, pattern)
// 	fmt.Printf("模式串首次出现的位置: %d\n", index)
// }

// package main

// import "fmt"

// func bruteForceSearch(text, pattern string) int {
// 	m := len(text)
// 	n := len(pattern)

// 	for i := 0; i <= m-n; i++ {
// 		j := 0
// 		for j < n && text[i+j] == pattern[j] {
// 			j++
// 		}
// 		if j == n {
// 			return i
// 		}
// 	}

// 	return -1
// }

// func buildNext(pattern string) []int {
// 	next := make([]int, len(pattern))
// 	j := 0

// 	for i := 1; i < len(pattern); i++ {
// 		for j > 0 && pattern[i] != pattern[j] {
// 			j = next[j-1]
// 		}
// 		if pattern[i] == pattern[j] {
// 			j++
// 		}
// 		next[i] = j
// 	}

// 	return next
// }

// func kmpSearch(text, pattern string) int {
// 	if len(pattern) == 0 {
// 		return 0
// 	}

// 	next := buildNext(pattern)
// 	j := 0

// 	for i := range len(text) {
// 		for j > 0 && text[i] != pattern[j] {
// 			j = next[j-1]
// 		}
// 		if text[i] == pattern[j] {
// 			j++
// 		}
// 		if j == len(pattern) {
// 			return i - j + 1
// 		}
// 	}

// 	return -1
// }

// func buildNext(pattern string) []int {
// 	next := make([]int, len(pattern))
// 	j := 0

// 	for i := 1; i < len(pattern); i++ {
// 		for j > 0 && pattern[i] != pattern[j] {
// 			j = next[j-1]
// 		}
// 		if pattern[i] == pattern[j] {
// 			j++
// 		}
// 		next[i] = j
// 	}

// 	return next
// }

// func kmpSearch(text, pattern string) int {
// 	if len(pattern) == 0 {
// 		return 0
// 	}

// 	next := buildNext(pattern)
// 	j := 0

// 	for i := range len(text) {
// 		for j > 0 && text[i] != pattern[j] {
// 			j = next[j-1]
// 		}
// 		if text[i] == pattern[j] {
// 			j++
// 		}
// 		if j == len(pattern) {
// 			return i - j + 1
// 		}
// 	}

// 	return -1
// }

// package main

// import "fmt"

// func buildNext(pattern string) []int {
// 	next := make([]int, len(pattern))
// 	j := 0

// 	for i := 1; i < len(pattern); i++ {
// 		for j > 0 && pattern[i] != pattern[j] {
// 			j = next[j-1]
// 		}

// 		if pattern[i] == pattern[j] {
// 			j++
// 		}

// 		next[i] = j
// 	}

// 	return next
// }

// func kmpSearch(text, pattern string) int {
// 	if len(pattern) == 0 {
// 		return 0
// 	}

// 	next := buildNext(pattern)
// 	j := 0

// 	for i := range len(text) {
// 		for j > 0 && text[i] != pattern[j] {
// 			j = next[j-1]
// 		}

// 		if text[i] == pattern[j] {
// 			j++
// 		}

// 		if j == len(pattern) {
// 			return i - j + 1
// 		}
// 	}

// 	return -1
// }

// func main() {
// 	text := "ABABCABCACBAB"
// 	pattern := "ABCAC"

// 	fmt.Println(kmpSearch(text, pattern))
// }

// package main

// import "fmt"

// // func add[T int | float32 | float64](x, y T) T {
// // 	return x + y
// // }

// // func sub[T int | float32 | float64](x, y T) T {
// // 	return x - y
// // }

// func add[T int | float32 | float64](nums ...T) T {
// 	var sum T
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }

// func main() {
// 	// x := 2
// 	// y := 3
// 	// fmt.Println(add(x, y))
// 	// m := 2.4
// 	// n := 6.8
// 	// fmt.Println(add(m, n))

// 	// fmt.Println(sub(x, y))
// 	// fmt.Println(sub(m, n))
// 	fmt.Println(add(10, 20, 30, 40))
// }

// package main

// import "fmt"

// func add(x int, y ...int) int {
// 	sum := x
// 	for _, v := range y {
// 		sum += v
// 	}
// 	return sum
// }

// func main() {
// 	fmt.Println(add(10, 20, 30, 40))
// }

// package main

// import "fmt"

// func calc[T int | float32 | float64](x, y T) (sum, diff T) {
// 	sum = x + y
// 	diff = x - y
// 	return
// }

// func main() {
// 	sum, diff := calc(12, 34)
// 	fmt.Println(sum)
// 	fmt.Println(diff)
// }

// package main

// import "fmt"

// func sayHello() {
// 	fmt.Println("Hello World!")
// }

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	sayHello()
// 	fmt.Println(add(10, 20))
// }

// package main

// import "fmt"

// func add(x int, y ...int) int {
// 	sum := x
// 	for _, v := range y {
// 		sum += v
// 	}
// 	return sum
// }

// func main() {
// 	fmt.Println(add(10))
// 	fmt.Println(add(10, 20))
// 	fmt.Println(add(10, 20, 30))
// 	fmt.Println(add(10, 20, 30, 40))
// }

// package main

// func quickSort(slice []int) {
// 	if len(slice) <= 1 {
// 		return
// 	}

// 	pivot := slice[len(slice)/2]
// 	left, right := 0, len(slice)-1

// 	for left <= right {
// 		for slice[left] < pivot {
// 			left++
// 		}
// 		for slice[right] > pivot {
// 			right--
// 		}
// 		if left <= right {
// 			slice[left], slice[right] = slice[right], slice[left]
// 			left++
// 			right--
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func mapSort(userinfo map[string]string) string {
// 	keys := make([]string, 0, len(userinfo))

// 	for key := range userinfo {
// 		keys = append(keys, key)
// 	}

// 	sort.Strings(keys)

// 	str := ""

// 	for _, key := range keys {
// 		str += fmt.Sprintf("%v=>%v ", key, userinfo[key])
// 	}

// 	return str
// }

// func main() {
// 	userinfo := map[string]string{
// 		"username": "王婷婷",
// 		"age":      "20",
// 		"sex":      "女",
// 		"height":   "165",
// 	}

// 	str := mapSort(userinfo)

// 	fmt.Println(str)
// }

// package main

// import "fmt"

// var a = "全局变量"

// func run() {
// 	a = "阿迪"
// 	fmt.Println(a)
// 	fmt.Println("run")
// }

// func main() {
// 	fmt.Println(a)
// 	run()
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// // 定义全局变量 num。
// var num = 10

// func testGlobal() {
// 	// 在函数中访问全局变量 num。
// 	fmt.Printf("num = %v\n", num)
// }

// func main() {
// 	// num = 10
// 	testGlobal()
// }

// package main

// import "fmt"

// // 定义全局变量 num。
// var num = 10

// func testNum() {
// 	num := 100
// 	// 函数中优先使用局部变量 num。
// 	fmt.Printf("num = %d\n", num)
// }

// func main() {
// 	testNum()                     // num = 100
// 	fmt.Printf("num = %d\n", num) // num = 10
// }

// package main

// import "fmt"

// type cala func(int, int) int

// type integer int

// func add(x, y int) int {
// 	return x + y
// }

// func sub(x, y int) int {
// 	return x - y
// }

// func test() {
// 	fmt.Println("test...")
// }

// func main() {
// 	var c cala = add
// 	fmt.Printf("c 的类型：%T\n", c)
// 	c = sub
// 	fmt.Printf("c 的类型：%T\n", c)
// 	fmt.Printf("test 的类型：%T\n", test)
// 	fmt.Println(c(10, 5))
// 	f := sub
// 	fmt.Printf("f 的类型：%T\n", f)
// 	fmt.Println(f(15, 5))
// 	var a int = 10
// 	var b integer = 20
// 	fmt.Printf("a 的类型：%T\n", a)
// 	fmt.Printf("b 的类型：%T\n", b)
// 	fmt.Println(a + int(b))
// }

// package main

// import (
// 	"fmt"
// )

// func buildNext(pattern string) []int {
// 	next := make([]int, len(pattern))
// 	j := 0

// 	for i := 1; i < len(pattern); i++ {
// 		for j > 0 && pattern[i] != pattern[j] {
// 			j = next[j-1]
// 		}

// 		if pattern[i] == pattern[j] {
// 			j++
// 		}

// 		next[i] = j
// 	}

// 	return next
// }

// func kmpSearch(text, pattern string) int {
// 	if len(pattern) == 0 {
// 		return 0
// 	}

// 	next := buildNext(pattern)
// 	j := 0

// 	for i := range len(text) {
// 		for j > 0 && pattern[i] != pattern[j] {
// 			j = next[j-1]
// 		}

// 		if pattern[i] == pattern[j] {
// 			j++
// 		}

// 		if j == len(pattern) {
// 			return i - j + 1
// 		}
// 	}

// 	return -1
// }

// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))
// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 		} else {
// 			result = append(result, right[j])
// 		}
// 	}

// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)

// 	return result
// }

// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2
// 	left := mergeSort(arr[:mid])
// 	right := mergeSort(arr[mid:])

// 	return merge(left, right)
// }

// func heapify(arr []int, n, i int) {
// 	largest := i
// 	left, right := 2*i+1, 2*i+2

// 	for left < n && arr[largest] < arr[left] {
// 		largest = left
// 	}

// 	for right < n && arr[largest] < arr[right] {
// 		largest = right
// 	}

// 	if largest != i {
// 		arr[i], arr[largest] = arr[largest], arr[i]
// 		heapify(arr, n, largest)
// 	}
// }

// func heapSort(arr []int) {
// 	n := len(arr)

// 	for i := n/2 - 1; i >= 0; i-- {
// 		heapify(arr, n, i)
// 	}

// 	for i := n - 1; i > 0; i-- {
// 		arr[0], arr[i] = arr[i], arr[0]
// 		heapify(arr, i, 0)
// 	}

// }

// func add(x, y int) int {
// 	return x + y
// }

// func sub(x, y int) int {
// 	return x - y
// }

// func calc(x, y int, fn func(int, int) int) int {
// 	return fn(x, y)
// }

// func do(option string) func(x, y int) int {
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
// 	// fmt.Println(calc(10, 20, add))
// 	// fmt.Println(calc(10, 20, sub))
// 	// fmt.Println(calc(10, 20, func(x, y int) int {
// 	// 	return x * y
// 	// }))
// 	fmt.Println(do("+")(10, 20))
// 	fmt.Println(do("-")(10, 20))
// 	fmt.Println(do("*")(10, 20))

// 	func() {
// 		fmt.Println("test...")
// 	}()
// }

// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// type calc func(int, int) int

// func main() {
// 	// 声明一个 calc 类型的变量 c，
// 	// 并将 add 变量赋值给它。
// 	var c calc = add
// 	fmt.Printf("Type of c: %T\n", c)
// 	// 像调用 add 一样调用 c。
// 	fmt.Println(c(1, 2))

// 	f := add // 将函数 add 赋值给变量 f。
// 	fmt.Printf("Type of f: %T\n", f)
// 	// 像调用 add 一样调用 f。
// 	fmt.Println(f(10, 20))

// 	num := 0o12
// 	fmt.Println(num)
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

// func main() {
// 	add := func(x, y int) {
// 		fmt.Println(x + y)
// 	}
// 	add(10, 20) // 通过变量调用匿名函数。

// 	// 自执行函数指匿名函数定义完加 () 直接执行。
// 	func(x, y int) {
// 		fmt.Println(x + y)
// 	}(10, 20)
// }

// package main

// import "fmt"

// func fn1() {
// 	fmt.Println("方法 1")
// }

// func fn2() {
// 	fn1()
// 	fmt.Println("方法 2")
// }

// func main() {
// 	fn2()
// }

// package main

// import "fmt"

// // 函数中嵌套一个函数，最后返回其中的函数。
// func adder() func() int {
// 	var i = 10

// 	return func() int {
// 		return i + 1
// 	}
// }

// func adder2() func(y int) int {
// 	// 常驻内存，不污染全局。
// 	var i = 10
// 	return func(y int) int {
// 		i += y
// 		return i
// 	}
// }

// func main() {
// 	var fn = adder()
// 	fmt.Println(fn())
// 	fmt.Println(fn())
// 	fmt.Println(fn())

// 	var fn2 = adder2()
// 	fmt.Println(fn2(10))
// 	fmt.Println(fn2(10))
// 	fmt.Println(fn2(10))
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
// 	var fn = adder()
// 	fmt.Println(fn(10)) // 10
// 	fmt.Println(fn(20)) // 20
// 	fmt.Println(fn(30)) // 60

// 	fn2 := adder()
// 	fmt.Println(fn2(40)) // 40
// 	fmt.Println(fn2(50)) // 90
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
// 	fmt.Println(fn(10)) // 20
// 	fmt.Println(fn(20)) // 40
// 	fmt.Println(fn(30)) // 70

// 	fn2 := adder(20)
// 	fmt.Println(fn2(40)) // 60
// 	fmt.Println(fn2(50)) // 110
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
// 	fmt.Println(jpgFn("test")) // test.jpg
// 	fmt.Println(txtFn("test")) // test.txt
// }

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

// func main() {
// 	fn, fn2 := calc(10)
// 	fmt.Println(fn(1), fn2(2))
// 	fmt.Println(fn(3), fn2(4))
// 	fmt.Println(fn(5), fn2(6))
// }

// package main

// import "fmt"

// func adder() func(int) int {
// 	x := 0
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

package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	fn1, fn2 := calc(10)
	fmt.Println(fn1(1), fn2(2))
	fmt.Println(fn1(1), fn2(2))
	fmt.Println(fn1(1), fn2(2))
}
