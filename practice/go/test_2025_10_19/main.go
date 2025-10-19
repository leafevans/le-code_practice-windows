// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())

// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		scanner.Scan()
// 		nums[i], _ = strconv.Atoi(scanner.Text())
// 	}

// 	prefix := make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		prefix[i+1] = prefix[i] + nums[i]
// 	}

// 	for {
// 		var a, b int
// 		scanner.Scan()
// 		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
// 		if err == io.EOF {
// 			return
// 		}

// 		sum := prefix[b+1] - prefix[a]
// 		fmt.Println(sum)
// 	}
// }
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())

// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		scanner.Scan()
// 		nums[i], _ = strconv.Atoi(scanner.Text())
// 	}

// 	prefix := make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		prefix[i+1] = prefix[i] + nums[i]
// 	}

// 	for {
// 		var a, b int
// 		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
// 		if err != nil {
// 			return
// 		}

//			sum := prefix[b+1] - prefix[a]
//			fmt.Println(sum)
//		}
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())

// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		scanner.Scan()
// 		nums[i], _ = strconv.Atoi(scanner.Text())
// 	}

// 	prefix := make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		prefix[i+1] = prefix[i] + nums[i]
// 	}

// 	for {
// 		var a, b int
// 		scanner.Scan()
// 		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
// 		if err != nil {
// 			return
// 		}

//			sum := prefix[b+1] - prefix[a]
//			fmt.Println(sum)
//		}
//	}
//
// ...existing code...
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n, m int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

	if n <= 1 || m <= 1 {
		fmt.Println(0)
		return
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		fields := strings.Fields(line)
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			x, _ := strconv.Atoi(fields[j])
			matrix[i][j] = x
		}
	}

	preMatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		preMatrix[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			preMatrix[i+1][j+1] = matrix[i][j] + preMatrix[i][j+1] + preMatrix[i+1][j] - preMatrix[i][j]
		}
	}

	total := preMatrix[n][m]
	minDiff := total

	for i := 1; i < n; i++ {
		topSum := preMatrix[i][m]
		bottomSum := total - topSum

		diff := int(math.Abs(float64(topSum - bottomSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	for j := 1; j < m; j++ {
		leftSum := preMatrix[n][j]
		rightSum := total - leftSum

		diff := int(math.Abs(float64(leftSum - rightSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	fmt.Println(minDiff)
}
