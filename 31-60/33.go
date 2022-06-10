package main

import "fmt"

// 以下代码输出什么

/**
答案：
showA
showB

结构体嵌套，teacher里嵌套了people
*/
type people struct {
}

func (p *people) showA() {
	fmt.Println("showA")
	p.showB()
}

func (p *people) showB() {
	fmt.Println("showB")
}

type teacher struct {
	people
}

func (t *teacher) showB() {
	fmt.Println("teacher showA")
}

func main() {
	t := teacher{}
	t.showA()
}
