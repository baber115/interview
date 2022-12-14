package main

/*
下面代码有什么问题？
func main() {
    x := []int{
        1,
        2
    }
    _ = x
}
答：编译错误
解析：第四行代码没有逗号。用字面量初始化数组、slice 和 map 时，最好是在每个元素后面加上逗号，即使是声明在一行或者多行都不会出错。
*/

// 修复代码：
func main() {
	x := []int{ // 多行
		1,
		2,
	}
	x = x

	y := []int{3, 4} // 一行 no error
	y = y
}
