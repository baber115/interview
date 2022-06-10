package main

import "fmt"

// 以下代码输出什么

/**
答案：0
删除map中不存在的键时不会发生任何事，获取map中不存在的键值对时，由于是int，所以返回0
*/
func main() {
	s := make(map[string]int) // int改成string返回空字符串
	delete(s, "h")
	fmt.Println(s["h"])
}
