package calc

import "fmt"

var aaa = 2  // 私有
var AAA = 3	 // 公有

func init() {
	fmt.Println("add_sub.go")
}

func Add(x int, y int) int {  // 公有方法
	return x + y
}

func Sub(x int, y int) int {
	return x - y
}