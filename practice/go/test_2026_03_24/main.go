// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // 处理器函数（HandlerFunc）
// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello 哈基米！")
// }

// func main() {
// 	// 注册路由
// 	http.HandleFunc("/", sayHello)

//		// 启动服务，监听端口
//		err := http.ListenAndServe(":9090", nil)
//		if err != nil {
//			fmt.Println("启动服务失败:", err)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello 哈基米!")
// }

//	func main() {
//		http.HandleFunc("/", sayHello)
//		err := http.ListenAndServe(":9090", nil)
//		if err != nil {
//			fmt.Println("启动服务失败:", err)
//		}
//	}
// package main

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()

// 	query := r.URL.Query()
// 	name := query.Get("name")

// 	r.ParseForm()
// 	age := r.PostForm.Get("age")

// 	body, _ := io.ReadAll(r.Body)
// 	var jsonData map[string]any
// 	json.Unmarshal(body, &jsonData)

//		w.Header().Set("Content-Type", "application/json")
//		resp, _ := json.Marshal(map[string]any{
//			"status": "ok",
//			"name":   name,
//			"age":    age,
//			"json":   jsonData,
//		})
//		w.Write(resp)
//	}
// package main

// import (
// 	"fmt"
// 	"math"
// )

//	func safeAdd(a, b int) (int, error) {
//		if a > math.MaxInt-b {
//			return 0, fmt.Errorf("整型加法溢出：%d + %d", a, b)
//		}
//		return a + b, nil
//	}
// package main

// import "fmt"

//	func main() {
//		var a int8 = 59
//		fmt.Printf("%b\n", a) // 输出二进制：111011
//		fmt.Printf("%d\n", a) // 输出十进制：59
//		fmt.Printf("%o\n", a) // 输出八进制：73
//		fmt.Printf("%O\n", a) // 输出八进制（带 0o 前缀）：0o73
//		fmt.Printf("%x\n", a) // 输出十六进制（小写）：3b
//		fmt.Printf("%X\n", a) // 输出十六进制（大写）：3B
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

// func main() {
// 	var a int = 10
// 	// 使用显式类型转换
// 	fmt.Printf("%f\n", float64(a))

// 	// 将 a 通过 Itoa 转换为 string，
// 	// 然后通过 ParseFloat 将字符串转换为 float64
// 	var b float64
// 	b, _ = strconv.ParseFloat(strconv.Itoa(a), 64)
// 	fmt.Printf("%f\n", b)

//		// 先通过 Sprintf 转换为 string，
//		// 然后通过 ParseFloat 将字符串转换为 float64
//		formattedString := fmt.Sprintf("%.1f", float64(a))
//		b, _ = strconv.ParseFloat(formattedString, 64)
//		fmt.Println("b TypeOf:", reflect.TypeOf(b))
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

// func main() {
// 	var a int = 10
// 	// FormatInt 的包装，默认转换后显示的格式是 10 进制
// 	str := strconv.Itoa(a)
// 	fmt.Printf("str: %s, type of: %s\n", str, reflect.TypeOf(a))

// 	// 通过 Sprintf 将整型转换为字符串
// 	str = fmt.Sprintf("%d", a)
// 	fmt.Printf("str: %s, type of: %s\n", str, reflect.TypeOf(a))

//		// 指定转换后的字符串显示的进制是 10 进制
//		str = strconv.FormatInt(int64(a), 10)
//		fmt.Printf("str: %s, type of: %s\n", str, reflect.TypeOf(a))
//	}
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var f1 float32 = 16777216.0
// 	var f2 float32 = 16777217.0
// 	// 直接比较可能因为精度限制而产生反直觉的结果
// 	fmt.Println(f1 == f2) // true，因为超出 float32 精度范围

//		// 推荐比较方式
//		a := 0.1
//		b := 0.2
//		c := 0.3
//		if math.Abs((a+b)-c) < 1e-9 {
//			fmt.Println("a + b 等于 c")
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

