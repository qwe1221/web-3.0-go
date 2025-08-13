package main

import "fmt"

/*
*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func zcggqz(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		//比较当前字符串和prefix公共前缀
		j := 0
		if j < len(prefix) && j < len(strs[i]) && strs[i][j] == prefix[j] {
			j++
		}
		prefix = prefix[:j]
		if prefix == "" {
			return ""
		}
	}
	return prefix
}
func main() {
	strs := []string{"flower", "flow", "fdight"}
	fmt.Print(zcggqz(strs))
}
