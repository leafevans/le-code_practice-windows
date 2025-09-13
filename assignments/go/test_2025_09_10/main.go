// package main

// import "fmt"

// func main() {
// 	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
// 	// fmt.Printf("%v %T\n", arr, arr)
// 	// var arr [10]int
// 	// for i := range 10 {
// 	// 	arr[i] = i
// 	// }
// 	// fmt.Println(arr)
// 	// array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	// fmt.Println(array)

// 	// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	// fmt.Println(arr, len(arr))

// 	// arr := [...]int{1: 3, 3: 5}
// 	// fmt.Println(arr)
// 	// for i := range len(arr) {
// 	// 	fmt.Println(arr[i])
// 	// }
// 	// for _, num := range arr {
// 	// 	fmt.Println(num)
// 	// }
// 	// var arr = [...]int{12, 23, 45, 67, 2, 5}
// 	// var sum = 0
// 	// for _, num := range arr {
// 	// 	sum += num
// 	// }
// 	// fmt.Printf("sum = %v, mean = %v\n", sum, float64(sum)/float64(len(arr)))
// 	// max := arr[0]
// 	// maxKey := 0
// 	// for key, num := range arr {
// 	// 	if num > max {
// 	// 		max = num
// 	// 		maxKey = key
// 	// 	}
// 	// }
// 	// fmt.Printf("最大值为 %v，其下标为 %v\n", max, maxKey)
// 	// arr := [...]int{1, 3, 5, 7, 8}
// 	// for i := range len(arr) {
// 	// 	for j := i + 1; j < len(arr); j++ {
// 	// 		if arr[i]+arr[j] == 8 {
// 	// 			fmt.Printf("(%v, %v)\n", i, j)
// 	// 		}
// 	// 	}
// 	// }
// 	arr := [...]int{1, 3, 5, 7, 8}
// 	fmt.Printf("%T\n", arr)
// }

// package main

// func main() {
// 	// 定义一个长度 3 元素类型为 int 的数组 arr 并赋值
// 	var arr [3]int
// 	arr[0] = 80
// 	arr[1] = 100
// 	arr[2] = 96
// }

// package main

// import "fmt"

// func main() {
// 	var testArray [3]int
// 	var numArray = [3]int{1, 2}
// 	var cityArray = [3]string{"北京", "上海", "深圳"}
// 	fmt.Println(testArray)
// 	fmt.Println(numArray)
// 	fmt.Println(cityArray)
// }

// package main

// import "fmt"

// func main() {
// 	var testArray [3]int
// 	var numArray = [...]int{1, 2}
// 	var cityArray = [...]string{"北京", "上海", "深圳"}
// 	fmt.Println(testArray)
// 	fmt.Println(numArray)
// 	fmt.Printf("Type of numArray: %T\n", numArray)
// 	fmt.Println(cityArray)
// 	fmt.Printf("Type of cityArray: %T\n", cityArray)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...]int{0: 1, 3: 5, 6: 8}
// 	fmt.Println(arr)
// 	fmt.Printf("Type of arr: %T\n", arr)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...]string{"北京", "上海", "深圳"}

// 	// 方法一：for 循环遍历
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(arr[i])
// 	}

// 	// 方法二：for range 遍历
// 	for index, value := range arr {
// 		fmt.Println(index, value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...]int{1: 3, 4: 5, 6: 8, 9: 10}
// 	fmt.Println(arr)
// }

// package main

// func main() {
// 	// var a = 10
// 	// b := a
// 	// fmt.Println(a, b)
// 	// a = 20
// 	// fmt.Println(a, b)
// 	// arr1 := [...]int{1, 2, 3}
// 	// arr2 := arr1
// 	// arr2[0] = 0
// 	// fmt.Println(arr1)
// 	// fmt.Println(arr2)
// 	// arr1 := []int{1, 2, 3}
// 	// arr2 := arr1
// 	// arr2[0] = 11
// 	// fmt.Println(arr1, arr2)
// 	// arr1 := []int{1, 2, 3, 4}
// 	// arr2 := arr1
// 	// arr2[0] = 11
// 	// arr := [...][3]int{{1, 2, 3}, {1, 2, 3}}
// 	// fmt.Println(arr)
// 	// arr := [3][2]string{
// 	// 	{"北京", "上海"},
// 	// 	{"广州", "深圳"},
// 	// 	{"成都", "重庆"},
// 	// }
// 	// fmt.Println(arr[0][1])
// }

// package main

// import "fmt"

// func main() {
// 	arr := [3][]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"成都", "重庆"},
// 	}
// 	for _, v1 := range arr {
// 		for _, v2 := range v1 {
// 			fmt.Println(v2)
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...][3]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [3]int{10, 20, 30}
// 	// arr2 是 arr 的副本
// 	arr2 := arr
// 	arr2[0] = 40
// 	fmt.Println(arr)
// 	fmt.Println(arr2)
// 	fmt.Println()

