package main

/*
下面代码输出什么？
答案：100，100

*/
func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()
}
