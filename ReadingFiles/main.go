package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//直接读取文件的全部内容
	file, err := os.ReadFile("dat.txt")
	check(err)
	fmt.Println(string(file))

	//对文件中的具体内容进行读取
	//打开文件
	f, err := os.Open("dat.txt")
	check(err)

	//读取5字节数据
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	//从指定位置开始读取数据
	b2 := make([]byte, 2)
	ret, err := f.Seek(6, io.SeekStart) //从起始位置往后第六个开始'g'
	check(err)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, ret)
	fmt.Printf("%v\n", string(b2[:n2]))

	//验证是Seek是否改变了起始位置  当前的位置是go_  理解：上段代码读取数据后，光标移动到末尾，所以current就是go后面的索引
	b3 := make([]byte, 3)
	ret2, err := f.Seek(-4, io.SeekCurrent)
	check(err)
	n3, err := f.Read(b3)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n3, ret2)
	fmt.Printf("%v\n", string(b3[:n3]))

	//使用函数读取
	//将光标定位到g
	ret3, err := f.Seek(6, io.SeekStart)
	check(err)
	b4 := make([]byte, 2)
	n4, err2 := io.ReadAtLeast(f, b4, 2)
	check(err2)
	fmt.Printf("%d bytes @ %d: %s\n", n4, ret3, string(b4))
	//调整光标
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f)
	b5, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b5))

	f.Close()

}
