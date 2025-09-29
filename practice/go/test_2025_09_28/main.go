// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	ticker := time.NewTicker(time.Millisecond * 20)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Println("fn: 小 Go，你好！\t", count)
// 		count++
// 	}

// 	wg.Done()
// }

// func test(wg *sync.WaitGroup) {
// 	ticker := time.NewTicker(time.Millisecond * 20)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Println("test: 小 Go，你好！\t", count)
// 		count++
// 	}

// 	wg.Done() // 协程结束器减 1。
// }

//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(1)  // 协程结束器加 1。
//		go fn(&wg) // 表示开启一个协程。
//		wg.Add(1)
//		go test(&wg)
//		for i := range 10 {
//			fmt.Println("main: 小 Go，你好！\t", i)
//			time.Sleep(time.Millisecond * 50)
//		}
//		wg.Wait() // 等待协程执行完毕。
//		fmt.Println("主线程退出……")
//	}
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

//	func main() {
//		cpuNum := runtime.NumCPU()
//		fmt.Println("cpuNum =", cpuNum)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func fn(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := range 5 {
// 		fmt.Printf("(%v, %v)\n", num, i)
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		for i := range 5 {
//			wg.Add(1)
//			go fn(i, &wg)
//		}
//		wg.Wait()
//		fmt.Println("关闭主线程！")
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done() // goroutine 结束就登记 1。
// 	ticker := time.NewTicker(50 * time.Millisecond)
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
//		var wg sync.WaitGroup // 定义一个 WaitGroup。
//		wg.Add(1)             // 启动一个 goroutine 就加 1。
//		go fn(&wg)
//		for i := range 10 {
//			fmt.Printf("main: 小 Go，你好！\t%v\n", i)
//			time.Sleep(10 * time.Millisecond)
//		}
//		wg.Wait() // 等待所有的 goroutine 都结束。
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

// 	func main() {
// 		var wg sync.WaitGroup
// 		for i := range 10 {
// 			wg.Add(1)
// 			go fn(i, &wg)
// 		}
// 		wg.Wait()
// 	}
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	// 获取当前计算机上的 CPU 个数。
// 	cpuNum := runtime.NumCPU()
// 	fmt.Println("cpuNum =", cpuNum)

//		// 可以自行设置使用 CPU 的数量。
//		runtime.GOMAXPROCS(cpuNum - 1)
//		fmt.Println("OK!")
//	}
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

//	func main() {
//		cpuNum := runtime.NumCPU()
//		runtime.GOMAXPROCS(cpuNum / 2)
//		fmt.Println("OK!")
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ticker := time.NewTicker(50 * time.Millisecond)
// 	defer ticker.Stop()

// 	count := 0
// 	for range ticker.C {
// 		if count >= 10 {
// 			break
// 		}
// 		fmt.Println("fn: 小平您好")
// 		count++
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		wg.Add(1)
//		go fn(&wg)
//		for range 10 {
//			fmt.Println("mian: 小平您好")
//			time.Sleep(50 * time.Millisecond)
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
// 	"time"
// )

//	func main() {
//		start := time.Now()
//		for num := 2; num < 120000; num++ {
//			isPrime := true
//			for i := 2; i < num; i++ {
//				if num%i == 0 {
//					isPrime = false
//					break
//				}
//			}
//			if isPrime {
//				// fmt.Print(num, " ")
//			}
//		}
//		elapsed := time.Since(start)
//		// 统计普通的循环需要的时间。
//		fmt.Println(elapsed)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

//	func main() {
//		start := time.Now()
//		for num := 2; num < 1.2e5; num++ {
//			isPrime := true
//			for i := 2; i < num; i++ {
//				if num%i == 0 {
//					isPrime = false
//					break
//				}
//			}
//			if isPrime {
//			}
//		}
//		elapsed := time.Since(start)
//		fmt.Println(elapsed)
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func fn(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	startNum := n * 3e4
// 	endNum := (n + 1) * 3e4

// 	if startNum < 2 {
// 		startNum = 2
// 	}

// 	for num := startNum; num < endNum; num++ {
// 		var isPrime = true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 		}
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	for i := range 4 {
// 		wg.Add(1)
// 		go fn(i, &wg)
// 	}

//		wg.Wait()
//		elapsed := time.Since(start)
//		fmt.Println(elapsed)
//	}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	primeCount := 0

// 	for num := 2; num < 1.2e5; num++ {
// 		isPrime := true
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			primeCount++
// 		}
// 	}
// 	elapsed := time.Since(start)

//		fmt.Printf("在 120000 内共找到 %d 个素数。\n", primeCount)
//		fmt.Println("耗时:", elapsed)
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
// 	parts := 4
// 	interval := 1.2e5 / parts
// 	results := make([]int, parts)

// 	for i := range parts {
// 		wg.Add(1)
// 		s := i * interval
// 		e := (i + 1) * interval
// 		if s < 2 {
// 			s = 2
// 		}
// 		go countPrimes(s, e, &wg, &results[i])
// 	}

// 	wg.Wait()
// 	total := 0
// 	for _, v := range results {
// 		total += v
// 	}

//		fmt.Printf("在 120000 内共找到 %d 个素数。\n", total)
//		fmt.Println("耗时:", time.Since(start))
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
// 	parts := 4
// 	interval := 1.2e5 / parts
// 	result := make([]int, 4)

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

//		fmt.Printf("在 120000 内共找到 %d 个素数。\n", total)
//		fmt.Println("耗时:", time.Since(start))
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

//		fmt.Printf("在 120000 内共找到 %d 个素数。\n", total)
//		fmt.Println("耗时:", time.Since(start))
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
// 	parts := 5
// 	interval := 1.2e5 / parts
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

//		fmt.Printf("在 120000 内共找到 %v 个素数。", total)
//		fmt.Printf("花费 %v。", time.Since(start))
//	}
// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 10)
// 	for i := range 10 {
// 		ch <- i
// 	}
// 	for range 10 {
// 		fmt.Print(<-ch, " ")
// 	}
// 	fmt.Println(len(ch), cap(ch))

// 	ch1 := make(chan int, 4)

// 	ch1 <- 34
// 	ch1 <- 54
// 	ch1 <- 64

// 	ch2 := ch1
// 	ch2 <- 74

// 	<-ch1
// 	<-ch1
// 	<-ch1

//		ch3 := make(chan int, 1)
//		ch3 <- 34
//		<-ch3
//		ch3 <- 44
//	}
// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 10)

// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)

//		for val := range ch {
//			fmt.Print(val, " ")
//		}
//		fmt.Println()
//	}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // 写数据
// func write(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 		fmt.Printf("写入数据：%v\n", i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// 	close(ch)
// }

// // 读数据
// func read(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for v := range ch {
// 		fmt.Printf("读取数据：%v\n", v)
// 		time.Sleep(50 * time.Millisecond)
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		ch := make(chan int, 10)
//		wg.Add(2)
//		go write(ch, &wg)
//		go read(ch, &wg)
//		wg.Wait()
//		fmt.Println("执行完毕")
//	}
// package main

// import "fmt"

//	func main() {
//		ch := make(chan int, 1)
//		ch <- 10
//		fmt.Println("发送成功")
//	}
// package main

// import "fmt"

//	func main() {
//		ch := make(chan int, 1)
//		ch <- 10
//		ch <- 12
//		fmt.Println("发送成功")
//	}
// package main

// import "fmt"

//	func main() {
//		ch := make(chan int, 1)
//		ch <- 10 // 放入
//		<-ch     // 取走
//		ch <- 12 // 放入
//		<-ch     // 取走
//		ch <- 17 // 放入
//		fmt.Println("发送成功")
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
// 		for i := 2; i*i < num; i++ {
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
// 	parts := 6
// 	interval := 1.2e6 / parts
// 	results := make([]int, parts)

// 	for i := range parts {
// 		wg.Add(1)
// 		s := i * interval
// 		e := (i + 1) * interval
// 		go countPrimes(s, e, &wg, &results[i])
// 	}
// 	wg.Wait()

//		total := 0
//		for _, v := range results {
//			total += v
//		}
//		fmt.Println(total)
//		fmt.Println(time.Since(start))
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
// 	"time"
// )

// func write(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := range 10 {
// 		ch <- i + 1
// 		fmt.Printf("写入数据：%v\n", i+1)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	close(ch)
// }

// func read(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for v := range ch {
// 		fmt.Printf("读取数据：%v\n", v)
// 		time.Sleep(time.Millisecond * 50)
// 	}
// }

//	func main() {
//		var wg sync.WaitGroup
//		ch := make(chan int)
//		wg.Add(2)
//		go write(ch, &wg)
//		go read(ch, &wg)
//		wg.Wait()
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
// 	for num := 2; num < 120; num++ {
// 		numChan <- num
// 	}
// 	close(numChan)
// }

// func filterPrimes(numChan chan int, primeChan chan int, exitChan chan bool, wg *sync.WaitGroup) {
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

// func printPrime(primeChan chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for prime := range primeChan {
// 		fmt.Println(prime)
// 	}
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workerCount := 4
// 	numChan := make(chan int, 1e3)
// 	primeChan := make(chan int, 2e4)
// 	exitChan := make(chan bool, 8)

// 	wg.Add(1)
// 	go generateNumbers(numChan, &wg)

// 	for range workerCount {
// 		wg.Add(1)
// 		go filterPrimes(numChan, primeChan, exitChan, &wg)
// 	}

// 	wg.Add(1)
// 	go printPrime(primeChan, &wg)

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

// func filterPrimes(numChan chan int, primeChan chan int, exitChan chan bool, wg *sync.WaitGroup) {
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

// func countPrimes(primeChan chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	count := 0
// 	for range primeChan {
// 		count++
// 	}
// 	fmt.Printf("在 120000 内共找到 %d 个素数。\n", count)
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	workerCount := 10
// 	numChan := make(chan int, 1000)
// 	primeChan := make(chan int, 1000)
// 	exitChan := make(chan bool, workerCount)

// 	wg.Add(1)
// 	go generateNumbers(numChan, &wg)

// 	for range workerCount {
// 		wg.Add(1)
// 		go filterPrimes(numChan, primeChan, exitChan, &wg)
// 	}

// 	wg.Add(1)
// 	go countPrimes(primeChan, &wg)

// 	wg.Go(func() {
// 		for range workerCount {
// 			<-exitChan
// 		}
// 		close(primeChan)
// 	})

//		wg.Wait()
//		fmt.Println("耗时:", time.Since(start))
//	}
package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := 2; num < 120; num++ {
		numChan <- num
	}
	close(numChan)
}

func filterPrimes(numChan chan int, primeChan chan int, exitChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numChan {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeChan <- num
		}
	}
	exitChan <- true
}

func printPrimes(primeChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for prime := range primeChan {
		fmt.Print(prime, " ")
	}
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	workerCount := 5
	numChan := make(chan int, 10)
	primeChan := make(chan int, 10)
	exitChan := make(chan bool, workerCount)

	wg.Add(1)
	go generateNumbers(numChan, &wg)

	for range workerCount {
		wg.Add(1)
		go filterPrimes(numChan, primeChan, exitChan, &wg)
	}

	wg.Add(1)
	go printPrimes(primeChan, &wg)

	wg.Go(func() {
		for range workerCount {
			<-exitChan
		}
		close(primeChan)
	})

	wg.Wait()

	fmt.Println()
	fmt.Println(time.Since(start))
}
