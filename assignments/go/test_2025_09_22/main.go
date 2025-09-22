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
// 	text := "ababcabcacbab"
// 	pattern := "abcac"
// 	pos := kmpSearch(text, pattern)
// 	fmt.Println("Pattern found at index:", pos)
// }

// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func bfs(graph map[int][]int, start int) {
// 	// visited := make(map[int]bool)
// 	// queue := list.New()
// 	// queue.PushBack(start)
// 	// visited[start] = true

// 	// for queue.Len() > 0 {
// 	// 	node := queue.Remove(queue.Front()).(int)
// 	// 	fmt.Printf("%d ", node)
// 	// 	for _, neighbor := range graph[node] {
// 	// 		if !visited[neighbor] {
// 	// 			visited[neighbor] = true
// 	// 			queue.PushBack(neighbor)
// 	// 		}
// 	// 	}
// 	// }
// 	visited := make(map[int]bool)
// 	queue := list.New()
// 	queue.PushBack(start)
// 	visited[start] = true

// 	for queue.Len() > 0 {
// 		node := queue.Remove(queue.Front()).(int)
// 		fmt.Printf("%d ", node)
// 		for _, neighbor := range graph[node] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue.PushBack(neighbor)
// 			}
// 		}
// 	}
// }

// func main() {
// 	graph := map[int][]int{
// 		0: {1, 2},
// 		1: {0, 3, 4},
// 		2: {0, 5},
// 		3: {1},
// 		4: {1},
// 		5: {2},
// 	}
// 	bfs(graph, 0)
// }

// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func bfs(graph map[int][]int, start int) {
// 	visited := make(map[int]bool)
// 	queue := list.New()
// 	queue.PushBack(start)
// 	visited[start] = true

// 	for queue.Len() > 0 {
// 		node := queue.Remove(queue.Front()).(int)
// 		fmt.Printf("%d ", node)
// 		for _, neighbor := range graph[node] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue.PushBack(neighbor)
// 			}
// 		}
// 	}
// }

// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func bfs(graph map[int][]int, start int) {
// 	visited := make(map[int]bool)
// 	queue := list.New()
// 	queue.PushBack(start)
// 	visited[start] = true

// 	for queue.Len() > 0 {
// 		node := queue.Remove(queue.Front()).(int)
// 		fmt.Printf("%d ", node)
// 		for _, neighbor := range graph[node] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue.PushBack(neighbor)
// 			}
// 		}
// 	}
// }

// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func bfs(graph map[int][]int, start int) {
// 	visited := make(map[int]bool)
// 	queue := list.New()
// 	queue.PushBack(start)
// 	visited[start] = true

// 	for queue.Len() > 0 {
// 		node := queue.Remove(queue.Front()).(int)
// 		fmt.Printf("%d ", node)
// 		for _, neighbor := range graph[node] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue.PushBack(neighbor)
// 			}
// 		}
// 	}
// }

// package main

// import "fmt"

// func dfs(graph map[int][]int, node int, visited map[int]bool) {
// 	visited[node] = true
// 	fmt.Printf("%d ", node)
// 	for _, neighbor := range graph[node] {
// 		if !visited[neighbor] {
// 			dfs(graph, neighbor, visited)
// 		}
// 	}
// }

// func main() {
// 	graph := map[int][]int{
// 		0: {1, 2},
// 		1: {0, 3, 4},
// 		2: {0, 5},
// 		3: {1},
// 		4: {1},
// 		5: {2},
// 	}
// 	visited := make(map[int]bool)
// 	dfs(graph, 0, visited)
// 	fmt.Println()
// }

// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func bfs(graph map[int][]int, start int) {
// 	visited := make(map[int]bool)
// 	queue := list.New()
// 	queue.PushBack(start)
// 	visited[start] = true

// 	for queue.Len() > 0 {
// 		node := queue.Remove(queue.Front()).(int)
// 		for _, neighbor := range graph[node] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue.PushBack(neighbor)
// 			}
// 		}
// 	}
// }

// func dfs(graph map[int][]int, node int, visited map[int]bool) {
// 	visited[node] = true
// 	for _, neighbor := range graph[node] {
// 		if !visited[neighbor] {
// 			dfs(graph, neighbor, visited)
// 		}
// 	}
// }

// func main() {
// 	graph := map[int][]int{
// 		0: {1, 2},
// 		1: {0, 3, 4},
// 		2: {0, 5},
// 		3: {1},
// 		4: {1},
// 		5: {2},
// 	}
// 	visited := make(map[int]bool)
// 	dfs(graph, 0, visited)
// 	fmt.Println()
// }

// package main

// import "container/list"

// func bfs(graph map[int][]int, start int) {
// 	visited := make(map[int]bool)
// 	queue := list.New()
// 	visited[start] = true
// 	queue.PushBack(start)

