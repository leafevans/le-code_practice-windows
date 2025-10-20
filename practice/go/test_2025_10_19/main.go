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
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var n, m int
// 	scanner.Scan()
// 	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

// 	matrix := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		matrix[i] = make([]int, m)
// 		scanner.Scan()
// 		line := scanner.Text()
// 		fields := strings.Fields(line)
// 		for j := 0; j < m; j++ {
// 			x, _ := strconv.Atoi(fields[j])
// 			matrix[i][j] = x
// 		}
// 	}

// 	prefix := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		prefix[i] = make([]int, m+1)
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			prefix[i+1][j+1] = matrix[i][j] + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
// 		}
// 	}

// 	total := prefix[n][m]
// 	minDiff := total

// 	for i := 1; i < n; i++ {
// 		top := prefix[i][m]
// 		bottom := total - top
// 		diff := int(math.Abs(float64(top - bottom)))
// 		if diff < minDiff {
// 			minDiff = diff
// 		}
// 	}

// 	for j := 1; j < m; j++ {
// 		left := prefix[n][j]
// 		right := total - left
// 		diff := int(math.Abs(float64(left - right)))
// 		if diff < minDiff {
// 			minDiff = diff
// 		}
// 	}

//		fmt.Println(minDiff)
//	}
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

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		for j := 0; j < m; j++ {
			x, _ := strconv.Atoi(fields[j])
			matrix[i][j] = x
		}
	}

	prefix := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		prefix[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			prefix[i+1][j+1] = matrix[i+1][j+1] + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
		}
	}

	total := prefix[n][m]
	minDiff := total

	for i := 1; i < n; i++ {
		top := prefix[i][m]
		bottom := total - top
		diff := int(math.Abs(float64(top - bottom)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	for j := 1; j < m; j++ {
		left := prefix[n][j]
		right := total - left
		diff := int(math.Abs(float64(left - right)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	fmt.Println(minDiff)
}
