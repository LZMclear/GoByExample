package main

import (
	"fmt"
	"sync"
)

// Con 将ops包装到结构体中使用互斥锁进行修改
type Con struct {
	mux sync.Mutex
	ops int64
}

func main() {
	o := Con{ops: 0}

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				o.mux.Lock()
				o.ops += 1
				o.mux.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", o.ops)
}
