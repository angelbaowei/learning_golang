package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
// 读写锁 读并发 写互斥
var rwmtx sync.RWMutex

func write() {
	rwmtx.Lock()  // 写互斥
	fmt.Println("Wtire---")
	time.Sleep(time.Second)
	rwmtx.Unlock()
	wg.Done()
}

func read() {
	rwmtx.RLock()  // 读共享
	fmt.Println("---Read")
	time.Sleep(time.Second)
	rwmtx.RUnlock()
	wg.Done()
}

func main() {

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	/*
	Wtire---
	---Read
	---Read
	---Read
	---Read
	---Read
	---Read
	---Read
	---Read
	---Read
	---Read
	Wtire---
	Wtire---
	Wtire---
	Wtire---
	Wtire---
	Wtire---
	Wtire---
	Wtire---
	Wtire---
	 */

}
