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

package main

import (
	"fmt"
	"sort"
)

func mapSort(userinfo map[string]string) string {
	keys := make([]string, 0, len(userinfo))

	for key := range userinfo {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	str := ""

	for _, key := range keys {
		str += fmt.Sprintf("%v=>%v ", key, userinfo[key])
	}

	return str
}

func main() {
	userinfo := map[string]string{
		"username": "王婷婷",
		"age":      "20",
		"sex":      "女",
		"height":   "165",
	}

	str := mapSort(userinfo)

	fmt.Println(str)
}
