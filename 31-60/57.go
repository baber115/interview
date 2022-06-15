package main

import "fmt"

/**
下面这段代码输出什么？

答案
0 zero
1 one
// 或者
1 one
0 zero

map 的输出是无序的
*/
func main() {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
