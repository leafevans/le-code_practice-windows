package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["孙连城"] = 90
	scoreMap["李达康"] = 100
	scoreMap["沙瑞金"] = 60
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
}
