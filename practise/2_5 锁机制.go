package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

    考察点 ： sync.Mutex 的使用、并发数据安全。

func main() {
	var (
		count int //共享计数器
		mutex sync.Mutex
		wg    sync.WaitGroup
	)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Print("最终计数器数值为:", count)
}

题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

    考察点 ：原子操作、并发数据安全。
*/

func main() {
	var (
		count int64
		wg    sync.WaitGroup
	)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Print("最终计数器数值为:", count)
}
