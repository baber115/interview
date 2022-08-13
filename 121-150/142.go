package main

/*
下面哪一行代码会 panic，请说明原因？
答案：s, s[0] = []int{1, 2}, 9
解析：左侧的s[0]中的s为nil
*/

func main() {
	var m map[int]bool // nil
	_ = m[123]
	var p *[5]string // nil
	for range p {
		_ = len(p)
	}
	var s []int // nil
	_ = s[:]
	s, s[0] = []int{1, 2}, 9
}
