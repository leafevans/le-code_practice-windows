// package main

// import "fmt"

// func main() {
// 	// fmt.Println("Hello World!")
// 	var slice []int

// 	for i := range 10 {
// 		slice = append(slice, i)
// 		fmt.Printf("val = %v, len = %d, cap = %d\n", slice, len(slice), cap(slice))
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// slice := []int{1, 2, 3, 4, 5}
// 	// slice2 := make([]int, 5)
// 	// copy(slice2, slice)
// 	// slice2[0] = 11
// 	// fmt.Println(slice, slice2)
// 	// slice2 := slice
// 	// slice2[0] = 11
// 	// fmt.Println(slice, slice2)
// 	arr := []int{30, 31, 32, 33, 34, 35, 36, 37}
// 	arr = append(arr[:5], arr[7:]...)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func main() {
// 	slice := make([]int, 2, 10)
// 	fmt.Printf("val = %v, len = %v, cap = %v\n", slice, len(slice), cap(slice))
// }

// package main

// import "fmt"

// func main() {
// 	var slice []int
// 	slice2 := []int{}
// 	slice3 := make([]int, 0)

// 	fmt.Printf("bool = %v, len = %v, cap = %v\n", slice == nil, len(slice), cap(slice))
// 	fmt.Printf("bool = %v, len = %v, cap = %v\n", slice2 == nil, len(slice2), cap(slice2))
// 	fmt.Printf("bool = %v, len = %v, cap = %v\n", slice3 == nil, len(slice3), cap(slice3))
// }

// package main

// import "fmt"

// func main() {
// 	slice := make([]int, 3)
// 	slice2 := slice
// 	slice2[0] = 100
// 	fmt.Println(slice)
// 	fmt.Println(slice2)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var numSlice []int
// 	for i := range 10 {
// 		numSlice = append(numSlice, i)
// 		fmt.Printf("%v\tlen: %v\tcap: %v\tptr: %p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var citySlice []string
// 	citySlice = append(citySlice, "北京")
// 	citySlice = append(citySlice, "上海", "广州", "深圳")
// 	cityArr := [2]string{"成都", "杭州"}
// 	citySlice = append(citySlice, cityArr[:]...)
// 	fmt.Println(citySlice)
// }

// package main

// import "fmt"

// func main() {
// 	slice := []int{100, 200, 300}
// 	slice2 := []int{400, 500, 600}
// 	slice3 := append(slice, slice2...)
// 	fmt.Println(slice3)
// }

// package main

// import "fmt"

// func main() {
// 	numSlice := []int{30, 31, 32, 33, 34, 35, 36, 37}
// 	numSlice = append(numSlice[:2], numSlice[3:]...)
// 	fmt.Println(numSlice)
// }

// package main

// import "fmt"

// func main() {
// 	var slice = make([]string, 5, 10)
// 	for i := range 12 {
// 		slice = append(slice, fmt.Sprintf("%v", i))
// 	}
// 	fmt.Println(slice)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	slice := []int{3, 7, 8, 9, 1}
// 	sort.Ints(slice)
// 	fmt.Println(slice)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...]int{1, 3, 5, 7, 8}
// 	for i := range len(arr) {
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[i]+arr[j] == 8 {
// 				fmt.Printf("(%v, %v)\n", i, j)
// 			}
// 		}
// 	}
// }

// package main

// import "fmt"

// func selectionSort(slice []int) {
// 	n := len(slice)
// 	for i := range n - 1 {
// 		minIndex := i
// 		for j := i + 1; j < n; j++ {
// 			if slice[j] < slice[minIndex] {
// 				minIndex = j
// 			}
// 		}

// 		if minIndex != i {
// 			slice[i], slice[minIndex] = slice[minIndex], slice[i]
// 		}
// 	}
// }

// func main() {
// 	slice := []int{64, 25, 12, 22, 11}
// 	fmt.Println("排序前:", slice)
// 	selectionSort(slice)
// 	fmt.Println("排序后:", slice)
// }

// package main

// import "fmt"

// func selectionSort(slice []int) {
// 	n := len(slice)
// 	for i := range n - 1 {
// 		minIndex := i
// 		for j := i + 1; j < n; j++ {
// 			if slice[j] < slice[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if i != minIndex {
// 			slice[i], slice[minIndex] = slice[minIndex], slice[i]
// 		}
// 	}
// }

// func main() {
// 	arr := [...]int{12, 32, 1, 89, 43, 55}
// 	selectionSort(arr[:])
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := range n - 1 {
// 		for j := range n - 1 - i {
// 			if slice[j] > slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// }

// func main() {
// 	arr := []int{45, 32, 56, 90, 12, 13}
// 	bubbleSort(arr)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := range n - 1 {
// 		for j := range n - 1 - i {
// 			if slice[j] > slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// }

// func main() {
// 	arr := []int{45, 32, 56, 90, 12, 13}
// 	bubbleSort(arr)
// 	fmt.Println(arr)
// }

// package main

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := range n - 1 {
// 		for j := i + 1; j < n-1-i; j++ {
// 			if slice[j] < slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
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

// package main

// import "fmt"

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
// 		for j >= 0 && slice[j] > key {
// 			slice[j+1] = slice[j]
// 			j--
// 		}
// 		slice[j+1] = key
// 	}
// }

// func main() {
// 	arr := []int{45, 32, 56, 90, 12, 13}
// 	insertionSort(arr)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-1-i; j++ {
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
// 		if i != minIndex {
// 			slice[minIndex], slice[i] = slice[i], slice[minIndex]
// 		}
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
// 	right := mergeSort(slice[mid+1:])
// 	return merge(left, right)
// }

// func main() {
// 	arr := []int{45, 32, 56, 90, 12, 13}
// 	mergeSort(arr)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

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
// 		minIndex := slice[0]
// 		for j := i + 1; j < n; j++ {
// 			if slice[j] < slice[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if i != minIndex {
// 			slice[i], slice[minIndex] = slice[minIndex], slice[i]
// 		}
// 	}
// }

// func insertionSort(slice []int) {
// 	n := len(slice)
// 	for i := 1; i < n; i++ {
// 		key := slice[len(slice)/2]
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
// 	slice := []int{54, 23, 98, 1023, 2, 433}
// 	sorted := mergeSort(slice)
// 	fmt.Println(sorted)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	slice := []int{3, 7, 8, 9, 1}
// 	sort.Sort(sort.Reverse(sort.IntSlice(slice))) // 降序排序
// 	fmt.Println(slice)                            // 输出: [9 8 7 3 1]
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
// }
