package main

/*
下面代码能编译通过吗？可以的话，输出什么？
答案：4 3
*/
func main() {
	println(DeferTest1(1))
	println(DeferTest2(1))
}

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}
