package main

import "fmt"

// 循环控制
//
//	func main() {
//		// 方式1
//		for i := 0; i < 10; i++ {
//			fmt.Println("方式1，第", i+1, "次循环")
//		}
//
//		// 方式2
//		b := 1
//		for b < 10 {
//			fmt.Println("方式2，第", b, "次循环")
//		}
//
//		// 方式3，无限循环
//		ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
//		var started bool
//		var stopped atomic.Bool
//		for {
//			if !started {
//				started = true
//				go func() {
//					for {
//						select {
//						case <-ctx.Done():
//							fmt.Println("ctx done")
//							stopped.Store(true)
//							return
//						}
//					}
//				}()
//			}
//			fmt.Println("main")
//			if stopped.Load() {
//				break
//			}
//		}
//
//		// 遍历数组
//		var a [10]string
//		a[0] = "Hello"
//		for i := range a {
//			fmt.Println("当前下标：", i)
//		}
//		for i, e := range a {
//			fmt.Println("a[", i, "] = ", e)
//		}
//
//		// 遍历切片
//		s := make([]string, 10)
//		s[0] = "Hello"
//		for i := range s {
//			fmt.Println("当前下标：", i)
//		}
//		for i, e := range s {
//			fmt.Println("s[", i, "] = ", e)
//		}
//
//		m := make(map[string]string)
//		m["b"] = "Hello, b"
//		m["a"] = "Hello, a"
//		m["c"] = "Hello, c"
//		for i := range m {
//			fmt.Println("当前key：", i)
//		}
//		for k, v := range m {
//			fmt.Println("m[", k, "] = ", v)
//		}
//	}
type A struct {
	i int
}

func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
//var function1 func(int) int
//
//// 声明闭包
//var squart2 func(int) int = func(p int) int {
//	p *= p
//	return p
//}
//
//func main() {
//	a := A{1}
//	// 把方法赋值给函数变量
//	function1 = a.add
//
//	// 声明一个闭包并直接执行
//	// 此闭包返回值是另外一个闭包（带参闭包）
//	returnFunc := func() func(int, string) (int, string) {
//		fmt.Println("this is a anonymous function")
//		return func(i int, s string) (int, string) {
//			return i, s
//		}
//	}()
//
//	// 执行returnFunc闭包并传递参数
//	ret1, ret2 := returnFunc(1, "test")
//	fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)
//
//	fmt.Println("a.i = ", a.i)
//	fmt.Println("after call function1, a.i = ", function1(1))
//	fmt.Println("a.i = ", a.i)
//}

// 全局变量----------------------------
var a int

func main() {
	//{
	//	fmt.Println("global variable, a = ", a)
	//	a = 3
	//	fmt.Println("global variable, a = ", a)
	//
	//	a := 10
	//	fmt.Println("local variable, a = ", a)
	//	a--
	//	fmt.Println("local variable, a = ", a)
	//}
	//fmt.Println("global variable, a = ", a)

	var b int = 4
	fmt.Println("local variable, b = ", b)
	if b := 3; b == 3 {
		fmt.Println("if statement, b = ", b)
		b--
		fmt.Println("if statement, b = ", b)
	}
	fmt.Println("local variable, b = ", b)
}