// 	for queue.Len() > 0 {
// 		node := queue.Remove(queue.Front()).(int)
// 		for _, neighbor := range graph[node] {
// 			if !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue.PushBack(neighbor)
// 			}
// 		}
// 	}
// }

// func dfs(graph map[int][]int, node int, visited map[int]bool) {
// 	visited[node] = true
// 	for _, neighbor := range graph[node] {
// 		if !visited[neighbor] {
// 			dfs(graph, neighbor, visited)
// 		}
// 	}
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

// func kmpSearch(text string, pattern string) int {
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

// func main() {
// 	a := new(int)
// 	b := new(bool)

// 	fmt.Printf("%T\n", a)
// 	fmt.Printf("%T\n", b)

// 	fmt.Println(*a)
// 	fmt.Println(*b)
// }

// package main

// import "fmt"

// type myInt int

// type myFn func(int, int) myInt

// type myFloat = float64

// func main() {
// 	var a myInt = 10
// 	var b myFloat = 10.0
// 	fmt.Printf("value = %v, type = %T\n", a, a)
// 	fmt.Printf("value = %v, type = %T\n", b, b)
// }

// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// 	sex  string
// }

// func main() {
// 	var person Person
// 	person.name = "Hachimi"
// 	person.sex = "男"
// 	person.age = 20
// 	fmt.Printf("%v %T %#v\n", person, person, person)
// }

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// 	Sex  string
// }

// func main() {
// 	// var p2 = new(Person)
// 	// p2.Name = "Hachimi"
// 	// p2.Age = 20
// 	// p2.Sex = "男"
// 	// fmt.Printf("%#v\n%T\n", p2, p2)
// 	p3 := Person{
// 		Name: "Hachimi",
// 		Age:  20,
// 		Sex:  "男",
// 	}
// 	fmt.Printf("%#v\n%T\n", p3, p3)
// }

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// 	Sex  string
// }

// func main() {
// 	var p6 = &Person{
// 		"王麻子",
// 		20,
// 		"男",
// 	}
// 	fmt.Printf("%#v\n", p6)
// }

// package main

// import "fmt"

// type newInt int  // 类型定义
// type myInt = int // 类型别名

// func main() {
// 	var a newInt
// 	var b myInt
// 	fmt.Printf("Type of a: %T\n", a)
// 	fmt.Printf("Type of b: %T\n", b)
// }

// package main

// type person struct {
// 	name, city string
// 	age        int8
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	var p1 person
// 	p1.name = "哈基米"
// 	p1.city = "北京"
// 	p1.age = 18
// 	fmt.Printf("p1 = %v\n", p1)
// 	fmt.Printf("p1 = %#v\n", p1)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	var p2 = new(person)
// 	p2.name = "曼波"
// 	p2.age = 18
// 	p2.city = "上海"
// 	fmt.Printf("%T\n", p2)
// 	fmt.Printf("p2 = %#v\n", p2)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	p3 := &person{}
// 	fmt.Printf("%T\n", p3)
// 	fmt.Printf("p3 = %#v\n", p3)

// 	p3.name = "哈基米"
// 	p3.age = 30
// 	p3.city = "杭州"
// 	(*p3).age = 40
// 	fmt.Printf("p3 = %#v\n", p3)
// }

// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

// func main() {
// 	p4 := person{
// 		name: "Mamba",
// 		city: "伊犁",
// 		age:  18,
// 	}

// 	fmt.Printf("p4 = %#v\n", p4)
// }

// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

// func main() {
// 	p5 := &person{
// 		name: "Kobe",
// 		city: "乌鲁木齐",
// 		age:  20,
// 	}
// 	fmt.Printf("p5 = %#v\n", p5)
// }

// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

// func main() {
// 	p6 := &person{
// 		city: "北京",
// 	}
// 	fmt.Printf("p6 = %#v\n", p6)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	p7 := &person{
// 		"LBJ",
// 		"New York",
// 		28,
// 	}
// 	fmt.Printf("p7 = %#v\n", p7)
// }

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// 	Sex  string
// }

// func main() {
// 	var p1 = Person{
// 		Name: "哈基米",
// 		Age:  20,
// 		Sex:  "男",
// 	}

// 	p2 := p1
// 	p2.Name = "曼波"
// 	fmt.Printf("%#v\n", p1)
// 	fmt.Printf("%#v\n", p2)
// }

package main

import "fmt"

type Person struct {
	Name, Sex   string
	Age, Height int
}

func (p Person) PrintInfo() {
	fmt.Printf("Name: %v\tAge: %v\n", p.Name, p.Age)
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {
	p1 := Person{
		"哈基米",
		"男",
		20,
		180,
	}
	p1.PrintInfo()
	p1.SetInfo("波奇", 10)
	p1.PrintInfo()

	p2 := Person{
		"曼波",
		"女",
		18,
		170,
	}
	p2.PrintInfo()
	p2.SetInfo("喜多", 15)
	p2.PrintInfo()
}
