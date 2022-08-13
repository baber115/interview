package main

import (
	"io/ioutil"
	"os"
)

/*
下面的代码有什么问题，请说明？
答案：defer应该放在if err != nil后面
*/

func main() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}
