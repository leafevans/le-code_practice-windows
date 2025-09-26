// package main

// import "fmt"

// func main() {
// 	scoreMap := make(map[string]int)
// 	scoreMap["孙连城"] = 90
// 	scoreMap["李达康"] = 100
// 	scoreMap["沙瑞金"] = 60
// 	for key, value := range scoreMap {
// 		fmt.Println(key, value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	scoreMap := make(map[string]int)
// 	scoreMap["孙连城"] = 90
// 	scoreMap["李达康"] = 100
// 	scoreMap["沙瑞金"] = 60
// 	delete(scoreMap, "李达康")
// 	for key := range scoreMap {
// 		fmt.Println(key)
// 	}
// }

// package main

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := range n - 1 {
// 		for j := range n - 1 - i {
// 			if slice[j+1] < slice[j] {
// 				slice[j+1], slice[j] = slice[j], slice[j+1]
// 			}
// 		}
// 	}
// }

// func selectionSort(slice []int) {
// 	n := len(slice)
// 	for i := range n {
// 		minIndex := i
// 		for j := i + 1; j < n; j++ {
// 			if slice[j] < slice[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		slice[i], slice[minIndex] = slice[minIndex], slice[i]
// 	}
// }

// func insertionSort(slice []int) {
// 	n := len(slice)
// 	for i := 1; i < n; i++ {
// 		key := slice[i]
// 		j := i - 1
// 		for j >= 0 && slice[j] < key {
// 			slice[j+1] = slice[j]
// 			j--
// 		}
// 		slice[j+1] = key
// 	}
// }

// func quickSort(slice []int) {
// 	if len(slice) <= 1 {
// 		return
// 	}

// 	pivot := slice[len(slice)/2]
// 	left, right := 0, len(slice)
// 	for left <= right {
// 		for slice[left] < slice[pivot] {
// 			left++
// 		}
// 		for slice[right] > slice[pivot] {
// 			right--
// 		}
// 		if left <= right {
// 			slice[left], slice[right] = slice[right], slice[left]
// 			left++
// 			right--
// 		}
// 	}

// 	quickSort(slice[left:])
// 	quickSort(slice[:right+1])
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

// func mergeSort(slice []int) []int {
// 	if len(slice) <= 1 {
// 		return slice
// 	}
// 	mid := len(slice) / 2
// 	left := mergeSort(slice[:mid])
// 	right := mergeSort(slice[mid:])
// 	return merge(left, right)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
// 	floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
// 	stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
// 	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
// 	sort.Sort(sort.Reverse(sort.Float64Slice(floatSlice)))
// 	sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
// 	fmt.Println(intSlice)
// 	fmt.Println(floatSlice)
// 	fmt.Println(stringSlice)
// 	sort.Ints(intSlice)
// 	sort.Float64s(floatSlice)
// 	sort.Strings(stringSlice)
// 	fmt.Println(intSlice)
// 	fmt.Println(floatSlice)
// 	fmt.Println(stringSlice)
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sort"
// )

// func main() {
// 	scoreMap := make(map[string]int, 200)
// 	for i := range 100 {
// 		scoreMap[fmt.Sprintf("stu%02d", i)] = rand.Intn(100)
// 	}

// 	keys := make([]string, 0, 200)
// 	for key := range scoreMap {
// 		keys = append(keys, key)
// 	}

// 	sort.Strings(keys)
// 	for _, key := range keys {
// 		fmt.Println(key, scoreMap[key])
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	slice := []int{}
// 	fmt.Println(slice, slice == nil)
// }

// package main

// import "fmt"

// func main() {
// 	userinfo := make([]map[string]string, 3)

// 	if userinfo[0] == nil {
// 		userinfo[0] = make(map[string]string)
// 		userinfo[0]["username"] = "哈基米"
// 		userinfo[0]["age"] = "18"
// 		userinfo[0]["height"] = "180"
// 	}

// 	if userinfo[1] == nil {
// 		userinfo[1] = make(map[string]string)
// 		userinfo[1]["username"] = "叮咚鸡"
// 		userinfo[1]["age"] = "17"
// 		userinfo[1]["height"] = "175"
// 	}

// 	if userinfo[2] == nil {
// 		userinfo[2] = make(map[string]string)
// 		userinfo[2]["username"] = "老爷爷"
// 		userinfo[2]["age"] = "80"
// 		userinfo[2]["height"] = "165"
// 	}

// 	for _, v := range userinfo {
// 		fmt.Println(v)
// 		for key, value := range v {
// 			fmt.Printf("key: %v, value: %v\n", key, value)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	// userinfo := make(map[string][]string)

// 	// userinfo["hobby"] = []string{
// 	// 	"哈基米",
// 	// 	"曼波",
// 	// 	"叮咚鸡",
// 	// 	"大狗叫",
// 	// }

// 	// userinfo["work"] = []string{
// 	// 	"野火哈基米",
// 	// 	"春风吹有咪",
// 	// }

// 	// for _, v := range userinfo {
// 	// 	for _, value := range v {
// 	// 		fmt.Println(value)
// 	// 	}
// 	// }

// 	// userinfo := make(map[string]string)

// 	// userinfo["username"] = "哈基米"
// 	// userinfo["age"] = "20"
// 	// userinfo2 := userinfo

// 	// fmt.Println(userinfo)
// 	// fmt.Println(userinfo2)

// 	// userinfo2["username"] = "曼波"
// 	// fmt.Println(userinfo)
// 	// fmt.Println(userinfo2)

// 	map1 := make(map[int]int, 10)
// 	map1[10] = 100
// 	map1[1] = 13
// 	map1[4] = 56
// 	map1[8] = 90
// 	fmt.Println(map1)

// 	// for key, value := range map1 {
// 	// 	fmt.Println(key, value)
// 	// }

// 	keys := []int{}

// 	for key := range map1 {
// 		keys = append(keys, key)
// 	}

// 	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

// 	for _, key := range keys {
// 		fmt.Println(key, map1[key])
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "how do you do"
// 	strSlice := strings.Split(str, " ")
// 	fmt.Println(strSlice)

// 	strMap := make(map[string]int)

// 	for _, v := range strSlice {
// 		strMap[v]++
// 	}

// 	fmt.Println(strMap)
// }

// package main

// import "fmt"

// func main() {
// 	mapSlice := make([]map[string]string, 3)
// 	for i, v := range mapSlice {
// 		fmt.Printf("i = %v, v = %v\n", i, v)
// 	}
// 	fmt.Println("After init...")
// 	// 对切片中的 map 元素进行初始化
// 	mapSlice[0] = make(map[string]string, 10)
// 	mapSlice[0]["name"] = "哈基米"
// 	mapSlice[0]["password"] = "123456"
// 	mapSlice[0]["address"] = "双流区"
// 	for i, v := range mapSlice {
// 		fmt.Printf("i = %v, v = %v\n", i, v)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	sliceMap := make(map[string][]string, 3)
// 	fmt.Println(sliceMap)
// 	fmt.Println("After init...")
// 	// 查找是否存在“中国”的键
// 	key := "中国"
// 	value, ok := sliceMap[key]
// 	// 不存在说明 ok 为 false，value 为 nil
// 	// 对 value 进行切片初始化
// 	if !ok {
// 		value = make([]string, 0, 2)
// 	}
// 	// 添加切片元素
// 	value = append(value, "北京", "上海")
// 	// 添加键值对
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
// 	str := "how  do you   do"
// 	for word := range strings.FieldsSeq(str) {
// 		wordMap[word]++
// 	}
// 	fmt.Println(wordMap)
// }

// package main

// import (
// 	"fmt"
// )

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := range n - 1 {
// 		for j := range n - 1 - i {
// 			if slice[j+1] < slice[j] {
// 				slice[j+1], slice[j] = slice[j], slice[j+1]
// 			}
// 		}
// 	}
// }

// func selectionSort(slice []int) {
// 	n := len(slice)
// 	for i := range n {
// 		minIndex := i
// 		for j := i + 1; j < n; j++ {
// 			if slice[j] < slice[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		slice[i], slice[minIndex] = slice[minIndex], slice[i]
// 	}
// }

// func insertionSort(slice []int) {
// 	n := len(slice)
// 	for i := 1; i < n; i++ {
// 		key := slice[i]
// 		j := i - 1
// 		for j >= 0 && slice[j] < key {
// 			slice[j+1] = slice[j]
// 			j--
// 		}
// 		slice[j+1] = key
// 	}
// }

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
// 	quickSort(slice[left:])
// 	quickSort(slice[:right+1])
// }

// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))
// 	i, j := 0, 0
// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}
// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)
// 	return result
// }

// func mergeSort(slice []int) []int {
// 	if len(slice) <= 1 {
// 		return slice
// 	}
// 	mid := len(slice) / 2
// 	left := mergeSort(slice[:mid])
// 	right := mergeSort(slice[mid:])
// 	return merge(left, right)
// }

// func main() {
// 	slice := []int{23, 5, 87, 12, 45, 3, 66, 29, 54, 10}
// 	fmt.Println(slice)
// 	slice = mergeSort(slice)
// 	fmt.Println(slice)
// }

// package main

// import "fmt"

// type Human struct {
// 	Name string
// }

// func main() {
// 	people := Human{Name: "哈基米"}
// 	fmt.Printf("%v\n%+v\n%#v\n%T\n%%\n", people, people, people, people)
// }

// package main

// import "fmt"

// type Human struct {
// 	Name string
// }

// func main() {
// 	people := Human{Name: "哈基米"}
// 	fmt.Printf("%v\n%+v\n%#v\n%T\n%%\n", people, people, people, people)
// 	fmt.Printf("%t\n", true)
// 	fmt.Printf("%b\n%c\n%d\n%o\n%x\n%X\n", 100, 100, 100, 100, 100, 100)
// }

// package main

// import "fmt"

// func main() {
// 	f := 314.15926
// 	fmt.Printf("%.2f\n%.2e\n%.3g\n%.100b\n", f, f, f, f)
// 	// str := "哈基米喔南北绿豆"
// 	// fmt.Printf("%-5.2s\n", str)
// }

// package main

// import "fmt"

// func main() {
// 	// num := 5.0
// 	// fmt.Printf("[%-6.2f]\n", num)
// 	str := "哈基米"
// 	fmt.Printf("[%-6.2s]\n", str)
// }

// package main

// import (
// 	"fmt"
// )

// func heapify(slice []int, n, i int) {
// 	largest := i
// 	left := 2*i + 1
// 	right := 2*i + 2
// 	if left < n && slice[left] > slice[largest] {
// 		largest = left
// 	}
// 	if right < n && slice[right] > slice[largest] {
// 		largest = right
// 	}
// 	if largest != i {
// 		slice[i], slice[largest] = slice[largest], slice[i]
// 		heapify(slice, n, largest)
// 	}
// }

// func heapSort(slice []int) {
// 	n := len(slice)
// 	for i := n/2 - 1; i >= 0; i-- {
// 		heapify(slice, n, i)
// 	}
// 	for i := n - 1; i >= 0; i-- {
// 		slice[0], slice[i] = slice[i], slice[0]
// 		heapify(slice, i, 0)
// 	}
// }

// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))
// 	i, j := 0, 0
// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}
// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)
// 	return result
// }

// func mergeSort(slice []int) []int {
// 	if len(slice) <= 1 {
// 		return slice
// 	}

// 	mid := len(slice) / 2
// 	left := mergeSort(slice[:mid])
// 	right := mergeSort(slice[mid:])
// 	return merge(left, right)
// }

// func shellSort(slice []int) {
// 	n := len(slice)
// 	for gap := n / 2; gap > 0; gap /= 2 {
// 		for i := gap; i < n; i++ {
// 			temp := slice[i]
// 			j := i
// 			for j >= gap && slice[j-gap] < temp {
// 				slice[j] = slice[j-gap]
// 				j -= gap
// 			}
// 			slice[j] = temp
// 		}
// 	}
// }

// func main() {
// 	slice := []int{23, 5, 87, 12, 45, 3, 66, 29, 54, 10}
// 	fmt.Println("Before:", slice)
// 	shellSort(slice)
// 	fmt.Println("After:", slice)
// }

// package main

// import (
// 	"fmt"

// 	"github.com/shopspring/decimal"
// )

// func main() {
// 	// const (
// 	// 	A = 0o111
// 	// 	B = 0x1p-2
// 	// 	C
// 	// 	D = iota
// 	// 	E = 13.23_2
// 	// )
// 	// fmt.Println(A, B, C, D, E)
// 	// fmt.Println(unsafe.Sizeof(A))
// 	// fmt.Println(math.MaxFloat32)
// 	// fmt.Println(math.MaxFloat64)
// 	// fmt.Println(math.Pi)
// 	// fmt.Println(math.E)
// 	d := decimal.NewFromFloat(1.23)
// 	d2 := decimal.NewFromInt(123)
// 	d3 := decimal.New(123, -2)
// 	fmt.Printf("%T %T %T\n", d, d2, d3)
// }

// package main

// import (
// 	"fmt"

// 	"github.com/shopspring/decimal"
// )

// func main() {
// 	a := decimal.NewFromFloat(1.23)
// 	b := decimal.NewFromFloat(4.56)

// 	sum := a.Add(b)
// 	diff := a.Sub(b)
// 	product := a.Mul(b)
// 	quotient := a.Div(b)

// 	fmt.Println(sum)
// 	fmt.Println(diff)
// 	fmt.Println(product)
// 	fmt.Println(quotient)

// 	if a.GreaterThan(b) {
// 		fmt.Println("a > b")
// 	}

// 	floatVal, _ := sum.Float64()
// 	intVal := sum.IntPart()

// 	fmt.Println(floatVal)
// 	fmt.Println(intVal)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// str := "123-456-789"
// 	// arr := strings.Split(str, "-")
// 	// fmt.Println(arr)
// 	// for _, v := range arr {
// 	// 	fmt.Println(v)
// 	// }
// 	// for word := range strings.SplitSeq(str, "-") {
// 	// 	fmt.Println(word)
// 	// }
// 	var str = "this is golang"
// 	// fmt.Println(strings.Contains(str, "golang"))
// 	// fmt.Println(strings.HasPrefix(str, "this"))
// 	// fmt.Println(strings.HasSuffix(str, "golang"))
// 	// str, _ = strings.CutPrefix(str, "this ")
// 	// str, _ = strings.CutSuffix(str, " golang")
// 	// fmt.Println(str)
// 	var index = strings.Index(str, "is")
// 	fmt.Println(index)

// 	index = strings.LastIndex(str, "is")
// 	fmt.Println(index)
// }

// package main

// import "fmt"

// func main() {
// 	// var str = "123-456-789"
// 	// var arr = strings.Split(str, "-")
// 	// var str2 = strings.Join(arr, "*")
// 	// fmt.Println(str2)
// 	// str := "基本数据类型ABC"
// 	// for _, ch := range str {
// 	// 	fmt.Printf("%v(%c)", ch, ch)
// 	// }
// 	ch := '营'
// 	ch2 := "营"
// 	fmt.Printf("%T %T\n", ch, ch2)

// 	str := "big"
// 	byteStr := []byte(str)
// 	byteStr[0] = 'p'
// 	fmt.Println(string(byteStr))

// 	str2 := "红萝卜"
// 	runeStr2 := []rune(str2)
// 	runeStr2[0] = '白'
// 	fmt.Println(string(runeStr2))
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	a, b := 3, 4
// 	c := int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }

// package main

// func main() {
// 	// num := 20
// 	// str := strconv.Itoa(num)
// 	// fmt.Printf("value = %v, type = %T\n", str, str)

// 	// str2 := strconv.FormatInt(int64(num), 12)
// 	// fmt.Printf("value = %v, type = %T\n", str2, str2)

// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	num := 3.1415926
// 	str := strconv.FormatFloat(num, 'G', 2, 64)
// 	fmt.Printf("%v %T\n", str, str)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	str := "1234"
// 	num, _ := strconv.ParseInt(str, 10, 32)
// 	fmt.Printf("%v %T\n", num, num)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	str := "3.1415926"
// 	num, _ := strconv.ParseFloat(str, 32)
// 	fmt.Println(num)
// 	num2, _ := strconv.ParseFloat(str, 64)
// 	fmt.Println(num2)
// }

// package main

// import "fmt"

// func test() bool {
// 	fmt.Println("test...")
// 	return true
// }

// func main() {
// 	var i int = 10
// 	// 前条件错，后面不检查
// 	if i < 9 && test() {
// 		fmt.Println("ok...")
// 	}
// 	// 前条件对，后面不检查
// 	if i > 9 || test() {
// 		fmt.Println("hello...")
// 	}
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

// func main() {
// 	if score := 56; score >= 90 {
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
// 	k := 0
// 	for {
// 		if k < 10 {
// 			fmt.Println("OK~~~", k)
// 		} else {
// 			break
// 		}
// 		k++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	count := 0
// 	sum := 0
// 	for i := range 100 {
// 		if i%9 == 0 {
// 			count++
// 			sum += i
// 		}
// 	}
// 	fmt.Println(sum, count)
// }

// package main

// import "fmt"

// func main() {
// 	n := 1
// 	for i := 1; i <= 5; i++ {
// 		n *= 1
// 	}
// 	fmt.Println(n)
// }

// package main

// import "fmt"

// func main() {
// 	for range 10 {
// 		for range 5 {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	n := 6
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
// 	for i := 1; i <= 9; i++ {
// 		for j := i; j <= 9; j++ {
// 			fmt.Printf("%v x %v = %v\t", i, j, i*j)
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	switch extname := ".html"; extname {
// 	case ".html":
// 		fmt.Println("text/html")
// 		fallthrough
// 	case ".css":
// 		fmt.Println("text/css")
// 		fallthrough
// 	case ".js":
// 		fmt.Println("text/javascript")
// 		fallthrough
// 	default:
// 		fmt.Println("格式错误")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// switch n := 2; n {
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

// 	for i := range 2 {
// 		for j := range 10 {
// 			if j == 3 {
// 				goto label
// 			}
// 			fmt.Printf("i = %v, j = %v\n", i, j)
// 		}
// 	}
// label:
// }

// package main

// import "fmt"

// func main() {
// 	// arr := [...]int{0: 1, 11: 100}
// 	// fmt.Println(arr, unsafe.Sizeof(arr))
// 	arr := [...][2]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"成都", "重庆"},
// 	}

// 	for _, v1 := range arr {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s\n", v2)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 3, 5, 7, 8}
// 	for i := range arr {
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[i]+arr[j] == 8 {
// 				fmt.Printf("(%v, %v)\n", i, j)
// 			}
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
// 	fmt.Printf("arr: %v type: %T len: %v cap: %v\n", arr, arr, len(arr), cap(arr))
// 	arr2 := arr[1:3]
// 	fmt.Printf("arr2: %v type: %T len: %v cap: %v\n", arr2, arr2, len(arr2), cap(arr2))
// 	arr3 := arr2[1:5]
// 	fmt.Printf("arr3: %v type: %T len: %v cap: %v\n", arr3, arr3, len(arr3), cap(arr3))
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(arr)
// 	fmt.Printf("len = %v, cap = %v\n", len(arr), len(arr))

// 	arr2 := arr[:2]
// 	fmt.Println(arr2)
// 	fmt.Printf("len = %v, cap = %v\n", len(arr2), cap(arr2))

// 	arr3 := arr[1:3]
// 	fmt.Println(arr3)
// 	fmt.Printf("len = %v, cap = %v\n", len(arr3), cap(arr3))
// }

// package main

// import "fmt"

// func main() {
// 	numSlice := []int{30, 31, 32, 33, 34, 35, 36, 37}
// 	numSlice = append(numSlice[:2], numSlice[3:]...)
// 	fmt.Println(numSlice)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func bubbleSort(arr []int) {
// 	n := len(arr)
// 	for i := range n - 1 {
// 		for j := range n - 1 - i {
// 			if arr[j+1] < arr[j] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// }

// func selectionSort(arr []int) {
// 	n := len(arr)
// 	for i := range n {
// 		minIndex := i
// 		for j := i + 1; j < n; j++ {
// 			if arr[j] < arr[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if minIndex != i {
// 			arr[i], arr[minIndex] = arr[minIndex], arr[i]
// 		}
// 	}
// }

// func insertionSort(arr []int) {
// 	n := len(arr)
// 	for i := 1; i < n; i++ {
// 		key := arr[i]
// 		j := i - 1
// 		for j >= 0 && arr[j] > key {
// 			arr[j+1] = arr[j]
// 			j--
// 		}
// 		arr[j+1] = key
// 	}
// }

// func shellSort(arr []int) {
// 	n := len(arr)
// 	for gap := n / 2; gap > 0; gap /= 2 {
// 		for i := gap; i < n; i++ {
// 			temp := arr[i]
// 			j := i
// 			for j >= gap && arr[j-gap] > temp {
// 				arr[j] = arr[j-gap]
// 				j -= gap
// 			}
// 			arr[j] = temp
// 		}
// 	}
// }

// func quickSort(arr []int) {
// 	if len(arr) <= 1 {
// 		return
// 	}

// 	pivot := arr[len(arr)/2]
// 	left, right := 0, len(arr)-1

// 	for left <= right {
// 		for arr[left] < pivot {
// 			left++
// 		}
// 		for arr[right] > pivot {
// 			right--
// 		}

// 		if left <= right {
// 			arr[left], arr[right] = arr[right], arr[left]
// 			left++
// 			right--
// 		}
// 	}

// 	quickSort(arr[left:])
// 	quickSort(arr[:right+1])
// }

// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))
// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
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
// 	largest, left, right := i, 2*i+1, 2*i+2

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
// 	arr := []int{3, 7, 8, 9, 1}
// 	fmt.Println(arr)
// 	sort.Ints(arr)
// 	fmt.Println(arr)
// 	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func main() {
// 	scoreMap := make(map[string]int, 8)
// 	fmt.Println(scoreMap == nil)
// 	scoreMap["哈基米"] = 90
// 	scoreMap["曼波"] = 100
// 	fmt.Println(scoreMap)
// 	fmt.Println(scoreMap["曼波"])
// 	fmt.Printf("Type of scoreMap: %T\n", scoreMap)
// }

// package main

// import "fmt"

// func main() {
// 	scoreMap := make(map[string]int, 8)
// 	scoreMap["哈基米"] = 90
// 	scoreMap["曼波"] = 100
// 	value, ok := scoreMap["哈基米"]
// 	if ok {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("查无此人")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	scoreMap := make(map[string]int)
// 	scoreMap["孙连城"] = 90
// 	scoreMap["李达康"] = 100
// 	scoreMap["沙瑞金"] = 60
// 	delete(scoreMap, "沙瑞金")
// 	for key := range scoreMap {
// 		fmt.Println(key)
// 	}
// }

package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

func main() {
	scoreMap := make(map[string]int, 100)

	for i := range 100 {
		scoreMap[fmt.Sprintf("stu%02d", i)] = rand.IntN(100)
	}

	keys := make([]string, 100)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
