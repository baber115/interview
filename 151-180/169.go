package main

/*
下面代码输出什么？
答案：
指针地址 2
指针地址 2
解析：
装饰延迟求值，for循环局部变量i，匿名函数每一次使用的都是同一个变量
*/
func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}
