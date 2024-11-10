package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	//var wg sync.WaitGroup

	// 发送者goroutine
	go func() {
		//defer wg.Done() // 确保在函数返回时调用wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i // 发送数据到channel
		}
		close(ch) // 发送完毕后关闭channel
	}()

	// 接收者goroutine
	go func() {
		//defer wg.Done()     // 确保在函数返回时调用wg.Done()
		for c := range ch { // 从channel接收数据
			fmt.Println(fmt.Sprintf("go %d", c))
		}
	}()

	//wg.Add(2) // 为两个goroutine增加计数
	//wg.Wait() // 等待所有goroutine完成

	//time.Sleep(time.Second)

	for c := range ch { // 从channel接收数据
		fmt.Println(c)
	}

	fmt.Println("end")
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1

	select {
	case <-c:
		x, y = y, x+y
	case <-quit:
		fmt.Println("quit")
		return
	}
}
