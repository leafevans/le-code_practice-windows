package main

import "fmt"

// print2DArrayAddrs 打印 2x3 二维数组中每个元素的地址
func print2DArrayAddrs() {
	arr := [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}

	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("arr[%d][%d] 地址：%p\n", i, j, &arr[i][j])
		}
	}
}

func main() {
	print2DArrayAddrs()
}
