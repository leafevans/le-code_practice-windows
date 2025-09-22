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

package main

import "fmt"

func main() {
	a := new(int)
	b := new(bool)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Println(*a)
	fmt.Println(*b)
}
