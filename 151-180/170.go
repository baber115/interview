package main

/*
下面的代码能编译通过吗？可以的话输出什么，请说明？
答案：10x
解析：f(i - 1)这里的f是用的之前var定义的f
*/
var f = func(i int) {
	print("x")
}

func main() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}
