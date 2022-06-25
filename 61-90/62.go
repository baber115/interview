package main

import "fmt"

/**
下面这段代码输出什么？为什么？

答案：s is nil ， p is not nil

解析：
仅当动态值和动态类型都为nil时，接口类型才为nil，这里动态值是nil，但是动态类型是*student
*/
type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func main() {

	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
