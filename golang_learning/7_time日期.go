package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取时间
	time1 := time.Now()
	fmt.Println(time1)
	// year = time1.Year()   年月日是分秒都有
	var str = time1.Format("2006-01-02 03:04:05")  // 这个2006 01 02 03 04 05 是固定的  12小时制  因为golang是2006年诞生的
	fmt.Println(str)
	str = time1.Format("2006-01-02 15:04:05")  // 这个2006 01 02 15 04 05 是固定的  24小时制
	fmt.Println(str)

	// 获取时间戳
	time2 := time.Now()
	str = time2.Format("2006-01-02 15:04:05")
	fmt.Println(str)
	unixtime := time2.Unix()  // 日期转时间戳 精确到秒
	fmt.Println(unixtime)
	time2 = time.Unix(unixtime, 0)  // 时间戳转日期
	str = time2.Format("2006-01-02 15:04:05")
	fmt.Println(str)

	// 字符串转时间戳
	var str1 = "2020/01/02 15:15:45"
	var mode = "2006/01/02 15:04:05"
	var time3, _ = time.ParseInLocation(mode, str1, time.Local)  // 字符串根据模版先转化成日期
	fmt.Println(time3)
	unixtime = time3.Unix() // 日期再转时间戳
	fmt.Println(unixtime)

	// 时间操作
	time4 := time.Now()
	fmt.Println(time4)
	time4 = time4.Add(time.Hour)
	fmt.Println(time4)

	//定时器
	ticker1 := time.NewTicker(time.Second)  // 每隔一秒的定时器
	n := 1
	for t := range ticker1.C {
		fmt.Println(t)    // 每隔一秒执行一次
		n++
		if n > 5 {
			ticker1.Stop()  // 停止定时器
			break  // 记得break
		}
	}

	for i := 1; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)  // 休眠一秒
	}

}
