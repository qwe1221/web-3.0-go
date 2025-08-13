package main

import (
	"fmt"
	"sync"
)

/**
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

    考察点 ：通道的缓冲机制。
*/

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	//创建缓冲区大小为10
	ch := make(chan int, 10)
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			ch <- i
			fmt.Printf("生产者发送: %d\n", i)
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
}
