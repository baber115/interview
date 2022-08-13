package main

import (
	"encoding/json"
	"fmt"
)

/*
下面这段代码输出什么？
答案：{}，空结构体
解析：因为people里的结构体是小写，其他包不能访问，所以返回空结构体
*/
type People struct {
	name string `json:"name"`
}

func main() {
	js := `{
        "name":"seekload"
    }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}