//	func main() {
//		var a float64 = 10.0
//		// 参数 1：必须是 float64 类型的待转换参数
//		// 参数 2：输出格式（支持 f、e、E、g、G 等）
//		// 参数 3：填 -1 表示自动精简，去掉无意义的末尾 0 和多余小数点
//		// 参数 4：指定浮点数的位宽（32=float32 / 64=float64）
//		// 因为第一个参数是 float64，所以这里填写 64 来匹配类型
//		strNum := strconv.FormatFloat(a, 'f', -1, 64)
//		fmt.Printf("strNum: %s, type of: %s\n", strNum, reflect.TypeOf(strNum))
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

//	func main() {
//		var a float64 = 10.0
//		strNum := strconv.FormatFloat(a, 'f', -1, 64)
//		fmt.Printf("strNum: %s, type of: %s\n", strNum, reflect.TypeOf(strNum))
//	}
// package main

// import (
// 	"fmt"
// 	"math"
// 	"reflect"
// )

// func main() {
// 	var f float64 = 3.14
// 	// 显式类型转换，直接截断小数部分
// 	b := int(f)
// 	fmt.Printf("str Num: %d, type of: %s\n", b, reflect.TypeOf(b))

//		// 使用 Round 四舍五入
//		i := int(math.Round(f))
//		fmt.Printf("i: %d, type of: %s\n", i, reflect.TypeOf(i))
//	}
// package main

// import "fmt"

//	func main() {
//		var c complex64 = 5 + 6i
//		var d = 0o123 + .12345e+5i
//		fmt.Println("c:", c)
//		fmt.Println("d:", d)
//	}
// package main

// import "fmt"

//	func main() {
//		var c = complex(5, 6)
//		var d = complex(0o123, .12345e+5)
//		fmt.Println("c:", c)
//		fmt.Println("d:", d)
//	}
// package main

// import "fmt"

//	func main() {
//		var c = complex(5, 6)
//		r := real(c)
//		i := imag(c)
//		fmt.Println("c:", c)
//		fmt.Println("r:", r)
//		fmt.Println("i:", i)
//	}
// package main

// import "fmt"

//	func main() {
//		for i := 1; i <= 9; i++ {
//			for j := 1; j <= i; j++ {
//				fmt.Printf("%d * %d = %d\t", j, i, j*i)
//			}
//			fmt.Println()
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

