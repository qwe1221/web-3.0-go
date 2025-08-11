package main

import (
	"fmt"
	"strconv"
)

// 判断一个整数是否是回文数
// 一个数正面的值和反面的值一样
func isHuiWen(number int) bool {
	if number < 0 {
		return false
	}
	if number < 10 {
		return true
	}
	if number%10 == 0 {
		return false
	}
	str := strconv.Itoa(number)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	num := 11112221111
	fmt.Print(num, map[bool]string{true: "是回文", false: "不是回文"}[isHuiWen(num)])
}
