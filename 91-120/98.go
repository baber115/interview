package main

/*
下面这段代码存在什么问题？

答案：存在两个问题

解析：
1、map 需要初始化才能使用；
2、指针不支持索引。修复代码如下：
*/
type Param map[string]interface{}

type Show struct {
	*Param
}

func main() {
	s := new(Show)
	s.Param["day"] = 2
}
