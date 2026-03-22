// package calc

// import "testing"

// func TestAdd(t *testing.T) {
// 	if ans := Add(1, 2); ans != 3 {
// 		t.Errorf("Add(1, 2) 预期结果为 3，实际得到 %d", ans)
// 	}

// 	if ans := Add(-10, -20); ans != -30 {
// 		t.Errorf("Add(-10, -20) 预期结果为 -30，实际得到 %d", ans)
// 	}

// 	if ans := Add(0, 5); ans != 5 {
// 		t.Errorf("Add(0, 5) 预期结果为 5，实际得到 %d", ans)
// 	}
// }

// func TestDiv(t *testing.T) {
// 	if ans := Div(10, 2); ans != 5 {
// 		t.Errorf("Div(10, 2) 预期结果为 5，实际得到 %d", ans)
// 	}

//		defer func() {
//			if err := recover(); err != nil {
//				t.Errorf("Div(10, 0) 未触发除数为 0 的 panic，测试失败")
//			}
//		}()
//		Div(10, 0)
//	}
// package calc

// import "testing"

// func TestAdd(t *testing.T) {
// 	if ans := Add(1, 2); ans != 3 {
// 		t.Errorf("Add(1, 2) 预期结果为 3，实际得到 %d", ans)
// 	}

// 	if ans := Add(-10, -20); ans != -30 {
// 		t.Errorf("Add(-10, -20) 预期结果为 -30，实际得到 %d", ans)
// 	}

// 	if ans := Add(0, 5); ans != 5 {
// 		t.Errorf("Add(0, 5) 预期结果为 5，实际得到 %d", ans)
// 	}
// }

// func TestDiv(t *testing.T) {
// 	if ans := Div(10, 2); ans != 5 {
// 		t.Errorf("Div(10, 2) 预期结果为 5，实际得到 %d", ans)
// 	}

// 	defer func() {
// 		if err := recover(); err != nil {
// 			t.Errorf("Div(10, 0) 未触发除数为 0 的 panic，测试失败")
// 		}
// 	}()

//		Div(10, 0)
//	}
// package calc

// import "testing"

// func TestAdd(t *testing.T) {
// 	t.Run("postive_nums", func(t *testing.T) {
// 		if ans := Add(2, 3); ans != 5 {
// 			t.Fatalf("2 + 3 预期 5，实际 %d", ans)
// 		}
// 	})

// 	t.Run("negative_num", func(t *testing.T) {
// 		if ans := Add(-2, -3); ans != -5 {
// 			t.Fatalf("-2 + (-3) 预期 -5，实际 %d", ans)
// 		}
// 	})

//		t.Run("with_zero", func(t *testing.T) {
//			if ans := Add(0, 3); ans != 3 {
//				t.Fatalf("0 + 3 预期 3，实际 %d", ans)
//			}
//		})
//	}
// package calc

// import "testing"

// func TestAdd(t *testing.T) {
// 	testCases := []struct {
// 		name     string
// 		a, b     int
// 		expected int
// 	}{
// 		{"postive: 2+3", 2, 3, 5},
// 		{"negative: -2+(-3)", -2, -3, -5},
// 		{"zero: 0+5", 0, 5, 5},
// 		{"max: 1e9+1", 1e9, 1, 1e9 + 1},
// 	}

//		for _, tc := range testCases {
//			t.Run(tc.name, func(t *testing.T) {
//				actual := Add(tc.a, tc.b)
//				if actual != tc.expected {
//					t.Errorf("%d+%d 预期 %d，实际 %d", tc.a, tc.b, tc.expected, actual)
//				}
//			})
//		}
//	}
// package calc

// import "testing"

// func TestAdd(t *testing.T) {
// 	testCases := []struct {
// 		name     string
// 		a, b     int
// 		expected int
// 	}{
// 		{"positive: 2+3", 2, 3, 5},
// 		{"negative: -2+(-3)", -2, -3, -5},
// 		{"zero: 0+5", 0, 5, 5},
// 		{"max: 1e9+1", 1e9, 1, 1e9 + 1},
// 	}

//		for _, tc := range testCases {
//			t.Run(tc.name, func(t *testing.T) {
//				actual := Add(tc.a, tc.b)
//				if actual != tc.expected {
//					t.Errorf("%d+%d 预期 %d，实际 %d", tc.a, tc.b, tc.expected, actual)
//				}
//			})
//		}
//	}
// package calc

// import (
// 	"fmt"
// 	"os"
// 	"testing"
// )

// func setup() {
// 	fmt.Println("=== 测试前初始化：建立数据库连接 ===")
// }

// func teardawn() {
// 	fmt.Println("=== 测试后清理：断开数据库连接 ===")
// }

// func Test1(t *testing.T) {
// 	fmt.Println("正在执行测试：Test1")
// }

// func Test2(t *testing.T) {
// 	fmt.Println("正在执行测试：Test2")
// }

//	func TestMain(m *testing.M) {
//		setup()
//		exitCode := m.Run()
//		teardawn()
//		os.Exit(exitCode)
//	}
package calc

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("=== 测试前初始化：建立数据库连接 ===")
}

func teardown() {
	fmt.Println("=== 测试后清理：断开数据库连接 ===")
}

func Test1(t *testing.T) {
	fmt.Println("正在执行测试：Test1")
}

func Test2(t *testing.T) {
	fmt.Println("正在执行测试：Test2")
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}
