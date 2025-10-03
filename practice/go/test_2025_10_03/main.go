// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name  string `json:"name" form:"username"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// func (s Student) GetInfo() string {
// 	return fmt.Sprintf("Name: %v\tAge: %v\tScore: %v",
// 		s.Name, s.Age, s.Score)
// }

// func (s *Student) SetInfo(name string, age, score int) {
// 	s.Name = name
// 	s.Age = age
// 	s.Score = score
// }

// func (s Student) Print() {
// 	fmt.Println("打印方法")
// }

// func PrintStructField(s any) {
// 	// 判断参数是否为结构体类型
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)
// 	if (t.Kind() != reflect.Struct) &&
// 		(t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的参数不是一个结构体")
// 		return
// 	}
// 	field0 := t.Elem().Field(0)
// 	fmt.Printf("%v\n", field0)
// 	fmt.Println(field0.Name)
// 	fmt.Println(field0.Type)
// 	fmt.Println(field0.Tag)
// 	fmt.Println(field0.Index)
// 	fmt.Println(field0.Offset)

// 	fmt.Println("------------------------------------")

// 	field1, ok := t.Elem().FieldByName("Age")
// 	if ok {
// 		fmt.Printf("%v\n", field1)
// 		fmt.Println(field1.Name)
// 		fmt.Println(field1.Type)
// 		fmt.Println(field1.Tag)
// 		fmt.Println(field1.Index)
// 		fmt.Println(field1.Offset)
// 	}

// 	fmt.Println("------------------------------------")

// 	fmt.Println(t.Elem().NumField())

// 	fmt.Println("------------------------------------")

// 	fmt.Println(v.Elem().FieldByName("Name"))
// 	fmt.Println(v.Elem().FieldByName("Age"))
// 	fmt.Println(v.Elem().NumField())

// 	fmt.Println("------------------------------------")

// 	for i := range t.Elem().NumField() {
// 		fmt.Printf("属性名称：%v 属性值：%v 属性类型：%v 属性 Tag：%v\n",
// 			t.Elem().Field(i).Name, v.Elem().Field(i),
// 			t.Elem().Field(i).Type, t.Elem().Field(i).Tag.Get("json"))
// 	}
// }

// func PrintStructFn(s any) {
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if (t.Kind() != reflect.Struct) &&
// 		(t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的参数不是一个结构体")
// 	}

// 	// 按照字母顺序来排列。
// 	method0 := t.Method(0)
// 	fmt.Println(method0.Name)
// 	fmt.Println(method0.Type)

// 	method1, ok := t.MethodByName("GetInfo")

// 	if ok {
// 		fmt.Println(method1.Name)
// 		fmt.Println(method1.Type)
// 	}
// 	fmt.Println("------------------------------------")

// 	v.Method(1).Call(nil)
// 	fmt.Println("------------------------------------")

// 	v.MethodByName("Print").Call(nil)
// 	fmt.Println("------------------------------------")

// 	fmt.Println(v.MethodByName("GetInfo").Call(nil))
// 	fmt.Println("------------------------------------")

// 	params := []reflect.Value{
// 		reflect.ValueOf("平泽唯"),
// 		reflect.ValueOf(23),
// 		reflect.ValueOf(99),
// 	}
// 	v.MethodByName("SetInfo").Call(params)
// 	fmt.Println(v.MethodByName("GetInfo").Call(nil))
// 	fmt.Println("------------------------------------")
// 	fmt.Println(t.NumMethod())
// }

// func reflectChangeStruct(s any) {
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if (t.Kind() != reflect.Pointer) ||
// 		(t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体指针类型")
// 		return
// 	}

// 	name := v.Elem().FieldByName("Name")
// 	name.SetString("后藤一里")
// }

//	func main() {
//		stu := Student{
//			"平泽唯",
//			16,
//			100,
//		}
//		// PrintStructField(&stu)
//		// PrintStructFn(&stu)
//		reflectChangeStruct(&stu)
//		fmt.Println(stu)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // Student 结构体
// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// // 获取 Student 结构体信息
// func (s Student) GetInfo() string {
// 	return fmt.Sprintf(
// 		"Name: %v\tAge: %v\tScore: %v",
// 		s.Name, s.Age, s.Score,
// 	)
// }

// // 设置 Student 结构体信息
// func (s *Student) SetInfo(name string, age, score int) {
// 	s.Name = name
// 	s.Age = age
// 	s.Score = score
// }

// func (s *Student) Print() {
// 	fmt.Println("打印方法")
// }

