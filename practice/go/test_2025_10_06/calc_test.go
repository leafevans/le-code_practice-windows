// package main

// // TestAdd 测试 Add 函数：覆盖正常值、负数、边界值（如 0）
// func TestAdd(t *testing.T) {
// 	// 用例 1：正常正数相加
// 	if ans := Add(1, 2); ans != 3 {
// 		t.Errorf("Add(1, 2) 预期结果为 3，实际得到 %d", ans)
// 	}

// 	// 用例 2：负数相加
// 	if ans := Add(-10, -20); ans != -30 {
// 		t.Errorf("Add(-10, -20 预期结果 -30，实际得到 %d", ans)
// 	}

// 	// 用例 3：边界值（含 0）
// 	if ans := Add(0, 5); ans != 5 {
// 		t.Errorf("Add(0, 5) 预期结果为 5，实际得到 %d", ans)
// 	}
// }

// TestDiv 测试 Div 函数：覆盖正常情况和异常情况（除数为 0）
// func TestDiv(t *testing.T) {
// 	// 正常用例：整除
// 	if ans := Div(10, 2); ans != 5 {
// 		t.Errorf("Div(10, 2) 预期结果为 5，实际得到 %d", ans)
// 	}

// 	// 异常用例：除数为 0（此处会触发 panic，需特殊处理）
// 	defer func() {
// 		if err := recover(); err == nil {
// 			t.Errorf("Div(10, 0) 未触发除数为 0 的 panic，测试失败")
// 		}
// 	}()
// 	_ = Div(10, 0)
// }

// func TestAdd(t *testing.T) {
// 	// 子测试 1：正数相加
// 	t.Run("positive_nums", func(t *testing.T) {
// 		if ans := Add(2, 3); ans != 5 {
// 			t.Fatalf("2 + 3 预期 5，实际 %d", ans)
// 		}
// 	})

// 	// 子测试 2：负数相加
// 	t.Run("negative_nums", func(t *testing.T) {
// 		if ans := Add(-2, -3); ans != -5 {
// 			t.Fatalf("-2 + (-3) 预期 -5，实际 %d", ans)
// 		}
// 	})

//		// 子测试 3：含 0 的相加
//		t.Run("with_zero", func(t *testing.T) {
//			if ans := Add(0, 3); ans != 3 {
//				t.Fatalf("0 + 3 预期 3，实际 %d", ans)
//			}
//		})
//	}
// package main

// import "testing"

// func TestAdd(t *testing.T) {
// 	// 定义测试用例表格：包含用例名、输入 A、输入 B、预期结果
// 	testCases := []struct {
// 		name     string // 用例名（唯一，便于识别）
// 		a, b     int    // 输入参数
// 		expected int    // 预期结果
// 	}{
// 		{"postive: 2+3", 2, 3, 5},
// 		{"negative: -2+(-3)", -2, -3, -5},
// 		{"zero: 0+5", 0, 5, 5},
// 		{"max: 1e9+1", 1e9, 1, 1e9 + 1}, // 边界值：大数相加
// 	}

//		// 循环执行所有用例
//		for _, tc := range testCases {
//			// 子测试：用例名作为子测试标识
//			t.Run(tc.name, func(t *testing.T) {
//				actual := Add(tc.a, tc.b)
//				if actual != tc.expected {
//					t.Errorf("%d+%d 预期 %d，实际 %d",
//						tc.a, tc.b, tc.expected, actual)
//				}
//			})
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// 	"testing"
// 	// 假设引入数据库驱动
// 	// _ "github.com/go-sql-driver/mysql"
// )

// // setup 测试前初始化：如连接数据库
// func setup() {
// 	fmt.Println("=== 测试前初始化：连接数据库 ===")
// 	// 实际代码：db, err := sql.Open(...)
// }

// // teardown 测试后清理：如关闭数据库
// func teardown() {
// 	fmt.Println("=== 测试后清理：关闭数据库 ===")
// 	// 实际代码：db.Close()
// }

// // Test1 普通测试函数1
// func Test1(t *testing.T) {
// 	fmt.Println("执行测试1：Test1")
// }

// // Test2 普通测试函数2
// func Test2(t *testing.T) {
// 	fmt.Println("执行测试2：Test2")
// }

// // TestMain 全局测试入口（必须命名为TestMain，参数为*m.testing.M）
// func TestMain(m *testing.M) {
// 	// 1. 测试前初始化
// 	setup()

// 	// 2. 执行所有测试函数，返回测试结果码（0=成功，非0=失败）
// 	exitCode := m.Run()

// 	// 3. 测试后清理
// 	teardown()

//		// 4. 退出程序（必须调用，否则测试会卡住）
//		os.Exit(exitCode)
//	}
// package main

// import "testing"

// func TestAdd(t *testing.T) {
// 	if ans := Add(1, 2); ans != 3 {
// 		t.Errorf("Add(1,2) 预期结果 3，实际得到 %d", ans)
// 	}

// 	if ans := Add(-10, -20); ans != -30 {
// 		t.Errorf("Add(-10,-20) 预期结果 -30，实际得到 %d", ans)
// 	}

// 	if ans := Add(0, 5); ans != 5 {
// 		t.Errorf("Add(0,5) 预期结果为 5，实际得到 %d", ans)
// 	}
// }

// func TestDiv(t *testing.T) {
// 	if ans := Div(10, 2); ans != 5 {
// 		t.Errorf("Div(10,2) 预期结果为 5，实际得到 %d", ans)
// 	}

//		defer func() {
//			if err := recover(); err == nil {
//				t.Errorf("Div(10,0) 未触发除数为 0 的 panic，测试失败")
//			}
//		}()
//		_ = Div(10, 0)
//	}
// package main

// import "testing"

// func TestAdd(t *testing.T) {
// 	t.Run("postive_nums", func(t *testing.T) {
// 		if ans := Add(2, 3); ans != 5 {
// 			t.Fatalf("2+3 预期 5，实际 %d", ans)
// 		}
// 	})

// 	t.Run("negative_nums", func(t *testing.T) {
// 		if ans := Add(-2, -3); ans != -5 {
// 			t.Fatalf("-2+(-3) 预期 -5，实际 %d", ans)
// 		}
// 	})

//		t.Run("with_zero", func(t *testing.T) {
//			if ans := Add(0, 3); ans != 3 {
//				t.Fatalf("0+3 预期 3，实际 %d", ans)
//			}
//		})
//	}
// package main

// import "testing"

// func TestAdd(t *testing.T) {
// 	testCases := []struct {
// 		name     string
// 		a, b     int
// 		expected int
// 	}{
// 		{"postive:2+3", 2, 3, 5},
// 		{"negative:-2+(-3)", -2, -3, -5},
// 		{"zero:0+5", 0, 5, 5},
// 		{"max:1e9+1", 1e9, 1, 1e9 + 1},
// 	}

//		for _, tc := range testCases {
//			t.Run(tc.name, func(t *testing.T) {
//				if actual := Add(tc.a, tc.b); actual != tc.expected {
//					t.Errorf("%d+%d 预期 %d，实际 %d", tc.a, tc.b, tc.expected, actual)
//				}
//			})
//		}
//	}
package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("=== 测试前初始化：连接数据库 ===")
}

func teardown() {
	fmt.Println("=== 测试后清理：关闭数据库 ===")
}

func Test1(t *testing.T) {
	fmt.Println("执行测试1：Test1")
}

func Test2(t *testing.T) {
	fmt.Println("执行测试2：Test2")
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}
