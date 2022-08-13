package main

/*
下面代码有什么问题？

type foo struct {
	bar int
}

func main() {
	var f foo
	f.bar, tmp := 1, 2
}


答案：短声明不能给结构体内部字段赋值
*/
