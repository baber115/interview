package main

/**
以下代码输出什么

答案：23  23  23
解析：
for循环中i,v 是每次重新赋值，而不是重新声明，地址不变，不停的赋值
各个goroutine中输出的i,v，输出的最后值，也就是循环最后的值，而不是各个goroutine启动时的值
*/
import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

/**
正确方法一
func main() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func(i,v int) {
			fmt.Println(i, v)
		}(i,v)
	}

	time.Sleep(time.Second * 3)
}
*/

/**
正确方法二
使用临时变量
func main() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		i := i
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}
*/
