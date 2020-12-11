package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func test2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生error  但是继续执行其他goroutine")
		}
	}()
	var maze map[int]int
	maze[0] = 1  // 没有给maze分配空间 所以这样写是错误的
	wg.Done()
}

func main() {
	wg.Add(1)
	go test1()
	wg.Add(1)
	go test2()
	wg.Wait()
	// 此时执行 因为test2()的错误  导致test1()方法也没有执行
	// 为了不影响其他函数执行 使用panic recover  这样就是test2() 报错 但是不影响test1() 正常执行



}
