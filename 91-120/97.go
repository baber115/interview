package main

/*
解析：有方向的channel不可以被关闭
*/
func Stop(stop <-chan bool) {
	close(stop)
}
