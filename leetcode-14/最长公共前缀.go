package main

import "fmt"

func main() {
	str_slice := []string{"flower", "flow", "flight"}
	s := longestCommonPrefix(str_slice)
	fmt.Printf("s:%v\n", s)
}

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(str_slice []string) string {
	if len(str_slice) == 0 {
		return ""
	}

	// 以第一个数的长度为基准
	for i := 0; i < len(str_slice[0]); i++ {
		for j := 1; j < len(str_slice); j++ {
			if i == len(str_slice[j]) || str_slice[j][i] != str_slice[0][i] {
				return str_slice[0][:i]
			}
		}
	}
	return str_slice[0]
}
