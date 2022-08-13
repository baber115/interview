package main

import "fmt"

/*
下面代码有什么不规范的地方吗？
答案：这样写条件永远成立，不存在的key对应的value也是空字符串
解析：判断map里是否有某一元素，不能用这种方式，应该用map的第二个返回值
*/
func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	if v := x["two"]; v == "" {
		fmt.Println("no entry")
	}

	// 修改后
	if _, ok := x["two"]; ok {
		fmt.Println("no entry")
	}
}
