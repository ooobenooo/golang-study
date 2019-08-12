package mytest

import (
	"testing"
)

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(4, 5)
	}
}

func BenchmarkDivideTimeConsumption(b *testing.B) {
	b.StopTimer() // 停止性能计时器

	// 初始化工作，这些时间不计算在性能消耗内

	b.StartTimer() // 开启性能计时器

	for i := 0; i < b.N; i++ {
		Divide(4, 5)
	}
}
