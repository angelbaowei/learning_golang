package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// wg.Add(1) wg.Done() wg.Wait()
var wg sync.WaitGroup

func test2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i, " test2")
		time.Sleep(time.Millisecond * 100)
	}
}

func test3() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i, " test3")
		time.Sleep(time.Millisecond * 150)
	}
}

func main() {
	// 获取CPU数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum = ", cpuNum)

	// 指定程序使用的CPU数量
	runtime.GOMAXPROCS(3)
	fmt.Println("OK")

	wg.Add(1)
	go test2()
	wg.Add(1)
	go test3()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()  // 主线程等待所有线程执行完
	/*
	0
	0  test
	1
	1  test
	2
	3
	4
	2  test
	5
	6
	3  test
	7
	4  test
	8
	9
	5  test
	6  test
	7  test
	8  test
	9  test
	 */

}
