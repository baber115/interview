package main

/*
下面的代码输出什么，请说明。
答案：312
解析：对比166题，本题的s.Add(1).Add(2)作为一个整体包在一个匿名函数中，会延迟执行
*/
import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSlice()
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}
