// package main

// import (
// 	"fmt"
// 	"time"
// )

// func fn() {
// 	ticker := time.NewTicker(time.Millisecond * 50)
// 	defer ticker.Stop()
// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		go fn()
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(time.Millisecond * 50)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ticker := time.NewTicker(time.Millisecond * 50)
// 	defer ticker.Stop()
// 	count := 0

// 	for range ticker.C {
// 		if count >= 10 {
// 			return
// 		}
// 		fmt.Printf("fn: 小 Go，你好！\t%v\n", count)
// 		count++
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(1)
//		go fn(&wg)
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(time.Millisecond * 100)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func fn(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello Goroutine!", num)
// }

//	func main() {
//		var wg sync.WaitGroup
//		for i := range 10 {
//			wg.Add(1)
//			go fn(i, &wg)
//		}
//		wg.Wait()
//	}
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	cpuNum := runtime.NumCPU()
// 	fmt.Println("cpuNum =", cpuNum)

//		runtime.GOMAXPROCS(cpuNum)
//		fmt.Println("OK!")
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	primeCount := 0

//		for num := 2; num < 1.2e5; num++ {
//			isPrime := true
//			for i := 2; i*i <= num; i++ {
//				if num%i == 0 {
//					isPrime = false
//					break
//				}
//			}
//			if isPrime {
//				primeCount++
//			}
//		}
//		elapsed := time.Since(start)
//		fmt.Println(primeCount, elapsed)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func countPrimes(start, end int, wg *sync.WaitGroup, result *int) {
// 	defer wg.Done()
// 	count := 0
// 	for num := start; num < end; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			count++
// 		}
// 	}
// 	*result = count
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	parts := 10
// 	interval := 1.2e5 / parts
// 	result := make([]int, parts)

// 	for i := range parts {
// 		wg.Add(1)
// 		s := i * interval
// 		e := (i + 1) * interval
// 		if s < 2 {
// 			s = 2
// 		}
// 		go countPrimes(s, e, &wg, &result[i])
// 	}

// 	wg.Wait()

// 	total := 0
// 	for _, v := range result {
// 		total += v
// 	}

//		fmt.Println(time.Since(start))
//		fmt.Println(total)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func countPrimes(start, end int, wg *sync.WaitGroup, result *int) {
// 	defer wg.Done()
// 	count := 0

// 	for num := start; num < end; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			count++
// 		}
// 	}

// 	*result = count
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	parts := 20
// 	interval := 1.2e6 / parts
// 	results := make([]int, parts)

// 	for i := range parts {
// 		wg.Add(1)
// 		s := i * interval
// 		e := (i + 1) * interval
// 		go countPrimes(s, e, &wg, &results[i])
// 	}

// 	wg.Wait()

// 	total := 0
// 	for _, v := range results {
// 		total += v
// 	}

//		fmt.Println(time.Since(start), total)
//	}
// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 5)

// 	for i := range 5 {
// 		ch <- i + 1
// 	}
// 	close(ch)

//		for v := range ch {
//			fmt.Print(v, " ")
//		}
//		fmt.Println()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func write(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := range 10 {
// 		fmt.Printf("写入数据：%v\n", i+1)
// 		ch <- i + 1
// 	}

// 	close(ch)
// }

// func read(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for v := range ch {
// 		fmt.Printf("读取数据：%v\n", v)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(2)
// 	go write(ch, &wg)
// 	go read(ch, &wg)
// 	wg.Wait()

//		fmt.Println("读取完毕……")
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func generateNumbers(numChan chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := 2; num < 1.2e5; num++ {
// 		numChan <- num
// 	}
// 	close(numChan)
// }

// func filterPrimes(numChan, primeChan chan int,
// 	exitChan chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range numChan {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			primeChan <- num
// 		}
// 	}
// 	exitChan <- true
// }

// func printPrimes(primeChan chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for prime := range primeChan {
// 		fmt.Println(prime, " ")
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workerCount := 5
// 	numChan := make(chan int, 10)
// 	primeChan := make(chan int, 10)
// 	exitChan := make(chan bool, workerCount)

// 	wg.Add(1)
// 	go generateNumbers(numChan, &wg)

// 	for range workerCount {
// 		wg.Add(1)
// 		go filterPrimes(numChan, primeChan, exitChan, &wg)
// 	}

// 	wg.Add(1)
// 	go printPrimes(primeChan, &wg)

// 	wg.Go(func() {
// 		for range workerCount {
// 			<-exitChan
// 		}
// 		close(primeChan)
// 	})

// 	wg.Wait()

//		fmt.Println()
//		fmt.Println(time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func generateNumbers(numChan chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := 2; num < 1.2e5; num++ {
// 		numChan <- num
// 	}
// 	close(numChan)
// }

// func filterPrimes(numChan <-chan int, primeChan chan<- int,
// 	exitChan chan<- bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range numChan {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			primeChan <- num
// 		}
// 	}
// 	exitChan <- true
// }

// func printPrimes(primeChan <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for prime := range primeChan {
// 		fmt.Println(prime)
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workerCount := 20
// 	numChan := make(chan int, workerCount)
// 	primeChan := make(chan int, workerCount)
// 	exitChan := make(chan bool, workerCount)

// 	wg.Add(1)
// 	go generateNumbers(numChan, &wg)

// 	for range workerCount {
// 		wg.Add(1)
// 		go filterPrimes(numChan, primeChan, exitChan, &wg)
// 	}

// 	wg.Add(1)
// 	go printPrimes(primeChan, &wg)

// 	wg.Go(func() {
// 		for range workerCount {
// 			<-exitChan
// 		}
// 		close(primeChan)
// 	})

// 	wg.Wait()

//		fmt.Println(time.Since(start))
//	}
// package main

// import "fmt"

// func main() {
// 	intChan := make(chan int, 10)
// 	for i := range 10 {
// 		intChan <- i
// 	}

// 	stringChan := make(chan string, 5)
// 	for i := range 5 {
// 		stringChan <- "Hello, " + fmt.Sprintf("%d", i)
// 	}

//		for {
//			select {
//			case v := <-intChan:
//				fmt.Printf("从 intChan 读取的数据：%d\n", v)
//			case v := <-stringChan:
//				fmt.Printf("从 stringChan 读取的数据：%s\n", v)
//			default:
//				fmt.Printf("取完所有数据")
//				return
//			}
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHello() {
// 	for range 10 {
// 		time.Sleep(time.Millisecond * 100)
// 		fmt.Println("Hello World!")
// 	}
// }

// func test() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Printf("发生错误：%v\n", err)
// 		}
// 	}()
// 	var myMap map[int]string
// 	myMap[0] = "Golang"
// }

// func main() {
// 	go sayHello()
// 	go test()

//		for i := range 10 {
//			fmt.Println("main() ok =", i)
//			time.Sleep(time.Millisecond * 100)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap map[int]int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	numMap[num] = res
// }

// func main() {
// 	numMap := make(map[int]int)
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 60; i++ {
// 		wg.Add(1)
// 		go factorial(i, numMap, &wg)
// 	}

// 	wg.Wait()

//		for i, v := range numMap {
//			fmt.Printf("map[%d] = %d\n", i, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func factorial(num int, numMap map[int]int,
// 	wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= i
// 	}
// 	mutex.Lock()
// 	numMap[num] = res
// 	mutex.Unlock()
// }

// func main() {
// 	numMap := make(map[int]int)
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex

// 	wg.Add(10)
// 	for i := 1; i <= 10; i++ {
// 		go factorial(i, numMap, &wg, &mutex)
// 	}
// 	wg.Wait()

//		for i, v := range numMap {
//			fmt.Printf("map[%d] = %d\n", i, v)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func write(wg *sync.WaitGroup, mutex *sync.RWMutex) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	fmt.Println("---执行写操作")
// 	time.Sleep(time.Millisecond * 100)
// 	mutex.Unlock()
// }

// func read(wg *sync.WaitGroup, mutex *sync.RWMutex) {
// 	defer wg.Done()
// 	mutex.RLock()
// 	fmt.Println("***执行读操作")
// 	time.Sleep(time.Millisecond * 100)
// 	mutex.RUnlock()
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	var mutex sync.RWMutex

// 	wg.Add(10)
// 	for range 10 {
// 		go write(&wg, &mutex)
// 	}

// 	wg.Add(10)
// 	for range 10 {
// 		go read(&wg, &mutex)
// 	}

// 	wg.Wait()

//		fmt.Println(time.Since(start))
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type myInt int

// type newInt = int

// type Person struct {
// 	name string
// 	age  int
// }

// func reflectFn(x any) {
// 	fmt.Printf("%v\n", reflect.TypeOf(x))
// }

//	func main() {
//		a := 10
//		b := 23.4
//		c := true
//		d := "你好"
//		e := Person{"平泽唯", 17}
//		var f myInt = 1
//		h := 25
//		var g newInt
//		reflectFn(a)
//		reflectFn(b)
//		reflectFn(c)
//		reflectFn(d)
//		reflectFn(e)
//		reflectFn(f)
//		reflectFn(&h)
//		fmt.Println(reflect.TypeOf(f).Name(), reflect.TypeOf(f).Kind())
//		fmt.Println(reflect.TypeOf(g).Name(), reflect.TypeOf(g).Kind())
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

//	func main() {
//		var intArray = [3]int{1, 2, 3}
//		var intSlice = []int{11, 22, 33}
//		fmt.Printf("%v, %v", reflect.TypeOf(intArray), reflect.TypeOf(intSlice))
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectValue(x any) {
// 	fmt.Println(x)
// 	v := reflect.ValueOf(x)
// 	var i int = 32
// 	num := int(v.Int()) + i
// 	fmt.Println(num)
// }

//	func main() {
//		var a = 13
//		reflectValue(a)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectValue(x any) {
// 	v := reflect.ValueOf(x)
// 	kind := v.Kind()

// 	switch kind {
// 	case reflect.Int64:
// 		fmt.Printf("int 类型：%v\n", v.Int()+10)
// 	case reflect.Float64:
// 		fmt.Printf("float64 类型：%v\n", v.Float()+10.1)
// 	case reflect.Float32:
// 		fmt.Printf("float32 类型：%v\n", v.Float()+10.1)
// 	case reflect.String:
// 		fmt.Printf("string 类型：%v\n", v.String())
// 	default:
// 		fmt.Println("无法判断类型")
// 	}
// }

//	func main() {
//		var a float32 = 3.14
//		var b int64 = 100
//		var c string = "你好 Golang"
//		reflectValue(a)
//		reflectValue(b)
//		reflectValue(c)
//	}
// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// type Student struct {
// 	ID     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

//	func main() {
//		s1 := Student{
//			ID:     1,
//			Gender: "男",
//			Name:   "李四",
//			Sno:    "s0001",
//		}
//		data, _ := json.Marshal(s1)
//		jsonStr := string(data)
//		fmt.Println(jsonStr)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectType(x any) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("Type: %v\n", t)
// }

//	func main() {
//		var f float32 = 12.5
//		reflectType(f)
//		var i int64 = 100
//		reflectType(i)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectType(x any) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("TypeOf: %v\tName: %v\tKind: %v\n",
// 		t, t.Name(), t.Kind())
// }

// type MyInt int64

// type Person struct {
// 	Name string
// 	Age  int
// }

// type Animal struct {
// 	Name string
// }

// func main() {
// 	var ptr *float32
// 	var mi MyInt
// 	var c rune
// 	var p = Person{
// 		Name: "平泽唯",
// 		Age:  17,
// 	}
// 	var a = Animal{"后藤一里"}
// 	var slice = []int{1, 2, 3, 4, 5}

//		reflectType(ptr)
//		reflectType(mi)
//		reflectType(c)
//		reflectType(p)
//		reflectType(a)
//		reflectType(slice)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectValue(x any) {
// 	v := reflect.ValueOf(x)
// 	c := v.Int() + 6
// 	fmt.Println(c)
// }

//	func main() {
//		var num int32 = 100
//		reflectValue(num)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectValue(x any) {
// 	v := reflect.ValueOf(x)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Int64:
// 		fmt.Printf("Type: int64\tValue: %d\n", v.Int())
// 	case reflect.Float32:
// 		fmt.Printf("Type: float32\tValue: %f\n", v.Float())
// 	case reflect.Float64:
// 		fmt.Printf("Type: float64\tValue: %f\n", v.Float())
// 	default:
// 		fmt.Println("Type: unknown")
// 	}
// }

// func main() {
// 	var a float32 = 3.14
// 	var b int64 = 100
// 	reflectValue(a)
// 	reflectValue(b)

//		// 将 int 类型的原始值转换为 reflect.Value 类型。
//		c := reflect.ValueOf(10)
//		fmt.Printf("Type(c): %T\n", c)
//	}
// package main

// import "reflect"

//	func main() {
//		// x := 42
//		// v := reflect.ValueOf(x)
//		// v.SetInt(100)
//		x := 42
//		// .Elem() 获取指针指向的可寻址值
//		v := reflect.ValueOf(&x).Elem()
//		// 成功修改 x 的值
//		v.SetInt(100)
//	}
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectSetValue(x any) {
// 	v := reflect.ValueOf(x)
// 	if v.Kind() == reflect.Int64 {
// 		// 修改的是副本，reflect 包会触发 panic。
// 		v.SetInt(200)
// 	}
// }

// func reflectSetVal(x any) {
// 	v := reflect.ValueOf(x)
// 	// 反射中使用 Elem() 方法获取指针对于的值。
// 	if v.Elem().Kind() == reflect.Int64 {
// 		v.Elem().SetInt(200)
// 	}
// }

//	func main() {
//		var a int64 = 100
//		// reflectSetValue(a)
//		reflectSetVal(&a)
//		fmt.Println(a)
//	}
// package main

// import "reflect"

// func reflectSetValue(x any) {
// 	v := reflect.ValueOf(x)
// 	if v.Elem().Kind() == reflect.Int64 {
	
// 	}
// }
