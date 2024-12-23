package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	//开启一个线程，执行选择，如果c1中有值，执行第一个，
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): //一秒之后time.After通道中有值可以取出，执行第二个case
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
