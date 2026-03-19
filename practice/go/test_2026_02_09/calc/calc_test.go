package calc

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("Add(1, 2) 预期结果为 3，实际得到 %d", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("Add(-10, -20) 预期结果为 -30，实际得到 %d", ans)
	}

	if ans := Add(0, 5); ans != 5 {
		t.Errorf("Add(0, 5) 预期结果为 5，实际得到 %d", ans)
	}
}

func TestDiv(t *testing.T) {
	if ans := Div(10, 2); ans != 5 {
		t.Errorf("Div(10, 2) 预期结果为 5，实际得到 %d", ans)
	}

	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Div(10, 0) 未触发除数为 0 的 panic，测试失败")
		}
	}()
	Div(10, 0)
}
