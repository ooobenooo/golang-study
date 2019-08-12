package mytest

// 文件名必须以test结尾
// import testing
import (
	"testing"
)

// 单元测试的方法名必须Test开头，后面跟着非小写字母开头的任意字符串
// testing.T 记录测试过程中的内容，包括错误信息，日志等
// go test 命令执行测试
// go test -v 可以看到详细的测试内容
func TestDivideOne(t *testing.T) {
	if _, err := Divide(3, 3); err != nil {
		t.Error("error")
	} else {
		t.Log("pass")
	}
}
