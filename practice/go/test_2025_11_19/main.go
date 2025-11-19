package main

import "fmt"

func main() {
	num := -4.2 + 101i
	str := "Hello World!"
	fmt.Printf("%.2E\n", num)
	fmt.Printf("%04.2s\n", str)
}
