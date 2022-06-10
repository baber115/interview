package main

/**
new()与make()的区别
new和make都是内置函数，用来分配内存，但是适用类型不同
new(T)是新建并返回一个空指针，此处为*T，适用于数组和结构体
make(T,3)是返回初始化之后的T值，不是指针，是初始化之后的引用，适用于slice,map和channel
*/
