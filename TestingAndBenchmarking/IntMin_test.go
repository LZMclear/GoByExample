package main

import (
	"fmt"
	"testing"
)

//测试比较整型的函数实现
//典型的，我们测试代码的源文件名字大概为intutils.go，它的测试文件就可以命名为iniutils_test.go

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestIntMinBasic 测试函数是通过创建一个以Test开头的函数
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		//t.Errorf会输出错误测试，并且继续执行test t.Fatal会输出错误停止测试
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// TestIntMinTableDriven 测试可以重复，习惯使用表驱动模式，其中，输入和输出放在表中单独进行测试
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// BenchmarkIntMin 基准测试测试_test.go典型的go,testing runner执行几次Benchmark函数，每次执行都会增加b.N直到收集精准的测量结果
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
