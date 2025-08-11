package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	//a := [5]int{6, 5, 4, 3, 2}
	//// 从数组下标2开始，直到数组的最后一个元素
	//s7 := a[2:]
	//// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	//s8 := a[1:3]
	//// 从0到下标2的元素，创建一个新的切片
	//s9 := a[:2]
	//fmt.Println(s7)
	//fmt.Println(s8)
	//fmt.Println(s9)
	//a[0] = 9
	//a[1] = 8
	//a[2] = 7
	//fmt.Println(s7)
	//fmt.Println(s8)
	//fmt.Println(s9)
	//复制切片
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)
}
