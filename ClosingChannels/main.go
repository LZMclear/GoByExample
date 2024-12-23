package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	//执行完发送数据后，无缓冲通道done阻塞程序，线程执行，获取完毕后向done通道传输数据，程序继续执行知道结束
	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
