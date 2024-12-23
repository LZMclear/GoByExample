package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu sync.Mutex
	a  int64
}

func (c *Container) inc() { //要修改a的值，传递的是结构体的指针
	//访问前开启互斥锁，访问结束后开锁。
	c.mu.Lock()
	defer c.mu.Unlock()
	c.a += 1
}

func main() {
	c := Container{

		a: 0,
	}

	var wg sync.WaitGroup

	doIncrement := func(n int) {
		for i := 0; i < n; i++ {
			c.inc()
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement(10000)
	go doIncrement(10000)
	go doIncrement(10000)

	wg.Wait()
	fmt.Println(c.a)
}