// // 字段
// func PrintStructField(s any) {
// 	fmt.Println("----结构体字段----")
// 	t := reflect.TypeOf(s)

// 	if (t.Kind() != reflect.Struct) &&
// 		(t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体")
// 		return
// 	}

// 	field0 := t.Field(0)
// 	fmt.Println(field0.Name)
// 	fmt.Println(field0.Type)
// 	fmt.Println(field0.Tag.Get("json"))
// 	fmt.Println("----")

// 	field1, ok := t.FieldByName("Age")
// 	if ok {
// 		fmt.Println(field1.Name)
// 		fmt.Println(field1.Type)
// 		fmt.Println(field1.Tag.Get("json"))
// 		fmt.Println("----")
// 	}

// 	fmt.Println("字段数量：", t.NumField())
// }

// // 方法
// func PrintStructFunc(s any) {
// 	fmt.Println("----结构体方法----")
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if (t.Kind() != reflect.Struct) &&
// 		(t.Elem().Kind() != reflect.Struct) {
// 		fmt.Println("传入的不是结构体")
// 		return
// 	}

// 	method0 := t.Method(0)
// 	fmt.Println(method0.Name)
// 	fmt.Println(method0.Type)
// 	fmt.Println("----")

// 	fmt.Println(t.NumMethod())
// 	fmt.Println("----")

// 	v.MethodByName("Print").Call(nil)
// 	fmt.Println("----")

// 	v.MethodByName("SetInfo").Call([]reflect.Value{
// 		reflect.ValueOf("后藤独"),
// 		reflect.ValueOf(16),
// 		reflect.ValueOf(60),
// 	})
// 	fmt.Println(v.MethodByName("GetInfo").Call(nil))
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		PrintStructField(stu)
//		PrintStructFunc(&stu)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // Student 结构体
// type Student struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

// // 获取 Student 结构体信息
// func (s Student) GetInfo() string {
// 	return fmt.Sprintf(
// 		"Name: %v\tAge: %v\tScore: %v",
// 		s.Name, s.Age, s.Score,
// 	)
// }

// // 反射修改结构体属性
// func reflectChangeStruct(s any) {
// 	t := reflect.TypeOf(s)
// 	v := reflect.ValueOf(s)

// 	if t.Elem().Kind() != reflect.Struct {
// 		fmt.Println("传入的不是结构体指针类型")
// 		return
// 	}

// 	name := v.Elem().FieldByName("Name")
// 	name.SetString("后藤一里")

// 	age := v.Elem().FieldByName("Age")
// 	age.SetInt(17)

// 	score := v.Elem().FieldByName("Score")
// 	score.SetInt(60)
// }

//	func main() {
//		stu := Student{
//			Name:  "平泽唯",
//			Age:   16,
//			Score: 100,
//		}
//		fmt.Println(stu.GetInfo())
//		reflectChangeStruct(&stu)
//		fmt.Println(stu.GetInfo())
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap *sync.Map, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	numMap.Store(num, res)
// }

// func main() {
// 	var numMap sync.Map
// 	var wg sync.WaitGroup

// 	wg.Add(10)
// 	for i := 1; i <= 10; i++ {
// 		go factorial(i, &numMap, &wg)
// 	}
// 	wg.Wait()

//		for k, v := range numMap.Range {
//			fmt.Printf("map[%d] = %d\n", k, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap *sync.Map, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	numMap.Store(num, res)
// }

// func main() {
// 	var numMap sync.Map
// 	var wg sync.WaitGroup

// 	wg.Add(10)
// 	for i := 1; i <= 10; i++ {
// 		go factorial(i, &numMap, &wg)
// 	}
// 	wg.Wait()

//		for k, v := range numMap.Range {
//			fmt.Printf("(%v, %v)\n", k, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		wordMap := make(map[string]int)
//		str := "how  do you do"
//		for word := range strings.SplitSeq(str, " ") {
//			wordMap[word]++
//		}
//		fmt.Println(wordMap)
//	}
// package main

// import "fmt"

// func toSlice[T any](args ...T) []T {
// 	return args
// }

//	func main() {
//		strings := toSlice("hello", "world")
//		nums := toSlice(1, 2, 3)
//		fmt.Println(strings)
//		fmt.Println(nums)
//	}
// package main

// import "fmt"

// type Stack[T any] struct {
// 	data []T
// }

// func (s *Stack[T]) Push(x T) {
// 	s.data = append(s.data, x)
// }

