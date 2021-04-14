package main

import (
	"fmt"
	"sync"
	"time"
)

type singleton struct {

}
var mu sync.Mutex
var wg sync.WaitGroup

var instance *singleton

// 单例  check-lock-check模式

func GetInstance() *singleton {
	// 1  错误的方式
	//if instance == nil {
	//	fmt.Println("确保这句话只出现了一次")
	//	instance = &singleton{}   // 不是并发安全的
	//}
	// 2  太严格了 效率不高  948.205µs
	//mu.Lock()                    // 如果实例存在没有必要加锁
	//defer mu.Unlock()
	//
	//if instance == nil {
	//	fmt.Println("确保这句话只出现了一次")
	//	instance = &singleton{}
	//}
	// 3 check-lock-check模式  683.022µs
	if instance == nil {     // 不太完美 因为这里不是完全原子的
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			fmt.Println("确保这句话只出现了一次")
			instance = &singleton{}
		}
	}
	defer wg.Done()
	return instance
}

func main() {
	start := time.Now()
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go GetInstance()
	}
	wg.Wait()
	dual := time.Since(start)
	fmt.Println(dual)
}
