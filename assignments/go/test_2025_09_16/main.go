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

package main

import "fmt"

func buildNext(pattern string) []int {
	next := make([]int, len(pattern))
	j := 0

	for i := 1; i < len(pattern); i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}

	return next
}

func kmpSearch(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}

	next := buildNext(pattern)
	j := 0

	for i := range len(text) {
		for j > 0 && text[i] != pattern[j] {
			j = next[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return i - j + 1
		}
	}

	return -1
}

func main() {
	text := "ABABCABCACBAB"
	pattern := "ABCAC"

	fmt.Println(kmpSearch(text, pattern))
}
