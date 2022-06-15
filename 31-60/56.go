package main

/**
下面选项正确的是？
A. 1 2
B. compilation error

答案： A
代码块和变量作用域
*/
func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}
