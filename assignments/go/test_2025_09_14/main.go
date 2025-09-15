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

package main

func main() {
	// num := 20
	// str := strconv.Itoa(num)
	// fmt.Printf("value = %v, type = %T\n", str, str)

	// str2 := strconv.FormatInt(int64(num), 12)
	// fmt.Printf("value = %v, type = %T\n", str2, str2)

}
