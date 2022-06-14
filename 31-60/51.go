package main

/**
A、B、C、D 哪些选项有语法错误？

答案：
BD错误
interface可以接收指针，但是也是用interface{}，并不是*interface{}
*/
type S struct {
}

func f51(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f51(s) //A
	g(s)   //B
	f51(p) //C
	g(p)   //D
}
