package main

/**
以下能否正常结束，如果能正常结束输出什么

答案：[1 2 3 0 1 2]
解析：循环次数在之前就确定了，循环往v里插入从0开始的数据

*/
import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}

	fmt.Println(v)
}
