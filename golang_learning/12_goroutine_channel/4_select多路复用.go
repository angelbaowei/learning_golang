package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	intChan := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 10)
	for i := 1; i <= 10; i++ {
		stringChan <- strconv.FormatInt(int64(i), 10) + " string"
	}

	// 并行执行  在一个方法里面 从多个chan里面并行取数据 时可以用select
	for {
		select {  // 使用select时一定不要close chan
		case v := <- intChan:
			fmt.Println("读int ", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Println("读string ", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("over")
			return  // 读取完毕就结束
		}
	}
	/*
	读string  1 string
	读int  1
	读string  2 string
	读int  2
	读string  3 string
	读int  3
	读int  4
	读string  4 string
	读string  5 string
	读string  6 string
	读int  5
	读string  7 string
	读int  6
	读string  8 string
	读int  7
	读int  8
	读string  9 string
	读int  9
	读int  10
	读string  10 string
	over
	*/
}
