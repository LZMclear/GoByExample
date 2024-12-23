package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	//测试模式是否与字符串匹配
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//在上面我们直接使用字符串模式，但是对于其他的正则表达式任务需要优化
	r, _ := regexp.Compile("p([a-z]+)ch")
	//和上面的一样，输出正则表达式与字符串是否匹配
	fmt.Println(r.MatchString("peach"))
	//找到正则表达式的匹配项  peach
	fmt.Println(r.FindString("peach punch"))
	//找到匹配项返回开始结束索引
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	//[peach ea]返回正则表达式的匹配项与子匹配项
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//返回匹配项与子匹配项的索引[0 5 1 3]  end不包括最后字符串的索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//查找所有匹配项示例
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	//查找所有匹配项和子匹配项返回其索引
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	//提供非负整数，限制匹配项数[peach punch]
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))
	//创建正则表达式的全局变量 你可以使用panic而不是返回一个error，这样会使使用全局变量更加安全
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	//用后面字符串代替正则表达式匹配到的字符串
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
