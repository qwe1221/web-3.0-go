package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

	考察点 ：协程原理、并发任务调度。
*/

type Task func()

type TaskResult struct {
	TaskIndex int           //任务索引
	Duration  time.Duration //执行时间
	Error     error
}

// 任务调度器
type Scheduler struct {
	tasks   []Task          // 待执行的任务列表
	results chan TaskResult // 任务结果通道
	wg      sync.WaitGroup  // 用于等待所有任务完成
	mutex   sync.Mutex      // 保护共享数据
}

// NewScheduler 创建新的任务调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		results: make(chan TaskResult, 10), // 缓冲通道防止阻塞
	}
}

// AddTask 添加任务到调度器
func (s *Scheduler) AddTask(task Task) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.tasks = append(s.tasks, task)
}

// Run 启动所有任务并返回结果通道
func (s *Scheduler) Run() <-chan TaskResult {
	s.wg.Add(len(s.tasks))

	for i, task := range s.tasks {
		go func(taskIndex int, taskFunc Task) {
			defer s.wg.Done()

			startTime := time.Now()
			defer func() {
				// 捕获panic并转换为错误
				if r := recover(); r != nil {
					s.results <- TaskResult{
						TaskIndex: taskIndex,
						Duration:  time.Since(startTime),
						Error:     fmt.Errorf("panic: %v", r),
					}
				}
			}()

			// 执行任务
			taskFunc()

			// 发送结果
			s.results <- TaskResult{
				TaskIndex: taskIndex,
				Duration:  time.Since(startTime),
				Error:     nil,
			}
		}(i, task)
	}

	// 启动一个goroutine在所有任务完成后关闭结果通道
	go func() {
		s.wg.Wait()
		close(s.results)
	}()

	return s.results
}

func main() {
	// 创建调度器
	scheduler := NewScheduler()

	// 添加一些示例任务
	for i := 0; i < 5; i++ {
		taskID := i // 创建局部变量避免闭包问题
		scheduler.AddTask(func() {
			// 模拟任务执行，耗时随机
			time.Sleep(time.Duration(taskID*100+100) * time.Millisecond)
			fmt.Printf("任务 %d 完成\n", taskID)
		})
	}

	// 添加一个会panic的任务
	scheduler.AddTask(func() {
		time.Sleep(100 * time.Millisecond)
		panic("模拟panic")
	})

	// 运行所有任务并获取结果
	startTime := time.Now()
	results := scheduler.Run()

	// 处理结果
	for result := range results {
		if result.Error != nil {
			fmt.Printf("任务 %d 执行失败: %v (耗时: %v)\n",
				result.TaskIndex, result.Error, result.Duration)
		} else {
			fmt.Printf("任务 %d 执行成功 (耗时: %v)\n",
				result.TaskIndex, result.Duration)
		}
	}

	fmt.Printf("所有任务完成，总耗时: %v\n", time.Since(startTime))
}
