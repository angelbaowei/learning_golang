package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写数据
func fn1(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入 %v len = %v cap = %v\n", i, len(ch), cap(ch))
		time.Sleep(time.Millisecond * 50)
		//time.Sleep(time.Millisecond * 400)
	}
	close(ch)
}
// 读数据
func fn2(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("读取 ", v)
		time.Sleep(time.Millisecond * 400)
		//time.Sleep(time.Millisecond * 50)
	}
}

func main() {

	// channel 管道  用于协程间的通信  先入先出  chan是引用数据类型
	ch1 := make(chan int, 3)

	ch1 <- 10
	ch1 <- 6

	a := <- ch1
	b := <- ch1
	fmt.Println(a, b)  // 10 6

	ch1 <- 1
	fmt.Printf("容量：%v, 长度：%v\n", cap(ch1), len(ch1))  // 3 1

	// 引用类型
	ch2 := ch1
	<-ch2
	fmt.Printf("ch1 容量：%v, 长度：%v\n", cap(ch1), len(ch1))  // 3 0
	fmt.Printf("ch2 容量：%v, 长度：%v\n", cap(ch2), len(ch2))  // 3 0

	// 管道阻塞
	ch3 := make(chan int, 1)
	ch3 <- 1
	//ch3 <- 2  // fatal error: all goroutines are asleep - deadlock!  放不下了

	<-ch3
	//<-ch3  // fatal error: all goroutines are asleep - deadlock! 已经没了

	ch4 := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch4 <- i
	}
	close(ch4)  // 关闭后就不会报错了 关闭chan 是为了防止 for range 死锁

	for v := range ch4 {  // 通过for range遍历管道 前 必须要关闭管道 否则会报错 fatal error: all goroutines are asleep - deadlock!
		fmt.Println(v)
	}

	ch5 := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch5 <- i
	}
	for i := 1; i <= 5; i++ {   // 通过for循环遍历就不用关闭管道就可以遍历
		fmt.Println(<-ch5)
	}
	fmt.Println("--------------------------")

	// 运用chan 一个写协程 一个读协程 并行执行
	// ch6为空时 再读 会等待   ch6满时 再写 也会等待  (*) 这点要好好利用
	ch6 := make(chan int, 3)
	wg.Add(1)
	go fn1(ch6)
	wg.Add(1)
	go fn2(ch6)
	wg.Wait()


	var ch7 = make(chan <- int, 1)  // 只写管道
	ch7 <- 7
	//<- ch7  // invalid operation: <-ch7 (receive from send-only type chan<- int)  报错 不能读
	//var ch8 = make(<- chan int, 1)  // 只读管道
	//ch8 <- 7  // invalid operation: ch8 <- 7 (send to receive-only type <-chan int)  报错 不能写


}
