package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	//数值常量没有类型，可以通过显式转换。
	fmt.Println(int64(d))
	//也可以通过函数调用赋值
	fmt.Println(math.Sin(n))
}
