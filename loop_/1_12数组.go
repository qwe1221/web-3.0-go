package main

import "fmt"

func main() {
	// 仅声明
	//var a [5]int
	//fmt.Println("a = ", a)
	//
	//var marr [2]map[string]string
	//fmt.Println("marr = ", marr)
	//// map的零值是nil，虽然打印出来是非空值，但真实的值是nil
	//// marr[0]["test"] = "1"
	//
	//// 声明以及初始化
	//var b [5]int = [5]int{1, 2, 3, 4, 5}
	//fmt.Println("b = ", b)
	//
	//// 类型推导声明方式
	//var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	//fmt.Println("c = ", c)
	//
	//d := [3]int{3, 2, 1}
	//fmt.Println("d = ", d)
	//
	//// 使用 ... 代替数组长度
	//autoLen := [...]string{"auto1", "auto2", "auto3"}
	//fmt.Println("autoLen = ", autoLen)
	//
	//// 声明时初始化指定下标的元素值
	//positionInit := [5]string{1: "position1", 3: "position3"}
	//fmt.Println("positionInit = ", positionInit)
	//
	//// 初始化时，元素个数不能超过数组声明的长度
	////overLen := [2]int{1, 2, 3}
	//2访问数组
	//a := [5]int{5, 4, 3, 2, 1}
	//
	//// 方式1，使用下标读取数据
	//element := a[2]
	//fmt.Println("element = ", element)
	//
	//// 方式2，使用range遍历
	//for i, v := range a {
	//	fmt.Println("index = ", i, "value = ", v)
	//}
	//
	//for i := range a {
	//	fmt.Println("only index, index = ", i)
	//}
	//
	//// 读取数组长度
	//fmt.Println("len(a) = ", len(a))
	//// 使用下标，for循环遍历数组
	//for i := 0; i < len(a); i++ {
	//	fmt.Println("use len(), index = ", i, "value = ", a[i])
	//}
	//b := [6]int{5, 4, 3, 2, 1, 2222}
	//for i := 0; i < len(b); i++ {
	//	fmt.Println("use len(), index = ", i, "value = ", b[i])
	//}
	//3、多维数组
	// 二维数组
	a := [3][2]int{
		{0, 1},
		{2, 3},
		{4, 5},
	}
	fmt.Println("a = ", a)

	// 三维数组
	b := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	fmt.Println("b = ", b)

	// 也可以省略各个位置的初始化,在后续代码中赋值
	c := [3][3][3]int{}
	c[2][2][1] = 5
	c[1][2][1] = 4
	fmt.Println("c = ", c)
}
