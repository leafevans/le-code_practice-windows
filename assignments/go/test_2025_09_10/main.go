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

package main

import "fmt"

func main() {
	arr := [...]int{1: 3, 4: 5, 6: 8, 9: 10}
	fmt.Println(arr)
}
