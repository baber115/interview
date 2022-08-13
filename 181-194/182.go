package main

import "fmt"

/*
下面的代码有什么问题？
答案：无限递归
解析，如果定义了String()方法，在使用print的时候会自动调用string方法
*/
type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func main() {
	c := &ConfigOne{}
	c.String()
}
