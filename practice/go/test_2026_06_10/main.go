package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range arr {
		if arr[i]%2 == 0 {
			fmt.Printf("%d is even\n", arr[i])
		} else {
			fmt.Printf("%d is odd\n", arr[i])
		}
	}
}
