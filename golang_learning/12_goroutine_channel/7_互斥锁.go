package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup
var mtx sync.Mutex  // 互斥锁

func test1() {
	mtx.Lock()
	count++
	fmt.Println(count)
	time.Sleep(time.Millisecond * 50)
	mtx.Unlock()
	wg.Done()
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go test1()
	}
	fmt.Println("finish")
	wg.Wait()
	/*
	1
	finish
	2
	3
	4
	5
	6
	7
	8
	9
	10
	 */

}