// 	arr3 := [3][2]int{
// 		{1, 1},
// 		{1, 1},
// 		{1, 1},
// 	}
// 	// arr4 是 arr3 的副本
// 	arr4 := arr3
// 	arr4[2][0] = 100
// 	fmt.Println(arr3)
// 	fmt.Println(arr4)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...][2]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"成都", "重庆"},
// 	}
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [...]int{1, -4, 232, 22, 54, 98, -34}
// 	sum := 0
// 	for _, val := range arr {
// 		sum += val
// 	}
// 	mean := float64(sum) / float64(len(arr))
// 	fmt.Printf("sum = %v, mean = %v\n", sum, mean)
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, -1, 12, 65, 11}
// 	max := arr[0]
// 	maxIndex := 0
// 	for i, v := range arr {
// 		if max < v {
// 			max = v
// 			maxIndex = i
// 		}
// 	}
// 	fmt.Printf("最大值为 %v，最大值的下标为 %v\n", max, maxIndex)
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
// 	// arr := []int{0, 1, 2, 3, 4}
// 	// fmt.Printf("%v %T %v\n", arr, arr, len(arr))
// 	// arr = append(arr, 5)
// 	// fmt.Println(arr)
// 	// arr := []int{0: 100, 1: 2, 3: 5, 9: 10}
// 	// fmt.Println(arr)
// 	// var arr []int
// 	// fmt.Println(arr, unsafe.Sizeof(arr))
// 	// var arr []int
// 	// fmt.Println(arr)
// 	// fmt.Println(arr == nil)
// 	// var num int
// 	// fmt.Println(num == 0)
// 	// var strSlice = []string{"php", "java", "nodejs", "golang"}
// 	// for k, v := range strSlice {
// 	// 	fmt.Println(k, v)
// 	// }
// 	// arr := [5]int{55, 56, 57, 58, 59}
// 	// slice := arr[:len(arr)-2]
// 	// fmt.Printf("%v %T\n", slice, slice)
// 	// arr := []string{"北京", "上海", "杭州", "苏州", "广州", "深圳"}
// 	// arr2 := arr[1:]
// 	// fmt.Printf("%v %T\n", arr2, arr2)
// 	arr := []int{2, 3, 5, 7, 11, 15}
// 	fmt.Printf("长度为 %d，容量为 %d\n", len(arr), cap(arr))
// 	arr2 := arr[1:3]
// 	fmt.Printf("长度为 %d，容量为 %d\n", len(arr2), cap(arr2))

// }

// package main

// func main() {
// 	// var a []string              // 声明字符串切片
// 	// var b = []int{}             // 声明整型切片并初始化
// 	// var c = []bool{false, true} // 声明布尔切片并初始化
// 	// var d = []bool{false, true} // 声明布尔切片并初始化
// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	// fmt.Println(c)
// 	// fmt.Println(a == nil)
// 	// fmt.Println(b == nil)
// 	// fmt.Println(c == nil)
// 	// 切片是引用类型，不支持直接比较，仅能和 nil 比较
// 	// fmt.Println(c == d)
// }

// package main

// import "fmt"

// func main() {
// 	arr := []string{"北京", "上海", "深圳"}
// 	for _, value := range arr {
// 		fmt.Println(value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{55, 56, 57, 58, 59}
// 	arr2 := arr[1:4]
// 	fmt.Println(arr2)
// 	fmt.Printf("Type of arr2: %T\n", arr2)
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

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	arr := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(arr)
// 	fmt.Printf("len = %v, cap = %v\n", len(arr), len(arr))
// 	fmt.Println(unsafe.Sizeof(arr))
// 	// arr2 := arr[:2]
// 	// fmt.Println(arr2)
// 	// fmt.Printf("len = %v, cap = %v\n", len(arr2), cap(arr2))

// 	// arr3 := arr[1:3]
// 	// fmt.Println(arr3)
// 	// fmt.Printf("len = %v, cap = %v\n", len(arr3), cap(arr3))
// }

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{55, 56, 57, 58, 59}
// 	arr2 := arr[1:]
// 	fmt.Println(arr2)
// 	fmt.Printf("Type of arr2: %T\n", arr2)
// 	fmt.Println(cap(arr2))
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

package main

import "fmt"

func main() {
	// arr := make([]int, 10, 20)
	// fmt.Println(arr, len(arr), cap(arr))
	// for i := range 10 {
	// 	arr[i] = i
	// }
	// fmt.Println(arr, len(arr), cap(arr))
	// var slice []int
	// fmt.Printf("val = %v, len = %v, cap = %v\n", slice, len(slice), cap(slice))
	// var arr []int
	// arr = append(arr, 10)
	// fmt.Println(arr)
	// var arr []int
	// arr = append(arr, 12, 24, 32, 486)
	// fmt.Println(arr)
	slice := []string{"php", "java"}
	slice2 := []string{"nodejs", "go"}
	slice = append(slice, slice2...)
	fmt.Println(slice)
}
