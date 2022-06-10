package main

// 下面这段代码有什么问题
func funcMui(x, y int) (sum int, error) {
	return x + y, nil
}

/*
解析：
1、有多个返回值，需要加括号()
2、有1个返回值，是命名返回值，也要加括号()
3、有多个返回值，需要都命名，或者都不命名，不能只命名其中一部分
*/
