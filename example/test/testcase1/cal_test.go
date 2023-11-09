package up

import (
	"testing"
)

// 编写测试用例, 测试 addUpper
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10)执行错误, 期望值=%v 实际值=%v\n", 55, res)
	}

	// 如果正确, 输出日志
	t.Logf("AddUpper(10) 执行正确")
}
