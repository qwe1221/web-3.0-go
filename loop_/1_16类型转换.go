package main

import "fmt"

func main() {
	//1、数字类型转换
	var i int32 = 17
	var b byte = 5
	var f float32

	// 数字类型可以直接强转
	f = float32(i) / float32(b)
	fmt.Printf("f 的值为: %f\n", f)

	// 当int32类型强转成byte时，高位被直接舍弃
	var i2 int32 = 256
	var b2 byte = byte(i2)
	fmt.Printf("b2 的值为: %d\n", b2)
	//字符串类型转换
	str := "hello, 123, 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes 的值为: %v \n", bytes)
	fmt.Printf("runes 的值为: %v \n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 的值为: %v \n", str2)
	fmt.Printf("str3 的值为: %v \n", str3)
}
