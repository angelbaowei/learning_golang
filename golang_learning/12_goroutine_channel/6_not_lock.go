package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup

func test1() {
	count++
	fmt.Println(count)
	time.Sleep(time.Millisecond * 50)
	wg.Done()
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go test1()
	}
	fmt.Println("finish")
	wg.Wait()
	// go build -race 6_not_lock.go  6_not_lock.exe  执行   Found 2 data race(s) 会报错 竞争关系

	/*
	8
	7
	finish
	3
	1
	4
	2
	5
	6
	9
	10
	 */

}
