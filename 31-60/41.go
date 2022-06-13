package main

import "fmt"

/**
下面代码中 A B 两处应该怎么修改才能顺利编译？
答案:
A处只声明了map，没有分配内存空间，无法赋值，推荐使用make
B处,v, k := m["b"],当key不存在时，v=0,k=false
*/
func main() {
	var m map[string]int //A
	m["a"] = 1
	if v := m["b"]; v != 0 { //B
		fmt.Println(v)
	}
}
