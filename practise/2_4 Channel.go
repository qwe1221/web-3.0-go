package main

import (
	"fmt"
	"sync"
)

/**
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

    考察点 ：通道的基本使用、协程间通信。

*/

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for i := 0; i < 110; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("接收到的数字", num)
		}
	}()

	wg.Wait()
	fmt.Println("主线程退出")
}
