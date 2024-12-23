package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//直接写入
	d1 := []byte("hello\ngo\n") //将字符串转换为byte类型
	err := os.WriteFile("dat1.txt", d1, 0644)
	check(err)

	//创建文件再写入
	file, err := os.Create("dat2.txt")
	check(err)
	defer file.Close()

	d2 := []byte{115, 111, 109, 101, 10} //阿斯克码表
	n, err := file.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
}
