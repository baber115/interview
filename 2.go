package main

import "fmt"

/**
官方解释：
for range 循环的时候会创建每个元素的副本，而不是每个元素的引用，所以 m[key] = &val 取的都是变量val的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出的都是3。
自我理解：
1、for range循环是拷贝，不是引用，不会对原数据进行修改
2、m[key]=&val，取的是val的地址，val的地址是同一个，for循环一直在给同一个地址赋值，所以最后一个值就是val所对应的地址的值

输出：
* map的顺序是乱序的
0->3
1->3
2->3
*/
func main() {
	s := []int{1, 2, 3}
	m := make(map[int]*int)

	for key, val := range s {
		/**
		此处打印的&val，是同一个地址
		0 0xc0000120a8
		1 0xc0000120a8
		2 0xc0000120a8
		*/
		fmt.Println(key, &val)
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
