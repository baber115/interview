package main

/*
下面哪一行代码会 panic，请说明原因？
答案：
	_ = y == y
切片是无法比较的，因为两个比较值的动态类型为同一个不可比较类型。
*/
func main() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	_ = y == y
}