// func (s *Stack[T]) Pop() T {
// 	n := len(s.data)
// 	x := s.data[n-1]
// 	s.data = s.data[:n-1]
// 	return x
// }

// func main() {
// 	var s Stack[int]
// 	s.Push(10)
// 	s.Push(20)
// 	s.Push(30)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	s.Push(40)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())

//		var strStack Stack[string]
//		strStack.Push("hello")
//		strStack.Push("world")
//		fmt.Println(strStack.Pop())
//		fmt.Println(strStack.Pop())
//	}
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
//		floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
//		stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}
//		sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
//		sort.Sort(sort.Reverse(sort.Float64Slice(floatSlice)))
//		sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//		sort.Ints(intSlice)
//		sort.Float64s(floatSlice)
//		sort.Strings(stringSlice)
//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		wordMap := make(map[string]int)
//		str := "how do  you do"
//		for word := range strings.FieldsSeq(str) {
//			wordMap[word]++
//		}
//		fmt.Println(wordMap)
//	}
// package main

// import "fmt"

// func toSlice[T any](args ...T) []T {
// 	return args
// }

//	func main() {
//		strings := toSlice("hello", "world")
//		nums := toSlice(1, 1.1, 321.3)
//		fmt.Println(strings)
//		fmt.Println(nums)
//	}
// package main

// import "fmt"

// type Stack[T any] struct {
// 	data []T
// }

// func (s *Stack[T]) Push(x T) {
// 	s.data = append(s.data, x)
// }

// func (s *Stack[T]) Pop() T {
// 	n := len(s.data)
// 	x := s.data[n-1]
// 	s.data = s.data[:n-1]
// 	return x
// }

// func main() {
// 	var s Stack[int]
// 	s.Push(10)
// 	s.Push(20)
// 	s.Push(30)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	s.Push(40)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())

//		var strStack Stack[string]
//		strStack.Push("hello")
//		strStack.Push("world")
//		fmt.Println(strStack.Pop())
//		fmt.Println(strStack.Pop())
//	}
// package main

// import "fmt"

// type Stack[T any] struct {
// 	data []T
// }

// func (s *Stack[T]) Push(x T) {
// 	s.data = append(s.data, x)
// }

// func (s *Stack[T]) Pop() T {
// 	n := len(s.data)
// 	x := s.data[n-1]
// 	s.data = s.data[:n-1]
// 	return x
// }

// func main() {
// 	s := Stack[int]{}
// 	s.Push(10)
// 	s.Push(20)
// 	s.Push(30)
// 	fmt.Println(s.Pop())
// 	fmt.Println(s.Pop())
// 	s.Push(40)

//		strStack := Stack[string]{}
//		strStack.Push("hello")
//		strStack.Push("world")
//		fmt.Println(strStack.Pop())
//	}
// package main

// import "fmt"

// type MyType[T fmt.Stringer] struct {
// 	data T
// }

//	func (m *MyType[T]) String() string {
//		return m.data.String()
//	}
// package main

// import (
// 	"cmp"
// 	"fmt"
// 	"slices"
// )

// func Sort[T cmp.Ordered](s []T) {
// 	slices.Sort(s)
// }

//	func main() {
//		numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
//		Sort(numbers)
//		fmt.Println(numbers)
//	}
// package main

// import (
// 	"slices"
// 	"sort"
// )

//	func main() {
//		nums := []int{5, 2, 8, 1, 9, 3, 7, 4, 6, 0}
//		sort.Ints(nums)
//		slices.Sort(nums)
//	}
// package main

// import (
// 	"cmp"
// 	"fmt"
// 	"slices"
// )

// func main() {
// 	intSlice := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
// 	floatSlice := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
// 	stringSlice := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

// 	slices.SortFunc(intSlice, func(a, b int) int { return cmp.Compare(b, a) })
// 	slices.SortFunc(floatSlice, func(a, b float64) int { return cmp.Compare(b, a) })
// 	slices.SortFunc(stringSlice, func(a, b string) int { return cmp.Compare(b, a) })

//		fmt.Println(intSlice)
//		fmt.Println(floatSlice)
//		fmt.Println(stringSlice)
//	}
// package main

// import (
// 	"cmp"
// 	"fmt"
// 	"slices"
// )

//	func main() {
//		nums := []int{6, 5, 4, 4, 3, 3, 3, 2, 1}
//		x := 3
//		index, ok := slices.BinarySearchFunc(nums, x, func(a, b int) int { return cmp.Compare(b, a) })
//		if ok {
//			fmt.Println(index)
//		}
//	}

