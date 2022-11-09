package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "i can do all things！"
	fmt.Println(wordCount(s))

}
func wordCount(s string) map[string]int {
	m := make(map[string]int)
	c := strings.Fields(s) //返回字符串切片
	// 	这里用到了内置的string包的Field方法来分隔单词，具体来讲，Field功能如下：
	// 1. Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身；
	// 2. 空白字符有：\t, \n, \v, \f, \r, ’ ', U+0085 (NEL), U+00A0 (NBSP)；
	// 3. 如果 s 中只包含空白字符，则返回一个空列表。
	// 由此看来，这种计词方法只能对一些简单的字符串起效，而如果遇上了"how do you do?"，你就会发现“do?”也被计词一次。

	for _, v := range c {
		m[v] += 1
	}
	return m

}