//	func main() {
//		str := "this is a string"
//		len1 := utf8.RuneCountInString(str)
//		fmt.Println("字符串的长度：", len1)
//		fmt.Println(len(str))
//		for i := 0; i < len1; i++ {
//			fmt.Printf("%s ", string(str[i]))
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		books := map[string]map[string]int{
//			"四书": {"论语": 80, "大学": 66, "中庸": 60, "孟子": 70},
//			"五经": {"周易": 90, "诗书": 80, "礼记": 88, "尚书": 78, "春秋": 99},
//			"书法": {"兰亭集序": 66, "九成宫碑": 68, "多宝塔": 56},
//		}
//		for key, value := range books {
//			slice := []string{}
//			for v := range value {
//				slice = append(slice, v)
//			}
//			fmt.Printf("%s: %s\n", key, strings.Join(slice, ", "))
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	str := "hello"
// 	fmt.Println("Before update:")
// 	fmt.Println("str =", str)
// 	updateString(str)
// 	fmt.Println("After update:")
// 	fmt.Println("str =", str)

// 	slice := []string{"hello", "world"}
// 	fmt.Println("Before update:")
// 	fmt.Println("slice =", slice)
// 	updateSlice(slice)
// 	fmt.Println("After update:")
// 	fmt.Println("slice =", slice)

// 	mp := map[string]string{"key": "value"}
// 	fmt.Println("Before update:")
// 	fmt.Println("mp =", mp)
// 	updateMap(mp)
// 	fmt.Println("After update:")
// 	fmt.Println("mp =", mp)
// }

// func updateString(s string) {
// 	s = "updated"
// }

// func updateSlice(s []string) {
// 	s[0] = "updated"
// }

//	func updateMap(m map[string]string) {
//		m["key"] = "updated"
//	}
// package main

// import "fmt"

// func main() {
// 	res := fn(add, 10, 5)
// 	fmt.Println("Addition Result:", res)

// 	res = fn(sub, 10, 5)
// 	fmt.Println("Subtraction Result:", res)
// }

// func fn(operation func(int, int) int, a, b int) int {
// 	return operation(a, b)
// }

// func add(a, b int) int {
// 	return a + b
// }

//	func sub(a, b int) int {
//		return a - b
//	}
// package main

// import "fmt"

// func main() {
// 	addOperation := getOperation("add")
// 	subOperation := getOperation("sub")

// 	res := addOperation(10, 5)
// 	fmt.Println(res)

// 	res = subOperation(10, 5)
// 	fmt.Println(res)
// }

// func getOperation(operationType string) func(int, int) int {
// 	switch operationType {
// 	case "add":
// 		return add
// 	case "sub":
// 		return sub
// 	default:
// 		return nil
// 	}
// }

// func add(a, b int) int {
// 	return a + b
// }

//	func sub(a, b int) int {
//		return a - b
//	}
// package main

// import "fmt"

// type MyInterface interface {
// 	MyMethod()
// }

// type MyType struct{}

// func (t MyType) MyMethod() {
// 	fmt.Println("MyMethod called")
// }

// type EmptyInterface interface{}

// func main() {
// 	var empty EmptyInterface
// 	fmt.Printf("Empty interface's method set is empty: %v\n", empty)

//		var t MyType
//		var i MyInterface = t
//		fmt.Printf("Type t implements interface I: %v\n", i)
//	}
// package main

// import "fmt"

// type Container[T any] interface {
// 	Add(elem T)
// 	Get() T
// }

// type MyContainer[T any] struct {
// 	value T
// }

// func (c *MyContainer[T]) Add(elem T) {
// 	c.value = elem
// }

// func (c *MyContainer[T]) Get() T {
// 	return c.value
// }

//	func main() {
//		var c Container[string] = &MyContainer[string]{}
//		c.Add("Hello")
//		fmt.Println(c.Get())
//	}
// package main

// import "fmt"

// type Container[T any] interface {
// 	Add(elem T)
// 	Get() T
// }

// type MyContainer[T any] struct {
// 	value T
// }

// func (c *MyContainer[T]) Add(elem T) {
// 	c.value = elem
// }

// func (c *MyContainer[T]) Get() T {
// 	return c.value
// }

//	func main() {
//		var c Container[string] = &MyContainer[string]{}
//		c.Add("Hello")
//		fmt.Println(c.Get())
//	}
// package main

// import "fmt"

// type Animal interface {
// 	Speak() string
// }

// type Cat struct{}

// func (c Cat) Speak() string {
// 	return "Meow"
// }

// type Dog struct{}

// func (d Dog) Speak() string {
// 	return "Woof"
// }

// func main() {
// 	var animal Animal

// 	cat := Cat{}
// 	dog := Dog{}

// 	animal = cat
// 	fmt.Println(animal.Speak())

//		animal = dog
//		fmt.Println(animal.Speak())
//	}
// package main

// type Writer interface {
// 	Write(data []byte) error
// }

// type Reader interface {
// 	Read() ([]byte, error)
// }

// type Logger interface {
// 	Log(msg string)
// }

// type FileWriter struct{}

// func (fw FileWriter) Write(data []byte) error {
// 	return nil
// }

// type ConsoleReader struct{}

// func (cr ConsoleReader) Read() ([]byte, error) {
// 	return nil, nil
// }

// type FileLogger struct{}

// func (fl FileLogger) Log(msg string) {}

// type FileProcessor struct {
// 	Writer
// 	Reader
// 	Logger
// }

// func main() {
// 	writer := FileWriter{}
// 	reader := ConsoleReader{}
// 	logger := FileLogger{}

// 	processor := FileProcessor{
// 		Writer: writer,
// 		Reader: reader,
// 		Logger: logger,
// 	}

//		processor.Write([]byte("Hello, world!"))
//		processor.Read()
//		processor.Log("This is a log message.")
//	}
// package main

// import "fmt"

// type Writer interface {
// 	Write(data []byte) error
// }

// type Reader interface {
// 	Read() ([]byte, error)
// }

// type Processor struct {
// 	Writer
// 	Reader
// }

// func ProcessData(w Writer, r Reader) {
// 	w.Write([]byte(""))
// 	data, _ := r.Read()
// 	fmt.Println(string(data))
// }

// type FileWriter struct{}

// func (fw FileWriter) Write(data []byte) error {
// 	return nil
// }

// type ConsoleReader struct{}

// func (cr ConsoleReader) Read() ([]byte, error) {
// 	return []byte("Data from reader"), nil
// }

// func main() {
// 	writer := FileWriter{}
// 	reader := ConsoleReader{}

// 	processor := Processor{
// 		Writer: writer,
// 		Reader: reader,
// 	}

//		ProcessData(processor.Writer, processor.Reader)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name   string  `json:"name"`
// 	Age    int     `json:"age,string"`
// 	Height float64 `json:"height"`
// 	Email  string  `json:"email,omitempty"`
// }

// func main() {
// 	p := Person{
// 		Name:   "Alice",
// 		Age:    25,
// 		Height: 170.5,
// 		Email:  "",
// 	}

// 	jsonData, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Println("JSON serialization error:", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))

//		var p2 Person
//		err = json.Unmarshal(jsonData, &p2)
//		if err != nil {
//			fmt.Println("JSON deserialization error:", err)
//			return
//		}
//		fmt.Println(p2)
//	}
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name   string  `json:"name"`
// 	Age    int     `json:"age,string"`
// 	Height float64 `json:"height"`
// 	Email  string  `json:"email,omitempty"`
// }

//	func main() {
//		p := Person{
//			Name:   "Alice",
//			Age:    25,
//			Height: 170.5,
//			Email:  "",
//		}
//		jsonData, err := json.Marshal(p)
//		if err != nil {
//			fmt.Println("JSON serialization error:", err)
//			return
//		}
//		fmt.Println(string(jsonData))
//		var p2 Person
//		err = json.Unmarshal(jsonData, &p2)
//		if err != nil {
//			fmt.Println("JSON deserialization error:", err)
//			return
//		}
//		fmt.Println(p2)
//	}
// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type Person struct {
// 	Name   string `xml:"name"`
// 	Age    int    `xml:"age,omitempty"`
// 	Height int    `xml:"height,attr"`
// }

//	func main() {
//		p := &Person{
//			Name:   "Alice",
//			Age:    0,
//			Height: 185,
//		}
//		xmlData, err := xml.Marshal(p)
//		if err != nil {
//			fmt.Println("XML serialization error:", err)
//			return
//		}
//		fmt.Println(string(xmlData))
//		var p2 Person
//		err = xml.Unmarshal(xmlData, &p2)
//		if err != nil {
//			fmt.Println("XML deserialization error:", err)
//			return
//		}
//		fmt.Println(p2)
//	}
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancelFunc := context.WithCancel(context.Background())
// 	go Working(ctx)

// 	time.Sleep(3 * time.Second)
// 	fmt.Println("主程序发送取消信号……")
// 	cancelFunc()

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("程序结束")
// }

//	func Working(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done():
//				fmt.Println("工作协程：接收到取消信号")
//				return
//			default:
//				fmt.Println("工作协程：用户正在工作")
//				time.Sleep(500 * time.Millisecond)
//			}
//		}
//	}
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func Working(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("工作协程：接收到取消信号")
// 			return
// 		default:
// 			fmt.Println("工作协程：用户正在工作中")
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancelFunc := context.WithCancel(context.Background())
// 	go Working(ctx)

// 	time.Sleep(3 * time.Second)
// 	fmt.Println("主程序发送取消信号")
// 	cancelFunc()

//		time.Sleep(time.Second)
//		fmt.Println("程序结束")
//	}
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	done := make(chan bool, 1)
// 	go func() {
// 		fmt.Println("开始执行耗时操作")
// 		time.Sleep(5 * time.Second)
// 		fmt.Println("耗时操作执行完毕")
// 		done <- true
// 	}()

//		select {
//		case <-done:
//			fmt.Println("操作完成")
//		case <-ctx.Done():
//			if ctx.Err() == context.DeadlineExceeded {
//				fmt.Println("操作已超时（3 秒限制）")
//			} else {
//				fmt.Printf("操作被取消：%v\n", ctx.Err())
//			}
//			return
//		}
//	}
// package main

// import "fmt"

//	func main() {
//		ch1 := make(chan int)
//		go func() {
//			for i := range 100 {
//				ch1 <- i
//			}
//			close(ch1)
//		}()
//		for {
//			select {
//			case i, ok := <-ch1:
//				if !ok {
//					return
//				}
//				fmt.Println(i)
//			default:
//				fmt.Println("default")
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ch <- 1
// 	}()

//		select {
//		case data, ok := <-ch:
//			if ok {
//				fmt.Println("接收到数据：", data)
//			} else {
//				fmt.Println("通道已关闭")
//			}
//		case <-time.After(2 * time.Second):
//			fmt.Println("超时了！")
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var x int64
// var wg sync.WaitGroup
// var lock sync.Mutex

// func add() {
// 	for range 5000 {
// 		lock.Lock()
// 		x++
// 		lock.Unlock()
// 	}
// 	wg.Done()
// }

//	func main() {
//		wg.Add(2)
//		go add()
//		go add()
//		wg.Wait()
//		fmt.Println(x)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	x      int64
// 	wg     sync.WaitGroup
// 	lock   sync.Mutex
// 	rwlock sync.RWMutex
// )

// func write() {
// 	rwlock.Lock()
// 	x++
// 	time.Sleep(10 * time.Millisecond)
// 	rwlock.Unlock()
// 	wg.Done()
// }

// func read() {
// 	rwlock.RLock()
// 	time.Sleep(time.Millisecond)
// 	rwlock.RUnlock()
// 	wg.Done()
// }

//	func main() {
//		start := time.Now()
//		for range 10 {
//			wg.Add(1)
//			go write()
//		}
//		for range 1000 {
//			wg.Add(1)
//			go read()
//		}
//		wg.Wait()
//		end := time.Now()
//		fmt.Println(end.Sub(start))
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

//	func main() {
//		var x float64 = 3.4
//		v := reflect.ValueOf(&x)
//		if v.Elem().CanSet() {
//			v.Elem().SetFloat(7.1)
//		}
//		fmt.Println(x)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

//	func main() {
//		var i any = "hello"
//		v := reflect.ValueOf(i)
//		if s, ok := v.Interface().(string); ok {
//			fmt.Println(s)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type MyStruct struct {
// 	privateField int
// }

//	func main() {
//		s := MyStruct{privateField: 1}
//		v := reflect.ValueOf(&s)
//		structVal := v.Elem()
//		field := structVal.FieldByName("privateField")
//		fmt.Println("Private field:", field.Int())
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type MyStruct struct {
// 	privateField int
// }

//	func main() {
//		s := MyStruct{privateField: 1}
//		v := reflect.ValueOf(&s)
//		structVal := v.Elem()
//		field := structVal.FieldByName("privateField")
//		fmt.Println("Private field:", field.Int())
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type MyMethods struct{}

// func (m *MyMethods) MyMethod() string {
// 	return "Hello, world!"
// }

//	func main() {
//		obj := &MyMethods{}
//		method := reflect.ValueOf(obj).MethodByName("MyMethod")
//		res := method.Call(nil)
//		fmt.Println("Method Result:", res[0].String())
//	}
// package main

// import "fmt"

//	func main() {
//		a := []int{1, 2, 3, 4, 5}
//		fmt.Println(len(a), cap(a), a)
//		b := a[1:3:4]
//		fmt.Println(len(b), cap(b), b)
//	}
// package main

// import "fmt"

//	func main() {
//		nums := []int{1, 2, 3, 4, 5}
//		ch := make(chan int)
//		for _, num := range nums {
//			go func(n int) {
//				ch <- n
//			}(num)
//		}
//		sum := 0
//		for range len(nums) {
//			sum += <-ch
//		}
//		fmt.Println("Sum:", sum)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

//	func main() {
//		var mu sync.Mutex
//		var count int
//		increment := func() {
//			mu.Lock()
//			defer mu.Unlock()
//			count++
//			fmt.Println("Count:", count)
//		}
//		for range 5 {
//			go increment()
//		}
//		time.Sleep(time.Second)
//	}
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutexA, mutexB sync.Mutex

	go func() {
		mutexA.Lock()
		fmt.Println("Goroutine 1: Locked mutexA")
		mutexB.Lock()
		fmt.Println("Goroutine 1: Locked mutexB")
		mutexB.Unlock()
		mutexA.Unlock()
	}()

	go func() {
		mutexB.Lock()
		fmt.Println("Goroutine 2: Locked mutexB")
		mutexA.Lock()
		fmt.Println("Goroutine 2: Locked mutexA")
		mutexA.Unlock()
		mutexB.Unlock()
	}()

	select {}
}
