package main

import "fmt"

/*
下面的代码输出什么？
答案：9 [{3} {9}]
解析：参考155，156，157
*/
type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}
