package main

import "fmt"

type person struct {
	name string
}

// 这里m的类型是map[person]init,因为m中不存在p，所以打印默认值
// 答案是0
func main() {
	var m map[person]int // int 默认是0，如果是string输出空字符串
	p := person{"make"}
	fmt.Println(m[p])
}
