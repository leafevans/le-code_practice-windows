// package main

// import "fmt"

// type Human struct {
// 	Name string
// 	Age  int
// }

//	func main() {
//		people := Human{Name: "张三", Age: 18}
//		fmt.Printf("%v %+v %#v %T %%\n", people, people, people, people)
//		fmt.Printf("%b", 5)
//		fmt.Printf("%c", 0x4E2D)
//		fmt.Printf("%d", 0b101)
//		fmt.Printf("%o", 0o122)
//		fmt.Printf("%s", "He said \"Hi\"")
//	}
package main

import "fmt"

func main() {
	num := 10.9 - 2.2i
	fmt.Printf("%g", num)
}
