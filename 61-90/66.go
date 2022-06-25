package main

import "fmt"

/**
以下代码会输出什么
答案：报错
变量作用域，在main方法里定义的p，会覆盖全局的p变量，运行到bar函数时，p还是nil,nil取指会报错
*/
var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func main() {
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}

/**
正确写法
func main() {
	var err  error
	p,err=foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
*/
