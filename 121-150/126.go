package main

/*
下面代码输出什么？
type T struct {
	n int
}

func main() {
	m := make(map[int]T)
	m[0].n = 1
	fmt.Println(m[0].n)
}

答案：报错
解析：map中的struct是无法寻址的，会报错
*/
