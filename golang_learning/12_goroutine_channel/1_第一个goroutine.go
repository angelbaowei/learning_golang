package main

import (
	"fmt"
	"time"
)

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " test")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	// goroutine 协程  用户级线程  开销小
	go test1()  // 方法前加一个 go  表示开启一个协程  并行

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 50)
	}
	// 主线程执行完毕后直接退出 test协程不论执行完没有都会强制退出  即 主线程不等待协程
	/*
	0
	0  test
	1
	1  test
	2
	3
	2  test
	4
	5
	3  test
	6
	7
	4  test
	8
	9
	5  test
	 */


}
