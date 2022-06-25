package main

import "fmt"

type direction int

/*
以下代码输出什么

答案：South
知识点：iota 的用法、类型的 String() 方法。
*/
const (
	North direction = iota
	East
	South
	West
)

func (d direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(South)
}
