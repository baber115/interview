package main

import (
	"fmt"
	"sync"
)

/*
下面的代码有什么问题？
答案：
1、在协程中使用wg.Add；
2、使用了sync.WaitGroup副本
*/
func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	fmt.Println("exit")

	/*
		修改后
		wg := sync.WaitGroup{}

		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(i int) {
				fmt.Printf("i:%d\n", i)
				wg.Done()
			}(i)
		}

		wg.Wait()

		fmt.Println("exit")
	*/
}
