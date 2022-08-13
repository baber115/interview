package main

import "fmt"

/*
下面代码能否编译通过？如果通过，输出什么？

答案：报错

解析：GetValue里返回的参数string,bool，不能返回nil
*/

func GetValue(m map[int]string, id int) (string, bool) {

	if _, exist := m[id]; exist {
		return "exist", true
	}
	return nil, false
}
func main() {
	intmap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}
