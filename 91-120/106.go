package main

/*
下面代码输出什么？
A. true true true
B. false true true
C. true true true
D. false true false
答案：D
解析：
知识点：类型断言。
类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败。
*/
func main() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c)
}
