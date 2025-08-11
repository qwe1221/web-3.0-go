package main

import "fmt"

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
// 其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
// 结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func main() {
	ch := [8]string{"a", "b", "c", "d", "e", "f", "g", "a"}
	fmt.Println(ch)
	countMap := make(map[string]int)
	for _, value := range ch {
		//如果key相同则value+1
		countMap[value]++
	}
	//找出值为1的记录
	fmt.Println("出现值为1的数为:")
	for key, value := range countMap {
		if value == 1 {
			fmt.Println(key, value)
		}
	}
}
