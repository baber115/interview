package main

import "fmt"

/**
下面这段代码输出什么？为什么？

答案
A. speak
B. compilation error

Student没有实现接口Speak方法，而是*Student实现了
*/
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	//var peo People = &Student{}
	var peo People = Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
