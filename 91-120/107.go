package main

/*
下面代码有几处错误的地方？请说明原因。
答案：1处错误
解析：不能对未初始化的map直接赋值，应该先make()初始化，但是切片可以用append直接增加元素
*/
func main() {
	var s []int
	s = append(s, 1)

	var m map[string]int
	m["one"] = 1
}
